
import React, { FC } from 'react';
import { useQuery } from '@tanstack/react-query';
import { Link } from 'react-router-dom';
import { getAllFeedArticles } from '../api/feedArticles';
import { FeedArticle } from '../types';
import '../styles/FeedArticleList.css'; // CSSをインポート

export const FeedArticleList: FC = () => {
  const { data: articles, isLoading, isError } = useQuery<FeedArticle[]>(
    ['feedArticles'], 
    getAllFeedArticles
  );

  if (isLoading) return <div className="loader">読み込み中...</div>;
  if (isError) return <div className="error">エラーが発生しました</div>;
  
  return (
    <div className="feed-article-list">
      <h2>フィード記事一覧</h2>
      {articles && articles.length > 0 ? (
        <div className="articles">
          {articles.map((article) => (
            <div key={article.id} className="article-card">
              <h3 className="article-title">{article.title}</h3>
              <p className="article-summary">
                {article.summary || 
                 (article.content ? article.content.substring(0, 150) : '内容なし')}
                ...
              </p>
              <div className="article-meta">
                <Link to={`${article.url}`} className="article-link">
                  詳細を見る
                </Link>
                <span>公開日: {article.published_at ? new Date(article.published_at).toLocaleDateString() : '不明'}</span>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <p className="no-articles">記事が見つかりませんでした</p>
      )}
    </div>
  );
};