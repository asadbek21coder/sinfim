<template>
  <div>
    <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Login</p>
    <h1 class="mt-2 font-display text-3xl font-bold text-primary">Enter your school workspace</h1>
    <p class="mt-3 text-sm leading-6 text-on-surface-variant">
      Use the phone number and password given by the platform admin or your school owner.
    </p>

    <form class="mt-8 space-y-5" @submit.prevent="submit">
      <div>
        <label class="form-label" for="phone">Phone number</label>
        <input id="phone" v-model="phone" class="form-input" placeholder="+998 90 000 00 00" type="tel">
      </div>
      <div>
        <label class="form-label" for="password">Password</label>
        <input id="password" v-model="password" class="form-input" placeholder="Temporary or personal password" type="password">
      </div>
      <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">
        {{ errorMessage }}
      </div>
      <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isSubmitting">
        {{ isSubmitting ? 'Signing in...' : 'Sign in' }}
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()
const phone = ref('')
const password = ref('')
const isSubmitting = ref(false)
const errorMessage = ref('')

async function submit() {
  isSubmitting.value = true
  errorMessage.value = ''
  try {
    await auth.login(phone.value, password.value)
    if (auth.user?.mustChangePassword) {
      router.push('/auth/change-password')
      return
    }
    const redirect = typeof route.query.redirect === 'string' ? route.query.redirect : '/app/dashboard'
    router.push(redirect)
  } catch {
    errorMessage.value = 'Phone number or password is incorrect.'
  } finally {
    isSubmitting.value = false
  }
}
</script>
