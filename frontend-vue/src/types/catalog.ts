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
