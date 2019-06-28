.. pythia-server documentation master file, created by
   sphinx-quickstart on Sun Jun  9 17:19:22 2019.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

Pythia-server: REST API to execute code and tasks on the Pythia platform
========================================================================

Pythia is a framework deployed as an online platform whose goal is to teach programming and algorithm design. The platform executes the code in a safe environment and its main advantage is to provide intelligent feedback to its users to suppor their learning. More details about the whole project can be found on the `official website of Pythia
<https://www.pythia-project.org/>`_.

Pythia-server is one frontend for the Pythia framework. It offers a REST API to execute code and tasks on the Pythia platform. It also includes functions to manage and create tasks and environments and to get health information about the Pythia backbone. Pythia-server is written in `Go
<https://golang.org>`_.



Quick install
-------------

The pythia-server frontend can be run on Linux, Windows and macOS.

Start by installing required dependencies:

* Go (1.11 or later)

Then, clone the Git repository, and launch the installation:

.. code-block:: none

   > git clone https://github.com/pythia-project/pythia-server.git
   > cd pythia-server
   > go build

Once successfully installed, you can launch the server:

.. code-block:: none

   > ./pythia-server

You can now try to execute a simple task (assuming that you launched the `pythia-core
<https://pythia-core.readthedocs.io>`_ framework with a queue listening on port 9000) by calling the ``/api/execute`` route, for example with the `curl tool
<https://curl.haxx.se/>`_:

.. code-block:: none

   > curl -d '{"tid": "hello-world"}' http://localhost:8080/api/execute

and you will see, among others, ``Hello world!`` printed in your terminal.



Contents
--------

This documentation is split into two parts: the first one is targeter to users and the second one is for developers. In any case, we recommend you to first read the user's documentation to understand how to use and test the server.


.. toctree::
   :maxdepth: 1
   :caption: User's Documentation
   
   general
   environments
   tasks


.. toctree::
   :maxdepth: 1
   :caption: Developer's Documentation
