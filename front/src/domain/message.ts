import { type MailAddress } from '@/domain/mailAddress'

export interface Message {
  Subject: string
  Id: number
  Mails: MailAddress[]
}

export interface UnseenMessagesResponse {
  valid: boolean
  token: string
  ismailconfigok: boolean
}
