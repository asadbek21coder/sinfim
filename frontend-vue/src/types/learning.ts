export interface LearningClassDto {
  id: string
  name: string
  course_id: string
  course_title: string
  access_status: string
  payment_status: string
}

export interface LearningLessonDto {
  lesson_id: string
  title: string
  description?: string | null
  status: 'available' | 'locked'
  available_at?: string | null
  order_number: number
  publish_day: number
  has_video: boolean
  has_materials: boolean
  material_count: number
  completed: boolean
}

export interface StudentDashboardResponse {
  student: { id: string }
  organization: { id: string; name: string; slug: string; logo_url?: string | null }
  class: LearningClassDto
  progress: { completed_lessons: number; total_lessons: number; percentage: number }
  locked: boolean
  lessons: LearningLessonDto[]
  classes: LearningClassDto[]
}

export interface LearningVideoDto {
  provider: 'telegram' | 'external'
  stream_ref?: string | null
  telegram_channel_id?: string | null
  telegram_message_id?: string | null
  embed_url?: string | null
  duration_seconds?: number | null
}

export interface LearningMaterialDto {
  id: string
  title: string
  material_type: string
  source_type: string
  url?: string | null
  file_id?: string | null
  order_number: number
}

export interface LearningLessonDetailResponse {
  lesson: {
    id: string
    title: string
    description?: string | null
    status: 'available' | 'locked'
    available_at?: string | null
    order_number: number
    publish_day: number
    completed: boolean
  }
  video?: LearningVideoDto | null
  materials: LearningMaterialDto[]
}

export interface MarkLessonCompletedResponse {
  lesson_id: string
  class_id: string
  student_user_id: string
  completed: boolean
  completed_at: string
}
