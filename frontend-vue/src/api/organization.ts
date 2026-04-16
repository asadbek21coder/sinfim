import client from './client'
import type {
  CreateSchoolRequestPayload,
  CreateSchoolRequestResponse,
  CreateOrganizationPayload,
  CreateOrganizationResponse,
  DemoAccessResponse,
  CreateLeadPayload,
  CreateLeadResponse,
  ListMyWorkspacesResponse,
  OwnerDashboardResponse,
  ListLeadsResponse,
  ListSchoolRequestsResponse,
  LeadStatus,
  PublicSchoolPageResponse,
  SchoolRequestStatus,
  UpdateLeadStatusResponse,
  UpdateOrganizationPayload,
  UpdateOrganizationResponse,
  UpdateSchoolRequestStatusResponse,
} from '@/types/organization'

export const organizationApi = {
  createSchoolRequest: (body: CreateSchoolRequestPayload) =>
    client.post<CreateSchoolRequestResponse>('/organization/create-school-request', body),

  getDemoAccess: () =>
    client.get<DemoAccessResponse>('/organization/get-demo-access'),

  listSchoolRequests: (params?: { status?: SchoolRequestStatus; limit?: number }) =>
    client.get<ListSchoolRequestsResponse>('/organization/list-school-requests', { params }),

  updateSchoolRequestStatus: (body: { id: string; status: SchoolRequestStatus }) =>
    client.post<UpdateSchoolRequestStatusResponse>('/organization/update-school-request-status', body),

  createOrganization: (body: CreateOrganizationPayload) =>
    client.post<CreateOrganizationResponse>('/organization/create-organization', body),

  listMyWorkspaces: () =>
    client.get<ListMyWorkspacesResponse>('/organization/list-my-workspaces'),

  getOwnerDashboard: (params?: { organization_id?: string }) =>
    client.get<OwnerDashboardResponse>('/organization/get-owner-dashboard', { params }),

  updateOrganization: (body: UpdateOrganizationPayload) =>
    client.post<UpdateOrganizationResponse>('/organization/update-organization', body),

  getPublicSchoolPage: (slug: string) =>
    client.get<PublicSchoolPageResponse>('/organization/get-public-school-page', { params: { slug } }),

  createLead: (body: CreateLeadPayload) =>
    client.post<CreateLeadResponse>('/organization/create-lead', body),

  listLeads: (params: { organization_id: string; status?: LeadStatus; limit?: number }) =>
    client.get<ListLeadsResponse>('/organization/list-leads', { params }),

  updateLeadStatus: (body: { id: string; status: LeadStatus }) =>
    client.post<UpdateLeadStatusResponse>('/organization/update-lead-status', body),
}
