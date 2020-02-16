FROM golang:alpine
WORKDIR /go/src/app
COPY . .
EXPOSE 5000
RUN apk add make
CMD ["make", "run"]