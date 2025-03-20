// Layoutモデルの型定義
export type Layout = {
  id: number;
  title: string;
  created_at?: string;
  updated_at?: string;
};

// 編集中のLayoutの状態を表す型
export type EditedLayout = {
  id: number;
  title: string;
};

// 初期状態
export const initialEditedLayout: EditedLayout = {
  id: 0,
  title: ''
};

// LayoutComponentの型定義
export type LayoutComponent = {
  id: number;
  title: string;
  layout_id: number;
  type?: string;
  content?: string;
};

// 編集中のLayoutComponentの状態
export type EditedLayoutComponent = {
  id: number;
  title: string;
  layout_id: number;
};

// 初期状態
export const initialEditedLayoutComponent: EditedLayoutComponent = {
  id: 0,
  title: '',
  layout_id: 0
};