FROM golang:alpine

COPY ./main.go /go

ENTRYPOINT ["go","run","main.go"]
