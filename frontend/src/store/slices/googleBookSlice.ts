import { StateCreator } from 'zustand';
import { GoogleBookState } from '../types';
import { GoogleBookVolume } from '../../types/models/googleBook';

/**
 * Google Books API関連のZustandストアスライス
 */
export const createGoogleBookSlice: StateCreator<GoogleBookState> = (set) => ({
  searchQuery: '',
  searchResults: [],
  selectedBook: null,
  updateSearchQuery: (query: string) => set({
    searchQuery: query
  }),
  setSearchResults: (results: GoogleBookVolume[]) => set({
    searchResults: results
  }),
  selectBook: (book: GoogleBookVolume | null) => set({
    selectedBook: book
  }),
  resetGoogleBookState: () => set({
    searchQuery: '',
    searchResults: [],
    selectedBook: null
  })
});
