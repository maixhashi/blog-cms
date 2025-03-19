import { EditedArticle } from '../../types/models/article';

export interface ArticleState {
  editedArticle: EditedArticle;
  updateEditedArticle: (payload: EditedArticle) => void;
  resetEditedArticle: () => void;
}
