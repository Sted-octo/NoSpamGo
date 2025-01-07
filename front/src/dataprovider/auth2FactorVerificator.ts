import { API_CONFIG } from './config'
import type { TwoFactorVerifyRequest, TwoFactorVerifyResponse } from '@/domain/auth'
import api from './api'

export class Auth2FactorVerificator {
  static async verify2FA(request: TwoFactorVerifyRequest): Promise<TwoFactorVerifyResponse> {
    try {
      const { data } = await api.post<TwoFactorVerifyResponse>(
        API_CONFIG.BASE_URL + API_CONFIG.ENDPOINTS.VERIFY_2FA,
        request,
      )
      return data
    } catch (error) {
      console.error('Erreur lors de la vérification 2FA:', error)
      throw error
    }
  }
}
