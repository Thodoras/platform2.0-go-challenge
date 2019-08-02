FROM golang:1.10.4

RUN mkdir -p /go/src/platform2.0-go-challenge
ADD . /go/src/platform2.0-go-challenge
WORKDIR /go/src/platform2.0-go-challenge/src

RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/subosito/gotenv
RUN go get golang.org/x/crypto/bcrypt

RUN go build -o main .

ENTRYpOINT /go/src/platform2.0-go-challenge/src/main