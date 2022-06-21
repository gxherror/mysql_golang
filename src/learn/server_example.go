package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

type greeting string

func WithLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  fmt.Printf("path:%s process start...\n", r.URL.Path)
	  defer func() {
	  fmt.Printf("path:%s process end...\n", r.URL.Path)
	  }()
	  handler.ServeHTTP(w, r)
	})
  }

  func PanicRecover(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  defer func() {
		if err := recover(); err != nil {
		  fmt.Println("Error")
		}
	  }()
  
	  handler.ServeHTTP(w, r)
	})
  }


  func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares)-1; i >= 0; i-- {
	  handler = middlewares[i](handler)
	}
  
	return handler
  }  

  func Metric(handler http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
	  start := time.Now()
	  defer func() {
		fmt.Printf("path:%s elapsed:%fs\n", r.URL.Path, time.Since(start).Seconds())
	  }()
	  time.Sleep(1 * time.Second)
	  handler.ServeHTTP(w, r)
	}
  }


func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g)
  }

  func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
  }
  



func begin()  {
	mux := http.NewServeMux()
	middlewares := []Middleware{
		PanicRecover,
		WithLogger,

	  }
	  mux.Handle("/", applyMiddlewares(http.HandlerFunc(index), middlewares...))
	  mux.Handle("/greeting", applyMiddlewares(greeting("welcome, dj"), middlewares...))
  
	server := &http.Server{
	  Addr:         ":8080",
	  Handler:      mux,
	  ReadTimeout:  20 * time.Second,
	  WriteTimeout: 20 * time.Second,
	}
	server.ListenAndServe()
}