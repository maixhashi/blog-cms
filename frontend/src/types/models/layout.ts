import { definitions } from '../api/generated';

// APIから返されるレスポンスの型
export type LayoutComponent = definitions['model.LayoutComponentResponse'];
export type Layout = definitions['model.LayoutResponse'];

// 編集中のレイアウトコンポーネント状態
export interface EditedLayoutComponent {
  id: number;
  name: string;
  type: string;
  content: string;
}

// 編集中のレイアウト状態
export interface EditedLayout {
  id: number;
  title: string;
  // 必要に応じて他のプロパティを追加
}

// 初期状態
export const initialEditedLayoutComponent: EditedLayoutComponent = {
  id: 0,
  name: '',
  type: '',
  content: ''
};

// レイアウトの初期状態
export const initialEditedLayout: EditedLayout = {
  id: 0,
  title: ''
  // 必要に応じて他のプロパティの初期値を設定
};
