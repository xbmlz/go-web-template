<script setup lang="js">
import NaiveProvider from '@/components/common/NaiveProvider.vue'
import { NConfigProvider } from 'naive-ui'
</script>

<template>
  <NConfigProvider class="w-full h-full">
    <NaiveProvider>
      <RouterView />
    </NaiveProvider>
  </NConfigProvider>
</template>
