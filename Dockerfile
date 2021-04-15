FROM golang:1.15.2-buster

COPY . /myip 

WORKDIR /myip

RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' -o myip 

FROM alpine:3.12

COPY --from=0 /myip/myip /bin/

RUN apk --no-cache add ca-certificates

EXPOSE 5578

CMD ["/bin/myip", "--bind", ":5578"]