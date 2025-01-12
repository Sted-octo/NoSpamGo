import { API_CONFIG } from './config'
import type { SpamDetectorResult } from '@/dataprovider/spamDetectorResult'
import type { Emails } from '@/dataprovider/emails'
import api from './api'

export class SpamDetector {
  static async Call(email: string): Promise<SpamDetectorResult[]> {
    try {
      const params = <Emails>{
        mails: [email],
      }
      const { data } = await api.post<SpamDetectorResult[]>(
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
