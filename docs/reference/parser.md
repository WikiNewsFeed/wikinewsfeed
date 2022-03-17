# Parser (Go) [![Go Reference](https://pkg.go.dev/badge/github.com/wikinewsfeed/wikinewsfeed/parser.svg)](https://pkg.go.dev/github.com/wikinewsfeed/wikinewsfeed/parser)

![Parsing progress](https://i.imgur.com/tlF6xbL.png)

Parse Wikipedia pages to extract events

## Install

Requires [Go](https://go.dev/dl/) v1.16 and greater

```sh:no-line-numbers
go get -u github.com/wikinewsfeed/wikinewsfeed/parser
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
    events, err := parser.Parse(wiki.Body, parser.ParserOptions{
		IncludeOriginal: options.IncludeOriginal,
	})
    if err != nil {
        panic(err)
    }

    fmt.Println(events[0])
}
```

Will return

```:no-line-numbers
{Three <a href="https://en.wikipedia.org/wiki/Palestine" class="mw-disambig" title="Palestine">Palestinians</a> are killed and nine more are injured by <a href="https://en.wikipedia.org/wiki/Israel_Defense_Forces" title="Israel Defense Forces">Israeli soldiers</a> during <a href="https://en.wikipedia.org/wiki/Raid" class="mw-disambig" title="Raid">raids</a> in the <a href="https://en.wikipedia.org/wiki/West_Bank" title="West Bank">West Bank</a>.   Three Palestinians are killed and nine more are injured by Israeli soldiers during raids in the West Bank.    [{Israeli–Palestinian conflict /wiki/Israeli%E2%80%93Palestinian_conflict https://en.wikipedia.org/wiki/Israeli%E2%80%93Palestinian_conflict}] {Israeli–Palestinian conflict /wiki/Israeli%E2%80%93Palestinian_conflict https://en.wikipedia.org/wiki/Israeli%E2%80%93Palestinian_conflict} [{Al Jazeera https://www.aljazeera.com/news/2022/3/15/several-palestinians-including-teen-killed-by-israeli-forces www.aljazeera.com}] {Al Jazeera https://www.aljazeera.com/news/2022/3/15/several-palestinians-including-teen-killed-by-israeli-forces www.aljazeera.com} [{Palestinians /wiki/Palestine https://en.wikipedia.org/wiki/Palestine} {Israeli soldiers /wiki/Israel_Defense_Forces https://en.wikipedia.org/wiki/Israel_Defense_Forces} {raids /wiki/Raid https://en.wikipedia.org/wiki/Raid} {West Bank /wiki/West_Bank https://en.wikipedia.org/wiki/West_Bank}] 2022-03-15 00:00:00 +0000 UTC  39a3957c19983045b75204aec9ff1781f79e1267}
```
