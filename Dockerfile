FROM golang:1.18-alpine as build
WORKDIR /go/src/app
COPY . .
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

EXPOSE 8022
CMD ["/go/bin/app"]