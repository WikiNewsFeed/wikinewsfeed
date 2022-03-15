<template>
  <ul>
    <li>
      <code>https://wikinewsfeed.org/feed/atom?subscribe={{subscriber_id}}</code>
    </li>
    <li>
      <code>https://wikinewsfeed.org/feed/rss?subscribe={{subscriber_id}}</code>
    </li>
    <li>
      <code>https://wikinewsfeed.org/feed/json?subscribe={{subscriber_id}}</code>
    </li>
  </ul>
</template>

<script>
import { customAlphabet } from 'nanoid'

export default {
  data() {
    return {
      subscriber_id: customAlphabet('1234567890abcdef', 10)()
    }
  },
  mounted(){
    // Override subscribe button with unique subscriber id
    const subscribeButton = document.querySelector('.action-button.primary')
    const subscribeButtonUrl = subscribeButton.getAttribute("href")
    subscribeButton.setAttribute("href", subscribeButtonUrl + `?subscribe=${this.subscriber_id}`)
  }
}
</script>
