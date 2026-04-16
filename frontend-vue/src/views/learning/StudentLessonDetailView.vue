<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Lesson</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">{{ detail?.lesson.title || 'Lesson detail' }}</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">Video, materials and completion state for this lesson.</p>
      </div>
      <RouterLink class="btn-secondary" to="/learn/dashboard">Back to lessons</RouterLink>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>
    <div v-if="successMessage" class="rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">{{ successMessage }}</div>

    <section v-if="detail?.lesson.status === 'locked'" class="rounded-lg border border-amber-200 bg-amber-50 p-5 text-amber-900">
      <h2 class="font-display text-xl font-bold">Lesson is locked</h2>
      <p class="mt-2 text-sm">This lesson opens {{ detail.lesson.available_at ? formatDate(detail.lesson.available_at) : 'later' }}.</p>
    </section>

    <section v-else class="grid gap-6 xl:grid-cols-[1fr_360px]">
      <div class="space-y-6">
        <div class="rounded-lg border border-outline-variant bg-primary p-8 text-white shadow-sm">
          <p class="text-sm uppercase tracking-wide text-white/70">Video player</p>
          <h2 class="mt-3 font-display text-2xl font-bold">{{ detail?.video ? detail.video.provider : 'No video attached' }}</h2>
          <p class="mt-3 break-all text-sm text-white/80">{{ detail?.video?.stream_ref || detail?.video?.embed_url || 'The teacher has not attached a video yet.' }}</p>
          <a v-if="detail?.video?.embed_url" class="mt-5 inline-flex rounded-lg bg-white px-4 py-2 text-sm font-bold text-primary" :href="detail.video.embed_url" target="_blank">Open video source</a>
        </div>

        <div class="card">
          <h2 class="font-display text-xl font-bold text-primary">About this lesson</h2>
          <p class="mt-2 text-sm leading-6 text-on-surface-variant">{{ detail?.lesson.description || 'No extra description yet.' }}</p>
        </div>

        <div class="card space-y-4">
          <div>
            <h2 class="font-display text-xl font-bold text-primary">Homework</h2>
            <p class="mt-1 text-sm text-on-surface-variant">{{ homeworkState.item ? homeworkState.item.instructions || 'Submit your answer for this lesson.' : 'No homework for this lesson yet.' }}</p>
          </div>
          <div v-if="homeworkState.submission" class="rounded-lg border border-emerald-200 bg-emerald-50 p-4 text-sm text-emerald-900">
            <p class="font-semibold">Submitted</p>
            <p class="mt-1">Attempt {{ homeworkState.submission.attempt_number }} · {{ homeworkState.submission.status }}</p>
            <p v-if="homeworkState.submission.score !== null && homeworkState.submission.score !== undefined" class="mt-1">Score: {{ homeworkState.submission.score }} / {{ homeworkState.submission.max_score }}</p>
            <p v-if="homeworkState.submission.feedback" class="mt-2 whitespace-pre-wrap rounded-lg bg-white/70 p-3 text-emerald-950">{{ homeworkState.submission.feedback }}</p>
          </div>
          <form v-if="homeworkState.item" class="space-y-3" @submit.prevent="submitHomework">
            <div v-if="homeworkState.item.submission_type === 'text'">
              <label class="form-label" for="text-answer">Text answer</label>
              <textarea id="text-answer" v-model.trim="submissionForm.text_answer" class="form-textarea" rows="5" placeholder="Write your answer here" />
            </div>
            <div v-else-if="homeworkState.item.submission_type === 'file'">
              <label class="form-label" for="file-url">File/photo URL</label>
              <input id="file-url" v-model.trim="submissionForm.file_url" class="form-input" placeholder="https://...">
            </div>
            <div v-else-if="homeworkState.item.submission_type === 'audio'">
              <label class="form-label" for="audio-url">Audio URL</label>
              <input id="audio-url" v-model.trim="submissionForm.audio_url" class="form-input" placeholder="https://...">
            </div>
            <div v-else class="space-y-4">
              <div v-for="question in homeworkState.quiz_questions" :key="question.id" class="rounded-lg border border-outline-variant p-4">
                <p class="font-semibold text-primary">{{ question.order_number }}. {{ question.prompt }}</p>
                <div class="mt-3 space-y-2">
                  <label v-for="option in question.options" :key="option.id" class="flex items-center gap-2 rounded-lg border border-outline-variant px-3 py-2 text-sm text-primary">
                    <input v-model="submissionForm.quiz_answers[question.id]" type="radio" :name="question.id" :value="option.id">
                    {{ option.label }}
                  </label>
                </div>
              </div>
            </div>
            <button class="btn-primary justify-center" type="submit" :disabled="isSubmittingHomework">{{ isSubmittingHomework ? 'Submitting...' : homeworkState.submission ? 'Submit again' : 'Submit homework' }}</button>
          </form>
        </div>
      </div>

      <aside class="space-y-4">
        <section class="card space-y-3">
          <h2 class="font-display text-xl font-bold text-primary">Materials</h2>
          <p v-if="!detail?.materials.length" class="text-sm text-on-surface-variant">No materials yet.</p>
          <a v-for="item in detail?.materials ?? []" :key="item.id" class="block rounded-lg border border-outline-variant p-3 text-sm font-semibold text-primary hover:bg-surface-container" :href="item.url || '#'" target="_blank">
            {{ item.order_number }}. {{ item.title }}
            <span class="mt-1 block text-xs font-medium uppercase text-on-surface-variant">{{ item.material_type }}</span>
          </a>
        </section>

        <section class="card space-y-3">
          <p class="text-sm text-on-surface-variant">Status</p>
          <p class="font-display text-xl font-bold text-primary">{{ detail?.lesson.completed ? 'Completed' : 'In progress' }}</p>
          <button class="btn-primary w-full justify-center py-3" type="button" :disabled="isCompleting || detail?.lesson.completed" @click="completeLesson">{{ isCompleting ? 'Saving...' : 'Mark completed' }}</button>
        </section>
      </aside>
    </section>
  </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { homeworkApi } from '@/api/homework'
import { learningApi } from '@/api/learning'
import type { StudentHomeworkResponse } from '@/types/homework'
import type { LearningLessonDetailResponse } from '@/types/learning'

const route = useRoute()
const detail = ref<LearningLessonDetailResponse | null>(null)
const errorMessage = ref('')
const successMessage = ref('')
const isCompleting = ref(false)
const isSubmittingHomework = ref(false)
const homeworkState = ref<StudentHomeworkResponse>({ item: null, quiz_questions: [], submission: null })
const submissionForm = ref({
  text_answer: '',
  file_url: '',
  audio_url: '',
  quiz_answers: {} as Record<string, string>,
})

const classId = String(route.query.class_id || '')
const lessonId = String(route.params.lessonId)

async function load() {
  if (!classId) {
    errorMessage.value = 'Class context is missing.'
    return
  }
  errorMessage.value = ''
  try {
    const response = await learningApi.getLessonDetail({ class_id: classId, lesson_id: lessonId })
    detail.value = response.data
    if (response.data.lesson.status !== 'locked') {
      await loadHomework()
    }
  } catch {
    errorMessage.value = 'Lesson could not be loaded.'
  }
}

async function loadHomework() {
  const response = await homeworkApi.getStudentHomework({ class_id: classId, lesson_id: lessonId })
  homeworkState.value = response.data
  if (response.data.submission) {
    submissionForm.value.text_answer = response.data.submission.text_answer ?? ''
    submissionForm.value.file_url = response.data.submission.file_url ?? ''
    submissionForm.value.audio_url = response.data.submission.audio_url ?? ''
  }
}

async function completeLesson() {
  if (!classId) return
  isCompleting.value = true
  errorMessage.value = ''
  successMessage.value = ''
  try {
    await learningApi.markLessonCompleted({ class_id: classId, lesson_id: lessonId })
    successMessage.value = 'Lesson marked as completed.'
    await load()
  } catch {
    errorMessage.value = 'Lesson could not be completed yet.'
  } finally {
    isCompleting.value = false
  }
}

async function submitHomework() {
  if (!homeworkState.value.item) return
  isSubmittingHomework.value = true
  errorMessage.value = ''
  successMessage.value = ''
  try {
    const quizAnswers = homeworkState.value.quiz_questions.map((question) => ({
      question_id: question.id,
      selected_option_id: submissionForm.value.quiz_answers[question.id],
    }))
    const response = await homeworkApi.submitHomework({
      definition_id: homeworkState.value.item.id,
      class_id: classId,
      text_answer: optional(submissionForm.value.text_answer),
      file_url: optional(submissionForm.value.file_url),
      audio_url: optional(submissionForm.value.audio_url),
      quiz_answers: homeworkState.value.item.submission_type === 'quiz' ? quizAnswers : [],
    })
    successMessage.value = response.data.quiz_score !== null && response.data.quiz_score !== undefined ? `Homework submitted. Score: ${response.data.quiz_score}` : 'Homework submitted.'
    await loadHomework()
  } catch {
    errorMessage.value = 'Homework could not be submitted yet.'
  } finally {
    isSubmittingHomework.value = false
  }
}

function formatDate(value: string) {
  return new Date(value).toLocaleDateString()
}

function optional(value?: string) {
  return value && value.trim() ? value.trim() : undefined
}

onMounted(load)
</script>
