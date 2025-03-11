import { StateCreator } from 'zustand';
import { LayoutState, LayoutComponentState } from '../types/layoutTypes';

export const createLayoutSlice: StateCreator<LayoutState> = (set) => ({
  editedLayout: {
    id: 0,
    title: ''
  },
  updateEditedLayout: (payload) => set({
    editedLayout: payload
  }),
  resetEditedLayout: () => set({ editedLayout: {
    id: 0,
    title: '' 
  } }),
});

export const createLayoutComponentSlice: StateCreator<LayoutComponentState> = (set) => ({
  editedLayoutComponent: {
    id: 0,
    title: '',
    layout_id: 0
  },
  updateEditedLayoutComponent: (payload) => set({
    editedLayoutComponent: payload
  }),
  resetEditedLayoutComponent: () => set({ editedLayoutComponent: {
    id: 0,
    title: '',
    layout_id: 0
  } }),
});