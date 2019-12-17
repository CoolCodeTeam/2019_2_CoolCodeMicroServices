FROM golang:1.13 AS build

COPY . /coolslack
WORKDIR /coolslack

RUN go build .

FROM alpine:latest

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=build /coolslack/ /coolslack/
WORKDIR /coolslack/

CMD [ "./messages" ]