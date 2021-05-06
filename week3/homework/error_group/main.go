package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func httpServer(ctx context.Context) error {
	defer func() {
		fmt.Println("http server done...")
	}()

	s := http.Server{
		Addr: ":8080",
	}

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("error group done...")
			s.Close()
		}
	}()
	log.Println("http server started...")
	return s.ListenAndServe()
}

func sigHandler(ctx context.Context) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	log.Println("signal handler started...")
	select {
	case sig := <-sigCh:
		return fmt.Errorf("get signal : %v", sig)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return httpServer(ctx)
	})

	group.Go(func() error {
		return sigHandler(ctx)
	})

	if err := group.Wait(); err != nil {
		log.Fatal(err)
	}
}
