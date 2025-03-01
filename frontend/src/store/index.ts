import { create } from 'zustand'
import { ExternalAPI, Feed } from '../types'

type State = {
  // Task関連
  editedTask: { id: number; title: string }
  updateEditedTask: (payload: { id: number; title: string }) => void
  resetEditedTask: () => void
  
  // ExternalAPI関連
  editedExternalAPI: {
    id: number
    name: string
    base_url: string
    description: string
  }
  updateEditedExternalAPI: (payload: {
    id: number
    name: string
    base_url: string
    description: string
  }) => void
  resetEditedExternalAPI: () => void
  
  // Feed関連
  editedFeed: {
    id: number
    title: string
    url: string
    site_url: string  // 追加
    description: string  // 追加
    last_fetched_at: Date | string  // 追加
  }
  updateEditedFeed: (payload: {
    id: number
    title: string
    url: string
    site_url: string  // 追加
    description: string  // 追加
    last_fetched_at: Date | string  // 追加
  }) => void
  resetEditedFeed: () => void
  selectedFeedId: number | null
  setSelectedFeedId: (id: number | null) => void
}

const useStore = create<State>((set) => ({
  // Task関連
  editedTask: { id: 0, title: '' },
  updateEditedTask: (payload) =>
    set({
      editedTask: payload,
    }),
  resetEditedTask: () =>
    set({ editedTask: { id: 0, title: '' } }),
  
  // ExternalAPI関連
  editedExternalAPI: { id: 0, name: '', base_url: '', description: '' },
  updateEditedExternalAPI: (payload) =>
    set({
      editedExternalAPI: payload,
    }),
  resetEditedExternalAPI: () =>
    set({
      editedExternalAPI: { id: 0, name: '', base_url: '', description: '' },
    }),
    
  // Feed関連
  editedFeed: { 
    id: 0, 
    title: '', 
    url: '', 
    site_url: '',  // 追加
    description: '',  // 追加
    last_fetched_at: new Date()  // 追加
  },
  updateEditedFeed: (payload) =>
    set({
      editedFeed: payload,
    }),
  resetEditedFeed: () =>
    set({
      editedFeed: { 
        id: 0, 
        title: '', 
        url: '', 
        site_url: '',  // 追加
        description: '',  // 追加
        last_fetched_at: new Date()  // 追加
      },
    }),
  selectedFeedId: null,
  setSelectedFeedId: (id) =>
    set({
      selectedFeedId: id,
    }),
}))

export default useStore