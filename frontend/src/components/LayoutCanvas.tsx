import React from 'react';
import { useDrop } from 'react-dnd';
import GridLayout from 'react-grid-layout';
import { definitions } from '../types/api/generated';

interface LayoutCanvasProps {
  layout: definitions['model.LayoutResponse'] | null;
  onDrop: (componentId: number) => void;
  onLayoutChange: (layout: any[]) => void;
}

const LayoutCanvas: React.FC<LayoutCanvasProps> = ({ layout, onDrop, onLayoutChange }) => {
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
              <div className="component-drag-handle" style={{ cursor: 'move', padding: '5px', backgroundColor: '#eee' }}>
                {item.component.name}
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
