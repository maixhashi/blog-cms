import { create } from 'zustand'

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
    site_url: string
    description: string
    last_fetched_at: Date | string
  }
  updateEditedFeed: (payload: {
    id: number
    title: string
    url: string
    site_url: string
    description: string
    last_fetched_at: Date | string
  }) => void
  resetEditedFeed: () => void
  selectedFeedId: number | null
  setSelectedFeedId: (id: number | null) => void
  
  // Article関連
  editedArticle: {
    id: number
    title: string
    content: string
    published: boolean
    tags: string
  }
  updateEditedArticle: (payload: {
    id: number
    title: string
    content: string
    published: boolean
    tags: string
  }) => void
  resetEditedArticle: () => void
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
    site_url: '',
    description: '',
    last_fetched_at: new Date()
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
        site_url: '',
        description: '',
        last_fetched_at: new Date()
      },
    }),
  selectedFeedId: null,
  setSelectedFeedId: (id) =>
    set({
      selectedFeedId: id,
    }),
    
  // Article関連
  editedArticle: {
    id: 0,
    title: '',
    content: '',
    published: false,
    tags: ''
  },
  updateEditedArticle: (payload) =>
    set({
      editedArticle: payload,
    }),
  resetEditedArticle: () =>
    set({
      editedArticle: {
        id: 0,
        title: '',
        content: '',
        published: false,
        tags: ''
      }
    }),
}))

export default useStore