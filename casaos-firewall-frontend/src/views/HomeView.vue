<script setup lang="ts">
import { onMounted, reactive } from 'vue';
const portList = reactive([])
const loadingState = reactive({
  loading: false,
  error: false,
  errorMessage: ''
})

onMounted(async () => {
  const response = await fetch('/ports')
  const data = await response.json()
  // portList.push(...data)
})
</script>

<template>
  <main>
    <h1>Firewall</h1>
    <table  class="table">
      <thead>
        <tr>
          <th>Port</th>
          <th>State</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>443</td>
          <td>Open</td>
          <td>
            <button>Close</button>
          </td>
        </tr>
      </tbody>
    </table>
  </main>
</template>
