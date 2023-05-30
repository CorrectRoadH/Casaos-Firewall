<script setup lang="ts">
import {reactive, ref } from 'vue';
import { useQuery } from "vue-query";
import axios from 'axios';
import CButton from './kit/CButton.vue';

interface port{
  port: string;
  protocol: string;
  action: string;
}

const PortMetaData = {
  fromJson(object: any):port{
    return {
      port: object.port ? object.port : "error port",
      protocol: object.protocol ? object.protocol : "error protocol",
      action: object.action ? object.port : "error action",
    }
  }
}

const baseHost = "http://127.0.0.1"


const portList = reactive([])
const getPort = async ():Promise<any> => {
  const promise = axios.get(baseHost+'/v2/firewall/port').then((res)=>{
    return res.data.data.map((item:any)=>PortMetaData.fromJson(item))
  })
  return promise
}
const openPort = async(port:string,protocol:string):Promise<any> => {
  const promise = axios.post(baseHost+'/v2/firewall/port',{
    port: port,
    protocol: protocol,
    action: "open"
  })
  return promise
}

const closePort = async(port:string,protocol:string):Promise<any> => {
  const promise = axios.post(baseHost+'/v2/firewall/nftables',{
    port: port,
    protocol: protocol,
    action: "close"
  })
  return promise
}

const handelSaveBtnClick = async ()=>{
  console.log("hello")
  console.log(portRef.value,protocolRef.value)
  await openPort(portRef.value,protocolRef.value)
  alert("success")
}

const handelCloseBtnClick = async (port:string,protocol:string)=>{
  await closePort(port,protocol)
  alert("success")
}

const portRef = ref("8080")
const protocolRef = ref("tcp")

const { isLoading, isFetching, isError, data, error } = useQuery(
  "getPort",
  getPort
)

defineExpose({
  isLoading,
  isFetching,
  isError,
  data,
  error,
  portList,
  portRef,
  protocolRef,
})
</script>

<template>
  <main class="controllerContainer">
    <h1 class="title">Firewall</h1>
    <div>
        <p>These ports that didn't list in below is prohibited by default.</p>
    </div>
    <div class="input">
      Port:
      <input v-model="portRef" type="text" />
      <select v-model="protocolRef" name="protocol" id="protocol">
        <option value="tcp">TCP</option>
        <option value="udp">UDP</option>
      </select>
      
      <CButton @click="handelSaveBtnClick">Open</CButton>
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
            <th>Protocol</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item) in data" :key="item">
            <td>{{item.port}}</td>
            <td>Open</td>
            <td>{{item.protocol}}</td>
            <td>
              <CButton @click="handelCloseBtnClick(item.port,item.protocol)">Close</CButton>
            </td>
          </tr>
        </tbody>
      </table>

    </div>
  </main>
</template>
<style scoped>
.title{
    font-size: 20px;
    color: white;
}
.controllerContainer{
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;

  padding: 25px;

  border-radius: var(--backDropBorderRadius);
  background: var(--backDropColor);
  backdrop-filter:var(--backDropBlur);
}
.input{
}
.table{
    width: 100%;
    border-collapse: collapse;
    margin: auto;
    margin-top: 20px;
    border-radius: 5px;
    flex-grow: 2;
}
</style>