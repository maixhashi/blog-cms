import { StateCreator } from 'zustand';
import { TaskState } from '../types';

export const createTaskSlice: StateCreator<TaskState> = (set) => ({
  editedTask: {
    id: 0,
    title: ''
  },
  updateEditedTask: (payload) => set({
    editedTask: payload
  }),
  resetEditedTask: () => set({ editedTask: {
    id: 0,
    title: '' 
  } }),
});
