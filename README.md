# books-search

## Command line instructions 

In the folder containing go files (**/server**)

* run the tests of the *app.go* file
`go test -v`

* run the server
`go run app.go`

* compile the app
```
go build -o main .
./main
```

The result of the server can be seen at http://localhost:9000/ 

## Docker instructions

To create the Docker image, go to the main directory then run
`docker build -t glauda/books-search .`

### Run terminal in container
If container does not exist
`docker run -p 8080:8080 -it -d --name bs-container glauda/books-search sh`

Otherwise
```
docker start bs-container
docker exec -it bs-container sh
```