import { create } from 'zustand'
import { State } from './types'
import { createTaskSlice } from './slices/taskSlice'
import { createAPISlice } from './slices/apiSlice'
import { createFeedSlice } from './slices/feedSlice'
import { createArticleSlice } from './slices/articleSlice'

const useStore = create<State>((...args) => ({
  ...createTaskSlice(...args),
  ...createAPISlice(...args),
  ...createFeedSlice(...args),
  ...createArticleSlice(...args),
}))

export default useStore