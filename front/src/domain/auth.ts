export interface TwoFactorSetupResponse {
  secret: string
  qr_code: string
}

export interface TwoFactorVerifyRequest {
  mail: string
  token: string
}

export interface TwoFactorVerifyResponse {
  valid: boolean
  token: string
  ismailconfigok: boolean
}

export interface User {
  email: string
}
