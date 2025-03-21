import axios from 'axios';
import { definitions } from '../types/api/generated';

const API_URL = process.env.REACT_APP_API_URL || '';

// レイアウトコンポーネント一覧を取得
export const fetchLayoutComponents = async (): Promise<definitions['model.LayoutComponentResponse'][]> => {
  const response = await axios.get(`${API_URL}/layout-components`);
  return response.data;
};

// 特定のレイアウトを取得
export const fetchLayout = async (layoutId: number): Promise<definitions['model.LayoutResponse']> => {
  const response = await axios.get(`${API_URL}/layouts/${layoutId}`);
  return response.data;
};

// コンポーネントをレイアウトに割り当て
export const assignComponentToLayout = async (
  componentId: number,
  layoutId: number,
  data: definitions['model.AssignLayoutRequest']
): Promise<void> => {
  await axios.post(`${API_URL}/layout-components/${componentId}/assign/${layoutId}`, data);
};

// コンポーネントの位置を更新
export const updateComponentPosition = async (
  componentId: number,
  position: definitions['model.PositionRequest']
): Promise<void> => {
  await axios.put(`${API_URL}/layout-components/${componentId}/position`, { position });
};

// コンポーネントをレイアウトから削除
export const removeComponentFromLayout = async (componentId: number): Promise<void> => {
  await axios.delete(`${API_URL}/layout-components/${componentId}/assign`);
};