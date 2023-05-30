<script setup lang="ts">
import {onMounted } from 'vue';
import { useFirewallStore } from '../../stores/firewall';
import PortEdit from './PortEdit.vue';
import PortTable from './PortTable.vue';

const firewallState = useFirewallStore();

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
    <PortTable />
    <!-- <div v-if="isLoading" >
      Loading...
    </div>
    <div v-if="isError">
      Fetching Error
    </div> -->
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