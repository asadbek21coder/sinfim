export type Role = 'PLATFORM_ADMIN' | 'OWNER' | 'TEACHER' | 'MENTOR' | 'STUDENT'

export interface UserDto {
  id: string
  phoneNumber: string
  fullName: string
  role: Role
  organizationId?: string | null
  isActive: boolean
  mustChangePassword?: boolean
  createdAt: string
}

export interface AuthResponse {
  accessToken: string
  refreshToken: string
  tokenType: string
  expiresIn: number
  user: UserDto
}
