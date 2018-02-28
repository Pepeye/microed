#final stage
FROM alpine:latest

RUN mkdir /app
WORKDIR /app
LABEL Name=consignments Version=0.0.1
ADD service /app/consignments
EXPOSE 50051
CMD ["./consignments"]
