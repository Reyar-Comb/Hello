<script setup>
import { ref } from 'vue'

const message = ref('Clock')

const getData = async() => {
  try {
    const response = await fetch('/api/hello')

    const json = await response.json()
    message.value = json.message
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.value = 'Clock'
  } catch(error){
    message.value = 'this is a bug'
  }
}
</script>

<template>
  <div class="box">
    <h1>{{ message }}</h1>
    <button @click="getData">What time is it?</button>
  </div>
</template>

<style scoped>
.box { text-align: center; margin-top: 50px; }
button { padding: 10px 20px; font-size: 16px; cursor: pointer;}
</style>
