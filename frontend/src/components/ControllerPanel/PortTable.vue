<script setup lang="ts">
import { useFirewallStore } from '../../stores/firewall';
import CButton from '../kit/CButton.vue';

const firewallState = useFirewallStore();

const handelCloseBtnClick = async (port:string,protocol:string)=>{
  await firewallState.setPort(port,protocol,"close")
  firewallState.getPort()
  alert("success")
}

</script>

<template>
    <div class="table" v-if="firewallState.Ports">
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
</template>
<style scoped>
.table{
    width: 100%;
    border-collapse: collapse;
    margin: auto;
    border-radius: 5px;
    flex-grow: 2;
}
</style>
