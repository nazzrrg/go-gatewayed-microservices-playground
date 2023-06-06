FROM golang:1.20 as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build
RUN CGO_ENABLED=0 GOOS=linux go build github.com/nazzrrg/go-gatewayed-microservice-playgroud/gateway

FROM scratch as app
COPY --from=build /app/gateway ./gateway