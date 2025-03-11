import { StateCreator } from 'zustand';
import { ArticleState } from '../types/articleTypes';

export const createArticleSlice: StateCreator<ArticleState> = (set) => ({
  editedArticle: {
    id: 0,
    title: '',
    content: '',
    published: false,
    tags: ''
  },
  updateEditedArticle: (payload) => set({
    editedArticle: payload
  }),
  resetEditedArticle: () => set({
    editedArticle: {
      id: 0,
      title: '',
      content: '',
      published: false,
      tags: ''
    }
  }),
});
