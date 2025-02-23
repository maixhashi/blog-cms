import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { Feed } from '../types'
import useStore from '../store'
import { useError } from '../hooks/useError'

export const useMutateFeed = () => {
  const queryClient = useQueryClient()
  const { switchErrorHandling } = useError()
  const resetEditedFeed = useStore((state) => state.resetEditedFeed)

  const createFeedMutation = useMutation(
    (feed: Omit<Feed, 'id' | 'created_at' | 'updated_at'>) =>
      axios.post<Feed>(`${process.env.REACT_APP_API_URL}/feeds`, feed),
    {
      onSuccess: (res) => {
        const previousFeeds = queryClient.getQueryData<Feed[]>(['feeds'])
        if (previousFeeds) {
          queryClient.setQueryData(['feeds'], [...previousFeeds, res.data])
        }
        resetEditedFeed()
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )
  const updateFeedMutation = useMutation(
    (feed: Omit<Feed, 'created_at' | 'updated_at'>) =>
      axios.put<Feed>(`${process.env.REACT_APP_API_URL}/feeds/${feed.id}`, {
        title: feed.title,
      }),
    {
      onSuccess: (res, variables) => {
        const previousFeeds = queryClient.getQueryData<Feed[]>(['feeds'])
        if (previousFeeds) {
          queryClient.setQueryData<Feed[]>(
            ['feeds'],
            previousFeeds.map((feed) =>
              feed.id === variables.id ? res.data : feed
            )
          )
        }
        resetEditedFeed()
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )
  const deleteFeedMutation = useMutation(
    (id: number) =>
      axios.delete(`${process.env.REACT_APP_API_URL}/feeds/${id}`),
    {
      onSuccess: (_, variables) => {
        const previousFeeds = queryClient.getQueryData<Feed[]>(['feeds'])
        if (previousFeeds) {
          queryClient.setQueryData<Feed[]>(
            ['feeds'],
            previousFeeds.filter((feed) => feed.id !== variables)
          )
        }
        resetEditedFeed()
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )
  return {
    createFeedMutation,
    updateFeedMutation,
    deleteFeedMutation,
  }
}