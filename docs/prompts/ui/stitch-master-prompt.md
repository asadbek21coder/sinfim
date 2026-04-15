# Stitch AI Master Prompt - Sinfim Web App UI

Use this prompt in Stitch AI to generate a polished web-based UI concept for Sinfim, an online school platform.

---

Design a modern, production-ready web application UI for Sinfim, a multi-tenant online school platform.

The product helps schools, education centers, individual teachers, and course-selling teams manage online cohort-based education. Each school has its own organization/brand space inside the platform. The platform is web-only for now.

Important UI format requirement:

Design this primarily as a desktop web SaaS dashboard and public website, not as a mobile app. Use wide desktop browser layouts as the primary canvas. Mobile responsiveness can be considered later, but do not generate phone-sized app screens as the main output.

## Product Concept

This is not a generic course marketplace. It is an operations platform for schools and teachers who currently sell courses manually and run education through Telegram channels/groups.

The platform should support:

- Public landing page for the platform
- Animated slow-moving logo strip of schools/brands using the platform
- A clear "How it works" guide section
- Entry point screen for different user intents
- Superadmin-controlled school creation
- School/brand workspace for owner/teacher/mentor
- Course content management
- Class/group cohort management
- Student access/payment confirmation, manually controlled
- Lead collection from public school/course pages
- Student learning dashboard
- Mentor homework review inbox
- Demo/fake school experience

Important product rule:

Real schools cannot self-create an account directly from public signup. A public visitor can submit a school request, but the real organization is created by platform superadmin. There should also be a demo school option so visitors can explore the product.

## Visual Direction

Create a trustworthy, clean, modern education SaaS web UI.

The UI should feel:

- Reliable
- Organized
- Modern
- Calm
- Educational
- Operationally efficient

Avoid a generic corporate landing page. The first screen should clearly communicate the actual product and lead into the web app experience.

Do not design native mobile app mockups. Do not use phone frames. Do not make iOS/Android style bottom tab bars. Use desktop web layouts: top navigation, sidebars, data tables, split panes, cards, forms, filters, and dashboard grids.

Use a balanced, professional color system. Avoid an overly purple/blue gradient look, avoid beige/brown themes, and avoid one-color monotony. Keep the interface accessible, spacious, and easy to scan.

Use rounded corners, but keep them modest. Buttons and cards should not feel overly pill-shaped.

## Main Information Architecture

Public routes:

- `/` - Platform landing
- `/enter` - Entry point
- `/apply-school` - School request form
- `/demo` - Demo school
- `/{schoolSlug}` - Public school page
- `/{schoolSlug}/courses/{courseSlug}` - Public course page
- `/{schoolSlug}/apply` - School/course lead form

Auth route:

- `/auth/login` - Phone number + password login

Superadmin routes:

- `/admin/organizations`
- `/admin/organizations/new`
- `/admin/organizations/{organizationId}`

Owner / Teacher / Mentor app routes:

- `/app/dashboard`
- `/app/settings/organization`
- `/app/courses`
- `/app/courses/{courseId}`
- `/app/classes/{classId}`
- `/app/mentors`
- `/app/students`
- `/app/leads`
- `/app/lessons/{lessonId}/edit`
- `/app/homework/review`

Student routes:

- `/learn/dashboard`
- `/learn/lessons/{lessonId}`

## Roles

Platform Admin:

- Creates organizations/schools
- Assigns owner user
- Can mark a school as demo/fake

Organization Owner:

- Manages school profile
- Manages courses, classes/groups, mentors, students, leads, access, lessons, materials, homework, and progress

Teacher:

- Creates and edits course content, lessons, materials, homework, and quizzes

Mentor:

- Can be assigned to one or more classes/groups
- Reviews homework submissions
- Can add students and update access if allowed

Student:

- Sees only enrolled classes and accessible lessons/materials
- Watches lessons, opens materials, submits homework, sees feedback

Lead / Potential Student:

- Not a real student yet
- Leaves name and phone number from public school/course page
- Can later be converted into a student by school admin/mentor

## Key Product Model

Use this mental model in the UI:

Course = reusable content package

- Lessons
- Videos
- PDFs/materials
- Homework tasks
- Quizzes/tests
- Public course page

Class/Group = live cohort operation

- Real students
- Mentors
- Access/payment status
- Start date
- Lesson release schedule
- Progress
- Homework submissions

Example:

- Course: Russian A1
- Classes/groups:
  - Russian A1 - May 2026
  - Russian A1 - VIP
  - Russian A1 - Weekend

This distinction should be obvious in the UI.

## Screens To Design

### 1. Platform Landing Page

Purpose:
Explain the platform and guide visitors to the right path.

Sections:

- Hero section with concise headline and CTA
- Slow-moving animated logo strip / logo marquee of schools/brands
- Product explanation: courses, classes, mentors, homework, tests, access, leads
- "How it works" guide:
  1. School is created by platform admin
  2. Owner creates courses
  3. Owner creates classes/groups
  4. Teacher adds lessons, videos, materials, homework
  5. Students get access
  6. Mentors review homework
  7. Owner tracks progress
- CTA buttons:
  - Enter platform
  - Try demo school
  - Request platform for my school

### 2. Entry Point Screen

Purpose:
Separate user intent.

Show 4 clear options:

- I want to enter my school
- I am a student
- Explore demo school
- I want this platform for my school

Each option should be a clean choice card with short explanatory copy.

### 3. School Request Form

Purpose:
Collect request from a school/teacher who wants to use the platform.

Fields:

- Full name
- Phone number
- School/brand name
- Category
- Approximate student/group count
- Short note

Show a friendly success state:

"Your request was received. Platform admin will contact you."

### 4. Superadmin Organization Create

Purpose:
Platform admin creates a school/organization and assigns the owner.

Fields:

- School/brand name
- Description
- Logo upload
- Slug preview: `sinfim.uz/{school-slug}`
- Owner full name
- Owner phone number
- Temporary password/code
- Demo/fake school toggle

Layout:
Use a practical admin form with a preview panel showing how the public school page will look.

### 5. Owner Dashboard

Purpose:
The school owner sees the whole school operation.

Content:

- School name, logo, public page link
- Summary cards:
  - Active courses
  - Active classes/groups
  - Students
  - Leads
  - Pending homework
  - Pending access confirmations
- Quick actions:
  - Create course
  - Create class/group
  - Add mentor
  - Add student
  - View leads
- Recent activity:
  - New lead
  - New homework submission
  - New student
  - Access status changed
- Pending work:
  - Unreviewed homework
  - Students waiting for access confirmation
  - Lessons scheduled soon
- Course/class progress list

### 6. Course Detail

Purpose:
Manage reusable course content.

Tabs:

- Overview
- Lessons
- Classes
- Settings

Overview:

- Course title
- Category
- Level
- Public course page link
- Short description
- Number of lessons
- Active classes using this course
- Quick actions: add lesson, create class, view public page

Lessons tab:

- Lesson list
- Each lesson row:
  - Order number
  - Title
  - Video status
  - Material status
  - Homework/test status
  - Draft/ready status
- Add lesson button
- Reorder affordance

Classes tab:

- List classes/groups using this course
- Group name
- Mentor
- Student count
- Start date
- Active/passive status

Settings:

- Course title
- Description
- Category
- Level
- Course slug
- Draft/active/archived status
- Public page active/draft

### 7. Class / Group Detail

Purpose:
Manage live cohort operation.

Tabs:

- Overview
- Students
- Homework
- Access

Overview:

- Group name, e.g. "Russian A1 - May 2026"
- Linked course
- Start date
- Lesson cadence: daily / every other day / 3 times a week / manual
- Assigned mentors
- Student count
- Progress summary
- Pending homework count
- Students waiting for access confirmation

Students tab:

- Student list
- Name
- Phone number
- Access status
- Progress
- Last activity
- Add student button
- Convert lead to student action
- Show or generate temporary password/code

Homework tab:

- Pending submissions
- Submission type: text / file-photo / quiz / audio
- Assigned mentor
- Status: pending / reviewed / needs revision
- Open review action

Access tab:

- Student-level access/payment status
- Manual access on/off
- Note field
- Bulk action: grant access to selected students

MVP rule:
Access is class-level. Lesson/material-level locking is handled by publish schedule, not payment status.

### 8. Lesson Editor

Purpose:
Create or edit one lesson inside a course.

Use one page divided into sections:

- Basic Info
- Video
- Materials
- Homework
- Quiz/Test
- Publish Rules

Basic Info:

- Lesson title
- Short description
- Order number
- Estimated duration
- Status: draft / ready / archived

Video:

- Telegram channel/message reference
- Stream reference
- Duration
- Preview/test video action
- Error state for invalid reference

Materials:

- Add PDF/material
- Material title
- Material type: PDF / image / link / other
- File size

Homework:

- Homework on/off toggle
- Homework type:
  - Written answer
  - File/photo
  - Oral/audio
- Instructions
- Due rule
- Mentor feedback required

Quiz/Test:

- Quiz on/off toggle
- Single choice / multiple choice questions
- Correct answer selection
- Automatic scoring

Publish Rules:

- Open by lesson order
- Open on day N after class start
- Manual publish option

### 9. Student Dashboard

Purpose:
Student sees what to do today.

Desktop web layout first, with a clean responsive web structure if needed. Do not create a phone-app dashboard.

Content:

- Greeting
- Active school/course/class
- Today's lesson card
- Pending homework cards
- Latest mentor feedback
- Course progress bar
- Next lesson unlock date

Lesson cards should show:

- Available / locked / completed
- Video exists
- Materials exist
- Homework exists
- Due date

### 10. Student Lesson Detail

Purpose:
Student watches one lesson and submits homework.

Mobile layout:

1. Lesson title and progress
2. Video player area
3. Materials list
4. Homework instructions
5. Submission form
6. Submission status and mentor feedback

Homework submission types:

- Written answer: text area
- File/photo: upload area
- Oral/audio: audio file upload for MVP
- Quiz/test: answer options

States:

- Locked
- No material
- No homework
- Submitted
- Reviewed
- Needs revision

### 11. Homework Review Inbox

Purpose:
Mentor reviews submissions efficiently.

Layout:

- Desktop: left filtered inbox list, right selected submission detail
- Mobile: list and detail as separate screens

Filters:

- Class/group
- Course
- Lesson
- Homework type
- Status
- Late submissions

Submission detail:

- Student info
- Lesson and homework instructions
- Student answer: text / file-photo / audio / quiz result
- Feedback form
- Score
- Status selector:
  - Approved
  - Needs revision
  - Rejected
  - Score only

Actions:

- Send feedback
- Give score
- Request revision
- Go to next submission

### 12. Organization Setup/Edit

Purpose:
Owner edits school/brand profile.

Sections:

- Brand info
- Public page
- Contact
- Danger zone

Fields:

- School/brand name
- Description
- Logo upload
- Category
- Slug preview
- Public page active/draft
- Phone number
- Telegram URL
- Support contact

Show public page preview.

### 13. Public School Page

Purpose:
A school-specific public page for students/leads.

Use route style:

`sinfim.uz/{school-slug}`

Content:

- School logo
- School description
- Course list
- Lead form
- Contact / Telegram link

### 14. Public Course Page

Purpose:
Show one public course and collect leads.

Content:

- Course title
- Description
- Level/category
- What student will learn
- Lead form prefilled with this course

### 15. Demo School

Purpose:
Let visitors explore without creating real data.

Demo mode should include:

- Owner dashboard demo
- Student view demo
- Clear demo banner:
  "Demo mode. Changes are not saved."

Do not let demo users damage or persist data.

## Authentication UX

Use one desktop web login screen:

- Phone number
- Password

No SMS OTP.
No Telegram login.

First login:

- User enters phone number + temporary password/code
- Then they are asked to set a new password
- Redirect by role:
  - Platform admin -> `/admin/organizations`
  - Owner -> `/app/dashboard`
  - Teacher -> `/app/courses`
  - Mentor -> `/app/homework/review`
  - Student -> `/learn/dashboard`

Forgot password:

For MVP, show message:

"Please contact your school admin or mentor."

## Navigation

Owner/Teacher/Mentor app sidebar:

- Dashboard
- Courses
- Classes / Groups
- Mentors
- Students
- Leads
- Homework Review
- Organization Settings

Student navigation:

- Dashboard
- My Lessons
- Homework
- Feedback

Superadmin navigation:

- Organizations
- School Requests
- Demo School

## UI States To Include

Every major screen should include:

- Loading state
- Empty state
- Error state
- Access locked state where relevant
- Draft/active/archived status where relevant
- Pending/reviewed/needs revision status for homework

## Copy Style

Use short product copy. Avoid marketing fluff.

Use a mix of English UI labels and Uzbek/Turkish-friendly wording where natural, but keep the UI clean and understandable.

Examples:

- "Online school operations in one place"
- "Create course"
- "Create class"
- "Add student"
- "Pending homework"
- "Access waiting"
- "Arizangiz qabul qilindi"
- "Demo mode. Changes are not saved."

## Output Expectations

Generate a coherent web app UI concept with:

- Public landing page
- Entry point screen
- School request form
- Superadmin organization creation page
- Owner dashboard
- Course detail page
- Class/group detail page
- Lesson editor page
- Student dashboard
- Student lesson detail
- Homework review inbox
- Organization settings page
- Public school/course page
- Demo school page
- Login and first-password screen

Prioritize clarity, usability, and consistency over decorative complexity.

The result should feel like a real desktop-first SaaS product ready to be implemented as a Vue 3 + TypeScript web application.
