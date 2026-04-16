<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Homework</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">Review inbox</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">Check student submissions, leave feedback, and return scores from one queue.</p>
      </div>
      <select v-model="statusFilter" class="form-select max-w-xs" @change="loadList">
        <option value="submitted">Pending review</option>
        <option value="reviewed">Reviewed</option>
        <option value="needs_revision">Needs revision</option>
        <option value="all">All submissions</option>
      </select>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>
    <div v-if="successMessage" class="rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">{{ successMessage }}</div>

    <section class="grid gap-6 xl:grid-cols-[420px_1fr]">
      <div class="card space-y-3">
        <div class="flex items-center justify-between gap-3">
          <h2 class="font-display text-xl font-bold text-primary">Submissions</h2>
          <button class="text-sm font-semibold text-secondary" type="button" @click="loadList">Refresh</button>
        </div>
        <p v-if="items.length === 0" class="rounded-lg border border-dashed border-outline-variant p-4 text-sm text-on-surface-variant">No submissions in this queue.</p>
        <button v-for="item in items" :key="item.id" class="w-full rounded-lg border border-outline-variant p-4 text-left hover:bg-surface-container" :class="selectedId === item.id ? 'border-secondary bg-surface-container' : ''" type="button" @click="selectSubmission(item.id)">
          <span class="text-xs font-bold uppercase text-on-surface-variant">{{ item.status }} · {{ item.submission_type }}</span>
          <span class="mt-1 block font-display text-lg font-bold text-primary">{{ item.student_full_name }}</span>
          <span class="mt-1 block text-sm text-on-surface-variant">{{ item.homework_title }} · {{ item.class_name }}</span>
          <span class="mt-1 block text-xs text-on-surface-variant">{{ item.lesson_title }}</span>
        </button>
      </div>

      <div v-if="detail" class="space-y-4">
        <section class="card space-y-3">
          <div class="flex flex-col gap-3 lg:flex-row lg:items-start lg:justify-between">
            <div>
              <p class="text-sm font-semibold uppercase tracking-wide text-secondary">{{ detail.item.course_title }}</p>
              <h2 class="mt-1 font-display text-2xl font-bold text-primary">{{ detail.item.homework_title }}</h2>
              <p class="mt-2 text-sm text-on-surface-variant">{{ detail.item.student_full_name }} · {{ detail.item.student_phone }}</p>
            </div>
            <span class="rounded-lg bg-surface-container px-3 py-2 text-sm font-bold uppercase text-primary">{{ detail.item.status }}</span>
          </div>
          <p class="text-sm leading-6 text-on-surface-variant">{{ detail.definition.instructions || 'No instructions.' }}</p>
        </section>

        <section class="card space-y-4">
          <h3 class="font-display text-xl font-bold text-primary">Student answer</h3>
          <p v-if="detail.item.text_answer" class="whitespace-pre-wrap rounded-lg border border-outline-variant p-4 text-sm text-primary">{{ detail.item.text_answer }}</p>
          <a v-if="detail.item.file_url" class="block rounded-lg border border-outline-variant p-4 text-sm font-semibold text-primary" :href="detail.item.file_url" target="_blank">Open file/photo URL</a>
          <a v-if="detail.item.audio_url" class="block rounded-lg border border-outline-variant p-4 text-sm font-semibold text-primary" :href="detail.item.audio_url" target="_blank">Open audio URL</a>
          <div v-if="detail.quiz_answers.length" class="space-y-2">
            <div v-for="answer in detail.quiz_answers" :key="answer.question_id" class="rounded-lg border border-outline-variant p-4 text-sm">
              <p class="font-semibold text-primary">{{ answer.question_prompt }}</p>
              <p class="mt-1 text-on-surface-variant">{{ answer.selected_label || 'No answer' }} · {{ answer.is_correct ? 'Correct' : 'Incorrect' }} · {{ answer.points }} pts</p>
            </div>
          </div>
          <p v-if="!detail.item.text_answer && !detail.item.file_url && !detail.item.audio_url && !detail.quiz_answers.length" class="text-sm text-on-surface-variant">No answer content saved.</p>
        </section>

        <form class="card space-y-4" @submit.prevent="review">
          <h3 class="font-display text-xl font-bold text-primary">Review</h3>
          <div class="grid gap-3 sm:grid-cols-2">
            <select v-model="reviewForm.status" class="form-select">
              <option value="reviewed">Reviewed</option>
              <option value="needs_revision">Needs revision</option>
            </select>
            <input v-model.number="reviewForm.score" class="form-input" min="0" type="number" :max="detail.definition.max_score" placeholder="Score">
          </div>
          <textarea v-model.trim="reviewForm.feedback" class="form-textarea" rows="5" placeholder="Feedback for the student" />
          <button class="btn-primary justify-center" type="submit" :disabled="isSaving">{{ isSaving ? 'Saving...' : 'Save review' }}</button>
        </form>
      </div>

      <div v-else class="card flex min-h-[320px] items-center justify-center text-center text-sm text-on-surface-variant">Select a submission to review.</div>
    </section>
  </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { homeworkApi } from '@/api/homework'
import type { ReviewSubmissionDetailResponse, ReviewSubmissionSummaryDto } from '@/types/homework'

const items = ref<ReviewSubmissionSummaryDto[]>([])
const detail = ref<ReviewSubmissionDetailResponse | null>(null)
const selectedId = ref('')
const statusFilter = ref('submitted')
const errorMessage = ref('')
const successMessage = ref('')
const isSaving = ref(false)
const reviewForm = ref({ status: 'reviewed' as 'reviewed' | 'needs_revision', score: undefined as number | undefined, feedback: '' })

async function loadList() {
  errorMessage.value = ''
  try {
    const response = await homeworkApi.listReviewSubmissions({ status: statusFilter.value, limit: 100 })
    items.value = response.data.items
    if (selectedId.value && !items.value.some((item) => item.id === selectedId.value)) {
      detail.value = null
      selectedId.value = ''
    }
  } catch {
    errorMessage.value = 'Submissions could not be loaded.'
  }
}

async function selectSubmission(id: string) {
  selectedId.value = id
  successMessage.value = ''
  errorMessage.value = ''
  try {
    const response = await homeworkApi.getReviewSubmission(id)
    detail.value = response.data
    reviewForm.value.status = response.data.item.status === 'needs_revision' ? 'needs_revision' : 'reviewed'
    reviewForm.value.score = response.data.item.score ?? undefined
    reviewForm.value.feedback = response.data.item.feedback ?? ''
  } catch {
    errorMessage.value = 'Submission detail could not be loaded.'
  }
}

async function review() {
  if (!detail.value) return
  isSaving.value = true
  successMessage.value = ''
  errorMessage.value = ''
  try {
    await homeworkApi.reviewSubmission({ id: detail.value.item.id, status: reviewForm.value.status, score: reviewForm.value.score, feedback: optional(reviewForm.value.feedback) })
    successMessage.value = 'Review saved.'
    await selectSubmission(detail.value.item.id)
    await loadList()
  } catch {
    errorMessage.value = 'Review could not be saved.'
  } finally {
    isSaving.value = false
  }
}

function optional(value?: string) {
  return value && value.trim() ? value.trim() : undefined
}

onMounted(loadList)
</script>
