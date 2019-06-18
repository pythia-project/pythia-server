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
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pythia-project/pythia-server/handler"
	"github.com/pythia-project/pythia-server/server"

	"github.com/gorilla/mux"
)

func loadEnvironments() {
	files, err := ioutil.ReadDir(server.EnvironmentsPath)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, f := range files {
		server.Environments = append(server.Environments, server.Environement{Name: f.Name()})
	}
}

func main() {
	loadEnvironments()
	log.Printf("available environments: %v", server.Environments)
	r := mux.NewRouter()
	r.HandleFunc("/api/health", handler.HealthHandler).
		Methods("GET")

	r.HandleFunc("/api/execute", handler.ExecuteHandler).
		Queries("async", "{async}").
		Methods("POST")

	r.HandleFunc("/api/execute", handler.ExecuteHandler).
		Methods("POST")

	r.HandleFunc("/api/environments", handler.EnvironementsHandler).
		Methods("GET")

	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
