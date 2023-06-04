<script setup lang="ts">
import { ref } from 'vue';
import { useFirewallStore } from '../../stores/firewall';
import CButton from '../kit/CButton.vue';

const firewallState = useFirewallStore();

const portRef = ref("8080")
const protocolRef = ref("tcp")
const handelSaveBtnClick = async ()=>{
  await firewallState.setPort(portRef.value,protocolRef.value,"open")
  firewallState.getPort()
  alert("success")
}

</script>

<template>
    <div class="input">
        Port:
        <input v-model="portRef" type="text" />
        <select v-model="protocolRef" name="protocol" id="protocol">
          <option value="tcp">TCP</option>
          <option value="udp">UDP</option>
        </select>
        
        <CButton @click="handelSaveBtnClick">Open</CButton>
      </div>
</template>

<style scoped>
.input{
  display: flex;
  gap:5px;
}
input{
  width: 100px;
}
</style>