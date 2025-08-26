'use client'

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { syncUser, getUserProfile, updateUserProfile, deleteUserAccount } from '../api'
import type { UpdateUserRequest, SyncUserRequest } from '../types'

export function useUserProfile() {
  return useQuery({
    queryKey: ['user-profile'],
    queryFn: getUserProfile,
    retry: false, // Don't retry if user doesn't exist yet
  })
}

export function useSyncUser() {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: (data: SyncUserRequest) => syncUser(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['user-profile'] })
    },
  })
}

export function useUpdateUserProfile() {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: (data: UpdateUserRequest) => updateUserProfile(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['user-profile'] })
    },
  })
}

export function useDeleteUserAccount() {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: deleteUserAccount,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['user-profile'] })
      // Could also redirect to login or show success message
    },
  })
}