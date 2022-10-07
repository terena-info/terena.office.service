FROM golang:alpine

WORKDIR /usr/app/

COPY ./ /usr/app/

RUN go build .

EXPOSE 10001

CMD ["./goginhandlers"]
