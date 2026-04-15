import client from './client'
import type {
  AddStudentPayload,
  AddStudentResponse,
  ClassDetailResponse,
  ClassSummaryResponse,
  CreateClassPayload,
  ListClassesResponse,
  UpdateAccessPayload,
  UpdateAccessResponse,
} from '@/types/classroom'

export const classroomApi = {
  createClass: (body: CreateClassPayload) =>
    client.post<ClassSummaryResponse>('/classroom/create-class', body),

  listClasses: (params: { organization_id: string; course_id?: string; limit?: number }) =>
    client.get<ListClassesResponse>('/classroom/list-classes', { params }),

  getClassDetail: (id: string) =>
    client.get<ClassDetailResponse>('/classroom/get-class-detail', { params: { id } }),

  addStudent: (body: AddStudentPayload) =>
    client.post<AddStudentResponse>('/classroom/add-student', body),

  updateAccess: (body: UpdateAccessPayload) =>
    client.post<UpdateAccessResponse>('/classroom/update-access', body),
}
