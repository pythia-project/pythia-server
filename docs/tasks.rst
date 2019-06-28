Managing tasks
==============

Several routes available on the API can be used to manage Pythia `tasks`. For example, it is possible to retrieve a list with all the available tasks, to get detailed information about a task or even to create a new one.



Task
----

The core service provided by Pythia is to execute tasks. Roughly speaking, it amounts at executing code in a given execution environment. In particular, a task contains a sequence of commands to be executed, each of which corresponding to code that can be executed in the execution environment. The following metadata are associated to every task:

* A `unique identifier` for the Pythia backbone to uniquely identify the task.
* A `friendly name` for the task that can be shown to the end user.
* The `list of authors` who contributed to the creation of the task.
* A detailed `description` of the task.
* The `environment` in which the task is executed.
* The `constraints` of the virtual machine in which the task is executed: time, memory, disk and output size.

**Example of a task**

.. sourcecode:: json

   {
     "taskid": "hello-world",
     "name": "Hello World (Shell)",
     "authors": ["John Doe"],
     "description": "This task prints 'Hello World!' to the standard output.",
     "environment": "busybox",
     "limits": {
       "time": 60,
       "memory": 32,
       "disk": 50,
       "output": 1024
     }
   }



Details
-------

The following table summarises the operations available on tasks.

.. list-table::
   :widths: 20 40 40
   :header-rows: 1

   * - Resource
     - Operation
     - Description
   * - Task
     - GET /tasks/(`string: taskid`)
     - Get task.
   * - 
     - POST /tasks
     - Create task.
   * - Tasks collection
     - GET /tasks
     - Get collection of tasks.



Get task
********

The following route can be used to get a list of all the tasks that are available on the Pythia backbone to be executed. It returns an array with each existing task described by its unique identifier, its name and list of authors. If no tasks are available, it returns an empty array.

.. http:get:: /api/tasks

   The tasks available on the Pythia backbone.

   **Example request**

   .. sourcecode:: http

      GET /api/tasks HTTP/1.1
      Host: example.com
      Accept: application/json

   **Example response**

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      [
        {
          "taskid": "hello-world",
          "name": "Hello World (Shell)",
          "authors": ["John Doe"],
        },
        {
          "taskid": "execute-python",
          "name": "Python executor (Python 3.5.3)",
          "authors": ["John Doe", "Jane Doe"],
        }
      ]

   :reqheader Accept: application/json

   :resheader Content-Type: application/json

   :resjsonarr string taskid: the unique identifier of the task
   :resjsonarr string name: the friendly name of the task
   :resjsonarr string[] authors: the list of authors of the task

   :statuscode 200: no error



Create task
***********

The following route can be used to create a new task on the Pythia backbone. It provides a way to create a new task in a “raw way”, directly providing the code to be executed in the task. It also makes it possible to create a task following one of the existing task templates, given a configuration.

.. http:post:: /api/tasks

   Creates a new task.

   **Example request**

   .. sourcecode:: http

      POST /api/tasks HTTP/1.1
      Host: example.com
      Content-Type: application/json

      {
        "taskid": "hello-go",
        "environment": "go1.12",
        "type": "raw",
        "config": {
          "taskfs": "..."
        }
      }

   **Example response**

   .. sourcecode:: http

      HTTP/1.1 200 OK

   :reqheader Content-Type: application/json

   :reqjson string taskid: the identifier of the task
   :reqjson string environment: the environment in which to execute the task
   :reqjson string type: the type of the task
   :reqjson object limits: the constraints of the virtual machine in which to execute the task: time, memory, disk and output size (optional, default: ``{"time": 60, "memory": 32, "disk": 50, "output": 1024}``)
   :reqjson string config: the configuration of the task (optional)

   :statuscode 200: no error



Get all tasks
*************

The following route can be used to get the detailed information about a single task given its unique identifier. All the fields described in the beginning of this page are returned if it exists on the Pythia backbone, and an error is raised otherwise.

.. http:get:: /api/tasks/(string:taskid)

   The task with the unique identifier (`taskid`).

   **Example request**

   .. sourcecode:: http

      GET /api/tasks/hello-world HTTP/1.1
      Host: example.com
      Accept: application/json

   **Example response**

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      {
        "taskid": "hello-world",
        "name": "Hello World (Shell)",
        "authors": ["John Doe"],
        "description": "This task prints 'Hello World!' to the standard output.",
        "environment": "busybox",
        "limits": {
          "time": 60,
          "memory": 32,
          "disk": 50,
          "output": 1024
        }
      }

   :reqheader Accept: application/json

   :resheader Content-Type: application/json

   :resjson string taskid: the unique identifier of the task (`taskid`)
   :resjson string name: the friendly name of the task
   :resjson string[] authors: the list of authors of the task
   :resjson string description: the description of the task
   :resjson string environment: the environment in which to execute the task
   :resjson object limits: the constraints of the virtual machine in which the task is executed: time, memory, disk and output size

   :statuscode 200: no error
   :statuscode 404: no task with the specified (`taskid`) has been found on the Pythia backbone
