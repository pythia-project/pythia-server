#!/usr/bin/python3
# -*- coding: utf-8 -*-
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

import json
import os
import sys

sys.path.append('/task/static')
from lib import pythia

if not sys.argv[1] in ['student', 'teacher']:
    sys.exit(1)

sys.path.append('/tmp/work/' + sys.argv[1])

# Read the specification of the function.
with open('/task/config/spec.json', 'r', encoding='utf-8') as file:
    spec = json.loads(file.read())

# Import the code to execute.
if sys.argv[1] == 'student':
    try:
        import program
    except SyntaxError as e:
        with open('/tmp/work/output/out.err', 'w', encoding='utf-8') as file:
            (head, tail) = os.path.split(e.filename)
            file.write('invalid syntax ({}, line {})'.format(tail, e.lineno - 3))
        sys.exit(0)
else:
    import program

# Create the specific runner for the code to execute.
inputfile = '/tmp/work/input/data.csv'
if sys.argv[1] == 'student':
    class TaskTestSuite(pythia.TestSuite):
        def __init__(self, spec):
            pythia.TestSuite.__init__(self, inputfile, spec)

        def code(self, data):
            return getattr(program, spec['name'])(*data)

    runner = TaskTestSuite(spec)
    outputfile = 'data.res'
else:
    class TaskRunner(pythia.Runner):
        def __init__(self, spec):
            pythia.Runner.__init__(self, inputfile, spec)

        def code(self, data):
            return getattr(program, spec['name'])(*data)

    runner = TaskRunner(spec)
    outputfile = 'solution.res'

runner.run('/tmp/work/output', outputfile)
