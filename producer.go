//+build ignore
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	for {
		topic := "messages_topic"
		partition := 0
		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		randomString := RandomString(50)
		// to produce messages

		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(randomString)},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)

		}

		fmt.Println("Produce: " + randomString)

		time.Sleep(1 * time.Second)
	}
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
