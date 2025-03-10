export type ExternalAPIState = {
  editedExternalAPI: {
    id: number;
    name: string;
    base_url: string;
    description: string;
  };
  updateEditedExternalAPI: (payload: {
    id: number;
    name: string;
    base_url: string;
    description: string;
  }) => void;
  resetEditedExternalAPI: () => void;
};
