import React, { FC, useState, useCallback, useEffect } from 'react';
import { Box, Grid, Paper, Typography, Button, IconButton } from '@mui/material';
import RGL, { WidthProvider, Layout } from 'react-grid-layout';
import { ComponentSelector } from './ComponentSelector';
import { ComponentPropertyEditor } from './ComponentPropertyEditor';
import { DefaultComponent, getComponentById } from './default-components';
import 'react-grid-layout/css/styles.css';
import 'react-resizable/css/styles.css';

// グリッドレイアウトの設定
const ReactGridLayout = WidthProvider(RGL);

interface LayoutItem {
  id: string;
  componentId: string;
  props: any;
  // レイアウト情報
  x: number;
  y: number;
  w: number;
  h: number;
}

interface LayoutEditorProps {
  layoutId?: number;
}

export const LayoutEditor: FC<LayoutEditorProps> = ({ layoutId }) => {
  const [layoutItems, setLayoutItems] = useState<LayoutItem[]>([]);
  const [selectedItemId, setSelectedItemId] = useState<string | null>(null);
  const [gridCols, setGridCols] = useState(12); // グリッドの列数
  const [gridRowHeight, setGridRowHeight] = useState(50); // 行の高さ
  const [editMode, setEditMode] = useState<'layout' | 'content'>('content');

  // 選択されたアイテムを取得
  const selectedItem = selectedItemId 
    ? layoutItems.find(item => item.id === selectedItemId) 
    : null;
  
  // 選択されたコンポーネントを取得
  const selectedComponent = selectedItem 
    ? getComponentById(selectedItem.componentId) 
    : null;
  
  // レイアウトデータの読み込み（layoutId がある場合）
  useEffect(() => {
    if (layoutId) {
      // ここでサーバーからレイアウトデータを取得する処理を追加
      // 例: API呼び出しやローカルストレージからの読み込み
      const savedLayouts = JSON.parse(localStorage.getItem('savedLayouts') || '[]');
      const layout = savedLayouts.find((l: any) => l.id === layoutId);
      
      if (layout) {
        setLayoutItems(layout.items);
      }
    }
  }, [layoutId]);
  
  // 新しいコンポーネントを追加
  const handleAddComponent = (component: DefaultComponent) => {
    // コンポーネントのデフォルトサイズを取得（コンポーネントに応じて調整）
    const defaultSize = getComponentDefaultSize(component.id);
    
    const newItem: LayoutItem = {
      id: `item-${Date.now()}`,
      componentId: component.id,
      props: { ...component.defaultProps },
      x: 0, // デフォルトのX位置
      y: Infinity, // 最下部に配置
      w: defaultSize.w, // コンポーネントに応じたデフォルト幅
      h: defaultSize.h, // コンポーネントに応じたデフォルト高さ
    };
    
    setLayoutItems([...layoutItems, newItem]);
    setSelectedItemId(newItem.id);
  };
  
  // コンポーネントのデフォルトサイズを取得
  const getComponentDefaultSize = (componentId: string) => {
    // コンポーネントの種類に応じてデフォルトサイズを返す
    // 実際のアプリケーションでは、コンポーネントの種類に応じて適切なサイズを設定
    switch (componentId) {
      case 'header':
        return { w: 12, h: 2 };
      case 'footer':
        return { w: 12, h: 2 };
      case 'sidebar':
        return { w: 3, h: 8 };
      case 'article':
        return { w: 9, h: 8 };
      case 'card':
        return { w: 4, h: 4 };
      default:
        return { w: 6, h: 4 };
    }
  };
  
  // コンポーネントのプロパティを更新
  const handlePropsChange = (newProps: any) => {
    if (!selectedItemId) return;
    
    const updatedItems = layoutItems.map(item => 
      item.id === selectedItemId 
        ? { ...item, props: newProps } 
        : item
    );
    
    setLayoutItems(updatedItems);
  };
  
  // コンポーネントを削除
  const handleRemoveComponent = (itemId: string) => {
    setLayoutItems(layoutItems.filter(item => item.id !== itemId));
    if (selectedItemId === itemId) {
      setSelectedItemId(null);
    }
  };
  
  // グリッドレイアウトの変更を処理
  const handleLayoutChange = (newLayout: Layout[]) => {
    // レイアウト情報を更新
    const updatedItems = layoutItems.map(item => {
      const layoutItem = newLayout.find(l => l.i === item.id);
      if (layoutItem) {
        return {
          ...item,
          x: layoutItem.x,
          y: layoutItem.y,
          w: layoutItem.w,
          h: layoutItem.h,
        };
      }
      return item;
    });
    
    setLayoutItems(updatedItems);
  };
  
  // コンポーネントをレンダリング
  const renderComponent = (item: LayoutItem) => {
    const component = getComponentById(item.componentId);
    if (!component) return null;
    
    const Component = component.component;
    return <Component {...item.props} />;
  };
  
  // レイアウトを保存する関数
  const handleSaveLayout = () => {
    const layoutName = prompt('レイアウト名を入力してください:', 'マイレイアウト');
    if (!layoutName) return;
    
    const layoutData = {
      id: layoutId || Date.now(),
      name: layoutName,
      items: layoutItems,
      createdAt: new Date().toISOString(),
    };
    
    // ローカルストレージに保存（実際のアプリではAPIを使用）
    const savedLayouts = JSON.parse(localStorage.getItem('savedLayouts') || '[]');
    const existingIndex = savedLayouts.findIndex((l: any) => l.id === layoutData.id);
    
    if (existingIndex >= 0) {
      savedLayouts[existingIndex] = layoutData;
    } else {
      savedLayouts.push(layoutData);
    }
    
    localStorage.setItem('savedLayouts', JSON.stringify(savedLayouts));
    alert('レイアウトを保存しました');
  };
  
  // react-grid-layout 用のレイアウト配列を生成
  const gridLayout = layoutItems.map(item => ({
    i: item.id,
    x: item.x,
    y: item.y,
    w: item.w,
    h: item.h,
  }));
  
  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" gutterBottom>
        レイアウトエディタ
      </Typography>
      
      <Box sx={{ mb: 2, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <Box>
          <Button 
            variant={editMode === 'content' ? 'contained' : 'outlined'} 
            onClick={() => setEditMode('content')}
            sx={{ mr: 1 }}
          >
            コンテンツ編集
          </Button>
          <Button 
            variant={editMode === 'layout' ? 'contained' : 'outlined'} 
            onClick={() => setEditMode('layout')}
          >
            レイアウト編集
          </Button>
        </Box>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={handleSaveLayout}
        >
          レイアウトを保存
        </Button>
      </Box>
      
      <Grid container spacing={3}>
        <Grid item xs={12} md={8}>
          <Paper 
            sx={{ 
              p: 2, 
              minHeight: 600, 
              backgroundColor: '#f9f9f9',
              position: 'relative',
            }}
          >
            {editMode === 'content' && (
              <Box sx={{ mb: 2 }}>
                <ComponentSelector onSelectComponent={handleAddComponent} />
              </Box>
            )}
            
            <Box sx={{ position: 'relative' }}>
              {layoutItems.length > 0 ? (
                <ReactGridLayout
                  className="layout"
                  layout={gridLayout}
                  cols={gridCols}
                  rowHeight={gridRowHeight}
                  onLayoutChange={handleLayoutChange}
                  compactType="vertical"
                  preventCollision={false}
                  isBounded
                  // 編集モードに応じてドラッグ・リサイズの可否を切り替え
                  isDraggable={true}
                  isResizable={true}
                  // グリッドの表示を調整
                  margin={[10, 10]}
                  containerPadding={[10, 10]}
                >
                  {layoutItems.map(item => (
                    <div key={item.id}>
                      <Box
                        sx={{
                          height: '100%',
                          width: '100%',
                          border: selectedItemId === item.id ? '2px dashed #2196f3' : '1px solid #e0e0e0',
                          borderRadius: 1,
                          p: 1,
                          position: 'relative',
                          overflow: 'hidden',
                          backgroundColor: '#fff',
                          // コンポーネントの内容に合わせてサイズを調整
                          display: 'flex',
                          flexDirection: 'column',
                        }}
                        onClick={(e) => {
                          e.stopPropagation();
                          setSelectedItemId(item.id);
                        }}
                      >
                        <Box 
                          sx={{ 
                            position: 'absolute', 
                            top: 5, 
                            right: 5, 
                            zIndex: 10,
                            display: 'flex',
                            gap: 1,
                          }}
                        >
                          <Button 
                            variant="contained" 
                            color="error" 
                            size="small"
                            onClick={(e) => {
                              e.stopPropagation();
                              handleRemoveComponent(item.id);
                            }}
                          >
                            削除
                          </Button>
                        </Box>
                        <Box sx={{ flexGrow: 1, overflow: 'auto' }}>
                          {renderComponent(item)}
                        </Box>
                      </Box>
                    </div>
                  ))}
                </ReactGridLayout>
              ) : (
                <Box 
                  sx={{ 
                    p: 4, 
                    textAlign: 'center', 
                    border: '2px dashed #ccc',
                    borderRadius: 2,
                  }}
                >
                  <Typography color="textSecondary">
                    コンポーネントを追加してレイアウトを作成してください
                  </Typography>
                </Box>
              )}
            </Box>
          </Paper>
        </Grid>
        
        <Grid item xs={12} md={4}>
          {editMode === 'content' ? (
            <Paper sx={{ p: 2 }}>
              {selectedComponent && selectedItem ? (
                <ComponentPropertyEditor
                  component={selectedComponent}
                  currentProps={selectedItem.props}
                  onPropsChange={handlePropsChange}
                />
              ) : (
                <Typography color="textSecondary" sx={{ p: 2, textAlign: 'center' }}>
                  コンポーネントを選択してプロパティを編集してください
                </Typography>
              )}
            </Paper>
          ) : (
            <Paper sx={{ p: 2 }}>
              <Typography variant="h6" gutterBottom>
                レイアウト設定
              </Typography>
              <Typography variant="body2" sx={{ mb: 2 }}>
                コンポーネントをドラッグして位置を変更したり、端をドラッグしてサイズを変更できます。
              </Typography>
              
              <Box sx={{ mb: 2 }}>
                <Typography variant="subtitle2">グリッド列数: {gridCols}</Typography>
                <input
                  type="range"
                  min="1"
                  max="24"
                  value={gridCols}
                  onChange={(e) => setGridCols(Number(e.target.value))}
                  style={{ width: '100%' }}
                />
              </Box>
              
              <Box>
                <Typography variant="subtitle2">行の高さ: {gridRowHeight}px</Typography>
                <input
                  type="range"
                  min="30"
                  max="100"
                  value={gridRowHeight}
                  onChange={(e) => setGridRowHeight(Number(e.target.value))}
                  style={{ width: '100%' }}
                />
              </Box>
              
              {selectedItem && (
                <Box sx={{ mt: 3, p: 2, border: '1px solid #e0e0e0', borderRadius: 1 }}>
                  <Typography variant="subtitle1" gutterBottom>
                    選択中のコンポーネント
                  </Typography>
                  <Typography variant="body2">
                    位置: X={selectedItem.x}, Y={selectedItem.y}
                  </Typography>
                  <Typography variant="body2">
                    サイズ: 幅={selectedItem.w}, 高さ={selectedItem.h}
                  </Typography>
                </Box>
              )}
            </Paper>
          )}
        </Grid>
      </Grid>
    </Box>
  );
};

export default LayoutEditor;