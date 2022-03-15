<template>
    <table style="width: 100%">
      <tr>
        <td>Subscribers</td>
        <td>{{stats.subscribers_total || 0}}</td>
      </tr>
    </table>
</template>

<script>
export default {
  data() {
    return {
      stats: {}
    }
  },
  async mounted(){
    const res = await fetch('https://wikinewsfeed.org/api/stats', {
      headers: {
        'Content-Type': 'application/json'
      }
    })

    this.stats = await res.json()
  }
}
</script>
