export interface MailConfig {
  mail: string
  username: string
  password: string
  server: string
  port: number
}

export interface MailConfigUpdateResponse {
  saved: boolean
}
