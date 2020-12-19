package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", mainPage)
	// http.HandleFunc("/bar", barPage)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":3000", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "userID", 48247)
	ctx = context.WithValue(ctx, "userName", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, "Timeout while accessing database", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		// Some task which take some time
		fmt.Println(ctx.Value("userID"))
		userID := ctx.Value("userID").(int)

		time.Sleep(10 * time.Second)

		// check for error
		if ctx.Err() != nil {
			return
		}

		ch <- userID
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case userID := <-ch:
		return userID, nil
	}
}
