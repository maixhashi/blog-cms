export type ArticleState = {
  editedArticle: {
    id: number;
    title: string;
    content: string;
    published: boolean;
    tags: string;
  };
  updateEditedArticle: (payload: {
    id: number;
    title: string;
    content: string;
    published: boolean;
    tags: string;
  }) => void;
  resetEditedArticle: () => void;
};
