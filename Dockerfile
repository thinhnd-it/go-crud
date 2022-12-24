FROM golang:1.19-alpine

# ENV GO111MODULE=on ## to develop go out size the ROOT_PATH/src

# Add Maintainer info
LABEL maintainer="Ted Nguyen <thinhnd194@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
EXPOSE 3000

ENTRYPOINT /go/bin/CompileDaemon --build="go build main.go" --command=./main