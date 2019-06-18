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

import (
	"os"

	"github.com/pythia-project/pythia-core/go/src/pythia"
)

// Global configuration
var (
	// The address on which this server listens.
	ServerAddr, _ = pythia.ParseAddr("0.0.0.0:8080")

	// The address on which the queue listens.
	QueueAddr, _ = pythia.ParseAddr("127.0.0.1:9000")

	// The path where to find the environments.
	EnvironmentsPath = os.Getenv("PYTHIAPATH") + "/vm"

	// The path where to find the tasks.
	TasksPath = os.Getenv("PYTHIAPATH") + "/tasks"

	// The environments available on the Pythia backbone.
	Environments = make([]Environement, 0)
)
