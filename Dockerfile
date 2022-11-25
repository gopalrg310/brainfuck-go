# syntax=docker/dockerfile:1
FROM golang:1.19.3-alpine3.16 as builder
ENV GOOS=linux
RUN apk add --update git openssh-client && rm -rf /var/cache/apk/* && \
    mkdir /root/.ssh && echo "StrictHostKeyChecking no" > /root/.ssh/config && \
    echo "${SSH_KEY}" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    git clone https://github.com/gopalrg310/brainfuck-go.git /brainfuck-go
WORKDIR /brainfuck-go
COPY . .
RUN go mod download && \
    go build -o brainfuck-go BF.go
