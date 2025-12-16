<script setup>
import { ref } from 'vue'

const message = ref('awaiting...')

const getData = async() => {
  try {
    const response = await fetch('https://localhost:8080/api/hello')

    const test = await response.text()
    message.value = text
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
