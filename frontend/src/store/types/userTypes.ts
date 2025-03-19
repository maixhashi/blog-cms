import { EditedUser } from '../../types/models/user';

export interface UserState {
  editedUser: EditedUser;
  updateEditedUser: (payload: EditedUser) => void;
  resetEditedUser: () => void;
}