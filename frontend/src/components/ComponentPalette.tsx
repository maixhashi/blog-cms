import React from 'react';
import { useDrag } from 'react-dnd';
import { definitions } from '../types/api/generated';

interface ComponentPaletteProps {
  components: definitions['model.LayoutComponentResponse'][];
}

// ドラッグ可能なコンポーネントアイテム
const DraggableComponent: React.FC<{
  component: definitions['model.LayoutComponentResponse'];
}> = ({ component }) => {
  const [{ isDragging }, drag] = useDrag(() => ({
    type: 'COMPONENT',
    item: { id: component.id },
    collect: (monitor) => ({
      isDragging: !!monitor.isDragging(),
    }),
  }));

  return (
    <div
      ref={drag}
      className="draggable-component"
      style={{
        opacity: isDragging ? 0.5 : 1,
        cursor: 'move',
        padding: '10px',
        margin: '5px',
        border: '1px solid #ccc',
        backgroundColor: '#f9f9f9',
      }}
    >
      <div>{component.name}</div>
      <div>Type: {component.type}</div>
    </div>
  );
};

const ComponentPalette: React.FC<ComponentPaletteProps> = ({ components }) => {
  // すべてのコンポーネントを表示する（フィルタリングを削除）
  return (
    <div className="component-palette">
      {components.length === 0 ? (
        <p>利用可能なコンポーネントがありません</p>
      ) : (
        components.map((component) => (
          <DraggableComponent key={component.id} component={component} />
        ))
      )}
    </div>
  );
};

export default ComponentPalette;