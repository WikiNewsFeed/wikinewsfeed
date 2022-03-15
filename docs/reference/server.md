# Server

## Installation

### Docker

```sh:no-line-numbers
docker run -p 8080:8080 ghcr.io/wikinewsfeed/wikinewsfeed:latest
```

## Configuration

Configuration is applied using Environment variables

| Variable   | Type        | Default               | Description                                  |
|------------|-------------|-----------------------|----------------------------------------------|
| PORT       | String      | 8080                  | Listen to port                               |
| WNF_URL    | String, URL | http://localhost:8080 | URL to be displayed in readers               |
| WNF_MAXAGE | String      | 1800                  | Cache-Control age (only useful behind a CDN) |
| WNF_DB     | Path        | stats.db              | Path where database file should be saved     |
| WNF_CORS   | String      | *                     | CORS Header                                  |

## Recommendations

Default response times from Wikipedia range 600-2000ms. Therefor it is **absolutely necessary** to run the server behind a caching proxy if you want a quick response

Here's example configuration for [varnish](https://varnish-cache.org), that will work with the default settings:

```hcl
vcl 4.1;

backend default {
    .host = "0.0.0.0";
    .port = "8080";
}
```

## Building

### Docker

```sh:no-line-numbers
docker build -t wikinewsfeed/wikinewsfeed .
```

### From source

- [Go](https://go.dev/dl/) v1.16 and greater
- [NodeJS](https://nodejs.org/en/) optionally if you want to build the docs

#### 1. Clone the repository

```sh:no-line-numbers
git clone https://github.com/WikiNewsFeed/wikinewsfeed.git
```

#### 2. Build the binary

Install the dependencies

```sh:no-line-numbers
go mod download
```

Build

```sh:no-line-numbers
go build
```

#### 3. Build the docs

Install the dependencies

```sh:no-line-numbers
npm i
```

(Optionally) Preview the changes

 ```sh:no-line-numbers
 npm run docs:dev
 ```

Build

```sh:no-line-numbers
npm run docs:build
```
