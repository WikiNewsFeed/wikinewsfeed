const path = require('path');

module.exports = {
  // use default theme
  theme: '@vuepress/theme-default',
  templateBuild: path.join(__dirname, 'templates', 'index.build.html'),
  title: "WikiNewsFeed",
  description: "News aggregator powered by Wikipedia",
  // configure default theme
  themeConfig: {
    repo: 'wikinewsfeed/wikinewsfeed',
    darkMode: false,
    navbar: [
      // NavbarItem
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
        }],
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
      },
    ],
  },
  plugins: [
    [
      '@vuepress/plugin-shiki',
      {
        theme: 'github-light',
      }
    ]
  ],
}
