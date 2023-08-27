package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	resultCh := make(chan string, 4)
	ctx, cancel := context.WithCancel(context.Background())
	services := []string{"Super", "Villagemobil", "Sett Taxi", "Index Go"}
	var (
		wg     sync.WaitGroup
		winner string
	)

	defer cancel()

	for _, v := range services {
		svc := v

		wg.Add(1)
		go func() {
			defer wg.Done()
			requestRide(ctx, svc, resultCh)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		winner = <-resultCh
		cancel()
	}()

	wg.Wait()
	log.Printf("found car in %q", winner)
}

func requestRide(ctx context.Context, serviceName string, resultCh chan<- string) {
	time.Sleep(3 * time.Second)

	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped the search in %q (%v)", serviceName, ctx.Err())
			return
		default:
			if rand.Float64() > 0.75 {
				resultCh <- serviceName
				return
			}
		}
	}
}
