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
