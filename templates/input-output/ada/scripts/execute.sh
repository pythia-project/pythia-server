#!/bin/sh

# Move to working directory.
cd /tmp/work/student

# Compile the code if not already compiled.
if [ ! -f program ]
then
	gnatmake /tmp/work/student/program.adb -D /tmp/work -o /tmp/work/student/program 2> /tmp/work/output/out.err
	if [ -s /tmp/work/output/out.err ]
	then
		cat /tmp/work/output/out.err
		exit 1
	else
		rm /tmp/work/output/out.err
	fi
fi

# Execute the code.
/tmp/work/program
