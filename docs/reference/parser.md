# Parser (Go) [![Go Reference](https://pkg.go.dev/badge/github.com/wikinewsfeed/wikinewsfeed/parser.svg)](https://pkg.go.dev/github.com/wikinewsfeed/wikinewsfeed/parser)

## Install

```sh:no-line-numbers
go get -u github.com/wikinewsfeed/parser
```

## Example

```go
package main

import (
    "net/http"
    "fmt"

    "github.com/wikinewsfeed/wikinewsfeed/parser"
)

func main() {
    wiki, _ := http.Get("https://en.wikipedia.org/wiki/Portal:Current_events")
    events, err := parser.Parse(wiki.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println(events[0])
}
```

Will return
