import { type Filter } from '@/domain/Filter'

export interface UserFilters {
  Mail: string
  Filters: Filter[]
}
