import { StateCreator } from 'zustand';
import { BookState } from '../types';
import { initialEditedBook } from '../../types/models/book';

/**
 * 書籍関連のZustandストアスライス
 */
export const createBookSlice: StateCreator<BookState> = (set) => ({
  editedBook: initialEditedBook,
  updateEditedBook: (payload) => set({
    editedBook: payload
  }),
  resetEditedBook: () => set({ 
    editedBook: initialEditedBook 
  }),
});
