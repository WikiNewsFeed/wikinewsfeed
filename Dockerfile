FROM golang:1.17-alpine AS build-env
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /wikinewsfeed

FROM scratch
WORKDIR /app
COPY --from=build-env /wikinewsfeed /wikinewsfeed
ENTRYPOINT ./wikinewsfeed
