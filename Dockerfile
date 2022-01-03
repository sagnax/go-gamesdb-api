FROM golang:latest

WORKDIR /usr/src/api

COPY . /usr/src/api/

RUN go get .

EXPOSE 3000

CMD ["go", "run", "."]