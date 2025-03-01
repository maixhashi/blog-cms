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
  last_fetched_at: Date | string
  created_at?: string
  updated_at?: string
}
export type CsrfToken = {
  csrf_token: string
}
export type Credential = {
  email: string
  password: string
}

export type ExternalAPI = {
  id: number
  name: string
  base_url: string
  description: string
  created_at: string
  updated_at: string
}

export type QiitaTag = {
  name: string
}

export type QiitaUser = {
  id: string
  profile_image_url: string
  name: string
}

export type QiitaArticle = {
  id: string
  title: string
  url: string
  likes_count: number
  tags: QiitaTag[]
  created_at: string
  user: QiitaUser
}
