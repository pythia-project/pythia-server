# Copyright 2018 The Pythia Authors.
# This file is part of Pythia.
#
# Pythia is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, version 3 of the License.
#
# Pythia is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with Pythia.  If not, see <http://www.gnu.org/licenses/>.

# Ruby 2.3
install_debs ruby2.3 libruby2.3 rake ruby-did-you-mean ruby-minitest ruby-net-telnet \
	ruby-power-assert ruby-test-unit rubygems-integration

# Base libraries
install_debs libc6 libgmp10 rubygems-integration libgcc1 libffi6 libgdbm3 libncurses5 libreadline7 \
       	libssl1.0.2 libtinfo5 libyaml-0-2  zlib1g ca-certificates dpkg readline-common debconf 
