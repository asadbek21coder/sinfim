import { createRouter, createWebHistory } from 'vue-router'
import type { Role } from '@/types/auth'

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
    roles?: Role[]
    layout?: 'public' | 'auth' | 'app' | 'student'
    eyebrow?: string
    title?: string
    description?: string
    status?: string
  }
}

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: () => import('@/views/public/LandingView.vue'),
      meta: { layout: 'public' },
    },
    {
      path: '/enter',
      component: () => import('@/views/public/EntryPointView.vue'),
      meta: { layout: 'public' },
    },
    {
      path: '/apply-school',
      component: () => import('@/views/public/ApplySchoolView.vue'),
      meta: { layout: 'public' },
    },
    {
      path: '/demo',
      component: () => import('@/views/public/DemoView.vue'),
      meta: { layout: 'public' },
    },
    {
      path: '/auth/login',
      component: () => import('@/views/auth/LoginView.vue'),
      meta: { layout: 'auth' },
    },
    {
      path: '/auth/change-password',
      component: () => import('@/views/auth/ChangePasswordView.vue'),
      meta: { layout: 'auth', requiresAuth: true },
    },
    {
      path: '/login',
      redirect: '/auth/login',
    },
    {
      path: '/admin/school-requests',
      component: () => import('@/views/admin/SchoolRequestsView.vue'),
      meta: { layout: 'app', requiresAuth: true, roles: ['PLATFORM_ADMIN'] },
    },
    {
      path: '/admin/organizations/new',
      component: () => import('@/views/admin/CreateOrganizationView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        roles: ['PLATFORM_ADMIN'],
        eyebrow: 'Platform admin',
        title: 'Create organization',
        description: 'Superadmin creates a school workspace, assigns an owner, and fixes the school slug.',
        status: 'Step 3',
      },
    },
    {
      path: '/app/dashboard',
      component: () => import('@/views/DashboardView.vue'),
      meta: { layout: 'app', requiresAuth: true },
    },
    {
      path: '/app/courses',
      component: () => import('@/views/courses/CoursesView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'Catalog',
        title: 'Courses',
        description: 'Reusable course packages live here before they are attached to classes or groups.',
        status: 'Step 5',
      },
    },
    {
      path: '/app/courses/:courseId',
      component: () => import('@/views/courses/CourseDetailView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'Catalog',
        title: 'Course detail',
        description: 'Lessons, classes, materials, publish rules, and settings will be managed from this screen.',
        status: 'Step 5-7',
      },
    },
    {
      path: '/app/classes',
      component: () => import('@/views/classes/ClassesView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'Classroom',
        title: 'Classes and groups',
        description: 'Live cohorts, mentors, students, access status, and class schedules will be controlled here.',
        status: 'Step 6',
      },
    },
    {
      path: '/app/classes/:classId',
      component: () => import('@/views/classes/ClassDetailView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'Classroom',
        title: 'Class detail',
        description: 'Student enrollment, mentor assignment, manual payment/access confirmation, and progress tracking.',
        status: 'Step 6',
      },
    },
    {
      path: '/app/mentors',
      component: () => import('@/views/PlaceholderView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'People',
        title: 'Mentors',
        description: 'Mentors can be responsible for one or more classes and review homework submissions.',
        status: 'Step 6 and Step 10',
      },
    },
    {
      path: '/app/students',
      component: () => import('@/views/PlaceholderView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'People',
        title: 'Students',
        description: 'Admins and mentors can manually add students, attach them to groups, and manage access.',
        status: 'Step 6',
      },
    },
    {
      path: '/app/leads',
      component: () => import('@/views/leads/LeadsView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'Growth',
        title: 'Leads',
        description: 'Potential students from public school and course pages will be collected here.',
        status: 'Step 4',
      },
    },
    {
      path: '/app/lessons/:lessonId/edit',
      component: () => import('@/views/lessons/LessonEditorView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'Lesson editor',
        title: 'Edit lesson',
        description: 'Telegram stream references, materials, homework blocks, and quiz setup will be edited here.',
        status: 'Step 7 and Step 9',
      },
    },
    {
      path: '/app/homework/review',
      component: () => import('@/views/homework/HomeworkReviewView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'Homework',
        title: 'Review inbox',
        description: 'Mentors review written, file/photo, audio, and quiz submissions from their assigned classes.',
        status: 'Step 10',
      },
    },
    {
      path: '/app/settings/organization',
      component: () => import('@/views/settings/OrganizationSettingsView.vue'),
      meta: {
        layout: 'app',
        requiresAuth: true,
        eyebrow: 'Settings',
        title: 'Organization settings',
        description: 'School name, public profile, logo, slug, and contact details will be managed here.',
        status: 'Step 3',
      },
    },
    {
      path: '/learn/dashboard',
      component: () => import('@/views/learning/StudentDashboardView.vue'),
      meta: {
        layout: 'student',
        requiresAuth: true,
        eyebrow: 'Learning',
        title: 'Student dashboard',
        description: 'Students see current lessons, locked lessons, homework status, and progress here.',
        status: 'Step 8',
      },
    },
    {
      path: '/learn/lessons/:lessonId',
      component: () => import('@/views/learning/StudentLessonDetailView.vue'),
      meta: {
        layout: 'student',
        requiresAuth: true,
        eyebrow: 'Learning',
        title: 'Lesson detail',
        description: 'Video, PDF/materials, lesson completion, and homework submission will live on this page.',
        status: 'Step 8 and Step 9',
      },
    },
    {
      path: '/:schoolSlug/courses/:courseSlug',
      component: () => import('@/views/public/PublicCourseView.vue'),
      meta: {
        layout: 'public',
        eyebrow: 'Public course page',
        title: 'Course profile',
        description: 'This route will collect course-specific leads for a school.',
        status: 'Step 5',
      },
    },
    {
      path: '/:schoolSlug',
      component: () => import('@/views/public/PublicSchoolView.vue'),
      meta: {
        layout: 'public',
        eyebrow: 'Public school page',
        title: 'School profile',
        description: 'This route will become the public page for each school at sinfim.uz/{school-slug}.',
        status: 'Step 4',
      },
    },
    { path: '/:pathMatch(.*)*', redirect: '/' },
  ],
})

export default router
