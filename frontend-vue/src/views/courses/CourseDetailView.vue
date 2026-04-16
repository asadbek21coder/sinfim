<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Catalog</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">{{ form.title || 'Course detail' }}</h1>
          <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">
          Course settings, lessons, Telegram video references, materials and classes are managed from here.
        </p>
      </div>
      <div class="flex flex-wrap gap-2">
        <RouterLink class="btn-secondary" to="/app/courses">Back to courses</RouterLink>
        <a v-if="publicHref" class="btn-primary" :href="publicHref" target="_blank">Open public page</a>
      </div>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>
    <div v-if="successMessage" class="rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">{{ successMessage }}</div>

    <div class="grid gap-6 xl:grid-cols-[1fr_360px]">
      <section class="card space-y-5">
        <div>
          <h2 class="font-display text-xl font-bold text-primary">Course information</h2>
          <p class="mt-1 text-sm text-on-surface-variant">Slug is fixed after creation so public links do not break.</p>
        </div>

        <form class="grid gap-5 sm:grid-cols-2" @submit.prevent="save">
          <div class="sm:col-span-2">
            <label class="form-label" for="title">Title</label>
            <input id="title" v-model.trim="form.title" class="form-input" required>
          </div>
          <div>
            <label class="form-label">Slug</label>
            <input class="form-input" :value="course?.slug ?? ''" disabled>
          </div>
          <div>
            <label class="form-label" for="status">Internal status</label>
            <select id="status" v-model="form.status" class="form-select">
              <option value="draft">Draft</option>
              <option value="active">Active</option>
              <option value="archived">Archived</option>
            </select>
          </div>
          <div>
            <label class="form-label" for="public-status">Public status</label>
            <select id="public-status" v-model="form.public_status" class="form-select">
              <option value="draft">Draft</option>
              <option value="public">Public</option>
              <option value="hidden">Hidden</option>
            </select>
          </div>
          <div>
            <label class="form-label" for="level">Level</label>
            <input id="level" v-model.trim="form.level" class="form-input" placeholder="A1">
          </div>
          <div>
            <label class="form-label" for="category">Category</label>
            <input id="category" v-model.trim="form.category" class="form-input" placeholder="language">
          </div>
          <div class="sm:col-span-2">
            <label class="form-label" for="description">Description</label>
            <textarea id="description" v-model.trim="form.description" class="form-textarea" rows="5" />
          </div>
          <div class="sm:col-span-2">
            <button class="btn-primary" type="submit" :disabled="isSaving || !course">{{ isSaving ? 'Saving...' : 'Save course' }}</button>
          </div>
        </form>
      </section>

      <aside class="space-y-4">
        <section class="card">
          <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Next</p>
          <h2 class="mt-2 font-display text-xl font-bold text-primary">Lessons and classes</h2>
          <p class="mt-2 text-sm leading-6 text-on-surface-variant">
            Create lessons, attach Telegram stream references and add PDF/material links before groups start learning.
          </p>
        </section>
        <section class="card">
          <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Public URL</p>
          <p class="mt-2 break-all text-sm font-semibold text-primary">{{ publicHref || 'Publish the course after school page is public.' }}</p>
        </section>
      </aside>
    </div>

    <section class="grid gap-6 xl:grid-cols-[1fr_360px]">
      <div class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
        <table class="data-table">
          <thead><tr><th>Lesson</th><th>Publish</th><th>Content</th><th>Status</th><th>Action</th></tr></thead>
          <tbody>
            <tr v-if="lessons.length === 0"><td colspan="5" class="text-center text-on-surface-variant">No lessons for this course yet.</td></tr>
            <tr v-for="item in lessons" :key="item.id">
              <td><p class="font-semibold text-primary">{{ item.order_number }}. {{ item.title }}</p><p class="mt-1 text-xs text-on-surface-variant">{{ item.description || 'No description yet' }}</p></td>
              <td>Day {{ item.publish_day }}</td>
              <td><span class="text-sm text-on-surface-variant">{{ item.has_video ? 'Video' : 'No video' }} · {{ item.material_count }} material</span></td>
              <td><span class="chip-approved">{{ item.status }}</span></td>
              <td><RouterLink class="btn-secondary" :to="`/app/lessons/${item.id}/edit`">Edit</RouterLink></td>
            </tr>
          </tbody>
        </table>
      </div>

      <section class="card space-y-4">
        <div>
          <h2 class="font-display text-xl font-bold text-primary">New lesson</h2>
          <p class="mt-1 text-sm text-on-surface-variant">Start with the lesson shell; video and materials are added in the editor.</p>
        </div>
        <form class="space-y-4" @submit.prevent="createLesson">
          <div>
            <label class="form-label" for="lesson-title">Title</label>
            <input id="lesson-title" v-model.trim="lessonForm.title" class="form-input" required placeholder="Lesson 1 - Alphabet">
          </div>
          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="form-label" for="lesson-order">Order</label>
              <input id="lesson-order" v-model.number="lessonForm.order_number" class="form-input" min="1" type="number">
            </div>
            <div>
              <label class="form-label" for="publish-day">Publish day</label>
              <input id="publish-day" v-model.number="lessonForm.publish_day" class="form-input" min="1" type="number">
            </div>
          </div>
          <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isCreatingLesson || !course">{{ isCreatingLesson ? 'Creating...' : 'Create lesson' }}</button>
        </form>
      </section>
    </section>

    <section class="grid gap-6 xl:grid-cols-[1fr_360px]">
      <div class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
        <table class="data-table">
          <thead><tr><th>Class/group</th><th>Students</th><th>Status</th><th>Action</th></tr></thead>
          <tbody>
            <tr v-if="classes.length === 0"><td colspan="4" class="text-center text-on-surface-variant">No classes for this course yet.</td></tr>
            <tr v-for="item in classes" :key="item.id">
              <td><p class="font-semibold text-primary">{{ item.name }}</p><p class="mt-1 text-xs text-on-surface-variant">{{ cadenceLabel(item.lesson_cadence) }}</p></td>
              <td>{{ item.student_count }}</td>
              <td><span class="chip-approved">{{ item.status }}</span></td>
              <td><RouterLink class="btn-secondary" :to="`/app/classes/${item.id}`">Open</RouterLink></td>
            </tr>
          </tbody>
        </table>
      </div>

      <section class="card space-y-4">
        <div>
          <h2 class="font-display text-xl font-bold text-primary">New class/group</h2>
          <p class="mt-1 text-sm text-on-surface-variant">A class is where real students, access and payment state live.</p>
        </div>
        <form class="space-y-4" @submit.prevent="createClass">
          <div>
            <label class="form-label" for="class-name">Name</label>
            <input id="class-name" v-model.trim="classForm.name" class="form-input" required placeholder="Russian A1 - May group">
          </div>
          <div>
            <label class="form-label" for="start-date">Start date</label>
            <input id="start-date" v-model="classForm.start_date" class="form-input" type="date">
          </div>
          <div>
            <label class="form-label" for="cadence">Lesson cadence</label>
            <select id="cadence" v-model="classForm.lesson_cadence" class="form-select">
              <option value="daily">Daily</option>
              <option value="every_other_day">Every other day</option>
              <option value="weekly_3">Three times a week</option>
              <option value="manual">Manual</option>
            </select>
          </div>
          <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isCreatingClass || !course || !workspace">{{ isCreatingClass ? 'Creating...' : 'Create class' }}</button>
        </form>
      </section>
    </section>
  </section>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { catalogApi } from '@/api/catalog'
import { classroomApi } from '@/api/classroom'
import { organizationApi } from '@/api/organization'
import type { CourseDto, CoursePublicStatus, CourseStatus, LessonSummaryDto } from '@/types/catalog'
import type { ClassSummaryDto, LessonCadence } from '@/types/classroom'
import type { WorkspaceDto } from '@/types/organization'

const route = useRoute()
const workspace = ref<WorkspaceDto | null>(null)
const course = ref<CourseDto | null>(null)
const isSaving = ref(false)
const isCreatingClass = ref(false)
const isCreatingLesson = ref(false)
const errorMessage = ref('')
const successMessage = ref('')
const classes = ref<ClassSummaryDto[]>([])
const lessons = ref<LessonSummaryDto[]>([])

const form = reactive({
  title: '',
  description: '',
  category: '',
  level: '',
  status: 'draft' as CourseStatus,
  public_status: 'draft' as CoursePublicStatus,
})

const classForm = reactive({
  name: '',
  start_date: '',
  lesson_cadence: 'every_other_day' as LessonCadence,
})

const lessonForm = reactive({
  title: '',
  order_number: 0,
  publish_day: 0,
})

const publicHref = computed(() => {
  if (!workspace.value || !course.value) return ''
  return `/${workspace.value.slug}/courses/${course.value.slug}`
})

async function load() {
  errorMessage.value = ''
  try {
    const [workspaces, detail] = await Promise.all([
      organizationApi.listMyWorkspaces(),
      catalogApi.getCourseDetail(String(route.params.courseId)),
    ])
    workspace.value = workspaces.data.items[0] ?? null
    course.value = detail.data.item
    form.title = course.value.title
    form.description = course.value.description ?? ''
    form.category = course.value.category ?? ''
    form.level = course.value.level ?? ''
    form.status = course.value.status
    form.public_status = course.value.public_status
    await Promise.all([loadLessons(), loadClasses()])
  } catch {
    errorMessage.value = 'Course detail could not be loaded.'
  }
}

async function loadLessons() {
  if (!course.value) return
  const response = await catalogApi.listLessons({ course_id: course.value.id, limit: 200 })
  lessons.value = response.data.items
}

async function createLesson() {
  if (!course.value) return
  isCreatingLesson.value = true
  errorMessage.value = ''
  successMessage.value = ''
  try {
    const response = await catalogApi.createLesson({
      course_id: course.value.id,
      title: lessonForm.title,
      order_number: lessonForm.order_number || undefined,
      publish_day: lessonForm.publish_day || undefined,
    })
    lessonForm.title = ''
    lessonForm.order_number = 0
    lessonForm.publish_day = 0
    successMessage.value = 'Lesson created.'
    await loadLessons()
    window.setTimeout(() => {
      window.location.href = `/app/lessons/${response.data.item.id}/edit`
    }, 250)
  } catch {
    errorMessage.value = 'Lesson could not be created. Check order number conflicts.'
  } finally {
    isCreatingLesson.value = false
  }
}

async function loadClasses() {
  if (!workspace.value || !course.value) return
  const response = await classroomApi.listClasses({ organization_id: workspace.value.id, course_id: course.value.id, limit: 100 })
  classes.value = response.data.items
}

async function createClass() {
  if (!workspace.value || !course.value) return
  isCreatingClass.value = true
  errorMessage.value = ''
  successMessage.value = ''
  try {
    const response = await classroomApi.createClass({
      organization_id: workspace.value.id,
      course_id: course.value.id,
      name: classForm.name,
      start_date: optional(classForm.start_date),
      lesson_cadence: classForm.lesson_cadence,
    })
    classes.value = [response.data.item, ...classes.value]
    classForm.name = ''
    classForm.start_date = ''
    classForm.lesson_cadence = 'every_other_day'
    successMessage.value = 'Class created.'
  } catch {
    errorMessage.value = 'Class could not be created.'
  } finally {
    isCreatingClass.value = false
  }
}

function cadenceLabel(value: LessonCadence) {
  return { daily: 'Daily', every_other_day: 'Every other day', weekly_3: 'Three times a week', manual: 'Manual' }[value]
}

async function save() {
  if (!course.value) return
  isSaving.value = true
  successMessage.value = ''
  errorMessage.value = ''
  try {
    const response = await catalogApi.updateCourse({
      id: course.value.id,
      title: form.title,
      description: optional(form.description),
      category: optional(form.category),
      level: optional(form.level),
      status: form.status,
      public_status: form.public_status,
    })
    course.value = response.data.item
    successMessage.value = 'Course saved.'
  } catch {
    errorMessage.value = 'Course could not be saved.'
  } finally {
    isSaving.value = false
  }
}

function optional(value?: string) {
  return value && value.trim() ? value.trim() : undefined
}

onMounted(load)
</script>
