package main

import (
	"github.com/stripe/stripe-go/v72/customer"
	"log"
	"net/http"
)


func HandleEvent(w http.ResponseWriter, req * http.Request)  {
	event, err := getEvent(w, req)

	if err != nil{
		log.Fatal(err)
	}

	log.Println(event.Type)

	if event.Type == "customer.subscription.created" {
		c, err := customer.Get(event.Data.Object["customer"].(string), nil)
		if err != nil {
			log.Fatal(err)
		}
		email := c.Metadata["FinalEmail"]
		log.Println("Subscription created by", email)
	}


}