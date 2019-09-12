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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/pythia-project/pythia-core/go/src/pythia"
	"github.com/pythia-project/pythia-server/server"
)

// HealthHandler handles route /api/health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	conn, err := pythia.Dial(server.Conf.Address.Queue)
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
	conn := pythia.DialRetry(server.Conf.Address.Queue)

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

// ListEnvironments lists all the available environments.
func ListEnvironments(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(server.Conf.Path.Environments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	environments := make([]server.Environment, 0)
	for _, f := range files {
		name := f.Name()
		if strings.HasSuffix(name, ".sfs") {
			environments = append(environments, server.Environment{Name: name[:len(name)-4]})
		}
	}

	data, err := json.Marshal(environments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetEnvironment retrieves one given environment.
func GetEnvironment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	envpath := fmt.Sprintf("%s/%s.env", server.Conf.Path.Environments, vars["envid"])
	if _, err := os.Stat(envpath); err == nil {
		if content, err := ioutil.ReadFile(envpath); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(content)
			return
		}

	} else if os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

// ListTasks lists all the available tasks.
func ListTasks(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(server.Conf.Path.Tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tasks := make([]server.Task, 0)
	for _, f := range files {
		name := f.Name()
		if strings.HasSuffix(name, ".task") {
			tasks = append(tasks, server.Task{Taskid: name[:len(name)-5]})
		}
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetTask retrieves one given task.
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskpath := fmt.Sprintf("%s/%s.task", server.Conf.Path.Tasks, vars["taskid"])
	if _, err := os.Stat(taskpath); err == nil {
		if content, err := ioutil.ReadFile(taskpath); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(content)
			return
		}

	} else if os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

// DeleteTask deletes one given task.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskdir := fmt.Sprintf("%s/%s", server.Conf.Path.Tasks, vars["taskid"])
	if _, err := os.Stat(taskdir); err == nil {
		_ = os.RemoveAll(taskdir)
		_ = os.Remove(taskdir + ".sfs")
		_ = os.Remove(taskdir + ".task")

		w.WriteHeader(http.StatusOK)
		return
	} else if os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

// CreateTask creates a new task.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	request := server.TaskCreationRequest{
		Type: "raw",
		Limits: server.Limits{
			Time:   60,
			Memory: 32,
			Disk:   50,
			Output: 1024,
		},
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Check whether a task with the same ID already exists
	taskDir := fmt.Sprintf("%s/%s", server.Conf.Path.Tasks, request.Taskid)
	taskFile := fmt.Sprintf("%s.task", taskDir)
	if _, err := os.Stat(fmt.Sprintf("%s.task", taskDir)); err == nil {
		log.Println("Task id", request.Taskid, "already exists.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the task directory
	if err := os.Mkdir(taskDir, 0755); err != nil {
		log.Println("Impossible to create task directory:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create the task file
	task := pythia.Task{
		Environment: request.Environment,
		TaskFS:      request.Taskid + ".sfs",
		Limits:      request.Limits,
	}
	file, _ := json.MarshalIndent(task, "", "  ")
	_ = ioutil.WriteFile(taskFile, file, 0644)

	// Copy the files from the template
	switch request.Type {
	case "input-output":
		_ = os.Mkdir(taskDir+"/config", 0755)
		_ = os.Mkdir(taskDir+"/scripts", 0755)
		_ = os.Mkdir(taskDir+"/skeleton", 0755)

		templateDir := "templates/input-output/" + request.Environment
		_ = copyFile(templateDir+"/control", taskDir+"/control", 0755)
		_ = copyFile(templateDir+"/scripts/pythia-iot", taskDir+"/scripts/pythia-iot", 0755)

		switch request.Environment {
		case "python":
			_ = copyFile(templateDir+"/skeleton/program.py", taskDir+"/skeleton/program.py", 0755)
		case "php7":
			_ = copyFile(templateDir+"/skeleton/program.php", taskDir+"/skeleton/program.php", 0755)
		case "nodejs":
			_ = copyFile(templateDir+"/skeleton/program.js", taskDir+"/skeleton/program.js", 0755)
		case "java":
			_ = copyFile(templateDir+"/scripts/execute.sh", taskDir+"/scripts/execute.sh", 0755)
			_ = copyFile(templateDir+"/skeleton/Program.java", taskDir+"/skeleton/Program.java", 0755)
		case "bash":
			_ = copyFile(templateDir+"/skeleton/program.sh", taskDir+"/skeleton/program.sh", 0755)
		case "rexx":
			_ = copyFile(templateDir+"/skeleton/program.rexx", taskDir+"/skeleton/program.rexx", 0755)
		}

		// Save the configuration
		config := server.InputOutputTaskConfig{}
		if mapstructure.Decode(request.Config, &config) == nil {
			file, _ = json.MarshalIndent(config, "", "  ")
			_ = ioutil.WriteFile(taskDir+"/config/test.json", file, 0644)
		}

	case "unit-testing":
		_ = os.Mkdir(taskDir+"/config", 0755)
		_ = os.Mkdir(taskDir+"/scripts", 0755)
		_ = os.Mkdir(taskDir+"/skeleton", 0755)
		_ = os.Mkdir(taskDir+"/static", 0755)
		_ = os.Mkdir(taskDir+"/static/lib", 0755)
		templateDir := "templates/unit-testing/python"
		_ = copyFile(templateDir+"/control", taskDir+"/control", 0755)
		_ = copyFile(templateDir+"/scripts/preprocess.py", taskDir+"/scripts/preprocess.py", 0755)
		_ = copyFile(templateDir+"/scripts/generate.py", taskDir+"/scripts/generate.py", 0755)
		_ = copyFile(templateDir+"/scripts/execute.py", taskDir+"/scripts/execute.py", 0755)
		_ = copyFile(templateDir+"/scripts/feedback.py", taskDir+"/scripts/feedback.py", 0755)
		_ = copyFile(templateDir+"/static/lib/__init__.py", taskDir+"/static/lib/__init__.py", 0755)
		_ = copyFile(templateDir+"/static/lib/pythia.py", taskDir+"/static/lib/pythia.py", 0755)

		// Save the configuration
		config := server.UnitTestingTaskConfig{}
		if mapstructure.Decode(request.Config, &config) == nil {
			file, _ := json.MarshalIndent(config.Spec, "", "  ")
			_ = ioutil.WriteFile(taskDir+"/config/spec.json", file, 0644)
			file, _ = json.MarshalIndent(config.Test, "", "  ")
			_ = ioutil.WriteFile(taskDir+"/config/test.json", file, 0644)

			// Create skeletons files
			params := make([]string, 0)
			for _, elem := range config.Spec.Args {
				params = append(params, elem.Name)
			}
			content := fmt.Sprintf("# -*- coding: utf-8 -*-\n\ndef %s(%s):\n@  @f1@@", config.Spec.Name, strings.Join(params, ", "))
			ioutil.WriteFile(taskDir+"/skeleton/program.py", []byte(content), 0755)

			// Create solution file
			ioutil.WriteFile(taskDir+"/config/solution", []byte(config.Solution), 0644)
		}
	}

	// Compile the SFS
	// mksquashfs TASK TASK.sfs -all-root -comp lzo -noappend
	wd, _ := os.Getwd()
	_ = os.Chdir(server.Conf.Path.Tasks)
	exec.Command("mksquashfs", request.Taskid, request.Taskid+".sfs", "-all-root", "-comp", "lzo", "-noappend").Run()
	_ = os.Chdir(wd)

	w.WriteHeader(http.StatusOK)
}

func copyFile(src string, dst string, perms os.FileMode) (err error) {
	var from, to *os.File
	if from, err = os.Open(src); err == nil {
		defer from.Close()
		if to, err = os.OpenFile(dst, os.O_RDWR|os.O_CREATE, perms); err == nil {
			defer to.Close()
			_, err = io.Copy(to, from)
		}
	}
	return
}
