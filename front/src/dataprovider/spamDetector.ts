import { API_CONFIG } from './config'
import type { SavedResponse } from '@/domain/SavedResponse'
import type { Email } from '@/dataprovider/email'
import api from './api'

export class SpamDetector {
  static async Call(email: string): Promise<SavedResponse> {
    try {
      const params = <Email>{
        mail: email,
      }
      const { data } = await api.post<SavedResponse>(
        API_CONFIG.BASE_URL + API_CONFIG.ENDPOINTS.SPAM_DETECTOR,
        params,
      )
      return data
    } catch (error) {
      console.error('Erreur lors de spam detector:', error)
      throw error
    }
  }
}
