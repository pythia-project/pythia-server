#!/bin/sh
#
# Pythia task template for unit testing-based tasks
# Author: Sébastien Combéfis <sebastien@combefis.be>
#
# Copyright (C) 2019, Computer Science and IT in Education ASBL
#
# This program is free software: you can redistribute it and/or modify
# under the terms of the GNU General Public License as published by
# the Free Software Foundation, version 2 of the License, or
#  (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
# General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

# Setup environment variables.
PATH=/usr/lib/jvm/java-8-openjdk-i386/bin:$PATH
CLASSPATH=/task/static/lib/json-20180813.jar:/task/static/lib/commons-csv-1.7.jar:/task/static/lib/pythia-1.0.jar:/tmp/work/$1:$CLASSPATH

# Move to working directory.
if [ "$1" = "student" ] || [ "$1" =  "teacher" ]
then
	cd /tmp/work/$1
fi

# Try to compile student/teacher code.
if [ "$1" = "student" ]
then
	javac -cp $CLASSPATH Program.java 2> /tmp/work/output/out.err
	if [ -s /tmp/work/output/out.err ]
	then
		exit 1
	else
		rm /tmp/work/output/out.err
		java -cp $CLASSPATH org.pythia.Execute $1 2> /tmp/work/output/out.err
		if [ -s /tmp/work/output/out.err ]
		then
			exit 1
		else
			rm /tmp/work/output/out.err
		fi
	fi
elif [ "$1" =  "teacher" ]
then
	javac -cp $CLASSPATH Program.java
	if [ $? -eq 0 ]
	then
		java -cp $CLASSPATH org.pythia.Execute $1
	else
		exit 1
	fi
fi
