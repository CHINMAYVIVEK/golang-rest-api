## We specify the base image we need for our
## go application

FROM golang:1.20.0-alpine

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git gcc libc-dev  ca-certificates tzdata && update-ca-certificates


## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app

## Add this go mod download command to pull in any dependencies
RUN go mod download
RUN go mod tidy

ARG TARGETOS
ARG TARGETARCH

## we run go build to compile the binary
## executable of our Go program

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]
