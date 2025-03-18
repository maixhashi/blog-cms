import { EditedTask } from '../../types';

export interface TaskState {
  editedTask: EditedTask;
  updateEditedTask: (payload: EditedTask) => void;
  resetEditedTask: () => void;
}
