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

package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pythia-project/pythia-core/go/src/pythia"
	"github.com/pythia-project/pythia-server/server"
)

// HealthHandler handles route /api/health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	info := server.HealthInfo{true}

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

	request := server.SubmisssionRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Connection to the pool and execution of the task
	conn := pythia.DialRetry(pythia.QueueAddr)
	defer conn.Close()

	taskData := fmt.Sprintf(`{"environment": "python",
  		"taskfs": "%v.sfs",
  		"limits" :{
  			"time":   60,
  			"memory": 32,
  			"disk":   50,
  			"output": 1024
  		}}`, request.Tid)

	var task pythia.Task
	err = json.Unmarshal([]byte(taskData), &task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	conn.Send(pythia.Message{
		Message: pythia.LaunchMsg,
		Id:      "test",
		Task:    &task,
		Input:   request.Input,
	})
	msg, ok := <-conn.Receive()

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	result := server.SubmisssionResult{request.Tid, string(msg.Status), msg.Output}

	data, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// EnvironementsHandler handles route /api/environements
func EnvironementsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(server.Environments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
