<template>
  <div class="flex min-h-screen bg-surface text-on-surface">
    <aside class="hidden w-[260px] shrink-0 flex-col bg-sidebar-bg text-white lg:flex">
      <RouterLink to="/app/dashboard" class="flex h-[72px] items-center gap-3 px-6">
        <span class="flex h-9 w-9 items-center justify-center rounded-lg bg-secondary font-display text-sm font-bold">S</span>
        <div>
          <p class="font-display text-base font-bold">Sinfim.uz</p>
          <p class="text-xs text-white/55">School workspace</p>
        </div>
      </RouterLink>

      <nav class="flex-1 space-y-1 overflow-y-auto px-3 py-4">
        <SidebarItem icon="D" label="Dashboard" to="/app/dashboard" />
        <SidebarItem icon="C" label="Courses" to="/app/courses" />
        <SidebarItem icon="G" label="Classes" to="/app/classes" />
        <SidebarItem icon="M" label="Mentors" to="/app/mentors" />
        <SidebarItem icon="S" label="Students" to="/app/students" />
        <SidebarItem icon="L" label="Leads" to="/app/leads" />
        <SidebarItem icon="R" label="School requests" to="/admin/school-requests" />
        <SidebarItem icon="H" label="Homework" to="/app/homework/review" />
        <SidebarItem icon="O" label="Organization" to="/app/settings/organization" />
      </nav>

      <div class="m-3 rounded-lg border border-white/10 bg-white/10 px-4 py-3">
        <p class="text-xs uppercase text-white/45">Active workspace</p>
        <p class="mt-1 truncate text-sm font-semibold">{{ activeWorkspace?.name ?? 'No workspace yet' }}</p>
        <p class="mt-0.5 text-xs text-white/55">{{ activeWorkspace?.role ?? 'Create or assign one' }}</p>
      </div>
    </aside>

    <div class="flex min-w-0 flex-1 flex-col">
      <header class="flex h-[72px] shrink-0 items-center justify-between border-b border-outline-variant/70 bg-surface-container-lowest px-5 lg:px-8">
        <div>
          <p class="text-xs font-semibold uppercase tracking-wide text-secondary">Sinfim.uz</p>
          <p class="text-sm text-on-surface-variant">{{ activeWorkspace?.name ?? 'Owner, teacher, mentor and admin operations' }}</p>
        </div>
        <div class="flex items-center gap-3">
          <RouterLink class="btn-secondary hidden sm:inline-flex" to="/learn/dashboard">Student view</RouterLink>
          <button class="btn-primary" type="button" @click="logout">Logout</button>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-5 lg:p-8">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import SidebarItem from '@/components/ui/SidebarItem.vue'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { organizationApi } from '@/api/organization'
import { useAuthStore } from '@/stores/auth'
import type { WorkspaceDto } from '@/types/organization'

const router = useRouter()
const auth = useAuthStore()
const activeWorkspace = ref<WorkspaceDto | null>(null)

async function loadWorkspace() {
  const response = await organizationApi.listMyWorkspaces().catch(() => null)
  activeWorkspace.value = response?.data.items[0] ?? null
}

async function logout() {
  await auth.logout()
  router.push('/auth/login')
}

onMounted(loadWorkspace)
</script>
