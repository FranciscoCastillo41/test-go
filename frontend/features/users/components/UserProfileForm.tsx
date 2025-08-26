'use client'

import { useState, useEffect } from 'react'
import { useUserProfile, useUpdateUserProfile, useSyncUser } from '../hooks/useUsers'
import { useAuth } from '@/features/auth'

export default function UserProfileForm() {
  const { user: authUser } = useAuth()
  const { data: userProfile, isLoading, error } = useUserProfile()
  const updateProfile = useUpdateUserProfile()
  const syncUser = useSyncUser()

  const [fullName, setFullName] = useState('')
  const [avatarUrl, setAvatarUrl] = useState('')

  // Sync user on first load if they don't have a profile
  useEffect(() => {
    if (authUser?.email && error && !userProfile && !syncUser.isPending) {
      syncUser.mutate({ email: authUser.email })
    }
  }, [authUser, error, userProfile, syncUser])

  // Update form when profile loads
  useEffect(() => {
    if (userProfile) {
      setFullName(userProfile.full_name || '')
      setAvatarUrl(userProfile.avatar_url || '')
    }
  }, [userProfile])

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    
    updateProfile.mutate({
      full_name: fullName.trim() || undefined,
      avatar_url: avatarUrl.trim() || undefined,
    })
  }

  if (isLoading || syncUser.isPending) {
    return (
      <div className="flex justify-center items-center h-32">
        <div className="text-lg">Loading profile...</div>
      </div>
    )
  }

  return (
    <div className="max-w-md mx-auto bg-white p-6 rounded-lg shadow-sm border">
      <h2 className="text-xl font-bold mb-4">User Profile</h2>
      
      {userProfile && (
        <div className="mb-4 p-3 bg-gray-50 rounded text-sm">
          <div><strong>Email:</strong> {userProfile.email}</div>
          <div><strong>Account created:</strong> {new Date(userProfile.created_at).toLocaleDateString()}</div>
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="fullName" className="block text-sm font-medium text-gray-700">
            Full Name
          </label>
          <input
            id="fullName"
            type="text"
            value={fullName}
            onChange={(e) => setFullName(e.target.value)}
            placeholder="Enter your full name"
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
          />
        </div>

        <div>
          <label htmlFor="avatarUrl" className="block text-sm font-medium text-gray-700">
            Avatar URL
          </label>
          <input
            id="avatarUrl"
            type="url"
            value={avatarUrl}
            onChange={(e) => setAvatarUrl(e.target.value)}
            placeholder="https://example.com/your-avatar.jpg"
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
          />
        </div>

        <button
          type="submit"
          disabled={updateProfile.isPending}
          className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
        >
          {updateProfile.isPending ? 'Updating...' : 'Update Profile'}
        </button>

        {updateProfile.isSuccess && (
          <div className="text-green-600 text-sm text-center">
            Profile updated successfully!
          </div>
        )}

        {updateProfile.error && (
          <div className="text-red-600 text-sm text-center">
            Error: {updateProfile.error.message}
          </div>
        )}
      </form>
    </div>
  )
}