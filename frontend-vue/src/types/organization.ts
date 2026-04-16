export type SchoolRequestStatus = 'new' | 'contacted' | 'approved' | 'rejected'

export interface CreateSchoolRequestPayload {
  full_name: string
  phone_number: string
  school_name: string
  category?: string
  student_count?: number
  note?: string
}

export interface CreateSchoolRequestResponse {
  id: string
  status: SchoolRequestStatus
  message: string
}

export interface SchoolRequestDto {
  id: string
  full_name: string
  phone_number: string
  school_name: string
  category?: string | null
  student_count?: number | null
  note?: string | null
  status: SchoolRequestStatus
  created_at: string
  updated_at: string
}

export interface ListSchoolRequestsResponse {
  items: SchoolRequestDto[]
}

export interface UpdateSchoolRequestStatusResponse {
  item: SchoolRequestDto
}

export interface CreateOrganizationPayload {
  name: string
  slug: string
  description?: string
  logo_url?: string
  category?: string
  contact_phone?: string
  telegram_url?: string
  is_demo: boolean
  owner: {
    full_name: string
    phone_number: string
    temporary_password: string
  }
}

export interface OrganizationDto {
  id: string
  name: string
  slug: string
  description?: string | null
  logo_url?: string | null
  category?: string | null
  contact_phone?: string | null
  telegram_url?: string | null
  public_status: 'draft' | 'public' | 'hidden'
  is_demo: boolean
}

export interface WorkspaceDto extends OrganizationDto {
  role: 'OWNER' | 'TEACHER' | 'MENTOR' | 'STUDENT'
}

export interface OrganizationOwnerDto {
  id: string
  full_name: string
  phone_number: string
  role: 'OWNER'
  must_change_password: boolean
}

export interface CreateOrganizationResponse {
  organization: OrganizationDto
  owner: OrganizationOwnerDto
}

export interface ListMyWorkspacesResponse {
  items: WorkspaceDto[]
}

export interface OwnerDashboardResponse {
  organization: { id: string; name: string; slug: string; role: string }
  metrics: {
    active_courses: number
    active_classes: number
    active_students: number
    new_leads: number
    pending_homework: number
    pending_access: number
    needs_revision: number
    completed_lessons: number
  }
  pending_homework: Array<{
    submission_id: string
    student_full_name: string
    class_name: string
    lesson_title: string
    homework_title: string
    submission_type: string
    submitted_at: string
  }>
  pending_access: Array<{
    class_id: string
    class_name: string
    student_user_id: string
    student_name: string
    phone_number: string
    access_status: string
    payment_status: string
    note?: string | null
  }>
  course_progress: Array<{
    course_id: string
    course_title: string
    class_count: number
    student_count: number
    lesson_count: number
    completion_count: number
    pending_homework: number
    progress_percent: number
  }>
  recent_activity: Array<{
    type: string
    title: string
    subtitle: string
    created_at: string
  }>
}

export interface UpdateOrganizationPayload {
  id: string
  name: string
  description?: string
  logo_url?: string
  category?: string
  contact_phone?: string
  telegram_url?: string
  public_status: 'draft' | 'public' | 'hidden'
  is_demo: boolean
}

export interface UpdateOrganizationResponse {
  item: OrganizationDto
}

export type LeadStatus = 'new' | 'contacted' | 'converted' | 'archived'

export interface PublicSchoolPageResponse {
  organization: OrganizationDto
  courses: unknown[]
  lead_form: {
    enabled: boolean
    required_fields: string[]
  }
}

export interface CreateLeadPayload {
  organization_id: string
  full_name: string
  phone_number: string
  note?: string
}

export interface CreateLeadResponse {
  id: string
  status: LeadStatus
  message: string
}

export interface LeadDto {
  id: string
  organization_id: string
  full_name: string
  phone_number: string
  note?: string | null
  source: string
  status: LeadStatus
  created_at: string
  updated_at: string
}

export interface ListLeadsResponse {
  items: LeadDto[]
}

export interface UpdateLeadStatusResponse {
  item: LeadDto
}

export interface DemoUserDto {
  phone_number: string
  password: string
  role: 'OWNER' | 'MENTOR' | 'STUDENT'
}

export interface DemoAccessResponse {
  organization_id: string
  school_slug: string
  course_id: string
  class_id: string
  lesson_id: string
  public_url: string
  owner_url: string
  student_url: string
  owner: DemoUserDto
  mentor: DemoUserDto
  student: DemoUserDto
  message: string
}
