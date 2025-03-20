import { EditedLayout, EditedLayoutComponent } from '../../types/models/layout';

export interface LayoutState {
  editedLayout: EditedLayout;
  updateEditedLayout: (payload: EditedLayout) => void;
  resetEditedLayout: () => void;
}

export interface LayoutComponentState {
  editedLayoutComponent: EditedLayoutComponent;
  updateEditedLayoutComponent: (payload: EditedLayoutComponent) => void;
  resetEditedLayoutComponent: () => void;
}