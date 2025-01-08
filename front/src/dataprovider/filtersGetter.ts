import { API_CONFIG } from './config'
import api from './api'
import { type Filter } from '@/domain/filter'

export class FiltersGetter {
  static async get(email: string): Promise<Filter[]> {
    try {
      const { data } = await api.get<Filter[]>(
        API_CONFIG.BASE_URL + API_CONFIG.ENDPOINTS.FILTERS + `/${email}`,
      )
      return data
    } catch (error) {
      console.error('Erreur lors de la lecture des filtres:', error)
      throw error
    }
  }
}
