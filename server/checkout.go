package main

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/customer"
	"log"
	"net/http"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

var PriceId = "price_1MuKSiSHZTnL3EyWDJ2Kmi1D" //variable that stores the ID of the subscription product on Stripe

func checkout(email string) (*stripe.CheckoutSession, error) {   //checkout function that creates a new Stripe Checkout session
	var discounts []*stripe.CheckoutSessionDiscountParams

	discounts = []*stripe.CheckoutSessionDiscountParams{ //creates a slice of stripe.CheckoutSessionDiscountParams objects with a coupon code, which is added to the Checkout session to apply a discount
		&stripe.CheckoutSessionDiscountParams{
			Coupon: stripe.String("832g70gk"),
		},
	}

	customerParams := &stripe.CustomerParams{  //creates a new Stripe customer with the given email address and adds a metadata key-value pair with the email address
		Email: stripe.String(email),
	}
	customerParams.AddMetadata("FinalEmail", email)
	newCustomer, err := customer.New(customerParams)

	if err != nil {
		return nil, err
	}

	meta := map[string]string{
		"FinalEmail" : email,
	}

	log.Println("Creating meta for user: ", meta)

	params := &stripe.CheckoutSessionParams{
		Customer: &newCustomer.ID,
		SuccessURL: stripe.String("https://githu.com/Karthika-Rajagopal"),
		CancelURL: stripe.String("https://github.com/Karthika-Rajagopal"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Discounts: discounts,
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(PriceId),
				Quantity: stripe.Int64(1),
			},
		},
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			TrialPeriodDays: stripe.Int64(7),
			Metadata: meta,
		},
	}
	return session.New(params)
}

type EmailInput struct {
	Email string `json:"email"`
}

type SessionOutput struct {
	Id string `json:"id"`
}


 func CheckoutCreator(w http.ResponseWriter, req * http.Request){
 	input := &EmailInput{}
 	err := json.NewDecoder(req.Body).Decode(input)
	if err != nil{
		log.Fatal(err)
 	}

	stripeSession, err := checkout(input.Email)
	if err != nil{
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(&SessionOutput{Id: stripeSession.ID})

	if err != nil{
		log.Fatal(err)
	}
}











