import client from './client'
import type {
  HomeworkDefinitionResponse,
  ListReviewSubmissionsResponse,
  ReviewSubmissionDetailResponse,
  ReviewSubmissionPayload,
  ReviewSubmissionResponse,
  SaveHomeworkDefinitionPayload,
  StudentHomeworkResponse,
  SubmitHomeworkPayload,
  SubmitHomeworkResponse,
} from '@/types/homework'

export const homeworkApi = {
  saveDefinition: (body: SaveHomeworkDefinitionPayload) =>
    client.post<HomeworkDefinitionResponse>('/homework/save-definition', body),

  getLessonHomework: (lessonId: string) =>
    client.get<HomeworkDefinitionResponse>('/homework/get-lesson-homework', { params: { lesson_id: lessonId } }),

  getStudentHomework: (params: { class_id: string; lesson_id: string }) =>
    client.get<StudentHomeworkResponse>('/homework/get-student-homework', { params }),

  submitHomework: (body: SubmitHomeworkPayload) =>
    client.post<SubmitHomeworkResponse>('/homework/submit-homework', body),

  listReviewSubmissions: (params?: { organization_id?: string; class_id?: string; status?: string; limit?: number }) =>
    client.get<ListReviewSubmissionsResponse>('/homework/list-review-submissions', { params }),

  getReviewSubmission: (id: string) =>
    client.get<ReviewSubmissionDetailResponse>('/homework/get-review-submission', { params: { id } }),

  reviewSubmission: (body: ReviewSubmissionPayload) =>
    client.post<ReviewSubmissionResponse>('/homework/review-submission', body),
}
