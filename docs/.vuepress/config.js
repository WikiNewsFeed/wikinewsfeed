const path = require('path');

module.exports = {
  theme: '@vuepress/theme-default',
  templateBuild: path.join(__dirname, 'templates', 'index.build.html'),
  title: "WikiNewsFeed",
  description: "News aggregator powered by Wikipedia",
  head: [
    ['meta', { name: 'application-name', content: 'WikiNewsFeed' }],
    ['meta', { name: 'apple-mobile-web-app-title', content: 'WikiNewsFeed' }],
    ['meta', { name: 'theme-color', content: '#f6f6f6' }],
    [
      'link',
      {
        rel: 'shortcut icon',
        href: `/favicon.ico`,
      },
    ],
    [
      'link',
      {
        rel: 'icon',
        type: 'image/png',
        sizes: '72x72',
        href: `/assets/favicon-72x72.png`,
      },
    ],
    [
      'link',
      {
        rel: 'apple-touch-icon',
        type: 'image/png',
        href: `/assets/apple-touch-icon.png`,
      },
    ],
    [
      'link',
      {
        rel: 'alternate',
        type: 'application/atom+xml',
        href: 'https://wikinewsfeed.org/feed/atom',
        title: 'WikiNewsFeed'
      }
    ],
    [
      'link',
      {
        rel: 'alternate',
        type: 'application/rss+xml',
        href: 'https://wikinewsfeed.org/feed/rss',
        title: 'WikiNewsFeed'
      }
    ],
    [
      'link',
      {
        rel: 'alternate',
        type: 'application/feed+json',
        href: 'https://wikinewsfeed.org/feed/json',
        title: 'WikiNewsFeed'
      }
    ]
  ],
  themeConfig: {
    repo: 'wikinewsfeed/wikinewsfeed',
    darkMode: false,
    navbar: [
      {
        text: 'Subscribe',
        children: [{
          text: 'Atom',
          link: 'https://wikinewsfeed.org/feed/atom',
        },
        {
          text: 'RSS',
          link: 'https://wikinewsfeed.org/feed/rss',
        },
        {
          text: 'JSON',
          link: 'https://wikinewsfeed.org/feed/json',
        }]
      },
      {
        text: 'Reference',
        children: ['/reference/api.md', '/reference/server.md', '/reference/client.md', '/reference/parser.md'],
      },
      {
        text: 'Contribute',
        link: '/contribute',
      },
      {
        text: 'Donate',
        link: '/donate',
      },
      {
        text: 'About',
        link: '/about',
      }
    ]
  },
  plugins: [
    [
      '@vuepress/plugin-shiki',
      {
        theme: 'github-light',
      }
    ]
  ]
}
