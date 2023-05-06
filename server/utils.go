package main

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72"
	"io/ioutil"
	"net/http"
)

func getEvent(w http.ResponseWriter, req * http.Request) (eventRes * stripe.Event, err error){  //getEvent(). This function takes an HTTP request as input and reads the request body up to a maximum size of 65536 bytes. It then reads the payload and unmarshals it into a stripe.Event struct using the json.Unmarshal() function
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	event := stripe.Event{}
	err = json.Unmarshal(payload, &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}



func CORSCheck(handler func (w http.ResponseWriter, req *http.Request)) func (w http.ResponseWriter, req *http.Request){  //This function takes an HTTP handler function as input and returns a new handler function with CORS headers set. The CORS headers are used to allow cross-origin requests to the API. The returned handler function checks if the HTTP method is OPTIONS
	res := func (w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		handler(w, req)
	}
	return res


}
