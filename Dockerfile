FROM golang:alpine
WORKDIR /go/src/app
COPY . .
EXPOSE 80
RUN apk add make
RUN apk add git
CMD ["make", "run"]