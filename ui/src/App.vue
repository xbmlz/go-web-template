<script setup lang="js">
import NaiveProvider from '@/components/common/NaiveProvider.vue'
import { NConfigProvider } from 'naive-ui'
</script>

<template>
  <NConfigProvider>
    <NaiveProvider>
      <RouterView />
    </NaiveProvider>
  </NConfigProvider>
</template>
