FROM golang:1.15.3-alpine3.12

COPY ./server /go/src 



# Download all the dependencies
# RUN go get -d -v ./...

# Install the package
# RUN go install -v ./...

WORKDIR /go/src

# Build the Go app
RUN go build -o main .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]