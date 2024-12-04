<script setup lang="js">
import { useUserStore } from '@/store'
import { NButton, NForm, NFormItem, NH2, NInput, NSpace, useMessage } from 'naive-ui'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const userStore = useUserStore()
const message = useMessage()
const router = useRouter()
const route = useRoute()

const formRef = ref(null)
const formRules = computed(() => {
  return {
    username: {
      required: true,
      trigger: 'blur',
      message: 'Please enter your username',
    },
    password: {
      required: true,
      trigger: 'blur',
      message: 'Please enter your password',
    },
  }
})
const formModel = ref({
  username: '',
  password: '',
})

const isLoading = ref(false)

function handleLogin() {
  formRef.value?.validate(async (errors) => {
    if (errors)
      return
    isLoading.value = true

    try {
      await userStore.login({
        username: formModel.value.username,
        password: formModel.value.password,
      })
      message.success('Login success')
      const redirect = route.query.redirect
      const redirectUrl = redirect ? decodeURIComponent(redirect) : '/'
      router.push(redirectUrl)
    }
    catch (e) {
      console.error(e)
      message.error(e.message)
    }
    finally {
      isLoading.value = false
    }
  })
}
</script>

<template>
  <div class="h-screen">
    <NH2 class="text-center">
      Login Page
    </NH2>

    <NForm ref="formRef" :rules="formRules" :model="formModel" :show-label="false">
      <NFormItem path="username">
        <NInput v-model:value="formModel.username" clearable placeholder="Username" />
      </NFormItem>
      <NFormItem path="password">
        <NInput
          v-model:value="formModel.password" type="password" clearable placeholder="Password"
          show-password-on="click"
        />
      </NFormItem>
      <NSpace vertical :size="20">
        <NButton type="primary" :loading="isLoading" block size="large" :disabled="isLoading" @click="handleLogin()">
          Login
        </NButton>
      </NSpace>
    </NForm>
  </div>
</template>
