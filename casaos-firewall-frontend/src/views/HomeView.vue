<script setup lang="ts">
import {reactive } from 'vue';
import { useQuery } from "@tanstack/vue-query";
import axios from 'axios';

interface port{
  port: string;
  protocol: string;
  action: string;
}

const PortMetaData = {
  fromJson(object: any):port{
    return {
      port: object.port ? object.port : "error port",
      protocol: object.protocol ? object.port : "error protocol",
      action: object.action ? object.port : "error action",
    }
  }
}

const portList = reactive([])
const getPort = async ():Promise<any> => {
  console.log("hello")
  const promise = axios.get('http://127.0.0.1/v2/firewall/port').then((res)=>{
    return PortMetaData.fromJson(res.data)
  })
  return promise
}

// console.log(getPort())
const { isLoading, isFetching, isError, data, error } = useQuery({
  queryKey: ['getPort'],
  queryFn: getPort,
})

</script>

<template>
  <main>
    <h1>Firewall</h1>
    <div>
      Port:
      <input type="text" />
      <button>Open</button>
    </div>
    <div v-if="isLoading" >
      Loading...
    </div>
    <div v-if="isError">
      Fetching Error
    </div>
    <div v-if="data">
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
  
      <table  class="table">
        <thead>
          <tr>
            <th>Port</th>
            <th>State</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item) in data" :key="item">
            <td>{{item.port}}</td>
            <td>Open</td>
            <td>
              <button>Close</button>
            </td>
          </tr>
        </tbody>
      </table>

    </div>
  </main>
</template>
