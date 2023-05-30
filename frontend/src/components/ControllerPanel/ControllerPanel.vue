<script setup lang="ts">
import {onMounted, ref } from 'vue';
import axios from 'axios';
import CButton from '../kit/CButton.vue';
import { useFirewallStore } from '../../stores/firewall';
import PortEdit from './PortEdit.vue';


const firewallState = useFirewallStore();
const baseHost = "http://127.0.0.1"

const closePort = async(port:string,protocol:string):Promise<any> => {
  const promise = axios.post(baseHost+'/v2/firewall/nftables',{
    port: port,
    protocol: protocol,
    action: "close"
  })
  return promise
}

const handelCloseBtnClick = async (port:string,protocol:string)=>{
  await closePort(port,protocol)
  firewallState.getPort()
  alert("success")
}

onMounted(()=>{
  firewallState.getPort()
})

defineExpose({
  firewallState
})
</script>

<template>
  <main class="backgroundContainer">
    <h1 class="title">Firewall</h1>
    <div>
        <p>These ports that didn't list in below is prohibited by default.</p>
    </div>
    <PortEdit/> 
    <!-- <div v-if="isLoading" >
      Loading...
    </div>
    <div v-if="isError">
      Fetching Error
    </div> -->
    <div v-if="firewallState.Ports">
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
          <tr v-for="(item) in firewallState.Ports" :key="item.port">
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
.table{
    width: 100%;
    border-collapse: collapse;
    margin: auto;
    margin-top: 20px;
    border-radius: 5px;
    flex-grow: 2;
}
</style>