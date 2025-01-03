import { API_CONFIG } from './config'
import type { TwoFactorSetupResponse } from '@/domain/auth'

export class Auth2FactorSetuper {
  static async setup2FA(mail: string): Promise<TwoFactorSetupResponse> {
    try {
      const response = await fetch(`${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.SETUP_2FA}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ mail }),
      })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      return await response.json()
    } catch (error) {
      console.error('Erreur lors de la configuration 2FA:', error)
      throw error
    }
  }
}
