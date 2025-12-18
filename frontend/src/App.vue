<script setup>
import { ref } from 'vue'
import bgImage from '../assets/developing.png'

const message = ref('Clock')
const isLoading = ref(false)
const UpdateTime = "     " + BUILD_TIME
const getData = async() => {
  if (isLoading.value) return
  isLoading.value = true
  try {
    const response = await fetch('/api/hello')

    const json = await response.json()
    message.value = json.message
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.value = 'Clock'
  } catch(error){
    message.value = 'this is a bug'
  } finally {
    isLoading.value = false
  }
}


</script>

<template>
  <!--Reyar Here-->
  <div class="min-h-screen flex flex-col justify-between bg-linear-to-br from-gray-50 to-gray-100">
    <!-- Background Scrolling Text -->
    <div class="absolute inset-0 pointer-events-none overflow-hidden opacity-20">
      <div class="diagonal-scroll-container">
        <div 
          class="diagonal-row animate-scroll-diagonal bg-repeat-x" 
          v-for="i in 10" 
          :key="i"
          :style="{
            backgroundImage: `url(${bgImage})`,
            backgroundSize: 'auto 96px' 
          }"
        >
        </div>
      </div>
    </div>
    <div class="flex-1 flex items-center justify-center">
      <div class="bg-white/70 backdrop-blur-xl border border-white/50 rounded-2xl shadow-2xl p-8 md:p-12 max-w-md w-full mx-4">
        <div class="text-center mt-20">
          <h1 class="text-center text-6xl font-bold">{{ message }}</h1>
          <button 
            @click="getData"
            class="mt-20 px-5 py-2.5 bg-blue-600 text-white rounded hover:bg-blue-700 hover:scale-103 active:scale-95 transition"
          >
            What time is it?
          </button>
        </div>
      </div>
    </div>
    <footer class="text-center mb-5 text-lg text-gray-500">
      <span>Last Updated:</span>
      <span class="ml-4 font-bold text-gray-600">{{ UpdateTime }}</span>
      <div>
        <span class="ml-4 font-normal text-blue-800"><a href="https://space.bilibili.com/238730660?spm_id_from=333.1007.0.0" target="blank" rel="noopener noreferrer">@ 2025 Reyar</a></span>
      </div>
      
    </footer>
  </div>
  

</template>

<style>
.diagonal-scroll-container {
  transform: rotate(15deg) scale(1.5);
  transform-origin: center;
  width: 220%;
  height: 220%;
  position: absolute;
  top: -60%;
  left: -60%;
  display: flex;
  flex-direction: column;
  gap: 2rem;
  will-change: transform;
}

.diagonal-row {
  width: 400%;
  height: 5rem;
  margin-bottom: 2rem;
  gap: 10rem;
  animation: scroll-diagonal 150s linear infinite;
  will-change: transform;
}


.diagonal-row:nth-child(2n) {
  animation-duration: 130s;
  animation-delay: -9s;
  animation-direction: reverse;
}

@keyframes scroll-diagonal {
  from { transform: translate3d(0, 0, 0); }
  to { transform: translate3d(-50%, 0, 0); }
}
</style>

