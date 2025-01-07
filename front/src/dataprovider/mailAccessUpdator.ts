import { API_CONFIG } from './config'
import type { MailConfig, MailConfigUpdateResponse } from '@/domain/mailConfig'

export class MailAccessUpdator {
  static async updateMailConfig(request: MailConfig): Promise<MailConfigUpdateResponse> {
    try {
      const user = {
        Mail: request.mail,
        ImapUsername: request.username,
        ImapPassword: request.password,
        ImapServerUrl: request.server,
        ImapServerPort: request.port,
      }
      const response = await fetch(
        `${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.UPDATE_MAIL_ACCESS}`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(user),
        },
      )

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      return await response.json()
    } catch (error) {
      console.error('Erreur lors de la mise à jour de la configuration de mail:', error)
      throw error
    }
  }
}
