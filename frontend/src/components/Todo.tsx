import { FormEvent } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
  ArrowRightStartOnRectangleIcon,
  ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import useStore from '../store'
import { useQueryTasks } from '../hooks/useQueryTasks'
import { useMutateTask } from '../hooks/useMutateTask'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { TaskItem } from './TaskItem'
import '../Todo.css'

export const Todo = () => {
  const queryClient = useQueryClient()
  const { editedTask } = useStore()
  const updateTask = useStore((state) => state.updateEditedTask)
  const { data, isLoading } = useQueryTasks()
  const { createTaskMutation, updateTaskMutation } = useMutateTask()
  const { logoutMutation } = useMutateAuth()

  const submitTaskHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (editedTask.id === 0)
      createTaskMutation.mutate({
        title: editedTask.title,
      })
    else {
      updateTaskMutation.mutate(editedTask)
    }
  }

  const logout = async () => {
    await logoutMutation.mutateAsync()
    queryClient.removeQueries(['tasks'])
  }

  return (
    <div className="todo-container">
      <div className="todo-header">
        <ShieldCheckIcon className="header-icon" />
        <span className="header-title">Task Manager</span>
      </div>
      <ArrowRightStartOnRectangleIcon
        onClick={logout}
        className="logout-icon"
      />
      <form onSubmit={submitTaskHandler} className="todo-form">
        <input
          className="form-input"
          placeholder="title ?"
          type="text"
          onChange={(e) => updateTask({ ...editedTask, title: e.target.value })}
          value={editedTask.title || ''}
        />
        <button className="submit-button" disabled={!editedTask.title}>
          {editedTask.id === 0 ? 'Create' : 'Update'}
        </button>
      </form>
      {isLoading ? (
        <p className="loading-text">Loading...</p>
      ) : (
        <ul className="task-list">
          {data?.map((task) => (
            // task.idとtask.titleが存在する場合のみTaskItemをレンダリング
            task.id !== undefined && task.title !== undefined ? (
              <li key={task.id} className="task-item">
                <TaskItem id={task.id} title={task.title} />
              </li>
            ) : null
          ))}
        </ul>
      )}
    </div>
  )
}