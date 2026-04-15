<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Classroom</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">{{ item?.name ?? 'Class detail' }}</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">Manage accepted students and their manual access/payment state.</p>
      </div>
      <RouterLink class="btn-secondary" to="/app/classes">Back to classes</RouterLink>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>
    <div v-if="successMessage" class="rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">{{ successMessage }}</div>

    <div class="grid gap-6 xl:grid-cols-[1fr_380px]">
      <section class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
        <table class="data-table">
          <thead><tr><th>Student</th><th>Access</th><th>Payment</th><th>Note</th><th>Action</th></tr></thead>
          <tbody>
            <tr v-if="isLoading"><td colspan="5" class="text-center text-on-surface-variant">Loading students...</td></tr>
            <tr v-else-if="students.length === 0"><td colspan="5" class="text-center text-on-surface-variant">No students yet.</td></tr>
            <tr v-for="student in students" v-else :key="student.student_user_id">
              <td><p class="font-semibold text-primary">{{ student.full_name }}</p><p class="mt-1 text-xs text-on-surface-variant">{{ student.phone_number }}</p></td>
              <td>
                <select class="form-select min-w-32" :value="student.access_status" @change="updateStudent(student, 'access_status', ($event.target as HTMLSelectElement).value)">
                  <option value="pending">Pending</option><option value="active">Active</option><option value="paused">Paused</option><option value="blocked">Blocked</option>
                </select>
              </td>
              <td>
                <select class="form-select min-w-36" :value="student.payment_status" @change="updateStudent(student, 'payment_status', ($event.target as HTMLSelectElement).value)">
                  <option value="unknown">Unknown</option><option value="pending">Pending</option><option value="confirmed">Confirmed</option><option value="rejected">Rejected</option>
                </select>
              </td>
              <td class="max-w-xs truncate">{{ student.note || 'No note' }}</td>
              <td><button class="btn-secondary" type="button" @click="saveAccess(student)">Save</button></td>
            </tr>
          </tbody>
        </table>
      </section>

      <section class="card space-y-5">
        <div>
          <h2 class="font-display text-xl font-bold text-primary">Add student</h2>
          <p class="mt-1 text-sm text-on-surface-variant">Use the phone number as login. The temporary password can be changed later.</p>
        </div>
        <form class="space-y-4" @submit.prevent="addStudent">
          <div><label class="form-label" for="name">Full name</label><input id="name" v-model.trim="form.full_name" class="form-input" required></div>
          <div><label class="form-label" for="phone">Phone number</label><input id="phone" v-model.trim="form.phone_number" class="form-input" required placeholder="+998 90 123 45 67"></div>
          <div><label class="form-label" for="pass">Temporary password</label><input id="pass" v-model="form.temporary_password" class="form-input" minlength="8" placeholder="TempPass123"></div>
          <div class="grid gap-4 sm:grid-cols-2 xl:grid-cols-1">
            <div><label class="form-label">Access</label><select v-model="form.access_status" class="form-select"><option value="pending">Pending</option><option value="active">Active</option><option value="paused">Paused</option><option value="blocked">Blocked</option></select></div>
            <div><label class="form-label">Payment</label><select v-model="form.payment_status" class="form-select"><option value="unknown">Unknown</option><option value="pending">Pending</option><option value="confirmed">Confirmed</option><option value="rejected">Rejected</option></select></div>
          </div>
          <div><label class="form-label" for="note">Note</label><textarea id="note" v-model.trim="form.note" class="form-textarea" rows="3" /></div>
          <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isSubmitting || !item">{{ isSubmitting ? 'Adding...' : 'Add student' }}</button>
        </form>
      </section>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { classroomApi } from '@/api/classroom'
import type { AccessStatus, ClassDto, PaymentStatus, StudentDto } from '@/types/classroom'

const route = useRoute()
const item = ref<ClassDto | null>(null)
const students = ref<StudentDto[]>([])
const isLoading = ref(true)
const isSubmitting = ref(false)
const errorMessage = ref('')
const successMessage = ref('')
const form = reactive({ full_name: '', phone_number: '', temporary_password: 'TempPass123', access_status: 'pending' as AccessStatus, payment_status: 'unknown' as PaymentStatus, note: '' })

async function load() {
  isLoading.value = true
  try {
    const response = await classroomApi.getClassDetail(String(route.params.classId))
    item.value = response.data.item
    students.value = response.data.students
  } catch { errorMessage.value = 'Class detail could not be loaded.' } finally { isLoading.value = false }
}

async function addStudent() {
  if (!item.value) return
  isSubmitting.value = true; errorMessage.value = ''; successMessage.value = ''
  try {
    await classroomApi.addStudent({ class_id: item.value.id, ...form, note: optional(form.note), temporary_password: optional(form.temporary_password) })
    successMessage.value = 'Student added.'
    form.full_name = ''; form.phone_number = ''; form.temporary_password = 'TempPass123'; form.access_status = 'pending'; form.payment_status = 'unknown'; form.note = ''
    await load()
  } catch { errorMessage.value = 'Student could not be added. It may already be in this class.' } finally { isSubmitting.value = false }
}

function updateStudent(student: StudentDto, key: 'access_status' | 'payment_status', value: string) {
  if (key === 'access_status') student.access_status = value as AccessStatus
  if (key === 'payment_status') student.payment_status = value as PaymentStatus
}

async function saveAccess(student: StudentDto) {
  if (!item.value) return
  errorMessage.value = ''; successMessage.value = ''
  try {
    await classroomApi.updateAccess({ class_id: item.value.id, student_user_id: student.student_user_id, access_status: student.access_status, payment_status: student.payment_status, note: student.note ?? undefined })
    successMessage.value = 'Access updated.'
    await load()
  } catch { errorMessage.value = 'Access could not be updated.' }
}

function optional(value?: string) { return value && value.trim() ? value.trim() : undefined }
onMounted(load)
</script>
