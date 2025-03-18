// モデル関連の型をすべてエクスポート
export * from './models';

// API関連の型をエクスポート（必要に応じて）
export * from './api/generated';


// 他の型定義...
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

export type Article = {
  id: number
  title: string
  content: string
  published: boolean
  tags: string
  created_at: Date | string
  updated_at: Date | string
}

export interface HatenaArticle {
  id: string;
  title: string;
  url: string;
  summary: string;
  categories: string[];
  published_at: string;
  author: string;
}

export type FeedArticle = {
  id: number
  title: string
  content: string
  summary?: string
  url: string
  published_at: string | Date
  feed_id: number
  feed_title?: string
  author?: string
  thumbnail_url?: string
  tags?: string[]
  likes_count?: number
  created_at?: string | Date
  updated_at?: string | Date
}

export type Layout = {
  id: number
  title: string
  created_at: string
  updated_at: string
}

export type LayoutComponent = {
  id: number
  name: string
  type: string
  content: string
  created_at: string
  updated_at: string
}
