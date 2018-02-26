// Copyright 2018 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func main() {
	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"OK": true}`)
	})

	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		var v Value
		if err := json.Unmarshal(data, &v); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		message := fmt.Sprintf("Hello %s", v.Value.Name)

		if err := json.NewEncoder(w).Encode(Response{Message: message}); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
