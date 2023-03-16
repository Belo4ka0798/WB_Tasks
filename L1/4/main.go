package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(ch chan interface{}, ctx context.Context, wg *sync.WaitGroup, i int) {
	for {
		select {
		case <-ctx.Done():
			data, ok := <-ch
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(data, "saved")
		default:
			fmt.Println(<-ch, "   ", i)
		}
	}
}

func main() {
	ch := make(chan interface{})
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatal()
	}
	for i := 0; i < a; i++ {
		go Worker(ch, ctx, nil, i)
	}
	for {
		select {
		case <-ctx.Done():
			cancel()
			log.Println("CANCEL")
			return
		default:
			ch <- time.Now().UTC()
			time.Sleep(time.Second)
		}
	}
}
