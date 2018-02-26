// Copyright 2018 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/openwhisk-go/events"
	"github.com/kelseyhightower/openwhisk-go/openwhisk"
)

// Value holds a value.
type Value struct {
	Value Request `json:"value"`
}

// Request holds a request.
type Request struct {
	Name string `json:"name"`
}

// Response holds a reponse.
type Response struct {
	Message string `json:"message"`
}

// Handler processes OpenWhisk events.
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Processing OpenWhisk request")

	var v Value
	if err := events.FromHTTPRequest(r, &v); err != nil {
		events.WriteHTTPErrorResponse(w, err.Error(), 500)
		return
	}

	message := fmt.Sprintf("Hello %s", v.Value.Name)
	events.WriteHTTPResponse(w, Response{Message: message}, 200)
}

func main() {
	openwhisk.Start(Handler)
}
