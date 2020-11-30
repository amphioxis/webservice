# webservice

To start the webservice there are two different ways, inside a container or blank on the maschine.

## start inside container

To start the webservice inside a container make sure you have docker and docker-compose installed and just type as root:

```docker-compose up -d```

to start the container in detached mode. This will build the image and run the container after success.

Or

```docker-compose up -d ; docker-compose logs -f```

if you want to see the output.

### change the commandline flags:

Change the commandline flags in the file /build/env_file.

Per default, the webservice is waiting for requests on port 8080 and shuts down after 5 requests.
To change the port change the .env file (this will change the port of the container).

Inside the container the git hash and projectname are automatically taken.

To shut down the container use

```docker-compose down -v```

which also deletes automatically created volumes.

## start blank on the maschine

Run the makefile:

```make```

