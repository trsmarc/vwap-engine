FROM golang:1.17-alpine
ENV TIMEZONE Asia/Bangkok

WORKDIR /vwap-engine
COPY . .

RUN apk add build-base
RUN go mod download
RUN go mod vendor