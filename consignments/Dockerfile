#build stage
FROM golang:alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN apk add --no-cache git
# RUN go-wrapper download   # "go get -d -v ./..."
# RUN go-wrapper install    # "go install -v ./..."

#final stage
FROM alpine:latest
WORKDIR /app
LABEL Name=consignments Version=0.0.1
ADD consignments /app/consignments
EXPOSE 50051
CMD ["./consignments"]
