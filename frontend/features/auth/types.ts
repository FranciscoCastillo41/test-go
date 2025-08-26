export interface AuthUser {
  id: string
  email: string
  created_at: string
}

export interface SignUpData {
  email: string
  password: string
}

export interface SignInData {
  email: string
  password: string
}