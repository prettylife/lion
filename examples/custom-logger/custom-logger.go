package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/celrenheit/lion"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := lion.C(c)
	fmt.Fprintf(w, "Hello "+ctx.Param("name"))
}

type logger struct{}

func (*logger) ServeNext(next lion.Handler) lion.Handler {
	return lion.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		fmt.Printf("Served %s in %s\n", r.URL.Path, time.Since(start))
	})
}

func main() {
	l := lion.New()
	l.Use(&logger{})
	l.GetFunc("/", home)
	l.GetFunc("/hello/:name", hello)
	l.Run(":3000")
}
