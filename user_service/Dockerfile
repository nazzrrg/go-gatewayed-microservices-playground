FROM golang:1.20 as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build
RUN CGO_ENABLED=1 GOOS=linux go build #github.com/nazzrrg/go-gatewayed-microservice-playgroud/user_service
#RUN CGO_ENABLED=0 GOOS=linux go build #github.com/nazzrrg/go-gatewayed-microservice-playgroud/user_service
#
#FROM scratch as app
#COPY --from=build /app/user_service ./user_service