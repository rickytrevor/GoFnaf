FROM --platform=linux/arm64 golang:latest 
RUN apt -y update
RUN apt -y install git
RUN git clone https://github.com/rickytrevor/GoFnaf
CMD go run /go/GoFnaf/clientside/client.go
