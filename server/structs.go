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

package server

// Environment is the description of an execution environment for the Pythia backbone.
type Environment struct {
	Envid       string   `json:"envid"`
	Name        string   `json:"name"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
}

// Task is the description of a task for the Pythia backbone.
type Task struct {
	Taskid      string   `json:"taskid"`
	Name        string   `json:"name,omitempty"`
	Authors     []string `json:"authors,omitempty"`
	Description string   `json:"description,omitempty"`
}

// HealthInfo are the informations about the health of the Pythia backend
type HealthInfo struct {
	Running bool `json:"running"`
}

// SubmisssionRequest are the informations about a submission request
type SubmisssionRequest struct {
	Tid      string `json:"tid"`
	Input    string `json:"input"`
	Async    bool   `json:"async"`
	Callback string `json:"callback"`
}

// SubmisssionResult are the informations about a result of a submission
type SubmisssionResult struct {
	Tid    string `json:"tid"`
	Status string `json:"status"`
	Output string `json:"output"`
}

// TaskCreationRequest is the description of a task creation request.
type TaskCreationRequest struct {
	Taskid      string      `json:"taskid"`
	Environment string      `json:"environment"`
	Type        string      `json:"type"`
	Limits      Limits      `json:"limits"`
	Config      interface{} `json:"config,omitempty"`
}

type Limits struct {
	Time   int `json:"time"`
	Memory int `json:"memory"`
	Disk   int `json:"disk"`
	Output int `json:"output"`
}

type UnitTestingTaskConfig struct {
	Spec struct {
		Name string `json:"name"`
		Args []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"args"`
	} `json:"spec"`
	Test struct {
		Predefined []struct {
			Data     string            `json:"data"`
			Feedback map[string]string `json:"feedback,omitempty"`
		} `json:"predefined,omitempty"`
		Random struct {
			N    int      `json:"n"`
			Args []string `json:"args"`
		} `json:"random"`
	} `json:"test"`
	Solution string `json:"solution"`
}

type InputOutputTaskConfig struct {
	Predefined []struct {
		Input   string `json:"input"`
		Output  string `json:"output"`
		Message string `json:"message,omitempty"`
	} `json:"predefined"`
}
