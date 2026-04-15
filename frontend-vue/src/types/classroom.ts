export type ClassStatus = 'active' | 'paused' | 'archived'
export type LessonCadence = 'daily' | 'every_other_day' | 'weekly_3' | 'manual'
export type AccessStatus = 'pending' | 'active' | 'paused' | 'blocked'
export type PaymentStatus = 'unknown' | 'pending' | 'confirmed' | 'rejected'

export interface ClassDto {
  id: string
  organization_id: string
  course_id: string
  name: string
  start_date?: string | null
  lesson_cadence: LessonCadence
  status: ClassStatus
  created_at: string
  updated_at: string
}

export interface ClassSummaryDto extends ClassDto {
  course_title: string
  mentor_count: number
  student_count: number
}

export interface StudentDto {
  enrollment_id: string
  student_user_id: string
  full_name: string
  phone_number: string
  status: string
  access_status: AccessStatus
  payment_status: PaymentStatus
  note?: string | null
  enrolled_at: string
  granted_at?: string | null
}

export interface MentorDto {
  id: string
  user_id: string
  full_name: string
  phone_number: string
}

export interface CreateClassPayload {
  organization_id: string
  course_id: string
  name: string
  start_date?: string
  lesson_cadence?: LessonCadence
  mentor_user_ids?: string[]
}

export interface AddStudentPayload {
  class_id: string
  full_name: string
  phone_number: string
  temporary_password?: string
  access_status?: AccessStatus
  payment_status?: PaymentStatus
  note?: string
}

export interface UpdateAccessPayload {
  class_id: string
  student_user_id: string
  access_status: AccessStatus
  payment_status: PaymentStatus
  note?: string
}

export interface ClassSummaryResponse { item: ClassSummaryDto }
export interface ListClassesResponse { items: ClassSummaryDto[] }
export interface ClassDetailResponse { item: ClassDto; mentors: MentorDto[]; students: StudentDto[] }
export interface AddStudentResponse { student: { id: string; full_name: string; phone_number: string; must_change_password: boolean }; enrollment: unknown; access: unknown; temporary_password_generated: boolean }
export interface UpdateAccessResponse { item: unknown }
