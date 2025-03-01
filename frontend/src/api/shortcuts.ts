import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL || '';

export const getShortcutConfig = async () => {
  const response = await axios.get(`${API_URL}/api/shortcuts`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem('token')}`
    }
  });
  return response.data;
};

export const saveShortcutConfig = async (shortcuts: Record<string, string>) => {
  const response = await axios.post(`${API_URL}/api/shortcuts`, {
    shortcuts
  }, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem('token')}`
    }
  });
  return response.data;
};
