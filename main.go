package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe(":8080",
			http.Handler(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						time.Sleep(10 * time.Second)
						w.WriteHeader(200)
						fmt.Println(r)
					},
				),
			),
		)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	cancel()

	httpClient := http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout:   0,
			ResponseHeaderTimeout: 0,
			ExpectContinueTimeout: 0,
		},
		Timeout: 10 * time.Second,
	}

	_, err = httpClient.Do(req)
	var netErr net.Error
	fmt.Println(errors.As(err, &netErr))
	fmt.Println(netErr.Timeout())

}
