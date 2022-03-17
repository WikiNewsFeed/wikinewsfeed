# API

## Endpoints

### Feed

`GET` [https://wikinewsfeed.org/feed/{type}](https://wikinewsfeed.org/feed/{type})

#### Path Params

| Param | Type                  | Description |
|-------|-----------------------|-------------|
| type  | "atom", "rss", "json" | Feed type   |

#### Query Params

| Param           | Type    | Description                         |
|-----------------|---------|-------------------------------------|
| page            | String  | Wikipedia Page                      |
| subscribe       | String  | Unique subscriber id                |

#### Response

The endpoints responds with XML or JSON with correct headers for each feed type

::: details Example

<CodeGroup>
  <CodeGroupItem title="cURL" active>

```bash:no-line-numbers
curl https://wikinewsfeed.org/feed/atom
```

  </CodeGroupItem>
  <CodeGroupItem title="JavaScript">

```js:no-line-numbers
fetch('https://wikinewsfeed.org/feed/atom')
```

  </CodeGroupItem>
  <CodeGroupItem title="NodeJS">

```js:no-line-numbers
const fetch = require('node-fetch')
fetch('https://wikinewsfeed.org/feed/atom')
```
  </CodeGroupItem>
  <CodeGroupItem title="Python">

```python:no-line-numbers
import requests
response = requests.get('https://wikinewsfeed.org/feed/atom')
```

  </CodeGroupItem>
</CodeGroup>

```xml
<?xml version="1.0" encoding="UTF-8"?>
<feed
  xmlns="http://www.w3.org/2005/Atom">
  <title>WikiNewsFeed</title>
  <id>http://localhost:8080</id>
  <updated></updated>
  <rights>Creative Commons Attribution-ShareAlike License 3.0</rights>
  <subtitle>News aggregator powered by Wikipedia</subtitle>
  <link href="http://localhost:8080"></link>
  <author>
    <name>Wikipedia contributors</name>
  </author>
  <entry>
    <title>COVID-19 pandemic in South Korea</title>
    <updated>2022-03-17T00:00:00Z</updated>
    <id>9e461cd86e3bf9635acea12388cfe9db1fe6283e</id>
    <content type="html">&lt;a href=&#34;https://en.wikipedia.org/wiki/South_Korea&#34; title=&#34;South Korea&#34;&gt;South Korea&lt;/a&gt; reports 621,328 new &lt;a href=&#34;https://en.wikipedia.org/wiki/COVID-19&#34; title=&#34;COVID-19&#34;&gt;COVID-19&lt;/a&gt; cases, a new single day record. </content>
    <link href="https://en.wikipedia.org/wiki/Portal:Current_events/2022_March_17" rel="alternate"></link>
    <summary type="html">South Korea reports 621,328 new COVID-19 cases, a new single day record. </summary>
    <author>
      <name>Wikipedia contributors</name>
    </author>
  </entry>
</feed>
```
:::

### Events

`GET` [https://wikinewsfeed.org/api/events](https://wikinewsfeed.org/api/events)

#### Query Params

| Param           | Type    | Description                         |
|-----------------|---------|-------------------------------------|
| page            | String  | Wikipedia Page                      |
| includeOriginal | Boolean | Include unmodified text, body, date |

#### Response

[EventResponse](https://pkg.go.dev/github.com/wikinewsfeed/wikinewsfeed/web#EventsResponse)

::: details Example

<CodeGroup>
  <CodeGroupItem title="cURL" active>

```bash:no-line-numbers
curl https://wikinewsfeed.org/api/events
```

  </CodeGroupItem>
  <CodeGroupItem title="JavaScript">

```js:no-line-numbers
fetch('https://wikinewsfeed.org/api/events')
```

  </CodeGroupItem>
  <CodeGroupItem title="NodeJS">

```js:no-line-numbers
const fetch = require('node-fetch')
fetch('https://wikinewsfeed.org/api/events')
```
  </CodeGroupItem>
  <CodeGroupItem title="Python">

```python:no-line-numbers
import requests
response = requests.get('https://wikinewsfeed.org/api/events')
```

  </CodeGroupItem>
</CodeGroup>

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
:::

### Metrics

`GET` [https://wikinewsfeed.org/metrics](https://wikinewsfeed.org/metrics)

This endpoint exposes Prometheus Metrics, see [Monitoring](server.md#monitoring) for more

## Limits

### Connections

20 concurrent connections are allowed simultaneously

## Cache

TTL is set to 30 minutes
