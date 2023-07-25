// Copyright 2023 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

var pingFn = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("."))
	return
}

var denyFn = func(w http.ResponseWriter, r *http.Request) {
	log.Printf("DMDEBUG url=%#v, scheme=%s, path=%s", r.URL, r.URL.Scheme, r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("http server supports only the /ping entrypoint"))
	return
}

// NewPingServer creates an http server for handling only pings.
func NewPingServer(host string, port int, readTimeout, writeTimeout time.Duration) *http.Server {
	http.HandleFunc("/ping", pingFn)
	http.HandleFunc("/", denyFn)

	return &http.Server{
		Addr:         host + ":" + strconv.Itoa(port),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		// Handler:      mux,
	}
}
