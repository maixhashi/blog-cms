import axios from 'axios';
import { FeedArticle } from '../types/index';

const API_URL = process.env.REACT_APP_API_URL;

// すべてのフィード記事を取得する
export const getAllFeedArticles = async (): Promise<FeedArticle[]> => {
  const token = localStorage.getItem('token');
  const response = await axios.get(`${API_URL}/feed-articles`, {
    headers: {
      Authorization: `Bearer ${token}`
    }
  });
  return response.data;
};

// 特定のフィードの記事を取得する
export const getFeedArticlesByFeedId = async (feedId: number): Promise<FeedArticle[]> => {
  const token = localStorage.getItem('token');
  const response = await axios.get(`${API_URL}/feed-articles/${feedId}`, {
    headers: {
      Authorization: `Bearer ${token}`
    }
  });
  return response.data;
};

// 特定の記事の詳細を取得する
export const getFeedArticleById = async (feedId: number, articleId: number): Promise<FeedArticle> => {
  const token = localStorage.getItem('token');
  const response = await axios.get(`${API_URL}/feed-articles/${feedId}/${articleId}`, {
    headers: {
      Authorization: `Bearer ${token}`
    }
  });
  return response.data;
};
