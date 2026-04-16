import client from './client'
import type { LearningLessonDetailResponse, MarkLessonCompletedResponse, StudentDashboardResponse } from '@/types/learning'

export const learningApi = {
  getStudentDashboard: (params?: { class_id?: string }) =>
    client.get<StudentDashboardResponse>('/learning/get-student-dashboard', { params }),

  getLessonDetail: (params: { class_id: string; lesson_id: string }) =>
    client.get<LearningLessonDetailResponse>('/learning/get-lesson-detail', { params }),

  markLessonCompleted: (body: { class_id: string; lesson_id: string }) =>
    client.post<MarkLessonCompletedResponse>('/learning/mark-lesson-completed', body),
}
