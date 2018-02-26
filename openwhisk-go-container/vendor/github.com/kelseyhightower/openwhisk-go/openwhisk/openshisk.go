// Copyright 2018 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package openwhisk

import (
	"fmt"
	"log"
	"net/http"
)

// Start starts a OpenWhisk function.
func Start(handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"OK": true}`)
	})
	http.HandleFunc("/run", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
