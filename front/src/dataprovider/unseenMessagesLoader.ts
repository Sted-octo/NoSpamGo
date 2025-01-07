import { API_CONFIG } from './config'
import api from './api'
import { type Message } from '@/domain/message'

export class UnseenMessageLoader {
  static async load(email: string): Promise<Message[]> {
    try {
      const { data } = await api.get<Message[]>(
        API_CONFIG.BASE_URL + API_CONFIG.ENDPOINTS.UNSEEN_MESSAGES_HEADERS + `/${email}`,
      )
      return data
    } catch (error) {
      console.error('Erreur lors de la lecture des messages non lus:', error)
      throw error
    }
  }
}
