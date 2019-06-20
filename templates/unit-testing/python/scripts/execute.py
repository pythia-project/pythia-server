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
import sys

sys.path.append('/task/static')
from lib import pythia

sys.path.append('/tmp/work')
import program

class TaskTestSuite(pythia.TestSuite):
  def __init__(self, spec):
    pythia.TestSuite.__init__(self, '/tmp/work/input/data.csv', spec)

  def studentCode(self, data):
    return getattr(program, spec['name'])(*data)

# Read function specification
with open('/task/config/spec.json', 'r', encoding='utf-8') as file:
  content = file.read()
  spec = json.loads(content)

TaskTestSuite(spec).run('/tmp/work/output', 'data.res')
