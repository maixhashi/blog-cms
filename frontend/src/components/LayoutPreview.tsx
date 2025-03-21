import React, { FC } from 'react';
import { Box, Paper, Typography, Button } from '@mui/material';
import { getComponentById } from './default-components';

interface LayoutPreviewProps {
  layoutItems: any[];
  onClose: () => void;
}

export const LayoutPreview: FC<LayoutPreviewProps> = ({ layoutItems, onClose }) => {
  // コンポーネントをレンダリング
  const renderComponent = (item: any) => {
    const component = getComponentById(item.componentId);
    if (!component) return null;
    
    const Component = component.component;
    return <Component {...item.props} />;
  };
  
  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 2 }}>
        <Typography variant="h5">レイアウトプレビュー</Typography>
        <Button variant="outlined" onClick={onClose}>
          エディタに戻る
        </Button>
      </Box>
      
      <Paper 
        sx={{ 
          p: 2, 
          minHeight: 600, 
          backgroundColor: '#ffffff',
        }}
      >
        {layoutItems.map(item => (
          <Box
            key={item.id}
            sx={{ mb: 2 }}
          >
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
              レイアウトが空です。コンポーネントを追加してください。
            </Typography>
          </Box>
        )}
      </Paper>
    </Box>
  );
};
