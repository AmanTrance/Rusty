package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/cespare/xxhash/v2"
	_ "github.com/dgryski/go-rendezvous"
	"github.com/redis/go-redis/v9"
)

type Message struct {
	response http.ResponseWriter
	request  *http.Request
	channel  *chan uint8
	client   *redis.Client
}

type ClientRequest struct {
	Name string     `json:"name"`
	Date *time.Time `json:"date"`
	Id   int        `json:"id"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	channel1 := make(chan Message)
	channel2 := make(chan Message)
	defer close(channel1)
	defer close(channel2)

	go Goroutine1(ctx, &channel1)
	go Goroutine2(ctx, &channel2)

	http.HandleFunc("/endpoint1", AddChannelAndClientToFunc(&channel1, client))
	http.HandleFunc("/endpoint2", AddChannelAndClientToFunc(&channel2, client))

	http.ListenAndServe(":8000", nil)
}

func Goroutine1(ctx context.Context, channel *chan Message) {
outer:
	for {
		select {
		case msg := <-*channel:
			if msg.request.Method != "POST" {
				msg.response.WriteHeader(http.StatusBadRequest)
				*msg.channel <- 0
				break
			}
			var body ClientRequest
			err := json.NewDecoder(msg.request.Body).Decode(&body)
			if err != nil {
				msg.response.WriteHeader(http.StatusBadRequest)
				*msg.channel <- 0
				break
			}
			*msg.channel <- 0
			break
		case <-ctx.Done():
			break outer
		}
	}
}

func Goroutine2(ctx context.Context, channel *chan Message) {
outer:
	for {
		select {
		case msg := <-*channel:
			if msg.request.Method != "GET" {
				msg.response.WriteHeader(http.StatusBadRequest)
				*msg.channel <- 0
				break
			}
			break
		case <-ctx.Done():
			break outer
		}
	}
}

func AddChannelAndClientToFunc(channel *chan Message, client *redis.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan uint8)
		defer close(ch)
		*channel <- Message{
			response: w,
			request:  r,
			channel:  &ch,
			client:   client,
		}
		select {
		case <-ch:
			return
		}
	}
}
