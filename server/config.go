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
	"errors"
	"net"
	"os"

	"github.com/pythia-project/pythia-core/go/src/pythia"
)

// Environments available on the Pythia backbone.
var Environments = make([]Environment, 0)

// Conf contains the configuration for this server.
var Conf Config = NewConfig()

// Config represents the configuration for this server.
type Config struct {
	Address struct {
		Server Address
		Queue  Address
	}

	Path struct {
		Environments string
		Tasks        string
	}
}

type Address struct {
	*net.TCPAddr
}

func (d *Address) UnmarshalText(text []byte) error {
	var err error
	addr, err := pythia.ParseAddr(string(text))
	if addr, ok := addr.(*net.TCPAddr); ok {
		d.TCPAddr = addr
	} else {
		err = errors.New("Invalid address: " + string(text))
	}
	return err
}

// NewConfig creates a new configuration with default values.
func NewConfig() Config {
	conf := Config{}
	conf.Address.Server.UnmarshalText([]byte("0.0.0.0:8080"))
	conf.Address.Queue.UnmarshalText([]byte("127.0.0.1:9000"))
	conf.Path.Environments = os.Getenv("PYTHIAPATH") + "/vm"
	conf.Path.Tasks = os.Getenv("PYTHIAPATH") + "/tasks"
	return conf
}
