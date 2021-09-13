FROM golang:1.17-alpine

RUN apk add gcc g++ musl-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build

EXPOSE 135

CMD [ "./khanh" ]