<script setup>
import { NButton } from 'naive-ui'
import { ref } from 'vue'

const count = ref(0)
</script>

<template>
  <div class="text-3xl font-bold underline bg-black">
    <NButton type="primary" @click="count++">
      +
    </NButton>
    {{ count }}
  </div>
</template>
