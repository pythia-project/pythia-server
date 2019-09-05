#!/bin/sh

# Setup environment variables.
PATH=/usr/lib/jvm/java-8-openjdk-i386/bin:$PATH

# Move to working directory.
cd /tmp/work/student

# Compile the code if not already compiled.
if [ ! -f /tmp/work/student/Program.class ]
then
	javac Program.java 2> /tmp/work/output/out.err
	if [ -s /tmp/work/output/out.err ]
	then
		cat /tmp/work/output/out.err
		exit 1
	else
		rm /tmp/work/output/out.err
	fi
fi

# Execute the code.
java Program
