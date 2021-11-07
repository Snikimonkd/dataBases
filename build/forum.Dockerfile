FROM golang:latest
RUN mkdir app
ADD . ./app
WORKDIR ./app
RUN go mod tidy

ENTRYPOINT go run main.go

EXPOSE 5000