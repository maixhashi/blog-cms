import { StateCreator } from 'zustand'
import { EditedLayoutComponent, initialEditedLayoutComponent } from '../../types/models/layout'
import { LayoutComponentState } from '../types/layoutComponentTypes'

export const createLayoutComponentSlice: StateCreator<LayoutComponentState> = (set) => ({
  editedLayoutComponent: initialEditedLayoutComponent,
  updateEditedLayoutComponent: (payload: EditedLayoutComponent) => set({
    editedLayoutComponent: payload
  }),
  resetEditedLayoutComponent: () => set({ 
    editedLayoutComponent: initialEditedLayoutComponent 
  }),
})