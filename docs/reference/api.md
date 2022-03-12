# API

## Endpoints

### Atom Feed

`GET` [https://wikinewsfeed.org/feed/atom](https://wikinewsfeed.org/feed/atom)

<!-- #### Response

#### Example

```

``` -->

### RSS Feed

`GET` [https://wikinewsfeed.org/feed/rss](https://wikinewsfeed.org/feed/rss)

### JSON Feed

`GET` [https://wikinewsfeed.org/feed/json](https://wikinewsfeed.org/feed/json)

### Events Feed

`GET` [https://wikinewsfeed.org/api/events](https://wikinewsfeed.org/api/events)

#### Example

```json
  [{
    "html": "A large convoy of \u003ca href=\"https://en.wikipedia.org/wiki/Russian_Armed_Forces\" title=\"Russian Armed Forces\"\u003eRussian military\u003c/a\u003e vehicles, including tanks and \u003ca href=\"https://en.wikipedia.org/wiki/Self-propelled_artillery\" title=\"Self-propelled artillery\"\u003eself-propelled artillery\u003c/a\u003e, begins \u0026#34;fanning out\u0026#34; to forests and towns near \u003ca href=\"https://en.wikipedia.org/wiki/Kyiv\" title=\"Kyiv\"\u003eKyiv\u003c/a\u003e as it prepares to advance on the capital. ",
    "text": "A large convoy of Russian military vehicles, including tanks and self-propelled artillery, begins \"fanning out\" to forests and towns near Kyiv as it prepares to advance on the capital. ",
    "category": "",
    "topics": [
      {
        "title": "Kyiv offensive (2022)",
        "uri": "/wiki/Kyiv_offensive_(2022)",
        "external_url": "https://en.wikipedia.org/wiki/Kyiv_offensive_(2022)"
      }
    ],
    "primary_topic": {
      "title": "Kyiv offensive (2022)",
      "uri": "/wiki/Kyiv_offensive_(2022)",
      "external_url": "https://en.wikipedia.org/wiki/Kyiv_offensive_(2022)"
    },
    "sources": [
      {
        "name": "BBC News",
        "url": "https://www.bbc.co.uk/news/world-europe-60702464",
        "domain": "www.bbc.co.uk"
      }
    ],
    "primary_source": {
      "name": "BBC News",
      "url": "https://www.bbc.co.uk/news/world-europe-60702464",
      "domain": "www.bbc.co.uk"
    },
    "references": [
      {
        "title": "Russian military",
        "uri": "/wiki/Russian_Armed_Forces",
        "external_url": "https://en.wikipedia.org/wiki/Russian_Armed_Forces"
      },
      {
        "title": "self-propelled artillery",
        "uri": "/wiki/Self-propelled_artillery",
        "external_url": "https://en.wikipedia.org/wiki/Self-propelled_artillery"
      },
      {
        "title": "Kyiv",
        "uri": "/wiki/Kyiv",
        "external_url": "https://en.wikipedia.org/wiki/Kyiv"
      }
    ],
    "date": "2022-03-11T00:00:00Z"
  }]
```

## Entities

## Limits

### Connections

20 concurrent connections are allowed simultaneously

### Cache

TTL is set to 30 minutes
