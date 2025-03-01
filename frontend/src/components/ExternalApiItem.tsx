import { FC, memo } from 'react'
import { PencilIcon, TrashIcon } from '@heroicons/react/24/solid'
import { useMutateExternalAPI } from '../hooks/useMutateExternalApi'
import useStore from '../store'
import '../ExternalApiItem.css'  // Change to import the correct CSS file

interface Props {
  id: number
  name: string
  base_url: string
  description: string
  onEdit: () => void
}

const ExternalApiItemMemo: FC<Props> = ({
  id,
  name,
  base_url,
  description,
  onEdit,
}) => {
  const { deleteExternalAPIMutation } = useMutateExternalAPI()
  const updateExternalAPI = useStore((state) => state.updateEditedExternalAPI)
  
  return (
    <li className="api-item">
      <div className="api-content">
        <h3 className="api-name">{name}</h3>
        <p className="api-url">{base_url}</p>
        <p className="api-description">{description}</p>
      </div>
      <div className="api-actions">
        <PencilIcon
          className="api-icon edit"
          onClick={() => {
            updateExternalAPI({
              id,
              name,
              base_url,
              description,
            })
            onEdit()
          }}
        />
        <TrashIcon
          className="api-icon delete"
          onClick={() => {
            deleteExternalAPIMutation.mutate(id)
          }}
        />
      </div>
    </li>
  )
}

export const ExternalAPIItem = memo(ExternalApiItemMemo)