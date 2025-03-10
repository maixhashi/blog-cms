import { StateCreator } from 'zustand';
import { FeedState } from '../types/feedTypes';

export const createFeedSlice: StateCreator<FeedState> = (set) => ({
  editedFeed: { 
    id: 0, 
    title: '', 
    url: '', 
    site_url: '',
    description: '',
    last_fetched_at: new Date()
  },
  updateEditedFeed: (payload) => set({
    editedFeed: payload
  }),
  resetEditedFeed: () => set({
    editedFeed: { 
      id: 0, 
      title: '', 
      url: '', 
      site_url: '',
      description: '',
      last_fetched_at: new Date()
    },
  }),
  selectedFeedId: null,
  setSelectedFeedId: (id) => set({
    selectedFeedId: id
  }),
});
