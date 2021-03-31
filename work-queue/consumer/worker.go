package main

import (
	"bytes"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	failOnError(err, "Failed to connect NATS server")

	nc.Subscribe("task", func(msg *nats.Msg) {
		log.Printf("Receive Task.")
		dotCount := bytes.Count(msg.Data, []byte("."))
		t := time.Duration(dotCount)
		time.Sleep(t * time.Second)
		log.Printf("Task Done")
	})
	nc.Flush()

	err = nc.LastError()
	failOnError(err, "nats server error")

	log.Printf("Listening on task")

	runtime.Goexit()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
