const path = require('path');

module.exports = {
  // use default theme
  theme: '@vuepress/theme-default',
  templateBuild: path.join(__dirname, 'templates', 'index.build.html'),
  title: "WikiNewsFeed",
  description: "Newsfeed based on Wikipedia's current events",
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
          link: 'https://example.com/feed.atom',
        },
        {
          text: 'RSS',
          link: 'https://example.com/feed.rss',
        },
        {
          text: 'JSON',
          link: 'https://example.com/feed.json',
        }],
      },
      {
        text: 'Reference',
        children: ['/reference/api.md', '/reference/client.md', '/reference/parser.md'],
      },
      {
        text: 'Contribute',
        link: '/contribute',
      },
      {
        text: 'About',
        link: '/about',
      },
      {
        text: 'Donate',
        link: '/donate',
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
