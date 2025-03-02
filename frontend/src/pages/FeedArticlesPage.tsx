import React, { FC } from 'react';
import { FeedArticleList } from '../components/FeedArticleList';

export const FeedArticlesPage: FC = () => {
  return (
    <div className="container">
      <h1>フィード記事一覧</h1>
      <FeedArticleList />
    </div>
  );
};
