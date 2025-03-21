import React, { FC, useState } from 'react';
import { Box, Grid, Paper, Typography, Button } from '@mui/material';
import { ComponentSelector } from './ComponentSelector';
import { ComponentPropertyEditor } from './ComponentPropertyEditor';
import { DefaultComponent, getComponentById } from './default-components';

interface LayoutItem {
  id: string;
  componentId: string;
  props: any;
  x: number;
  y: number;
  width: number;
  height: number;
}

interface LayoutEditorProps {
  layoutId?: number;
}

export const LayoutEditor: FC<LayoutEditorProps> = ({ layoutId }) => {
  const [layoutItems, setLayoutItems] = useState<LayoutItem[]>([]);
  const [selectedItemId, setSelectedItemId] = useState<string | null>(null);
  
  // 選択されたアイテムを取得
  const selectedItem = selectedItemId 
    ? layoutItems.find(item => item.id === selectedItemId) 
    : null;
  
  // 選択されたコンポーネントを取得
  const selectedComponent = selectedItem 
    ? getComponentById(selectedItem.componentId) 
    : null;
  
  // 新しいコンポーネントを追加
  const handleAddComponent = (component: DefaultComponent) => {
    const newItem: LayoutItem = {
      id: `item-${Date.now()}`,
      componentId: component.id,
      props: { ...component.defaultProps },
      x: 0,
      y: layoutItems.length * 100, // 縦に並べる
      width: 12, // フル幅
      height: 100,
    };
    
    setLayoutItems([...layoutItems, newItem]);
    setSelectedItemId(newItem.id);
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
  
  // コンポーネントをレンダリング
  const renderComponent = (item: LayoutItem) => {
    const component = getComponentById(item.componentId);
    if (!component) return null;
    
    const Component = component.component;
    return <Component {...item.props} />;
  };
  
  // レイアウトを保存する関数
  const handleSaveLayout = () => {
    const layoutData = {
      name: 'マイレイアウト', // 実際にはユーザーに名前を入力してもらう
      items: layoutItems,
      createdAt: new Date().toISOString(),
    };
    
    // ローカルストレージに保存（実際のアプリではAPIを使用）
    const savedLayouts = JSON.parse(localStorage.getItem('savedLayouts') || '[]');
    savedLayouts.push(layoutData);
    localStorage.setItem('savedLayouts', JSON.stringify(savedLayouts));
    
    alert('レイアウトを保存しました');
  };

  // レイアウトを読み込む関数
  const handleLoadLayout = (layoutData: any) => {
    setLayoutItems(layoutData.items);
    setSelectedItemId(null);
  };
  
  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" gutterBottom>
        レイアウトエディタ
      </Typography>
      
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
            <Box sx={{ mb: 2 }}>
              <ComponentSelector onSelectComponent={handleAddComponent} />
            </Box>
            
            <Box sx={{ position: 'relative' }}>
              {layoutItems.map(item => (
                <Box
                  key={item.id}
                  sx={{
                    position: 'relative',
                    mb: 2,
                    border: selectedItemId === item.id ? '2px dashed #2196f3' : '1px solid #e0e0e0',
                    p: 1,
                    borderRadius: 1,
                    cursor: 'pointer',
                  }}
                  onClick={() => setSelectedItemId(item.id)}
                >
                  <Box 
                    sx={{ 
                      position: 'absolute', 
                      top: 5, 
                      right: 5, 
                      zIndex: 10 
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
                  {renderComponent(item)}
                </Box>
              ))}
              
              {layoutItems.length === 0 && (
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
        </Grid>
      </Grid>
      
      <Box sx={{ mt: 2, display: 'flex', justifyContent: 'flex-end' }}>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={handleSaveLayout}
        >
          レイアウトを保存
        </Button>
      </Box>
    </Box>
  );
};

export default LayoutEditor;
