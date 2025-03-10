export type FeedState = {
  editedFeed: {
    id: number;
    title: string;
    url: string;
    site_url: string;
    description: string;
    last_fetched_at: Date | string;
  };
  updateEditedFeed: (payload: {
    id: number;
    title: string;
    url: string;
    site_url: string;
    description: string;
    last_fetched_at: Date | string;
  }) => void;
  resetEditedFeed: () => void;
  selectedFeedId: number | null;
  setSelectedFeedId: (id: number | null) => void;
};
