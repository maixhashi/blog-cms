import { definitions } from '../api/generated';

/**
 * 書籍モデルの型定義
 * 
 * このファイルでは書籍関連のすべての型を定義します。
 * APIから自動生成された型を基に、アプリケーション固有の型を定義します。
 */

// APIから生成された型が存在しないため、独自に定義
export type BookResponse = {
  id?: number;
  title?: string;
  author?: string;
  isbn?: string;
  description?: string;
  thumbnail_url?: string;
  page_count?: number;
  created_at?: string;
  updated_at?: string;
};

// APIから生成された型が存在しないため、独自に定義
export type BookRequest = {
  title: string;
  author: string;
  isbn?: string;
  description?: string;
  thumbnail_url?: string;
  page_count?: number;
};

/**
 * 編集用の書籍型
 * APIのレスポンス型から必要なプロパティを取り出し、必須にしたもの
 */
export type EditedBook = {
  id: number;
  title: string;
  author: string;
  isbn: string;
  description: string;
  thumbnail_url: string;
  page_count: number;
};

/**
 * 編集用書籍の初期値
 */
export const initialEditedBook: EditedBook = {
  id: 0,
  title: '',
  author: '',
  isbn: '',
  description: '',
  thumbnail_url: '',
  page_count: 0
};

/**
 * 書籍状態の型定義（Zustandストア用）
 */
export interface BookState {
  editedBook: EditedBook;
  updateEditedBook: (payload: EditedBook) => void;
  resetEditedBook: () => void;
}
