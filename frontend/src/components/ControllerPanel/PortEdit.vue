<script setup lang="ts">
import { ref } from 'vue';

const portRef = ref("8080")
const protocolRef = ref("tcp")

</script>

<template>
    <div class="input">
        Port:
        <input v-model="portRef" type="text" />
        <select v-model="protocolRef" name="protocol" id="protocol">
          <option value="tcp">TCP</option>
          <option value="udp">UDP</option>
        </select>
        
        <CButton>Open</CButton>
      </div>
</template>
