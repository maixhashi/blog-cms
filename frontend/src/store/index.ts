import { create } from 'zustand'

type EditedTask = {
  id: number
  title: string
}

type EditedFeed = {
  id: number
  title: string
  url: string
  site_url: string
  description: string
  last_fetched_at: Date
}

type State = {
  editedTask: EditedTask
  updateEditedTask: (payload: EditedTask) => void
  resetEditedTask: () => void
  editedFeed: EditedFeed
  updateEditedFeed: (payload: EditedFeed) => void
  resetEditedFeed: () => void
  selectedFeedId: number | null // 修正: selectedFeedId に変更
  setSelectedFeedId: (feedId: number | null) => void // 修正: 選択されたFeedのIDを更新
}

const useStore = create<State>((set) => ({
  editedTask: { id: 0, title: '' },
  updateEditedTask: (payload) =>
    set({
      editedTask: payload,
    }),
  resetEditedTask: () => set({ editedTask: { id: 0, title: '' } }),
  editedFeed: { 
    id: 0,
    title: '',
    url: '',
    site_url: '',
    description: '',
    last_fetched_at: new Date(),
  },
  updateEditedFeed: (payload) =>
    set({
      editedFeed: payload,
    }),
  resetEditedFeed: () => set({ editedFeed: { 
    id: 0,
    title: '',
    url: '',
    site_url: '',
    description: '',
    last_fetched_at: new Date(),
  } }),
  selectedFeedId: null, // 修正: 初期状態は null
  setSelectedFeedId: (feedId) => set({ selectedFeedId: feedId }) // 修正: 選択されたFeedのIDを更新
}))

export default useStore