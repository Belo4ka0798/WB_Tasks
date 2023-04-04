// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func Worker(ch chan interface{}, ctx context.Context, ctx2 context.Context, i int) {
	for {
		select {
		case <-ctx2.Done():
			return
		case <-ctx.Done():
			data, ok := <-ch
			if !ok {
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
	var b int
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatal()
	}
	ctx2, sig := context.WithTimeout(context.Background(), time.Duration(b)*time.Second)
	for i := 0; i < a; i++ {
		go Worker(ch, ctx, ctx2, i)
	}
	for {
		select {
		case <-ctx2.Done():
			sig()
			log.Println("TIMEOUT1!")
			return
		case <-ctx.Done():
			cancel()
			log.Println("CANCEL1!")
			return
		default:
			ch <- time.Now().UTC()
			time.Sleep(time.Second)
		}
	}
}
