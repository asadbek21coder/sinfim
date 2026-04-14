# Stitch AI Prompt - Desktop First Sinfim Website / Web App

Use this prompt in a fresh Stitch AI project.

---

Design a desktop-first website and web-based SaaS application UI for Sinfim, a multi-tenant online school platform.

Very important:

- This is a website and browser-based web application.
- The primary design target is desktop/laptop web, opened in Chrome or a similar browser.
- Use wide desktop web artboards, ideally 1440px wide.
- Responsive design is important, but desktop web is the priority.
- Do not start with mobile screens.
- Do not generate phone-sized app screens as the main output.
- Do not use mobile app proportions.
- Do not use iOS/Android style bottom navigation.
- Do not use phone frames.

The UI should look like a real desktop SaaS web product with:

- Full-width website landing pages
- Desktop top navigation
- Desktop sidebar navigation for dashboards
- Wide content areas
- Data tables
- Filter bars
- Tabs
- Split panes
- Multi-column dashboard grids
- Form sections
- Status badges
- Progress bars
- Loading, empty, error, and locked states

## Product

Sinfim is a multi-tenant online school platform for schools, education centers, individual teachers, and course-selling teams.

These users currently sell courses manually and run education through Telegram channels/groups. This platform gives them a structured web workspace for courses, classes/groups, videos, materials, homework, tests, mentor review, student access, leads, and progress tracking.

This is not a generic marketplace. It is an operations platform for schools and course teams.

## Core Product Rules

- Each school has its own organization/brand workspace.
- Real schools cannot create themselves through public signup.
- School creation is controlled by platform superadmin.
- Public visitors can submit a school request form.
- There is a demo/fake school experience for visitors.
- Login uses phone number + password.
- No SMS OTP in MVP.
- No Telegram login in MVP.
- Video is represented by Telegram stream references.
- Payment is outside the product in MVP, but access/payment status is manually confirmed inside the platform.
- Access is class-level in MVP.

## Key Domain Model

Make this distinction clear in the UI:

Course = reusable content package

- Lessons
- Video references
- PDF/materials
- Homework tasks
- Quiz/test
- Public course page

Class / Group = live cohort operation

- Real students
- Assigned mentors
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

## Visual Style

Create a trustworthy, organized, modern education SaaS website.

The look should be:

- Clean
- Professional
- Calm
- Operational
- Easy to scan
- Suitable for school owners and education teams

Avoid:

- Mobile app UI
- Purple/blue gradient-heavy startup look
- Beige/brown themes
- Overly playful visuals
- One-color monotony
- Native app bottom navigation

Use a balanced color palette, clear typography, modest border radius, accessible contrast, and structured spacing.

## Screens To Generate

Generate desktop web screens for the following. Use 1440px wide desktop artboards where possible.

### 1. Platform Landing Page

This is a full-width website page.

Content:

- Top navigation
- Hero section explaining the product
- CTA buttons:
  - Enter platform
  - Try demo school
  - Request platform for my school
- Slow-moving school/brand logo marquee
- "How it works" guide:
  1. School is created by platform admin
  2. Owner creates courses
  3. Owner creates classes/groups
  4. Teacher adds lessons, videos, materials, homework
  5. Students get access
  6. Mentors review homework
  7. Owner tracks progress
- Product benefit sections
- Demo CTA
- School request CTA

### 2. Entry Point Page

Desktop web page with four large choice cards:

- I want to enter my school
- I am a student
- Explore demo school
- I want this platform for my school

### 3. School Request Page

Desktop web form for schools that want to use the platform.

Fields:

- Full name
- Phone number
- School/brand name
- Category
- Approximate student/group count
- Short note

Show a success state:

"Your request was received. Platform admin will contact you."

### 4. Superadmin Organization Create

Desktop admin form.

Layout:

- Left: form fields
- Right: public school page preview panel

Fields:

- School/brand name
- Description
- Logo upload
- Slug preview: `sinfim.uz/{school-slug}`
- Owner full name
- Owner phone number
- Temporary password/code
- Demo/fake school toggle

### 5. Owner Dashboard

This must be a wide desktop SaaS dashboard.

Layout:

- Left sidebar navigation, about 260px wide
- Topbar across the main content
- Main content area using full desktop width

Content:

- School logo/name and public URL
- Metric cards in a horizontal row:
  - Active Courses
  - Active Classes
  - Students
  - Leads
  - Pending Homework
  - Access Waiting
- Quick actions:
  - Create Course
  - Create Class
  - Add Mentor
  - Add Student
  - View Leads
- Two-column section:
  - Recent Activity
  - Pending Work
- Wide table:
  - Course / Class Progress
  - Columns: Course, Class, Mentor, Students, Progress, Pending Homework, Access Waiting, Status

### 6. Course Detail

Desktop web page with sidebar and topbar.

Use tabs:

- Overview
- Lessons
- Classes
- Settings

Lessons should be a desktop data table, not mobile cards.

Overview:

- Course title
- Category
- Level
- Public course page link
- Short description
- Number of lessons
- Active classes using this course
- Quick actions: Add lesson, Create class, View public page

Lessons table:

- Order
- Lesson title
- Video status
- Material status
- Homework/test status
- Draft/ready status
- Actions

Classes table:

- Group name
- Mentor
- Student count
- Start date
- Status

### 7. Class / Group Detail

Desktop web page with tabs:

- Overview
- Students
- Homework
- Access

Overview:

- Group name
- Linked course
- Start date
- Lesson cadence
- Assigned mentors
- Student count
- Progress summary
- Pending homework count
- Access waiting count

Students tab:

- Desktop table with name, phone number, access status, progress, last activity, actions

Homework tab:

- Desktop table with student, lesson, submission type, status, mentor, submitted time, action

Access tab:

- Student-level access/payment table
- Bulk action to grant access

### 8. Lesson Editor

Desktop editor page.

Use a wide form/editor layout with sections:

- Basic Info
- Video
- Materials
- Homework
- Quiz/Test
- Publish Rules

Video section:

- Telegram channel/message reference
- Stream reference
- Duration
- Preview/test video action

Homework section:

- Homework on/off
- Written answer
- File/photo
- Oral/audio
- Instructions
- Due rule

Quiz/Test:

- Single choice / multiple choice
- Correct answer selection
- Automatic scoring

### 9. Homework Review Inbox

Desktop split-pane layout.

Left panel:

- Filter bar
- Submission inbox list
- Filters: class/group, course, lesson, homework type, status, late submissions

Right panel:

- Selected submission detail
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
- Actions:
  - Send feedback
  - Request revision
  - Next submission

### 10. Student Dashboard

This is still a web portal page, not a phone app.

Desktop web layout:

- Topbar
- Main content area
- Today's lesson
- Pending homework
- Latest feedback
- Course progress
- Next lesson unlock date

It can be simpler and more student-friendly than the admin dashboard, but still make it a browser-based web page.

### 11. Student Lesson Detail

Desktop web page:

- Video player area
- Materials panel
- Homework instructions
- Submission form
- Submission status
- Mentor feedback

For desktop, video can be the main left area and homework/materials can be in a right panel.

### 12. Organization Settings

Desktop form page:

- Brand info
- Logo
- Description
- Public page slug
- Contact phone
- Telegram URL
- Public page preview

### 13. Public School Page

Full-width website page:

- School logo
- School description
- Course list
- Lead form
- Contact / Telegram link

Route concept:

`sinfim.uz/{school-slug}`

### 14. Public Course Page

Full-width website page:

- Course title
- Description
- Level/category
- What students will learn
- Lead form prefilled with this course

### 15. Demo School

Desktop demo page:

- Demo mode banner:
  "Demo mode. Changes are not saved."
- Owner dashboard demo
- Student view demo
- No persistent destructive actions

### 16. Login and First Password

Desktop web login page:

- Phone number
- Password

First-login flow:

- User enters phone number + temporary password/code
- Then sets a new password
- Redirects by role:
  - Platform admin -> organizations
  - Owner -> dashboard
  - Teacher -> courses
  - Mentor -> homework review
  - Student -> learning dashboard

Forgot password MVP message:

"Please contact your school admin or mentor."

## Navigation

Owner/Teacher/Mentor sidebar:

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

## Required UI States

Include examples of:

- Loading
- Empty
- Error
- Access locked
- Draft / active / archived
- Pending / reviewed / needs revision
- Access waiting

## Final Instruction

Again: prioritize desktop website and desktop web app layouts.

Generate wide 1440px desktop web screens. Responsive behavior can be implied, but do not make mobile the primary design.

The final result should look like a real web SaaS application ready to implement in Vue 3 + TypeScript.
