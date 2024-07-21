package appilication

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/redis/v9"
)

type App struct {
	router http.Handler
	rdb *redis.Client
}

func New() *App {
	
	app := &App{
		router: loadRoutes(),
		rdb: redis.NewClient(&redis.Options{}),
	}

	return app

}

// to start the redis serverin windows --> run on the wsl follow redis documentation

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: ":3000",
		Handler: a.router,
	}

	err:= a.rdb.Ping(ctx).Err()

	return fmt.Errorf("failed to ping redis: %w", err)
	}

	defer func ()  {
		err := a.rdb.Close()
		if err != nil {
			fmt.Println("failed to close redis connection", err)
		}
	}()

	fmt.Println("Starting the server")

	ch := make(chan error, 1)

	go func ()  {

		err= server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to listen to the server: %w", err) //publishin the error to the channel
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(timeout)
	}


	// return nil

}