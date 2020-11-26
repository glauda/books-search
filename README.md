# books-search

Personnal projet. The goal is to build a search engine to find content in books using a go server and a postgres database.

## Command line instructions 

In the folder containing go files (**/server**)

* run the tests of the *app.go* file
```go test -v```

* run the server
```go run app.go```

* compile the app
```
go build -o main .
./main
```

The result of the server can be seen at http://localhost:8080/ 

## Docker instructions

To create the Docker image, go to the main directory then run
```docker build -t glauda/books-search .```

### Run terminal in container
If container does not exist
```docker run -p 8080:8080 -it -d --name bs-container glauda/books-search sh```

Otherwise
```
docker start bs-container
docker exec -it bs-container sh
```

Docker monting volume example
```
docker run -p 8080:8080 -it -v "$(pwd)"/deploy:/deploy glauda/athletes_classification python3 /deploy/app.py
```
## Database instructions 

The DBMS used is postgres running on Docker.

### Use the base postgres image

Start the postgres container and open a terminal on the container
```
docker run --name postgres-docker -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
docker exec -it postgres-docker bash
```
These instructions:
* set the `POSTGRES_PASSWORD` environment variable value to `postgres`,
* name (`--name`) the Docker container to be `postgres-docker`,
* map container’s internal `5432` port to external `5432` port, so we’ll be able to enter it from outside,
* enable to run the Docker container in the background (`-d`).
* open a `bash` shell in the container

To login to the database as a `postgres` user:
```
psql -U postgres
```

### Use a custom postgres image

The files mentionned are in the folder `database`.

To create the docker image and then run it as a container
```
cd database
docker build -t my-postgres-image .
docker run -d --name my-postgres-container -p 5555:5432 my-postgres-image
```

### Data storage

If I:
* add data into the container (using `AddData.sql`), 
* stop the container (using `docker stop my-postgres-container`),
* and then re-run it (using `docker container start my-postgres-container`).

**Caution**: With `docker run` command we create a new container from an image `my-postgres-image` so all changes made in `my-postgres-container` are not saved in new one.

To check what *Volume* is assigned to the container : `docker container inspect my-postgres-container` 