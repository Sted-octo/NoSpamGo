import { API_CONFIG } from './config'
import type { MailConfig, MailConfigUpdateResponse } from '@/domain/mailConfig'
import type { User } from '@/dataprovider/User.ts'
import api from './api'

export class MailAccessUpdator {
  static async updateMailConfig(request: MailConfig): Promise<MailConfigUpdateResponse> {
    try {
      const user = <User>{
        Mail: request.mail,
        ImapUsername: request.username,
        ImapPassword: request.password,
        ImapServerUrl: request.server,
        ImapServerPort: request.port,
      }

      const { data } = await api.post<MailConfigUpdateResponse>(
        API_CONFIG.BASE_URL + API_CONFIG.ENDPOINTS.UPDATE_MAIL_ACCESS,
        user,
      )
      return data
    } catch (error) {
      console.error('Erreur lors de la mise à jour de la configuration de mail:', error)
      throw error
    }
  }
}
