import axios from 'axios';
import { LayoutComponent } from '../types';
import { definitions } from '../types/api/generated';

// APIリクエスト型を使用
type LayoutComponentRequest = definitions['model.LayoutComponentRequest'];

export const fetchLayoutComponents = async (): Promise<LayoutComponent[]> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/layout-components`, {
    withCredentials: true
  });
  return response.data;
};

export const fetchLayoutComponent = async (id: number): Promise<LayoutComponent> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/layout-components/${id}`, {
    withCredentials: true
  });
  return response.data;
};

export const createLayoutComponent = async (component: LayoutComponentRequest): Promise<LayoutComponent> => {
  const response = await axios.post(`${process.env.REACT_APP_API_URL}/layout-components`, component, {
    withCredentials: true
  });
  return response.data;
};

export const updateLayoutComponent = async (id: number, component: LayoutComponentRequest): Promise<LayoutComponent> => {
  const response = await axios.put(`${process.env.REACT_APP_API_URL}/layout-components/${id}`, component, {
    withCredentials: true
  });
  return response.data;
};

export const deleteLayoutComponent = async (id: number): Promise<void> => {
  await axios.delete(`${process.env.REACT_APP_API_URL}/layout-components/${id}`, {
    withCredentials: true
  });
};