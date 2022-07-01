<template>
  <div class="container">
    你好, <span style="color: #f56c6c">{{ authStore.userInfo.account }}</span>！
  </div>
</template>


<script lang="ts" setup>
import useAuthStore from '@/sotre/auth'


const authStore = useAuthStore()
</script>
