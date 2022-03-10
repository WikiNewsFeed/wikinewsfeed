FROM golang:1.17-alpine AS build-env
COPY go.mod ./
COPY go.sum ./
COPY *.go ./

RUN go mod download
RUN go build -o /wikinewsfeed

FROM scratch
WORKDIR /app
COPY --from=build-env /wikinewsfeed /app/
ENTRYPOINT ./wikinewsfeed
