<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Catalog</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">Courses</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">
          Create reusable course packages before attaching them to classes and student groups.
        </p>
      </div>
      <a v-if="workspace" class="btn-secondary" :href="`/${workspace.slug}`" target="_blank">Open school page</a>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>

    <div class="grid gap-6 xl:grid-cols-[1fr_380px]">
      <section class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
        <table class="data-table">
          <thead>
            <tr>
              <th>Course</th>
              <th>Status</th>
              <th>Public</th>
              <th>Updated</th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="isLoading"><td colspan="5" class="text-center text-on-surface-variant">Loading courses...</td></tr>
            <tr v-else-if="courses.length === 0"><td colspan="5" class="text-center text-on-surface-variant">No courses yet.</td></tr>
            <tr v-for="course in courses" v-else :key="course.id">
              <td>
                <p class="font-semibold text-primary">{{ course.title }}</p>
                <p class="mt-1 text-xs text-on-surface-variant">/{{ workspace?.slug }}/courses/{{ course.slug }}</p>
              </td>
              <td><span :class="statusClass(course.status)">{{ labelStatus(course.status) }}</span></td>
              <td><span :class="publicClass(course.public_status)">{{ labelPublicStatus(course.public_status) }}</span></td>
              <td>{{ formatDate(course.updated_at) }}</td>
              <td><RouterLink class="btn-secondary" :to="`/app/courses/${course.id}`">Open</RouterLink></td>
            </tr>
          </tbody>
        </table>
      </section>

      <section class="card space-y-5">
        <div>
          <h2 class="font-display text-xl font-bold text-primary">New course</h2>
          <p class="mt-1 text-sm text-on-surface-variant">Keep this as the content package. Classes and students come later.</p>
        </div>

        <form class="space-y-4" @submit.prevent="createCourse">
          <div>
            <label class="form-label" for="title">Title</label>
            <input id="title" v-model.trim="form.title" class="form-input" placeholder="Russian A1" required @input="syncSlug">
          </div>
          <div>
            <label class="form-label" for="slug">Slug</label>
            <input id="slug" v-model.trim="form.slug" class="form-input" placeholder="russian-a1" required>
          </div>
          <div>
            <label class="form-label" for="description">Description</label>
            <textarea id="description" v-model.trim="form.description" class="form-textarea" rows="4" placeholder="Short public and internal course description" />
          </div>
          <div class="grid gap-4 sm:grid-cols-2 xl:grid-cols-1">
            <div>
              <label class="form-label" for="category">Category</label>
              <input id="category" v-model.trim="form.category" class="form-input" placeholder="language">
            </div>
            <div>
              <label class="form-label" for="level">Level</label>
              <input id="level" v-model.trim="form.level" class="form-input" placeholder="A1">
            </div>
          </div>
          <div>
            <label class="form-label" for="public-status">Public status</label>
            <select id="public-status" v-model="form.public_status" class="form-select">
              <option value="draft">Draft</option>
              <option value="public">Public</option>
              <option value="hidden">Hidden</option>
            </select>
          </div>

          <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isSubmitting || !workspace">
            {{ isSubmitting ? 'Creating...' : 'Create course' }}
          </button>
        </form>
      </section>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { catalogApi } from '@/api/catalog'
import { organizationApi } from '@/api/organization'
import type { CourseDto, CoursePublicStatus, CourseStatus } from '@/types/catalog'
import type { WorkspaceDto } from '@/types/organization'

const workspace = ref<WorkspaceDto | null>(null)
const courses = ref<CourseDto[]>([])
const isLoading = ref(true)
const isSubmitting = ref(false)
const errorMessage = ref('')

const form = reactive({
  title: '',
  slug: '',
  description: '',
  category: '',
  level: '',
  public_status: 'draft' as CoursePublicStatus,
})

async function load() {
  isLoading.value = true
  errorMessage.value = ''
  try {
    const workspaces = await organizationApi.listMyWorkspaces()
    workspace.value = workspaces.data.items[0] ?? null
    if (!workspace.value) {
      courses.value = []
      return
    }
    const response = await catalogApi.listCourses({ organization_id: workspace.value.id, limit: 100 })
    courses.value = response.data.items
  } catch {
    errorMessage.value = 'Courses could not be loaded.'
  } finally {
    isLoading.value = false
  }
}

async function createCourse() {
  if (!workspace.value) return
  isSubmitting.value = true
  errorMessage.value = ''
  try {
    const response = await catalogApi.createCourse({
      organization_id: workspace.value.id,
      title: form.title,
      slug: slugify(form.slug),
      description: optional(form.description),
      category: optional(form.category),
      level: optional(form.level),
      public_status: form.public_status,
    })
    courses.value = [response.data.item, ...courses.value]
    form.title = ''
    form.slug = ''
    form.description = ''
    form.category = ''
    form.level = ''
    form.public_status = 'draft'
  } catch {
    errorMessage.value = 'Course could not be created. Check slug and fields.'
  } finally {
    isSubmitting.value = false
  }
}

function syncSlug() {
  if (form.slug) return
  form.slug = slugify(form.title)
}

function slugify(value: string) {
  return value.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/^-+|-+$/g, '')
}

function optional(value?: string) {
  return value && value.trim() ? value.trim() : undefined
}

function labelStatus(status: CourseStatus) {
  return { draft: 'Draft', active: 'Active', archived: 'Archived' }[status]
}

function labelPublicStatus(status: CoursePublicStatus) {
  return { draft: 'Draft', public: 'Public', hidden: 'Hidden' }[status]
}

function statusClass(status: CourseStatus) {
  return { draft: 'chip-pending', active: 'chip-approved', archived: 'chip-rejected' }[status]
}

function publicClass(status: CoursePublicStatus) {
  return { draft: 'chip-pending', public: 'chip-approved', hidden: 'chip-info' }[status]
}

function formatDate(value: string) {
  return new Intl.DateTimeFormat('uz-UZ', { dateStyle: 'medium', timeStyle: 'short' }).format(new Date(value))
}

onMounted(load)
</script>
