FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . ./

RUN apk add --no-cache libgit2 libgit2-dev git gcc g++ pkgconfig

RUN CGO_ENABLED=1 GOOS=linux go build -o server

FROM alpine:latest AS production

ENV GIN_MODE=release
ENV DB_DRIVER="sqlite3"
ENV DB_DATA_SOURCE="./quotes.db"

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/database/migrations ./database/migrations
RUN touch quotes.db
RUN chown -R nobody:nobody /app

EXPOSE 8080

USER nobody:nobody

ENTRYPOINT ["/app/server"]
