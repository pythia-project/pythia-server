// Copyright 2019 The Pythia Authors.
// This file is part of Pythia.
//
// Pythia is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// Pythia is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Pythia.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// HealthInfo are the informations about the health of the Pythia backend
type HealthInfo struct {
	Running bool `json:"running"`
}

// SubmisssionRequest are the informations about a submission request
type SubmisssionRequest struct {
	Tid   string `json:"tid"`
	Input string `json:"input"`
}

// SubmisssionResult are the informations about a result of a submission
type SubmisssionResult struct {
	Tid    string `json:"tid"`
	Status string `json:"status"`
	Output string `json:"output"`
}

// HealthHandler handles route /api/health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	info := HealthInfo{true}

	data, err := json.Marshal(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// ExecuteHandler handles route /api/execute
func ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := SubmisssionRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: send to pythia core

	w.Header().Set("Content-Type", "application/json")
	result := SubmisssionResult{request.Tid, "success", "Hello Pythia"}

	data, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/health", HealthHandler)
	r.HandleFunc("/api/execute", ExecuteHandler)
	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
