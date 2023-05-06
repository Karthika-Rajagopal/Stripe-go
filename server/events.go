package main

import (
	"github.com/stripe/stripe-go/v72/customer"
	"log"
	"net/http"
)


func HandleEvent(w http.ResponseWriter, req * http.Request)  { //handles Stripe webhook events related to subscriptions
	event, err := getEvent(w, req)  //getEvent function is called to retrieve the Stripe webhook event from the request body

	if err != nil{
		log.Fatal(err)
	}

	log.Println(event.Type)

	if event.Type == "customer.subscription.created" {  //The code logs the type of the event received. If the event type is "customer.subscription.created", the code retrieves the customer information associated with the subscription using the customer ID from the event data. It then logs the email associated with the customer's metadata.
		c, err := customer.Get(event.Data.Object["customer"].(string), nil)
		if err != nil {
			log.Fatal(err)
		}
		email := c.Metadata["FinalEmail"]
		log.Println("Subscription created by", email)
	}


}
