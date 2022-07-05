package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getHello1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	io.WriteString(w, "Hello1!\n")
}
func getHello2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	io.WriteString(w, "Hello2\n")
}
func getHello3(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	io.WriteString(w, "Hello3\n")
}
func getHello4(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	io.WriteString(w, "Hello4\n")
}
func getHello5(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	io.WriteString(w, "Hello5\n")
}

func main() {
	port := getEnv("PORT", "8080")
	http.HandleFunc("/hello1", getHello1)
	http.HandleFunc("/hello2", getHello2)
	http.HandleFunc("/hello3", getHello3)
	http.HandleFunc("/hello4", getHello4)
	http.HandleFunc("/hello5", getHello5)
	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
