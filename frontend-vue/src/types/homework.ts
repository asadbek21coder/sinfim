export type HomeworkSubmissionType = 'text' | 'file' | 'audio' | 'quiz'
export type HomeworkStatus = 'draft' | 'published' | 'archived'

export interface HomeworkDefinitionDto {
  id: string
  organization_id: string
  course_id: string
  lesson_id: string
  title: string
  instructions?: string | null
  submission_type: HomeworkSubmissionType
  status: HomeworkStatus
  max_score: number
  due_days_after_publish?: number | null
  allow_resubmission: boolean
  created_at?: string | null
  updated_at?: string | null
}

export interface QuizOptionDto {
  id: string
  label: string
  is_correct: boolean
  order_number: number
}

export interface QuizQuestionDto {
  id: string
  prompt: string
  order_number: number
  points: number
  options: QuizOptionDto[]
}

export interface QuizOptionPayload {
  label: string
  is_correct: boolean
  order_number?: number
}

export interface QuizQuestionPayload {
  prompt: string
  order_number?: number
  points?: number
  options: QuizOptionPayload[]
}

export interface SaveHomeworkDefinitionPayload {
  lesson_id: string
  title: string
  instructions?: string
  submission_type: HomeworkSubmissionType
  status?: HomeworkStatus
  max_score?: number
  due_days_after_publish?: number
  allow_resubmission: boolean
  quiz_questions?: QuizQuestionPayload[]
}

export interface HomeworkDefinitionResponse {
  item?: HomeworkDefinitionDto | null
  quiz_questions: QuizQuestionDto[]
}

export interface HomeworkSubmissionDto {
  id: string
  organization_id: string
  definition_id: string
  lesson_id: string
  class_id: string
  student_user_id: string
  submission_type: HomeworkSubmissionType
  status: 'submitted' | 'reviewed' | 'needs_revision'
  attempt_number: number
  text_answer?: string | null
  file_url?: string | null
  audio_url?: string | null
  score?: number | null
  max_score: number
  auto_scored: boolean
  submitted_at?: string | null
  reviewed_at?: string | null
  reviewer_user_id?: string | null
  feedback?: string | null
}

export interface ReviewSubmissionSummaryDto extends HomeworkSubmissionDto {
  organization_name: string
  course_title: string
  lesson_title: string
  class_name: string
  homework_title: string
  student_full_name: string
  student_phone: string
}

export interface QuizAnswerDto {
  question_id: string
  question_prompt: string
  selected_option_id?: string | null
  selected_label?: string | null
  is_correct: boolean
  points: number
}

export interface ListReviewSubmissionsResponse {
  items: ReviewSubmissionSummaryDto[]
}

export interface ReviewSubmissionDetailResponse {
  item: ReviewSubmissionSummaryDto
  definition: HomeworkDefinitionDto
  quiz_answers: QuizAnswerDto[]
}

export interface ReviewSubmissionPayload {
  id: string
  status: 'reviewed' | 'needs_revision'
  score?: number
  feedback?: string
}

export interface ReviewSubmissionResponse {
  item: HomeworkSubmissionDto
}

export interface StudentHomeworkResponse {
  item?: HomeworkDefinitionDto | null
  quiz_questions: QuizQuestionDto[]
  submission?: HomeworkSubmissionDto | null
}

export interface SubmitHomeworkPayload {
  definition_id: string
  class_id: string
  text_answer?: string
  file_url?: string
  audio_url?: string
  quiz_answers?: Array<{ question_id: string; selected_option_id?: string }>
}

export interface SubmitHomeworkResponse {
  submission: HomeworkSubmissionDto
  quiz_score?: number | null
}
