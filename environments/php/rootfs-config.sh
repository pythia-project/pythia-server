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
#
# PHP 7.0
install_busybox

install_debs php7.0-cli php7.0-common php7.0-json php7.0-opcache php7.0-readline

# Depending libraries
install_debs libc6 libedit2 libpcre3 libssl1.1 libxml2 mime-support tzdata ucf zlib1g \
	libgcc1 libbsd0 libncurses5 libtinfo5 libmagic-mgc multiarch-support debconf \
	libicu57 liblzma5 php-common coreutils libstdc++6 libselinux1
