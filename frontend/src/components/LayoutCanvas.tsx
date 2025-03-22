import React from 'react';
import { useDrop } from 'react-dnd';
import GridLayout from 'react-grid-layout';
import { definitions } from '../types/api/generated';
import { TrashIcon } from '@heroicons/react/24/solid';
import { DefaultArticleCard } from './default-components/DefaultArticleCard';
import { DefaultCalendar } from './default-components/DefaultCalendar';
import { DefaultFooter } from './default-components/DefaultFooter';
import { DefaultHeader } from './default-components/DefaultHeader';
import { DefaultSidebar } from './default-components/DefaultSidebar';

interface LayoutCanvasProps {
  layout: definitions['model.LayoutResponse'] | null;
  onDrop: (componentId: number) => void;
  onLayoutChange: (layout: any[]) => void;
  onRemoveComponent: (componentId: number) => void;
}

// コンポーネントタイプとコンポーネントのマッピング
const componentMap: Record<string, React.FC<any>> = {
  'DefaultArticleCard': DefaultArticleCard,
  'DefaultCalendar': DefaultCalendar,
  'DefaultFooter': DefaultFooter,
  'DefaultHeader': DefaultHeader,
  'DefaultSidebar': DefaultSidebar,
};

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

  // コンポーネントを描画する関数
  const renderComponent = (component: any, width: number, height: number) => {
    // コンポーネントタイプが指定されている場合
    if (component.type && componentMap[component.type]) {
      const Component = componentMap[component.type];
      // コンポーネントのプロパティを解析
      let props = {};
      try {
        if (component.properties) {
          props = typeof component.properties === 'string' 
            ? JSON.parse(component.properties) 
            : component.properties;
        }
      } catch (e) {
        console.error('Failed to parse component properties:', e);
      }
      
      // コンポーネントの種類に応じてスタイルを調整
      const componentStyle: React.CSSProperties = {
        width: '100%',
        height: '100%',
      };

      // 特定のコンポーネントに対する特別なスタイル調整
      if (component.type === 'DefaultHeader' || component.type === 'DefaultFooter') {
        componentStyle.width = '100%';
      } else if (component.type === 'DefaultSidebar') {
        componentStyle.height = '100%';
        // サイドバーの幅を動的に設定
        props = { ...props, width: '100%' };
      } else if (component.type === 'DefaultArticleCard') {
        componentStyle.maxWidth = '100%';
      } else if (component.type === 'DefaultCalendar') {
        componentStyle.maxWidth = '100%';
      }
      
      return (
        <div style={componentStyle}>
          <Component {...props} />
        </div>
      );
    }
    
    // 従来の方法（HTML文字列として表示）
    return (
      <div 
        dangerouslySetInnerHTML={{ __html: component.content || '' }} 
        style={{ width: '100%', height: '100%' }}
      />
    );
  };

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
        overflow: 'auto'
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
          isDraggable={true}
          isResizable={true}
        >
          {gridItems.map(item => {
            const width = item.w * (1000 / 12); // グリッドの幅を計算
            const height = item.h * 30; // グリッドの高さを計算
            
            return (
              <div 
                key={item.i} 
                style={{ 
                  position: 'relative',
                  padding: 0,
                  margin: 0,
                  overflow: 'hidden'
                }}
              >
                {/* コンポーネントを直接描画 */}
                <div 
                  className="component-content" 
                  style={{ 
                    width: '100%', 
                    height: '100%',
                    position: 'relative'
                  }}
                >
                  {renderComponent(item.component, width, height)}
                </div>
                
                {/* 削除ボタンは小さく右上に配置 */}
                <button
                  onClick={() => onRemoveComponent(parseInt(item.i))}
                  style={{
                    position: 'absolute',
                    top: '2px',
                    right: '2px',
                    background: 'rgba(255, 255, 255, 0.7)',
                    border: 'none',
                    borderRadius: '50%',
                    width: '20px',
                    height: '20px',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    cursor: 'pointer',
                    padding: '2px',
                    zIndex: 10,
                    opacity: 0.7
                  }}
                  title="削除"
                  onMouseOver={(e) => e.currentTarget.style.opacity = '1'}
                  onMouseOut={(e) => e.currentTarget.style.opacity = '0.7'}
                >
                  <TrashIcon style={{ width: '14px', height: '14px', color: '#ff4d4f' }} />
                </button>
              </div>
            );
          })}
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