import { StateCreator } from 'zustand';
import { ArticleState } from '../types/articleTypes';
import { EditedArticle, initialEditedArticle } from '../../types/models/article';

/**
 * 記事関連のZustandストアスライス
 */
export const createArticleSlice: StateCreator<ArticleState> = (set) => ({
  editedArticle: initialEditedArticle,
  updateEditedArticle: (payload: EditedArticle) => set({
    editedArticle: payload
  }),
  resetEditedArticle: () => set({ 
    editedArticle: initialEditedArticle 
  }),
});