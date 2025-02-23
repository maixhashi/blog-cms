export type Task = {
  id: number
  title: string
  created_at: Date
  updated_at: Date
}
export type Feed = {
  id: number
  title: string
  url: string
  site_url: string
  description: string
  last_fetched_at: Date
  created_at: Date
  updated_at: Date
}
export type CsrfToken = {
  csrf_token: string
}
export type Credential = {
  email: string
  password: string
}