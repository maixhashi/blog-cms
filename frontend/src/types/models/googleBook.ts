/**
 * Google Books APIから取得する書籍の型定義
 */

// 業界識別子（ISBN等）の型定義
export interface IndustryIdentifier {
  type: string;
  identifier: string;
}

// バックエンドのモデルに合わせた型定義
export type GoogleBookVolume = {
  id: string;
  title: string;
  authors?: string[];
  description?: string;
  isbn?: string;
  image_url?: string;
  published_date?: string;
};

// バックエンドのレスポンス型に合わせる
export type GoogleBookSearchResponse = {
  items: GoogleBookVolume[];
  total_items: number; // または totalItems
};

export type GoogleBookState = {
  searchQuery: string;
  searchResults: GoogleBookVolume[];
  selectedBook: GoogleBookVolume | null;
  updateSearchQuery: (query: string) => void;
  setSearchResults: (results: GoogleBookVolume[]) => void;
  selectBook: (book: GoogleBookVolume | null) => void;
  resetGoogleBookState: () => void;
};
