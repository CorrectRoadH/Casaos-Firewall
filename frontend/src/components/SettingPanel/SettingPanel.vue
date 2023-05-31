<script setup lang="ts">
import CButton from '../kit/CButton.vue'
</script>
<template>
    <main class="backgroundContainer">
        <div class="version">        
            <p>firewall version: 0.1.0</p>
            <CButton>check update</CButton>
        </div>
    </main>
</template>
<style scoped>
.version{
    display: flex;
}
</style>