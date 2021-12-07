FROM --platform=linux/arm64 golang:bullseye AS build
RUN apt -y install git
RUN git clone https://github.com/rickytrevor/GoFnaf
RUN apt autoremove
CMD go run /go/GoFnaf/clientside/client.go
#CMD bash
