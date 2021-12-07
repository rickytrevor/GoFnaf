FROM golang:latest
FROM linux/arm64

RUN apt update
RUN apt install git
RUN git clone https://github.com/rickytrevor/GoFnaf
CMD go run /go/GoFnaf/clientside/client.go
