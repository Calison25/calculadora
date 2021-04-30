FROM golang:1.16-alpine

WORKDIR /go/src

COPY . .

EXPOSE 1500

CMD ["top"]