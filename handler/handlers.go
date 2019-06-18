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
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pythia-project/pythia-core/go/src/pythia"
	"github.com/pythia-project/pythia-server/server"
)

// HealthHandler handles route /api/health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	conn, err := pythia.Dial(pythia.QueueAddr)
	if err == nil {
		conn.Close()
	}
	info := server.HealthInfo{err == nil}

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
	request := server.SubmisssionRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var async bool
	if r.FormValue("async") == "" {
		async = false
	} else {
		async, err = strconv.ParseBool(r.FormValue("async"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if async && request.Callback == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Connection to the pool and execution of the task
	conn := pythia.DialRetry(pythia.QueueAddr)

	var task pythia.Task

	file, err := os.Open(fmt.Sprintf("%v/%v.task", server.Conf.Path.Tasks, request.Tid))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(file).Decode(&task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	conn.Send(pythia.Message{
		Message: pythia.LaunchMsg,
		Id:      "test",
		Task:    &task,
		Input:   request.Input,
	})

	receive := func() (res []byte, err error) {
		msg, ok := <-conn.Receive()

		if !ok {
			err = errors.New("Pythia request failed")
			return
		}

		result := server.SubmisssionResult{request.Tid, string(msg.Status), msg.Output}

		res, err = json.Marshal(result)
		if err != nil {
			return
		}
		return
	}

	if async {
		go func() {
			byteData, err := receive()
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			conn.Close()
			data := strings.NewReader(string(byteData))
			postResponse, err := http.Post(request.Callback, "application/json", data)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(postResponse)
		}()
	} else {
		byteData, err := receive()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		conn.Close()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteData)
	}

}

// EnvironementsHandler handles route /api/environements
func EnvironementsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(server.Environments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
