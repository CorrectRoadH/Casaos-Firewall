<script setup lang="ts">
import {reactive } from 'vue';
import { useQuery } from "@tanstack/vue-query";

// const portList = reactive([])
const getPort = async () => {
  console.log("hello")
  const res = await fetch('http://127.0.0.1/v2/firewall/port')
  if(!res.ok){
    throw new Error('Network response was not ok')
  }
  return res.json()
}

console.log(getPort())
// const { isLoading, isFetching, isError, data, error } = useQuery({
//   queryKey: ['todos'],
//   queryFn: getPort,
// })

</script>

<template>
  <main>
    <h1>Firewall</h1>
    <div>
      Port:
      <input type="text" />
      <button>Open</button>
    </div>
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
