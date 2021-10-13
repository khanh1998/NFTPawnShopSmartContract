##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /NFTPawningShopBackend

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /NFTPawningShopBackend /NFTPawningShopBackend

EXPOSE 4000

USER nonroot:nonroot

ENTRYPOINT ["/NFTPawningShopBackend"]