FROM golang:1.8

RUN mkdir /app
ADD ./SPA_Lab_Final /app/
WORKDIR /app
RUN go get -u github.com/gorilla/mux
RUN go build -o main .
CMD ["/app/main"]
