export interface HealthResult {
  ok: boolean
  status?: string
  error?: string
}

export async function checkBackendHealth(): Promise<HealthResult> {
  try {
    const response = await fetch('/health')
    if (!response.ok) {
      return { ok: false, error: `HTTP ${response.status}` }
    }

    const data = await response.json() as { status?: string }
    return { ok: data.status === 'ok', status: data.status }
  } catch (error) {
    return {
      ok: false,
      error: error instanceof Error ? error.message : 'Backend is not reachable',
    }
  }
}
