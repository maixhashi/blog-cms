import { StateCreator } from 'zustand'
import { State, LayoutComponentState } from '../types'

export const createLayoutComponentSlice: StateCreator<
  State,
  [],
  [],
  LayoutComponentState
> = (set) => ({
  editedLayoutComponent: { id: 0, name: '', type: '', content: '' },
  updateEditedLayoutComponent: (payload) =>
    set({
      editedLayoutComponent: payload,
    }),
  resetEditedLayoutComponent: () =>
    set({
      editedLayoutComponent: { id: 0, name: '', type: '', content: '' },
    }),
})