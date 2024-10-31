package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Worker")

	nc, err := nats.Connect("nats://queue:4222", nats.Token("qwerty"))
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := nc.Drain(); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan *nats.Msg, 1000)

	sub, err := nc.ChanSubscribe("test", ch)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := sub.Drain(); err != nil {
			log.Println(err)
		}
	}()

	for n := range 10 {
		go func() {
			for msg := range ch {
				// if err := msg.Ack(); err != nil {
				// 	log.Println(err)
				// }

				log.Println(n)

				if err := msg.Respond(msg.Data); err != nil {
					log.Println(err)
				}
			}
		}()
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan
}
