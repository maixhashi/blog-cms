import { create } from 'zustand'
import { State } from './types'
import { createTaskSlice } from './slices/taskSlice'
import { createAPISlice } from './slices/apiSlice'
import { createFeedSlice } from './slices/feedSlice'
import { createArticleSlice } from './slices/articleSlice'
import { createLayoutSlice } from './slices/layoutSlice'
import { createLayoutComponentSlice } from './slices/layoutComponentSlice'
import { createUserSlice } from './slices/userSlice'

const useStore = create<State>((...args) => ({
  ...createTaskSlice(...args),
  ...createAPISlice(...args),
  ...createFeedSlice(...args),
  ...createArticleSlice(...args),
  ...createLayoutSlice(...args),
  ...createLayoutComponentSlice(...args),
  ...createUserSlice(...args),
}))

export default useStore