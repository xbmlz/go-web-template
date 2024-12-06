<script setup lang="js">
import { NLayout, NLayoutContent, NLayoutHeader, NLayoutSider, NScrollbar } from 'naive-ui'
import { RouterView } from 'vue-router'
import Menu from './components/Menu.vue'
</script>

<template>
  <NLayout has-sider class="w-full h-full">
    <NLayoutSider>
      <NScrollbar class="flex-1">
        <Menu />
      </NScrollbar>
    </NLayoutSider>

    <NLayout class="h-full flex flex-col">
      <NLayoutHeader>
        header
      </NLayoutHeader>

      <NLayoutContent class="flex-1 p-4 flex flex-col">
        <RouterView />
      </NLayoutContent>
    </NLayout>
  </NLayout>
</template>
