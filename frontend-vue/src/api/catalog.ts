import client from './client'
import type {
  CoursePublicStatus,
  CourseResponse,
  CreateCoursePayload,
  CreateLessonPayload,
  LessonResponse,
  ListCoursesResponse,
  ListLessonsResponse,
  PublicCoursePageResponse,
  UpdateCoursePayload,
  UpdateLessonPayload,
} from '@/types/catalog'

export const catalogApi = {
  createCourse: (body: CreateCoursePayload) =>
    client.post<CourseResponse>('/catalog/create-course', body),

  updateCourse: (body: UpdateCoursePayload) =>
    client.post<CourseResponse>('/catalog/update-course', body),

  listCourses: (params: { organization_id: string; public_status?: CoursePublicStatus; limit?: number }) =>
    client.get<ListCoursesResponse>('/catalog/list-courses', { params }),

  getCourseDetail: (id: string) =>
    client.get<CourseResponse>('/catalog/get-course-detail', { params: { id } }),

  getPublicCoursePage: (schoolSlug: string, courseSlug: string) =>
    client.get<PublicCoursePageResponse>('/catalog/get-public-course-page', {
      params: { school_slug: schoolSlug, course_slug: courseSlug },
    }),

  createLesson: (body: CreateLessonPayload) =>
    client.post<LessonResponse>('/catalog/create-lesson', body),

  updateLesson: (body: UpdateLessonPayload) =>
    client.post<LessonResponse>('/catalog/update-lesson', body),

  listLessons: (params: { course_id: string; status?: string; limit?: number }) =>
    client.get<ListLessonsResponse>('/catalog/list-lessons', { params }),

  getLessonDetail: (id: string) =>
    client.get<LessonResponse>('/catalog/get-lesson-detail', { params: { id } }),
}
