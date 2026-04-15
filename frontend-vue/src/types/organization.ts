export type SchoolRequestStatus = 'new' | 'contacted' | 'approved' | 'rejected'

export interface CreateSchoolRequestPayload {
  full_name: string
  phone_number: string
  school_name: string
  category?: string
  student_count?: number
  note?: string
}

export interface CreateSchoolRequestResponse {
  id: string
  status: SchoolRequestStatus
  message: string
}

export interface SchoolRequestDto {
  id: string
  full_name: string
  phone_number: string
  school_name: string
  category?: string | null
  student_count?: number | null
  note?: string | null
  status: SchoolRequestStatus
  created_at: string
  updated_at: string
}

export interface ListSchoolRequestsResponse {
  items: SchoolRequestDto[]
}

export interface UpdateSchoolRequestStatusResponse {
  item: SchoolRequestDto
}

export interface CreateOrganizationPayload {
  name: string
  slug: string
  description?: string
  logo_url?: string
  category?: string
  contact_phone?: string
  telegram_url?: string
  is_demo: boolean
  owner: {
    full_name: string
    phone_number: string
    temporary_password: string
  }
}

export interface OrganizationDto {
  id: string
  name: string
  slug: string
  description?: string | null
  logo_url?: string | null
  category?: string | null
  contact_phone?: string | null
  telegram_url?: string | null
  public_status: 'draft' | 'public' | 'hidden'
  is_demo: boolean
}

export interface WorkspaceDto extends OrganizationDto {
  role: 'OWNER' | 'TEACHER' | 'MENTOR' | 'STUDENT'
}

export interface OrganizationOwnerDto {
  id: string
  full_name: string
  phone_number: string
  role: 'OWNER'
  must_change_password: boolean
}

export interface CreateOrganizationResponse {
  organization: OrganizationDto
  owner: OrganizationOwnerDto
}

export interface ListMyWorkspacesResponse {
  items: WorkspaceDto[]
}

export interface UpdateOrganizationPayload {
  id: string
  name: string
  description?: string
  logo_url?: string
  category?: string
  contact_phone?: string
  telegram_url?: string
  public_status: 'draft' | 'public' | 'hidden'
  is_demo: boolean
}

export interface UpdateOrganizationResponse {
  item: OrganizationDto
}

export type LeadStatus = 'new' | 'contacted' | 'converted' | 'archived'

export interface PublicSchoolPageResponse {
  organization: OrganizationDto
  courses: unknown[]
  lead_form: {
    enabled: boolean
    required_fields: string[]
  }
}

export interface CreateLeadPayload {
  organization_id: string
  full_name: string
  phone_number: string
  note?: string
}

export interface CreateLeadResponse {
  id: string
  status: LeadStatus
  message: string
}

export interface LeadDto {
  id: string
  organization_id: string
  full_name: string
  phone_number: string
  note?: string | null
  source: string
  status: LeadStatus
  created_at: string
  updated_at: string
}

export interface ListLeadsResponse {
  items: LeadDto[]
}

export interface UpdateLeadStatusResponse {
  item: LeadDto
}
