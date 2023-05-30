<script setup lang="ts">
import ControllerArea from '../components/ControllerPanel';
import MontiorArea from '../components/MontiorPanel.vue';
import AIArea from '../components/AIArea.vue';

// 下面做个三栏布局
</script>

<template>
  <main class="div">
    <div class="second">
      <MontiorArea />
    </div>
    <div class="main">
      <ControllerArea />
    </div>
    <div class="third">
      <AIArea />
    </div>
  </main>
</template>

<style scoped>
.div{
  display: flex;  
  background: url("http://127.0.0.1/img/wallpaper02.20ce0a96.jpg");
  height: 100vh;
}
.main{
  width: 70%;
  height: 80%;
  border-radius: 25px;
  padding: 25px;
  margin: 25px;
  color: white;
}
.second{
  margin: 25px;
  padding: 25px;
  height: 80%;
  widows: 25%;
}
</style>
