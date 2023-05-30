<script setup lang="ts">
import {  RouterView } from 'vue-router'
import { useQueryProvider } from "vue-query";

useQueryProvider();

</script>

<template>
  <RouterView />
</template>

<style>
:root{
  --backDropColor: rgba(90, 108, 124, 0.4);
  --backDropBlur: none ;
  --backDropBorderRadius: 10px;  
}

.backgroundContainer{
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;

  padding: 25px;

  border-radius: var(--backDropBorderRadius);
  background: var(--backDropColor);
  backdrop-filter:var(--backDropBlur);
}

</style>
