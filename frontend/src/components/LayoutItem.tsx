import { FC, memo } from 'react'
import { PencilIcon, TrashIcon } from '@heroicons/react/24/solid'
import useStore from '../store'
import { Layout } from '../types'
import { useMutateLayout } from '../hooks/useMutateLayout'
import '../LayoutItem.css' // ← CSSファイルを作成する必要があります

interface LayoutItemProps {
  id: number
  title: string
  onEdit: () => void
}

export const LayoutItem = ({ id, title, onEdit }: LayoutItemProps) => {
  const { deleteLayoutMutation } = useMutateLayout()
  return (
    <li className="layout-item">
      <span className="layout-id">{id}</span>
      <span className="layout-title">{title}</span>
      <div className="layout-actions">
        <button 
          className="edit-button" 
          onClick={onEdit}
        >
          編集
        </button>
        <TrashIcon
          className="layout-icon delete"
          onClick={() => {
            deleteLayoutMutation.mutate(id)
          }}
        />
      </div>
    </li>
  )
}
