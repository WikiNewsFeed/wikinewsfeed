# Build go binary
FROM golang:1.17-alpine AS build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o wikinewsfeed

# Build docs
FROM node:lts-alpine AS docs
WORKDIR /build
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run docs:build

FROM scratch
WORKDIR /app
COPY --from=build /build/wikinewsfeed /wikinewsfeed
COPY --from=docs /build/docs/.vuepress/dist/ /docs/.vuepress/dist/
ENTRYPOINT /wikinewsfeed
