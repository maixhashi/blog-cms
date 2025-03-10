export type TaskState = {
  editedTask: {
    id: number;
    title: string
  };
  updateEditedTask: (payload: {
    id: number;
    title: string
  }) => void;
  resetEditedTask: () => void;
};
