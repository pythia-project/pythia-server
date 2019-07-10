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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pythia-project/pythia-server/handler"
	"github.com/pythia-project/pythia-server/server"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

func loadConfig() {
	rawcfg, err := ioutil.ReadFile("config.toml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Warning: unable to read configuration file:", err)
		return
	}
	if _, err := toml.Decode(string(rawcfg), &server.Conf); err != nil {
		fmt.Println(os.Stderr, "Error: malformed configuration file:", err)
		return
	}
}

func loadEnvironments() {
	files, err := ioutil.ReadDir(server.Conf.Path.Environments)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, f := range files {
		server.Environments = append(server.Environments, server.Environment{Name: f.Name()})
	}
}

func main() {
	loadConfig()
	loadEnvironments()
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

	r.HandleFunc("/api/tasks", handler.ListTasks).Methods("GET")
	r.HandleFunc("/api/tasks", handler.CreateTask).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         server.Conf.Address.Server.String(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening to", server.Addr)
	log.Fatal(server.ListenAndServe())
}
