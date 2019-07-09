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

# Read function specification
with open('/task/config/solution', 'r', encoding='utf-8') as file:
  os.makedirs('/tmp/work/scripts')
  pythia.fillSkeletons('/task/skeleton', '/tmp/work/scripts', {'f1': file.read()})
  os.rename('/tmp/work/scripts/program.py', '/tmp/work/scripts/solution.py')

sys.path.append('/tmp/work/scripts')
import solution

class TaskFeedbackSuite(pythia.FeedbackSuite):
  def __init__(self, config, spec):
    pythia.FeedbackSuite.__init__(self, '/tmp/work/input/data.csv', '/tmp/work/output/data.res', config, spec)

  def teacherCode(self, data):
    return getattr(solution, spec['name'])(*data)

# Read function specification
with open('/task/config/spec.json', 'r', encoding='utf-8') as file:
  content = file.read()
  spec = json.loads(content)

# Read test configuration
config = []
with open('/task/config/test.json', 'r', encoding='utf-8') as file:
  content = file.read()
  config = json.loads(content)
  config = config['predefined'] if 'predefined' in config else []
(verdict, feedback) = TaskFeedbackSuite(config, spec).generate()

# Retrieve task id
with open('/tmp/work/tid', 'r', encoding='utf-8') as file:
  tid = file.read()
print(json.dumps({'tid': tid, 'status': 'success' if verdict else 'failed', 'feedback': feedback}))
