// Copyright 2018 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package events

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ErrorResponse holds an.
type ErrorResponse struct {
	Error string
}

// FromHTTPRequest extracts.
func FromHTTPRequest(r *http.Request, v interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body.Close()
	return json.Unmarshal(data, &v)
}

// WriteHTTPResponse writes an HTTP reponse.
func WriteHTTPResponse(w http.ResponseWriter, v interface{}, code int) error {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	w.WriteHeader(code)
	return nil
}

// WriteHTTPErrorResponse writes an HTTP reponse.
func WriteHTTPErrorResponse(w http.ResponseWriter, e string, code int) {
	if err := json.NewEncoder(w).Encode(ErrorResponse{Error: e}); err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
}
