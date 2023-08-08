FROM golang:1.19-alpine3.15 AS build-env
WORKDIR /src

COPY . .
RUN go mod download

RUN go build -o /app ./main.go

FROM alpine:3.15
COPY --from=build-env app /
ENTRYPOINT ["/app"]