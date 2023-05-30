<script setup lang="ts">
import {ControllerPanel} from '../components/ControllerPanel';
import {MonitorPanel} from '../components/MonitorPanel';
import {SettingPanel} from '../components/SettingPanel';
// 下面做个响应式的三栏布局
</script>

<template>
  <main class="div">
    <div class="side">
      <div class="second">
        <MonitorPanel />
      </div>
      <div class="third">
        <SettingPanel />
  
      </div>
  
    </div>
    <div class="main">
      <ControllerPanel />
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
.side{
  width: 25%;
  height: 80%;
}
.second{
  margin: 25px;
  padding: 25px;
  height: 55%;
}
.third{
  margin: 25px;
  padding: 25px;
  height: 25%;
}
</style>
