import {
    apiGet,
    apiPost,
    apiPatch,
    apiPut,
    apiDelete
} from '../../../lib/apiWithAuth'
import type { User, UpdateUserRequest, SyncUserRequest } from '../types'

// Sync user with backend (create if not exists)
export function syncUser(data: SyncUserRequest) {
  return apiPost<User>('/users/sync', data)
}

// Get current user profile
export function getUserProfile() {
  return apiGet<User>('/users/profile')
}

// Update user profile
export function updateUserProfile(data: UpdateUserRequest) {
  return apiPut<User>('/users/profile', data)
}

// Delete user account
export function deleteUserAccount() {
  return apiDelete<{ message: string }>('/users/profile')
}