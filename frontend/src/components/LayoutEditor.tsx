import React, { useState, useEffect } from 'react';
import { DndProvider } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';
import GridLayout from 'react-grid-layout';
import 'react-grid-layout/css/styles.css';
import 'react-resizable/css/styles.css';
import { definitions } from '../types/api/generated';
import ComponentPalette from './ComponentPalette';
import LayoutCanvas from './LayoutCanvas';
import { useQueryLayoutComponents } from '../hooks/useQueryLayoutComponents';
import { fetchLayout, assignComponentToLayout, updateComponentPosition } from '../api/layoutApi';
import axios from 'axios';

interface LayoutEditorProps {
  layoutId: number;
}

const LayoutEditor: React.FC<LayoutEditorProps> = ({ layoutId }) => {
  // useQueryLayoutComponentsフックを使用してコンポーネントを取得
  const { data: components, isLoading: componentsLoading, refetch } = useQueryLayoutComponents();
  const [layout, setLayout] = useState<definitions['model.LayoutResponse'] | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadLayout = async () => {
      try {
        setLoading(true);
        // レイアウトを取得
        const layoutData = await fetchLayout(layoutId);
        setLayout(layoutData);
      } catch (error) {
        console.error('Failed to load layout:', error);
      } finally {
        setLoading(false);
      }
    };

    loadLayout();
  }, [layoutId]);

  // コンポーネントをレイアウトに追加する処理
  const handleDrop = async (componentId: number) => {
    try {
      // 初期位置とサイズを設定
      const position = {
        x: 0,
        y: 0,
        width: 2,
        height: 2
      };

      // APIを呼び出してコンポーネントをレイアウトに割り当て
      await assignComponentToLayout(componentId, layoutId, {
        layout_id: layoutId,
        position
      });

      // レイアウトを再取得して表示を更新
      const updatedLayout = await fetchLayout(layoutId);
      setLayout(updatedLayout);
      // コンポーネント一覧も更新
      refetch();
    } catch (error) {
      console.error('Failed to assign component:', error);
    }
  };

  // レイアウト内のコンポーネント位置変更処理
  const handleLayoutChange = async (newLayout: any[]) => {
    if (!layout || !layout.components) return;

    // 変更されたコンポーネントの位置を更新
    for (const item of newLayout) {
      const component = layout.components.find(c => c.id === parseInt(item.i));
      if (component) {
        try {
          await updateComponentPosition(component.id!, {
            x: item.x,
            y: item.y,
            width: item.w,
            height: item.h
          });
        } catch (error) {
          console.error(`Failed to update position for component ${item.i}:`, error);
        }
      }
    }

    // レイアウトを再取得
    const updatedLayout = await fetchLayout(layoutId);
    setLayout(updatedLayout);
  };

  // 新しいコンポーネントを作成する関数
  const handleCreateComponent = async () => {
    try {
      // APIを呼び出して新しいコンポーネントを作成
      await axios.post(`${process.env.REACT_APP_API_URL}/layout-components`, {
        name: `新しいコンポーネント ${Date.now()}`,
        type: 'text',
        content: '<p>ここにコンテンツを入力</p>'
      });
      
      // コンポーネント一覧を更新
      refetch();
    } catch (error) {
      console.error('Failed to create component:', error);
    }
  };

  if (loading || componentsLoading) {
    return <div>Loading...</div>;
  }

  return (
    <DndProvider backend={HTML5Backend}>
      <div className="layout-editor">
        <h2>{layout?.title} エディタ</h2>
        
        <div className="layout-editor-container">
          {/* コンポーネントパレット（ドラッグ元） */}
          <div className="component-palette">
            <h3>コンポーネント</h3>
            <button 
              onClick={handleCreateComponent}
              style={{
                padding: '8px 12px',
                margin: '10px 0',
                backgroundColor: '#4CAF50',
                color: 'white',
                border: 'none',
                borderRadius: '4px',
                cursor: 'pointer'
              }}
            >
              新規コンポーネント作成
            </button>
            <ComponentPalette components={components || []} />
          </div>
          
          {/* レイアウトキャンバス（ドロップ先） */}
          <div className="layout-canvas">
            <h3>レイアウト</h3>
            <LayoutCanvas 
              layout={layout} 
              onDrop={handleDrop}
              onLayoutChange={handleLayoutChange}
            />
          </div>
        </div>
      </div>
    </DndProvider>
  );
};

export default LayoutEditor;