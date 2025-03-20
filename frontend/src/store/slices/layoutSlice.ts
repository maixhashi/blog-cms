import { StateCreator } from 'zustand';
import { LayoutState, LayoutComponentState } from '../types/layoutTypes';
import { EditedLayout, EditedLayoutComponent, initialEditedLayout, initialEditedLayoutComponent } from '../../types/models/layout';

/**
 * レイアウト関連のZustandストアスライス
 */
export const createLayoutSlice: StateCreator<LayoutState> = (set) => ({
  editedLayout: initialEditedLayout,
  updateEditedLayout: (payload: EditedLayout) => set({
    editedLayout: payload
  }),
  resetEditedLayout: () => set({ 
    editedLayout: initialEditedLayout 
  }),
});

export const createLayoutComponentSlice: StateCreator<LayoutComponentState> = (set) => ({
  editedLayoutComponent: initialEditedLayoutComponent,
  updateEditedLayoutComponent: (payload: EditedLayoutComponent) => set({
    editedLayoutComponent: payload
  }),
  resetEditedLayoutComponent: () => set({ 
    editedLayoutComponent: initialEditedLayoutComponent 
  }),
});