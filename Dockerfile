FROM golang:1.27.4-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o /out/api ./cmd/api && go build -o /out/cron ./cmd/cron && go build -o /out/migrate ./cmd/migrate

FROM alpine:3.24.4
WORKDIR /app
COPY --from=build /out/* ./
COPY migrations ./migrations
CMD ["./api"]