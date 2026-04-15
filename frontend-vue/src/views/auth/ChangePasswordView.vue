<template>
  <div>
    <p class="text-sm font-semibold uppercase tracking-wide text-secondary">First login</p>
    <h1 class="mt-2 font-display text-3xl font-bold text-primary">Set your password</h1>
    <p class="mt-3 text-sm leading-6 text-on-surface-variant">
      Temporary passwords are only for the first entry. Choose a personal password before opening the workspace.
    </p>

    <form class="mt-8 space-y-5" @submit.prevent="submit">
      <div>
        <label class="form-label" for="current-password">Current password</label>
        <input id="current-password" v-model="currentPassword" class="form-input" type="password" required>
      </div>
      <div>
        <label class="form-label" for="new-password">New password</label>
        <input id="new-password" v-model="newPassword" class="form-input" minlength="8" type="password" required>
      </div>
      <div>
        <label class="form-label" for="confirm-password">Confirm new password</label>
        <input id="confirm-password" v-model="confirmPassword" class="form-input" minlength="8" type="password" required>
      </div>
      <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">
        {{ errorMessage }}
      </div>
      <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isSubmitting">
        {{ isSubmitting ? 'Saving...' : 'Save password' }}
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const auth = useAuthStore()
const currentPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const errorMessage = ref('')
const isSubmitting = ref(false)

async function submit() {
  errorMessage.value = ''
  if (newPassword.value !== confirmPassword.value) {
    errorMessage.value = 'New passwords do not match.'
    return
  }

  isSubmitting.value = true
  try {
    await auth.changeMyPassword(currentPassword.value, newPassword.value)
    router.push('/app/dashboard')
  } catch {
    errorMessage.value = 'Password could not be changed. Check the current password and try again.'
  } finally {
    isSubmitting.value = false
  }
}
</script>
