FROM golang:1.24 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./build/todoapp ./cmd/main.go

FROM alpine:3.21
WORKDIR /app
RUN mkdir ./openapi
COPY --from=build /app/build/todoapp ./
COPY .env ./
COPY openapi/swagger.html ./openapi
EXPOSE 6969
ENTRYPOINT [ "/app/todoapp" ]
