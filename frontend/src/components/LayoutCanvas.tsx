import React from 'react';
import { useDrop } from 'react-dnd';
import GridLayout from 'react-grid-layout';
import { definitions } from '../types/api/generated';
import { TrashIcon } from '@heroicons/react/24/solid'; // HeroIconsを使用

interface LayoutCanvasProps {
  layout: definitions['model.LayoutResponse'] | null;
  onDrop: (componentId: number) => void;
  onLayoutChange: (layout: any[]) => void;
  onRemoveComponent: (componentId: number) => void; // 削除機能を追加
}

const LayoutCanvas: React.FC<LayoutCanvasProps> = ({ 
  layout, 
  onDrop, 
  onLayoutChange,
  onRemoveComponent 
}) => {
  // ドロップ領域の設定
  const [{ isOver }, drop] = useDrop(() => ({
    accept: 'COMPONENT',
    drop: (item: { id: number }) => {
      onDrop(item.id);
      return { name: 'LayoutCanvas' };
    },
    collect: (monitor) => ({
      isOver: !!monitor.isOver(),
    }),
  }));

  // GridLayoutで使用するレイアウト設定を生成
  const gridItems = layout?.components?.map(component => ({
    i: component.id!.toString(),
    x: component.x || 0,
    y: component.y || 0,
    w: component.width || 2,
    h: component.height || 2,
    component
  })) || [];

  return (
    <div
      ref={drop}
      className="layout-canvas"
      style={{
        width: '100%',
        height: '600px',
        border: `2px dashed ${isOver ? 'green' : '#ccc'}`,
        backgroundColor: isOver ? 'rgba(0, 255, 0, 0.1)' : '#f0f0f0',
        padding: '10px',
        position: 'relative',
      }}
    >
      {layout?.components && layout.components.length > 0 ? (
        <GridLayout
          className="layout"
          layout={gridItems}
          cols={12}
          rowHeight={30}
          width={1000}
          onLayoutChange={onLayoutChange}
          draggableHandle=".component-drag-handle"
        >
          {gridItems.map(item => (
            <div key={item.i} className="grid-item">
              <div className="component-header" style={{ 
                display: 'flex', 
                justifyContent: 'space-between', 
                alignItems: 'center',
                padding: '5px', 
                backgroundColor: '#eee' 
              }}>
                <div className="component-drag-handle" style={{ cursor: 'move' }}>
                  {item.component.name}
                </div>
                <button
                  onClick={() => onRemoveComponent(parseInt(item.i))}
                  style={{
                    background: 'none',
                    border: 'none',
                    cursor: 'pointer',
                    color: '#ff4d4f'
                  }}
                  title="コンポーネントを削除"
                >
                  <TrashIcon style={{ width: '20px', height: '20px' }} />
                </button>
              </div>
              <div className="component-content" dangerouslySetInnerHTML={{ __html: item.component.content || '' }} />
            </div>
          ))}
        </GridLayout>
      ) : (
        <div className="empty-layout">
          <p>コンポーネントをここにドラッグしてください</p>
        </div>
      )}
    </div>
  );
};

export default LayoutCanvas;