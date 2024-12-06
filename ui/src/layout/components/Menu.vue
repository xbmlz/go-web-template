<script setup lang="js">
import { BookOutline } from '@vicons/ionicons5'
import { NIcon, NMenu } from 'naive-ui'
import { h } from 'vue'

function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

/** @type {import('naive-ui').MenuOption[]} */
const menuOptions = [
  {
    label: '且听风吟',
    key: 'hear-the-wind-sing',
    icon: renderIcon(BookOutline),
  },
]
</script>

<template>
  <NMenu
    :indent="20"
    :collapsed-width="64"
    :options="menuOptions"
  />
</template>
