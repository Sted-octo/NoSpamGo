import { API_CONFIG } from './config';
import type {TwoFactorVerifyRequest, TwoFactorVerifyResponse } from '@/domain/auth';

export class Auth2FactorVerificator {
    static async verify2FA(request: TwoFactorVerifyRequest): Promise<TwoFactorVerifyResponse> {
    try {
      const response = await fetch(`${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.VERIFY_2FA}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(request),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      return await response.json();
    } catch (error) {
      console.error('Erreur lors de la vérification 2FA:', error);
      throw error;
    }
  }
}