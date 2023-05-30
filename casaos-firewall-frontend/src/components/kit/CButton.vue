<script setup lang="ts">
// props is title and click
import { defineProps, defineEmits } from 'vue';
const props = defineProps({
    click: {
        type: Function,
        required: true
    }
});
const emits = defineEmits(['click']);

const handleClick = () => {
    emits.click();
};
</script>
<template>
    <button class="button" @click="handleClick"><slot></slot></button>
</template>
<style scoped>
.button{
    background-color: #2075f3;
    border-color: transparent;
    color: #fff;
    font-size: .75rem;
    padding-bottom: calc(0.1875rem - 1px);
    padding-top: calc(0.1875rem - 1px);
    border-radius: 9999px;
    padding-left: 1.25em;
    padding-right: 1.25em;
}
</style>