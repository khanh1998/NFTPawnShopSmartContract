##
## Build
##
FROM golang:1.17-alpine AS build

RUN apk add gcc g++ musl-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /listener

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /listener /listener

USER nonroot:nonroot

ENTRYPOINT ["/listener"]