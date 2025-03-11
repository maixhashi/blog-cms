import axios from 'axios';
import { Layout } from '../types';

// Layoutの作成、更新時に使用するインターフェース
interface LayoutInput {
  id?: number;
  title: string;
}

/**
 * 全てのレイアウトを取得
 */
export const fetchLayouts = async (): Promise<Layout[]> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/layouts`, {
    withCredentials: true
  });
  return response.data;
};

/**
 * 指定IDのレイアウトを取得
 */
export const fetchLayout = async (id: number): Promise<Layout> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/layouts/${id}`, {
    withCredentials: true
  });
  return response.data;
};

/**
 * 新規レイアウトを作成
 */
export const createLayout = async (layout: LayoutInput): Promise<Layout> => {
  const response = await axios.post(`${process.env.REACT_APP_API_URL}/layouts`, layout, {
    withCredentials: true
  });
  return response.data;
};

/**
 * レイアウトを更新
 */
export const updateLayout = async (id: number, layout: LayoutInput): Promise<Layout> => {
  const response = await axios.put(`${process.env.REACT_APP_API_URL}/layouts/${id}`, layout, {
    withCredentials: true
  });
  return response.data;
};

/**
 * レイアウトを削除
 */
export const deleteLayout = async (id: number): Promise<void> => {
  await axios.delete(`${process.env.REACT_APP_API_URL}/layouts/${id}`, {
    withCredentials: true
  });
};
