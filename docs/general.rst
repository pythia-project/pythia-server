General routes
==============

Some `general routes` available on the API can be used to get a “raw” low-level access to the Pythia platform. For example, it is possible to get health information about the Pythia backbone or to directly execute a task.



Health information
------------------

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

.. http:post:: /api/execute

   Execute a task on the Pythia backbone.

   **Example request**

   .. sourcecode:: http

      POST /api/execute HTTP/1.1
      Host: example.com
      Accept: application/json

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

   :query boolean async: whether results are directly sent in the response of the request or later to a callback URL

   :reqjson string tid: the identifier of the task to execute
   :reqjson string input: the input for the task (optional, default: empty string)
   :reqjson string callback: the callback URL to use for `async` mode (optional, default: false)

   :resheader Content-Type: application/json

   :resjson string tid: the identifier of the executed task
   :resjson string status: the Pythia status of the execution
   :resjson string output: the output created by the task execution

   :statuscode 200: task submitted for execution
   :statuscode 400: bad request
