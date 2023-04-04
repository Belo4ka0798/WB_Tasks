// Реализовать все возможные способы остановки выполнения горутины.
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

func Worker(ch chan interface{}, ctx context.Context, chi chan int, wg *sync.WaitGroup, i int) {
	for {
		select {
		case tmp := <-chi:
			if tmp == 1 {
				return
			}
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
	chi := make(chan int, a)
	var b int
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatal()
	}
	for i := 0; i < a; i++ {
		go Worker(ch, ctx, chi, nil, i)
	}

	for i := 0; i < a; i++ {
		chi <- b
	}
	for {
		select {
		case tmp := <-chi:
			if tmp == 1 {
				log.Println("CHANNEL PUT 1!")
				return
			}
		case <-ctx.Done():
			cancel()
			log.Println("CANCEL WITH CONTEXT")
			return
		default:
			ch <- time.Now().UTC()
			time.Sleep(time.Second)
			break
		}
	}
	//log.Println("CANCEL WITH END MAIN FUNC!")
}
