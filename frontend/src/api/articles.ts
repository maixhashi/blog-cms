// You'll need to customize this based on your backend API
import axios from 'axios';

interface Article {
  id?: string;
  title: string;
  content: string;
  tags: string[];
  published: boolean
}

export const fetchArticle = async (id: string): Promise<Article> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/articles/${id}`);
  return response.data;
};

export const createArticle = async (article: Article): Promise<Article> => {
  const response = await axios.post(`${process.env.REACT_APP_API_URL}/articles`, article);
  return response.data;
};

export const updateArticle = async (id: string, article: Article): Promise<Article> => {
  const response = await axios.put(`${process.env.REACT_APP_API_URL}/articles/${id}`, article);
  return response.data;
};
