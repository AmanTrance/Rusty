package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/cespare/xxhash/v2"
	_ "github.com/dgryski/go-rendezvous"
	"github.com/redis/go-redis/v9"
)

const (
	SECRET_KEY = "abcdefghijklmno1"
)

type Message struct {
	response http.ResponseWriter
	request  *http.Request
	channel  *chan uint8
	client   *redis.Client
}

type ClientRequest struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
	Id   int       `json:"id"`
}

type GetRequestData struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
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
				msg.response.WriteHeader(http.StatusMethodNotAllowed)
				*msg.channel <- 0
				break
			}
			var body ClientRequest
			jsonErr := json.NewDecoder(msg.request.Body).Decode(&body)
			if jsonErr != nil {
				msg.response.WriteHeader(http.StatusBadRequest)
				*msg.channel <- 0
				break
			}
			byteData, MarshalErr := json.Marshal(body)
			if MarshalErr != nil {
				msg.response.WriteHeader(http.StatusInternalServerError)
				*msg.channel <- 0
				break
			}
			stringifyData, EncryptErr := EncryptData(byteData)
			if EncryptErr != nil {
				msg.response.WriteHeader(http.StatusInternalServerError)
				*msg.channel <- 0
				break
			}
			RedisErr := msg.client.RPush(ctx, "RandomData", stringifyData).Err()
			if RedisErr != nil {
				msg.response.WriteHeader(http.StatusInternalServerError)
				*msg.channel <- 0
				break
			}
			msg.response.WriteHeader(http.StatusOK)
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
				msg.response.WriteHeader(http.StatusMethodNotAllowed)
				*msg.channel <- 0
				break
			}
			data, err := msg.client.LRange(ctx, "RandomData", 0, -1).Result()
			if err != nil {
				msg.response.WriteHeader(http.StatusInternalServerError)
				*msg.channel <- 0
				break
			}
			var body GetRequestData
			jsonError := json.NewDecoder(msg.request.Body).Decode(&body)
			if jsonError != nil {
				msg.response.WriteHeader(http.StatusBadRequest)
				*msg.channel <- 0
				break
			}
			for _, i := range data {
				originalByteData, decryptError := DecryptData([]byte(i))
				if decryptError != nil {
					println(decryptError.Error())
				}
				reader := bytes.NewReader(originalByteData)
				var clientData ClientRequest
				unmarshalError := json.NewDecoder(reader).Decode(&clientData)
				if unmarshalError != nil {
					println(unmarshalError.Error())
				}
				if body.From.Compare(clientData.Date) == -1 && body.To.Compare(clientData.Date) == 1 {
					println(clientData.Name, clientData.Id, clientData.Date.String())
				} else {
					continue
				}
			}
			msg.response.WriteHeader(http.StatusOK)
			*msg.channel <- 0
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

func EncryptData(data []byte) (string, error) {
	block, err := aes.NewCipher([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	cfgEncrypter := cipher.NewCFBEncrypter(block, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	cipherText := make([]byte, len(data))
	cfgEncrypter.XORKeyStream(cipherText, data)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptData(data []byte) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	block, cipherError := aes.NewCipher([]byte(SECRET_KEY))
	if cipherError != nil {
		return nil, cipherError
	}
	cfgDecrypter := cipher.NewCFBDecrypter(block, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	originalData := make([]byte, len(data))
	cfgDecrypter.XORKeyStream(originalData, cipherText)
	return originalData, nil
}
