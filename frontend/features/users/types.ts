export interface User {
  id: string
  auth_user_id: string
  email: string
  full_name?: string
  avatar_url?: string
  created_at: string
  updated_at: string
}

export interface UpdateUserRequest {
  full_name?: string
  avatar_url?: string
}

export interface SyncUserRequest {
  email: string
}