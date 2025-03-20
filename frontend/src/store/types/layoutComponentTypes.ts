import { EditedLayoutComponent } from '../../types/models/layout';

export interface LayoutComponentState {
  editedLayoutComponent: EditedLayoutComponent;
  updateEditedLayoutComponent: (payload: EditedLayoutComponent) => void;
  resetEditedLayoutComponent: () => void;
}