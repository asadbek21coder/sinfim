<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Lesson editor</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">{{ form.title || 'Edit lesson' }}</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">
          Prepare the lesson, Telegram stream reference and materials that will later open for students by publish day.
        </p>
      </div>
      <RouterLink v-if="lesson" class="btn-secondary" :to="`/app/courses/${lesson.course_id}`">Back to course</RouterLink>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>
    <div v-if="successMessage" class="rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">{{ successMessage }}</div>

    <form class="grid gap-6 xl:grid-cols-[1fr_380px]" @submit.prevent="save">
      <section class="card space-y-5">
        <div>
          <h2 class="font-display text-xl font-bold text-primary">Lesson details</h2>
          <p class="mt-1 text-sm text-on-surface-variant">Order controls sequence; publish day controls drip timing from class start date.</p>
        </div>
        <div>
          <label class="form-label" for="title">Title</label>
          <input id="title" v-model.trim="form.title" class="form-input" required>
        </div>
        <div class="grid gap-4 sm:grid-cols-3">
          <div>
            <label class="form-label" for="order-number">Order</label>
            <input id="order-number" v-model.number="form.order_number" class="form-input" min="1" required type="number">
          </div>
          <div>
            <label class="form-label" for="publish-day">Publish day</label>
            <input id="publish-day" v-model.number="form.publish_day" class="form-input" min="1" required type="number">
          </div>
          <div>
            <label class="form-label" for="status">Status</label>
            <select id="status" v-model="form.status" class="form-select">
              <option value="draft">Draft</option>
              <option value="published">Published</option>
              <option value="archived">Archived</option>
            </select>
          </div>
        </div>
        <div>
          <label class="form-label" for="description">Description</label>
          <textarea id="description" v-model.trim="form.description" class="form-textarea" rows="5" placeholder="Short lesson note for teacher and student context" />
        </div>

        <div class="rounded-lg border border-outline-variant p-4">
          <div class="flex items-center justify-between gap-4">
            <div>
              <h3 class="font-display text-lg font-bold text-primary">Video reference</h3>
              <p class="mt-1 text-sm text-on-surface-variant">MVP stores a Telegram or external stream reference; upload/transcoding is later.</p>
            </div>
            <label class="flex items-center gap-2 text-sm font-semibold text-primary"><input v-model="video.enabled" type="checkbox"> Enabled</label>
          </div>
          <div v-if="video.enabled" class="mt-4 grid gap-4 sm:grid-cols-2">
            <div>
              <label class="form-label" for="provider">Provider</label>
              <select id="provider" v-model="video.provider" class="form-select">
                <option value="telegram">Telegram</option>
                <option value="external">External</option>
              </select>
            </div>
            <div>
              <label class="form-label" for="duration">Duration seconds</label>
              <input id="duration" v-model.number="video.duration_seconds" class="form-input" min="0" type="number">
            </div>
            <div class="sm:col-span-2">
              <label class="form-label" for="stream-ref">Stream ref</label>
              <input id="stream-ref" v-model.trim="video.stream_ref" class="form-input" placeholder="telegram-stream-ref or internal playback token">
            </div>
            <div>
              <label class="form-label" for="channel-id">Telegram channel id</label>
              <input id="channel-id" v-model.trim="video.telegram_channel_id" class="form-input" placeholder="-100...">
            </div>
            <div>
              <label class="form-label" for="message-id">Telegram message id</label>
              <input id="message-id" v-model.trim="video.telegram_message_id" class="form-input" placeholder="123">
            </div>
            <div class="sm:col-span-2">
              <label class="form-label" for="embed-url">Embed URL</label>
              <input id="embed-url" v-model.trim="video.embed_url" class="form-input" placeholder="https://...">
            </div>
          </div>
        </div>
      </section>

      <aside class="space-y-4">
        <section class="card space-y-4">
          <div>
            <h2 class="font-display text-xl font-bold text-primary">Materials</h2>
            <p class="mt-1 text-sm text-on-surface-variant">Add PDF, image, doc or link metadata. File upload will plug into the same list later.</p>
          </div>
          <div v-if="materials.length === 0" class="rounded-lg border border-dashed border-outline-variant p-4 text-sm text-on-surface-variant">No materials yet.</div>
          <div v-for="(item, index) in materials" :key="index" class="rounded-lg border border-outline-variant p-3">
            <div class="flex items-center justify-between gap-3">
              <p class="text-sm font-semibold text-primary">Material {{ index + 1 }}</p>
              <button class="text-sm font-semibold text-red-700" type="button" @click="removeMaterial(index)">Remove</button>
            </div>
            <div class="mt-3 space-y-3">
              <input v-model.trim="item.title" class="form-input" required placeholder="Workbook PDF">
              <div class="grid grid-cols-2 gap-3">
                <select v-model="item.material_type" class="form-select">
                  <option value="pdf">PDF</option>
                  <option value="image">Image</option>
                  <option value="doc">Doc</option>
                  <option value="link">Link</option>
                  <option value="other">Other</option>
                </select>
                <input v-model.number="item.order_number" class="form-input" min="1" type="number">
              </div>
              <input v-model.trim="item.url" class="form-input" placeholder="https://...">
            </div>
          </div>
          <button class="btn-secondary w-full justify-center" type="button" @click="addMaterial">Add material</button>
        </section>

        <section class="card space-y-4">
          <div>
            <h2 class="font-display text-xl font-bold text-primary">Homework</h2>
            <p class="mt-1 text-sm text-on-surface-variant">Create one homework block for this lesson. Quiz answers are auto-scored.</p>
          </div>
          <label class="flex items-center gap-2 text-sm font-semibold text-primary"><input v-model="homework.enabled" type="checkbox"> Enabled</label>
          <div v-if="homework.enabled" class="space-y-3">
            <input v-model.trim="homework.title" class="form-input" placeholder="Practice after lesson" required>
            <textarea v-model.trim="homework.instructions" class="form-textarea" rows="4" placeholder="What should the student submit?" />
            <div class="grid grid-cols-2 gap-3">
              <select v-model="homework.submission_type" class="form-select">
                <option value="text">Text</option>
                <option value="file">File/photo URL</option>
                <option value="audio">Audio URL</option>
                <option value="quiz">Quiz</option>
              </select>
              <select v-model="homework.status" class="form-select">
                <option value="draft">Draft</option>
                <option value="published">Published</option>
                <option value="archived">Archived</option>
              </select>
            </div>
            <div class="grid grid-cols-2 gap-3">
              <input v-model.number="homework.max_score" class="form-input" min="0" type="number" placeholder="Max score">
              <input v-model.number="homework.due_days_after_publish" class="form-input" min="0" type="number" placeholder="Due days">
            </div>
            <label class="flex items-center gap-2 text-sm font-semibold text-primary"><input v-model="homework.allow_resubmission" type="checkbox"> Allow resubmission</label>

            <div v-if="homework.submission_type === 'quiz'" class="space-y-3 rounded-lg border border-outline-variant p-3">
              <div class="flex items-center justify-between gap-3">
                <p class="text-sm font-semibold text-primary">Quiz questions</p>
                <button class="text-sm font-semibold text-secondary" type="button" @click="addQuizQuestion">Add question</button>
              </div>
              <div v-for="(question, questionIndex) in homework.quiz_questions" :key="questionIndex" class="rounded-lg border border-outline-variant p-3">
                <div class="flex items-center justify-between gap-3">
                  <p class="text-xs font-bold uppercase text-on-surface-variant">Question {{ questionIndex + 1 }}</p>
                  <button class="text-sm font-semibold text-red-700" type="button" @click="removeQuizQuestion(questionIndex)">Remove</button>
                </div>
                <input v-model.trim="question.prompt" class="form-input mt-3" placeholder="Question prompt">
                <input v-model.number="question.points" class="form-input mt-3" min="0" type="number" placeholder="Points">
                <div class="mt-3 space-y-2">
                  <div v-for="(option, optionIndex) in question.options" :key="optionIndex" class="grid grid-cols-[1fr_auto] gap-2">
                    <input v-model.trim="option.label" class="form-input" placeholder="Answer option">
                    <label class="flex items-center gap-2 text-xs font-semibold text-primary"><input v-model="option.is_correct" type="checkbox"> Correct</label>
                  </div>
                </div>
              </div>
            </div>
            <button class="btn-secondary w-full justify-center" type="button" :disabled="isSavingHomework || !lesson" @click="saveHomework">{{ isSavingHomework ? 'Saving homework...' : 'Save homework' }}</button>
          </div>
        </section>

        <section class="card space-y-3">
          <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isSaving || !lesson">{{ isSaving ? 'Saving...' : 'Save lesson' }}</button>
          <p class="text-xs leading-5 text-on-surface-variant">Published lessons appear in student learning by class access and publish day.</p>
        </section>
      </aside>
    </form>
  </section>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { catalogApi } from '@/api/catalog'
import { homeworkApi } from '@/api/homework'
import type { LessonDto, LessonMaterialPayload, LessonStatus, VideoProvider } from '@/types/catalog'
import type { HomeworkStatus, HomeworkSubmissionType, QuizQuestionPayload } from '@/types/homework'

const route = useRoute()
const lesson = ref<LessonDto | null>(null)
const isSaving = ref(false)
const errorMessage = ref('')
const successMessage = ref('')
const isSavingHomework = ref(false)

const form = reactive({
  title: '',
  description: '',
  order_number: 1,
  publish_day: 1,
  status: 'draft' as LessonStatus,
})

const video = reactive({
  enabled: false,
  provider: 'telegram' as VideoProvider,
  stream_ref: '',
  telegram_channel_id: '',
  telegram_message_id: '',
  embed_url: '',
  duration_seconds: 0,
})

const materials = ref<LessonMaterialPayload[]>([])

const homework = reactive({
  enabled: false,
  title: '',
  instructions: '',
  submission_type: 'text' as HomeworkSubmissionType,
  status: 'draft' as HomeworkStatus,
  max_score: 100,
  due_days_after_publish: undefined as number | undefined,
  allow_resubmission: true,
  quiz_questions: [] as QuizQuestionPayload[],
})

async function load() {
  errorMessage.value = ''
  try {
    const response = await catalogApi.getLessonDetail(String(route.params.lessonId))
    lesson.value = response.data.item
    form.title = lesson.value.title
    form.description = lesson.value.description ?? ''
    form.order_number = lesson.value.order_number
    form.publish_day = lesson.value.publish_day
    form.status = lesson.value.status
    if (response.data.video) {
      video.enabled = true
      video.provider = response.data.video.provider
      video.stream_ref = response.data.video.stream_ref ?? ''
      video.telegram_channel_id = response.data.video.telegram_channel_id ?? ''
      video.telegram_message_id = response.data.video.telegram_message_id ?? ''
      video.embed_url = response.data.video.embed_url ?? ''
      video.duration_seconds = response.data.video.duration_seconds ?? 0
    }
    materials.value = (response.data.materials ?? []).map((item) => ({
      title: item.title,
      material_type: item.material_type,
      source_type: item.source_type,
      url: item.url ?? '',
      file_id: item.file_id ?? undefined,
      order_number: item.order_number,
    }))
    await loadHomework()
  } catch {
    errorMessage.value = 'Lesson detail could not be loaded.'
  }
}

async function loadHomework() {
  if (!lesson.value) return
  const response = await homeworkApi.getLessonHomework(lesson.value.id)
  if (!response.data.item) {
    homework.enabled = false
    homework.title = ''
    homework.instructions = ''
    homework.submission_type = 'text'
    homework.status = 'draft'
    homework.max_score = 100
    homework.due_days_after_publish = undefined
    homework.allow_resubmission = true
    homework.quiz_questions = []
    return
  }
  const item = response.data.item
  homework.enabled = true
  homework.title = item.title
  homework.instructions = item.instructions ?? ''
  homework.submission_type = item.submission_type
  homework.status = item.status
  homework.max_score = item.max_score
  homework.due_days_after_publish = item.due_days_after_publish ?? undefined
  homework.allow_resubmission = item.allow_resubmission
  homework.quiz_questions = response.data.quiz_questions.map((question) => ({
    prompt: question.prompt,
    order_number: question.order_number,
    points: question.points,
    options: question.options.map((option) => ({ label: option.label, is_correct: option.is_correct, order_number: option.order_number })),
  }))
}

function addMaterial() {
  materials.value.push({ title: '', material_type: 'pdf', source_type: 'url', url: '', order_number: materials.value.length + 1 })
}

function removeMaterial(index: number) {
  materials.value.splice(index, 1)
}

function addQuizQuestion() {
  homework.quiz_questions.push({
    prompt: '',
    order_number: homework.quiz_questions.length + 1,
    points: 1,
    options: [
      { label: '', is_correct: true, order_number: 1 },
      { label: '', is_correct: false, order_number: 2 },
    ],
  })
}

function removeQuizQuestion(index: number) {
  homework.quiz_questions.splice(index, 1)
}

async function save() {
  if (!lesson.value) return
  isSaving.value = true
  errorMessage.value = ''
  successMessage.value = ''
  try {
    const response = await catalogApi.updateLesson({
      id: lesson.value.id,
      title: form.title,
      description: optional(form.description),
      order_number: form.order_number,
      publish_day: form.publish_day,
      status: form.status,
      video: {
        enabled: video.enabled,
        provider: video.provider,
        stream_ref: optional(video.stream_ref),
        telegram_channel_id: optional(video.telegram_channel_id),
        telegram_message_id: optional(video.telegram_message_id),
        embed_url: optional(video.embed_url),
        duration_seconds: video.duration_seconds || undefined,
      },
      materials: materials.value.map((item, index) => ({ ...item, title: item.title.trim(), url: optional(item.url), order_number: item.order_number || index + 1 })),
    })
    lesson.value = response.data.item
    successMessage.value = 'Lesson saved.'
    await load()
  } catch {
    errorMessage.value = 'Lesson could not be saved. Check required material titles and order conflicts.'
  } finally {
    isSaving.value = false
  }
}

async function saveHomework() {
  if (!lesson.value || !homework.enabled) return
  isSavingHomework.value = true
  errorMessage.value = ''
  successMessage.value = ''
  try {
    await homeworkApi.saveDefinition({
      lesson_id: lesson.value.id,
      title: homework.title,
      instructions: optional(homework.instructions),
      submission_type: homework.submission_type,
      status: homework.status,
      max_score: homework.max_score,
      due_days_after_publish: homework.due_days_after_publish || undefined,
      allow_resubmission: homework.allow_resubmission,
      quiz_questions: homework.submission_type === 'quiz' ? homework.quiz_questions : [],
    })
    successMessage.value = 'Homework saved.'
    await loadHomework()
  } catch {
    errorMessage.value = 'Homework could not be saved. Check quiz questions and required fields.'
  } finally {
    isSavingHomework.value = false
  }
}

function optional(value?: string) {
  return value && value.trim() ? value.trim() : undefined
}

onMounted(load)
</script>
