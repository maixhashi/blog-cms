import { UserState } from './types/userTypes';
// 衝突している型のインポートを削除
// import { GoogleBookState } from './types/googleBookTypes';
// import { BookState } from './types/bookTypes';

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

export type LayoutState = {
  editedLayout: { id: number; title: string }
  updateEditedLayout: (payload: { id: number; title: string }) => void
  resetEditedLayout: () => void
}

export type LayoutComponentState = {
  editedLayoutComponent: { id: number; name: string; type: string; content: string }
  updateEditedLayoutComponent: (payload: { id: number; name: string; type: string; content: string }) => void
  resetEditedLayoutComponent: () => void
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

// 書籍関連の型定義
export type BookState = {
  editedBook: { 
    id: number; 
    title: string;
    author: string;
    isbn: string;
    description: string;
    thumbnail_url: string;
    page_count: number;
  }
  updateEditedBook: (payload: { 
    id: number; 
    title: string;
    author: string;
    isbn: string;
    description: string;
    thumbnail_url: string;
    page_count: number;
  }) => void
  resetEditedBook: () => void
}

// Google Books関連の型定義
export type GoogleBookState = {
  searchQuery: string;
  searchResults: any[];
  selectedBook: any | null;
  updateSearchQuery: (query: string) => void;
  setSearchResults: (results: any[]) => void;
  selectBook: (book: any | null) => void;
  resetGoogleBookState: () => void;
}

// 全体のアプリケーション状態の型を更新
export type State = UserState & TaskState & ExternalAPIState & FeedState & ArticleState & LayoutState & LayoutComponentState & BookState & GoogleBookState