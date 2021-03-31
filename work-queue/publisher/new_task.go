// Copyright 2012-2019 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
)

func main() {

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	failOnError(err, "Failed to connect to NATS server")
	defer nc.Close()

	msg := []byte(bodyFrom(os.Args))

	nc.Publish("task", msg)
	nc.Flush()

	err = nc.LastError()
	failOnError(err, "nats server error")
	log.Printf("Published [task] : '%s'\n", msg)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
