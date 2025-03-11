import { create } from 'zustand'
import { State } from './types'
import { createTaskSlice } from './slices/taskSlice'
import { createAPISlice } from './slices/apiSlice'
import { createFeedSlice } from './slices/feedSlice'
import { createArticleSlice } from './slices/articleSlice'
import { createLayoutSlice } from './slices/layoutSlice'
import { createLayoutComponentSlice } from './slices/layoutSlice'

const useStore = create<State>((...args) => ({
  ...createTaskSlice(...args),
  ...createAPISlice(...args),
  ...createFeedSlice(...args),
  ...createArticleSlice(...args),
  ...createLayoutSlice(...args),
  ...createLayoutComponentSlice(...args),
}))

export default useStore