import { definitions } from '../api/generated';

// OpenAPIから生成された型を利用
export type Article = definitions['model.ArticleResponse'];
export type ArticleRequest = definitions['model.ArticleRequest'];

// 編集用の型定義
export type EditedArticle = {
  id: number;
  title: string;
  content: string;
  published: boolean;
  tags: string;
};

// 初期値
export const initialEditedArticle: EditedArticle = {
  id: 0,
  title: '',
  content: '',
  published: false,
  tags: '',
};
