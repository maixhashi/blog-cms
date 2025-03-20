import axios from 'axios';
import { Layout } from '../types';

// レイアウト一覧を取得
export const fetchLayouts = async (): Promise<Layout[]> => {
  const { data } = await axios.get<Layout[]>(
    `${process.env.REACT_APP_API_URL}/layouts`,
    { withCredentials: true }
  );
  return data;
};

// 新規レイアウト作成
export const createLayout = async (layout: { title: string }): Promise<Layout> => {
  const { data } = await axios.post<Layout>(
    `${process.env.REACT_APP_API_URL}/layouts`,
    layout,
    { withCredentials: true }
  );
  return data;
};

// レイアウト更新
export const updateLayout = async (id: number, layout: { title: string }): Promise<Layout> => {
  const { data } = await axios.put<Layout>(
    `${process.env.REACT_APP_API_URL}/layouts/${id}`,
    layout,
    { withCredentials: true }
  );
  return data;
};

// レイアウト削除
export const deleteLayout = async (id: number): Promise<void> => {
  await axios.delete(
    `${process.env.REACT_APP_API_URL}/layouts/${id}`,
    { withCredentials: true }
  );
};