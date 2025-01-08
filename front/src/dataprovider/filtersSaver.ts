import { API_CONFIG } from './config'
import type { Filter } from '@/domain/filter'
import type { SavedResponse } from '@/domain/SavedResponse'
import type { UserFilters } from '@/dataprovider/UserFilters'
import api from './api'

export class FiltesSaver {
  static async save(filters: Filter[], email: string): Promise<SavedResponse> {
    try {
      const userFilters = <UserFilters>{
        Mail: email,
        Filters: filters,
      }
      const { data } = await api.post<SavedResponse>(
        API_CONFIG.BASE_URL + API_CONFIG.ENDPOINTS.FILTERS_SAVER,
        userFilters,
      )
      return data
    } catch (error) {
      console.error('Erreur lors de la sauvegarde des filtres:', error)
      throw error
    }
  }
}
