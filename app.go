package main

import (
	"fmt"
	"gopkg.in/redis.v3"
	"log"
	"net/http"
)

const (
	Port     = ":8080"
	CountKey = "counter"
)

var rd *redis.Client

func redisNewClient() {
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func incrFunc(w http.ResponseWriter, r *http.Request) {
	if err := rd.Incr(CountKey).Err(); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "incr")
}

func countFunc(w http.ResponseWriter, r *http.Request) {
	val, err := rd.Get(CountKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, val)
}

func main() {
	http.HandleFunc("/incr", incrFunc)
	http.HandleFunc("/count", countFunc)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
