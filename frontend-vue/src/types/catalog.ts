export type CourseStatus = 'draft' | 'active' | 'archived'
export type CoursePublicStatus = 'draft' | 'public' | 'hidden'

export interface CourseDto {
  id: string
  organization_id: string
  title: string
  slug: string
  description?: string | null
  category?: string | null
  level?: string | null
  status: CourseStatus
  public_status: CoursePublicStatus
  created_at: string
  updated_at: string
}

export interface CreateCoursePayload {
  organization_id: string
  title: string
  slug: string
  description?: string
  category?: string
  level?: string
  public_status?: CoursePublicStatus
}

export interface UpdateCoursePayload {
  id: string
  title: string
  description?: string
  category?: string
  level?: string
  status: CourseStatus
  public_status: CoursePublicStatus
}

export interface CourseResponse {
  item: CourseDto
}

export interface ListCoursesResponse {
  items: CourseDto[]
}

export interface PublicCoursePageResponse {
  organization: {
    id: string
    name: string
    slug: string
    description?: string | null
    logo_url?: string | null
    category?: string | null
    contact_phone?: string | null
    telegram_url?: string | null
    is_demo: boolean
  }
  course: CourseDto
  lessons: unknown[]
  lead_form: {
    enabled: boolean
    required_fields: string[]
    source: string
  }
}

export type LessonStatus = 'draft' | 'published' | 'archived'
export type VideoProvider = 'telegram' | 'external'
export type MaterialType = 'pdf' | 'image' | 'doc' | 'link' | 'other'
export type MaterialSourceType = 'url' | 'filevault'

export interface LessonDto {
  id: string
  organization_id: string
  course_id: string
  title: string
  description?: string | null
  order_number: number
  publish_day: number
  status: LessonStatus
  created_at: string
  updated_at: string
}

export interface LessonSummaryDto extends LessonDto {
  has_video: boolean
  material_count: number
}

export interface LessonVideoDto {
  id: string
  organization_id: string
  lesson_id: string
  provider: VideoProvider
  stream_ref?: string | null
  telegram_channel_id?: string | null
  telegram_message_id?: string | null
  embed_url?: string | null
  duration_seconds?: number | null
  created_at: string
  updated_at: string
}

export interface LessonMaterialDto {
  id: string
  organization_id: string
  lesson_id: string
  title: string
  material_type: MaterialType
  source_type: MaterialSourceType
  url?: string | null
  file_id?: string | null
  order_number: number
  created_at: string
  updated_at: string
}

export interface CreateLessonPayload {
  course_id: string
  title: string
  description?: string
  order_number?: number
  publish_day?: number
  status?: LessonStatus
}

export interface LessonVideoPayload {
  enabled: boolean
  provider?: VideoProvider
  stream_ref?: string
  telegram_channel_id?: string
  telegram_message_id?: string
  embed_url?: string
  duration_seconds?: number
}

export interface LessonMaterialPayload {
  title: string
  material_type?: MaterialType
  source_type?: MaterialSourceType
  url?: string
  file_id?: string
  order_number?: number
}

export interface UpdateLessonPayload {
  id: string
  title: string
  description?: string
  order_number: number
  publish_day: number
  status: LessonStatus
  video?: LessonVideoPayload
  materials: LessonMaterialPayload[]
}

export interface LessonResponse {
  item: LessonDto
  video?: LessonVideoDto | null
  materials?: LessonMaterialDto[]
}

export interface ListLessonsResponse {
  items: LessonSummaryDto[]
}
