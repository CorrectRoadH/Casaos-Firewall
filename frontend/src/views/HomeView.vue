<script setup lang="ts">
import {ControllerPanel} from '../components/ControllerPanel';
import {MonitorPanel} from '../components/MonitorPanel';
import {SettingPanel} from '../components/SettingPanel';

</script>

<template>
  <main class="div">
    <div class="main">
      <ControllerPanel />
    </div>
    <div class="side">
      <div class="monitor">
        <MonitorPanel />
      </div>
      <div class="setting">
        <SettingPanel />
      </div>
    </div>
  </main>
</template>

<style scoped>
.div{
  background: url("http://127.0.0.1/img/wallpaper02.20ce0a96.jpg");
  display: flex;
  flex-direction: column;
  gap: 25px;
  padding: 25px;
  height: 100vh;
}
/* 这里做个响应式，如果大于550，那么main的order就是2 */

/* mobile */

.main{
  width: 100%;
  display: flex;
  flex-direction: column;
}
.side{
  width: 100%;

  display: flex;
  gap: 25px;
}
.monitor{
  width: 70%;
}
.setting{
  width: 30%;
}
/* desktop */

@media screen and (min-width: 550px) {
  .div{
    display: flex;
    flex-direction: row;
    height: 100vh;
    gap: 25px;
    padding: 25px;
  }
  .main{
    padding: 0px;
    width: 65%;
    height: 80%;
    order: 2;
  }
  .side{
    padding: 0px;
    display: flex;
    gap: 25px;
    flex-direction: column;
    width: 25%;
    height: 80%;
    order: 1;
  }
  .monitor{
    width: 100%;
    flex-grow: 1;
  }
  .setting{
    width: 100%;
    height: 30%;
  }
}
/*
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
}*/
</style>
