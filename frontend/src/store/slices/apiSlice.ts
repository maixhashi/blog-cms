import { StateCreator } from 'zustand';
import { ExternalAPIState } from '../types/apiTypes';

export const createAPISlice: StateCreator<ExternalAPIState> = (set) => ({
  editedExternalAPI: { 
    id: 0,
    name: '',
    base_url: '',
    description: ''
  },
  updateEditedExternalAPI: (payload) => set({
    editedExternalAPI: payload
  }),
  resetEditedExternalAPI: () => set({
    editedExternalAPI: {
      id: 0,
      name: '',
      base_url: '',
      description: ''
    },
  }),
});
