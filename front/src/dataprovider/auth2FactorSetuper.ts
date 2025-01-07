import { API_CONFIG } from './config'
import type { TwoFactorSetupResponse } from '@/domain/auth'
import api from './api'

export class Auth2FactorSetuper {
  static async setup2FA(mail: string): Promise<TwoFactorSetupResponse> {
    try {
      const { data } = await api.post<TwoFactorSetupResponse>(
        API_CONFIG.BASE_URL + API_CONFIG.ENDPOINTS.SETUP_2FA,
        {
          mail,
        },
      )
      return data
    } catch (error) {
      console.error('Erreur lors de la configuration 2FA:', error)
      throw error
    }
  }
}
