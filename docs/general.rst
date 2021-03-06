General routes
==============

Some `general routes` available on the API can be used to get a “raw” low-level access to the Pythia platform. For example, it is possible to get health information about the Pythia backbone or to directly execute a task.



Health information
------------------

One route is available to get information about the health of the Pythia backbone and this API server itself. Basic information about the running status of different services can be obtained.

.. http:get:: /api/health

   Health information about the Pythia backbone and this API server.

   **Example request**

   .. sourcecode:: http

      GET /api/health HTTP/1.1
      Host: example.com
      Accept: application/json

   **Example response**

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      {
        "running": true
      }

   :reqheader Accept: application/json

   :resheader Content-Type: application/json

   :resjson boolean running: Pythia backbone running status

   :statuscode 200: no error


Task execution
--------------

One route is available to ask to the Pythia backbone for the execution of a single task. The input and the output of the tasks are raw strings whose precise specification and format are defined for each task. The result of the execution, containing the Pythia status and the output, is either directly sent in the body of the response of this request (*sync mode*) or is sent later on to a callback URL that has been specified (*async mode*). For the latter case, the callback URL must point to a publicly available POST route.

.. http:post:: /api/execute

   Execute a task on the Pythia backbone.

   **Example request**

   .. sourcecode:: http

      POST /api/execute HTTP/1.1
      Host: example.com
      Accept: application/json
      Content-Type: application/json

      {
        "tid": "hello-world"
      }

   **Example response**

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      {
        "tid": "hello-world",
        "status": "success",
        "output": "Hello World!"
      }

   :reqheader Accept: application/json
   :reqheader Content-Type: application/json

   :query boolean async: whether results are directly sent in the response of the request or later to a callback URL (optional, default: false)

   :reqjson string tid: the identifier of the task to execute
   :reqjson string input: the input for the task (optional)
   :reqjson string callback: the callback URL to use for `async` mode (optional)

   :resheader Content-Type: application/json

   :resjson string tid: the identifier of the executed task
   :resjson string status: the Pythia status of the execution
   :resjson string output: the output created by the task execution

   :statuscode 200: task submitted for execution
   :statuscode 400: bad request (missing parameters, wrong parameter type, task not existing)
