// 各スライスの型定義
export type TaskState = {
  editedTask: { id: number; title: string }
  updateEditedTask: (payload: { id: number; title: string }) => void
  resetEditedTask: () => void
}

export type ExternalAPIState = {
  editedExternalAPI: {
    id: number
    name: string
    base_url: string
    description: string
  }
  updateEditedExternalAPI: (payload: {
    id: number
    name: string
    base_url: string
    description: string
  }) => void
  resetEditedExternalAPI: () => void
}

export type FeedState = {
  editedFeed: {
    id: number
    title: string
    url: string
    site_url: string
    description: string
    last_fetched_at: Date | string
  }
  updateEditedFeed: (payload: {
    id: number
    title: string
    url: string
    site_url: string
    description: string
    last_fetched_at: Date | string
  }) => void
  resetEditedFeed: () => void
  selectedFeedId: number | null
  setSelectedFeedId: (id: number | null) => void
}

export type ArticleState = {
  editedArticle: {
    id: number
    title: string
    content: string
    published: boolean
    tags: string
  }
  updateEditedArticle: (payload: {
    id: number
    title: string
    content: string
    published: boolean
    tags: string
  }) => void
  resetEditedArticle: () => void
}

// 全体のアプリケーション状態の型
export type State = TaskState & ExternalAPIState & FeedState & ArticleState
