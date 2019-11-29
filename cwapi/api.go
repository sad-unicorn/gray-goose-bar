package cwapi

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var apiConnection *amqp.Connection
var outbound *amqp.Channel
var apiLogin string
var apiInitialized bool

func InitAPI(login, password string) {
	dial(login, password)
	inbound := connectToApi()

	go func() {
		for d := range inbound {
			log.Println("At API: ", string(d.Body))
		}
	}()
	apiInitialized = true
}


func dial(login, password string) {
	apiLogin = login
	rabbitUrl := fmt.Sprintf("amqps://%s:%s@api.chatwars.me:5673/", login, password)
	api, err := amqp.Dial(rabbitUrl)
	panicOnError(err, "Cannot connect to api")
	apiConnection = api
}

var mutex sync.Mutex
func queryApi(payload string) {
	if !apiInitialized {
		log.Println("Not published: ", payload)
		return
	}
	if outbound == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if outbound == nil {
			ch, err := apiConnection.Channel()
			if err != nil {
				return // what do?
			}
			outbound = ch
		}
	}
	err := outbound.Publish(apiLogin +"_ex", apiLogin+ "_o", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(payload),
	})

	if err != nil {
		log.Println("Error while publishing on API: ", err)
		outbound = nil
	} else {
		log.Println("Published: ", payload)
	}
}

func connectToApi() <-chan amqp.Delivery {
	ch, err := apiConnection.Channel()
	panicOnError(err, "Error while oppenning channel")
	msgs, err := ch.Consume(
		apiLogin + "_i", // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	panicOnError(err, "Error while trying to connect to queue")
	log.Println("Consuming messages from queue")
	return msgs
}


func panicOnError(err error, reason string) {
	if err != nil {
		log.Panic(reason, ":", err)
	}
}

func Publish() {
	queryApi("{}")
}