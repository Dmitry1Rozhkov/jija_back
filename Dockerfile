FROM golang:1.21 AS build-stage

WORKDIR /app
COPY . .

RUN go mod download

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest AS build-release-stage
WORKDIR /

RUN mkdir -p /app/logs
COPY --from=build-stage /app/cmd/app /app/internal/config/config.yaml .
COPY --from=build-stage /app/atms.json /app/offices.json .

ENV CONFIG_PATH=/config.yaml
ENV GIN_MODE=release

EXPOSE 8181


CMD ["./app"]