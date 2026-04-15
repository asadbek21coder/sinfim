import client from './client'
import type { AuthResponse, UserDto } from '@/types/auth'

export const authApi = {
  login: (body: { phoneNumber: string; password: string }) =>
    client.post<AuthResponse>('/auth/admin-login', { phone_number: body.phoneNumber, password: body.password }),

  refresh: (refreshToken: string) =>
    client.post<AuthResponse>('/auth/refresh-token', { refreshToken }),

  logout: () =>
    client.post<Record<string, never>>('/auth/logout'),

  me: () =>
    client.get<UserDto>('/auth/me'),

  changeMyPassword: (body: { currentPassword: string; newPassword: string }) =>
    client.post<Record<string, never>>('/auth/change-my-password', {
      current_password: body.currentPassword,
      new_password: body.newPassword,
    }),
}
