import { StateCreator } from 'zustand';
import { TaskState } from '../types/taskTypes';

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
