export type LayoutState = {
  editedLayout: {
    id: number;
    title: string
  };
  updateEditedLayout: (payload: {
    id: number;
    title: string
  }) => void;
  resetEditedLayout: () => void;
};

export type LayoutComponentState = {
  editedLayoutComponent: {
    id: number;
    title: string;
    layout_id: number;
  };
  updateEditedLayoutComponent: (payload: {
    id: number;
    title: string;
    layout_id: number;
  }) => void;
  resetEditedLayoutComponent: () => void;
};
