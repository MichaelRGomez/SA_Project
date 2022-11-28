FROM golang:1.19

WORKDIR /test

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/api/*.go ./

RUN go build -o /randstring

EXPOSE 4000

CMD [ "/randstring" ]