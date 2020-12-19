package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/bar", barPage)

	http.ListenAndServe(":3000", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "userID", 6392)
	ctx = context.WithValue(ctx, "fname", "Bond")

	userID := dbAccess(ctx)

	fmt.Fprintln(w, userID)
}

func dbAccess(ctx context.Context) int {
	userID := ctx.Value("userID").(int)
	return userID
}

func barPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Fprintln(w, ctx)
}
