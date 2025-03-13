import axios from 'axios';
import { LayoutComponent } from '../types';

// LayoutComponentの作成、更新時に使用するインターフェース
interface LayoutComponentInput {
  id?: number;
  name: string;
  type: string;
  content?: string;
}

/**
 * 全てのレイアウトコンポーネントを取得
 */
export const fetchLayoutComponents = async (): Promise<LayoutComponent[]> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/layout-components`, {
    withCredentials: true
  });
  return response.data;
};

/**
 * 指定IDのレイアウトコンポーネントを取得
 */
export const fetchLayoutComponent = async (id: number): Promise<LayoutComponent> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/layout-components/${id}`, {
    withCredentials: true
  });
  return response.data;
};

/**
 * 新規レイアウトコンポーネントを作成
 */
export const createLayoutComponent = async (component: LayoutComponentInput): Promise<LayoutComponent> => {
  const response = await axios.post(`${process.env.REACT_APP_API_URL}/layout-components`, component, {
    withCredentials: true
  });
  return response.data;
};

/**
 * レイアウトコンポーネントを更新
 */
export const updateLayoutComponent = async (id: number, component: LayoutComponentInput): Promise<LayoutComponent> => {
  const response = await axios.put(`${process.env.REACT_APP_API_URL}/layout-components/${id}`, component, {
    withCredentials: true
  });
  return response.data;
};

/**
 * レイアウトコンポーネントを削除
 */
export const deleteLayoutComponent = async (id: number): Promise<void> => {
  await axios.delete(`${process.env.REACT_APP_API_URL}/layout-components/${id}`, {
    withCredentials: true
  });
};
