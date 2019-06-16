Managing environments
=====================

Several routes available on the API can be used to manage Pythia `environments`. For example, it is possible to retrieve a list with all the available environments, to get detailed information about an environment or even to create a new one.



Environment
-----------

With Pythia, a task is always executed within an execution environment. In particular, an environment contains all the executables and other tools that are necessary to execute and grade the task properly. The following metadata are associated to every environment:

* A `unique identifier` for the Pythia backbone to uniquely identify the environment.
* A `friendly name` for the environment that can be shown to the end user.
* The `list of authors` who contributed to the creation of the environment.
* A detailed `description` of the environment.

**Example of an environment**

.. sourcecode:: json

   {
     "envid": "python3",
     "name": "Python 3.7.3",
     "authors": ["John Doe", "Jane Doe"],
     "description": "This environment contains the standard installation of the Python 3.7.3 interpreter without any additional packages."
   }



Summary
-------

The following table summarises the operations available on environments.

.. list-table::
   :widths: 20 40 40
   :header-rows: 1

   * - Resource
     - Operation
     - Description
   * - Environment
     - GET /environments/(string:envid)
     - Get environment.
   * - Environments collection
     - GET /environments
     - Get collection of environments.



Details
-------

The following route can be used to get a list of all the environments that are available on the Pythia backbone to execute tasks. It returns an array with each existing environment described by its unique identifier, its name and list of authors. If no environments are available, it returns an empty array.

.. http:get:: /api/environments

   The environments available on the Pythia backbone.

   **Example request**

   .. sourcecode:: http

      GET /api/environments HTTP/1.1
      Host: example.com
      Accept: application/json

   **Example response**

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      [
        {
          "envid": "python3",
          "name": "Python 3.7.3",
          "authors": ["John Doe", "Jane Doe"]
        },
        {
          "envid": "c+make",
          "name": "C11 + Make",
          "authors": ["Smith"]
        }
      ]

   :reqheader Accept: application/json

   :resheader Content-Type: application/json

   :resjsonarr string envid: the unique identifier of the environment
   :resjsonarr string name: the friendly name of the environment
   :resjsonarr string[] authors: the list of authors of the environment

   :statuscode 200: no error



The following route can be used to get the detailed information about a single environment given its unique identifier. All the fields described in the beginning of this page are returned if it exists on the Pythia backbone, and an error is raised otherwise.

.. http:get:: /api/environments/(string:envid)

   The environment with the unique identifier (`envid`).

   **Example request**

   .. sourcecode:: http

      GET /api/environments/c+make HTTP/1.1
      Host: example.com
      Accept: application/json

   **Example response**

   .. sourcecode:: http

      HTTP/1.1 200 OK
      Content-Type: application/json

      {
        "envid": "c+make",
        "name": "C + Make",
        "authors": ["Smith"],
        "description": "This environment contains a C compiler (compatible with C11 specification) and the Make 4.2 tool, that both support the C11 specifications."
      }

   :reqheader Accept: application/json

   :resheader Content-Type: application/json

   :resjson string envid: the unique identifier of the environment (`envid`)
   :resjson string name: the friendly name of the environment
   :resjson string[] authors: the list of authors of the environment
   :resjson string description: the description of the environment

   :statuscode 200: no error
   :statuscode 404: no environment with the specified (`envid`) has been found on the Pythia backbone
