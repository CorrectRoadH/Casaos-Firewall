<script setup lang="ts">
import axios from 'axios';
import { useQuery } from 'vue-query';

const baseHost = "http://127.0.0.1"

const getState = async ():Promise<any> => {
  const promise = axios.get(baseHost+'/v2/firewall/firewall').then((res)=>{
    return res.data.data
  })
  return promise
}

const { isLoading, isFetching, isError, data, error } = useQuery(
  "getState",
  getState
)
defineExpose({
  isLoading,
  isFetching,
  isError,
  data,
  error
})
</script>
<template>
    <main class="backgroundContainer">
        <div v-if="data">
            Firewall Status: 🟢 Running
        </div>
        <div v-if="!data">
            Firewall Status: 🔴 Not Running
        </div>
    </main>
</template>
<style scoped>
</style>