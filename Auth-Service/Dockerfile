# use Go image
FROM golang:1.25-alpine

# install curl for healthcheck
RUN apk add --no-cache curl

# set working directory inside the container
WORKDIR /app

# copy go mod and sum files for caching
COPY go.mod go.sum ./
RUN go mod download

# copy the source code into the container
COPY . .

# build the Go application
RUN go build -o main .

# expose http and grpc to the outside world
EXPOSE 8080 5000

# command to run the executable
CMD ["./main"]