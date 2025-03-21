import React, { FC } from 'react';
import { Box, Typography, List, ListItem, ListItemText, Divider } from '@mui/material';

interface DefaultSidebarProps {
  title?: string;
  items?: Array<{ label: string; url: string }>;
  backgroundColor?: string;
  textColor?: string;
  width?: number | string;
}

export const DefaultSidebar: FC<DefaultSidebarProps> = ({
  title = 'サイドバー',
  items = [
    { label: '最新の記事', url: '/latest' },
    { label: '人気の記事', url: '/popular' },
    { label: 'カテゴリー1', url: '/category/1' },
    { label: 'カテゴリー2', url: '/category/2' },
    { label: 'アーカイブ', url: '/archive' },
  ],
  backgroundColor = '#f5f5f5',
  textColor = '#333333',
  width = 250,
}) => {
  return (
    <Box sx={{ width, backgroundColor, p: 2 }}>
      <Typography variant="h6" sx={{ mb: 2, color: textColor }}>
        {title}
      </Typography>
      <Divider />
      <List>
        {items.map((item, index) => (
          <ListItem key={index} sx={{ color: textColor, cursor: 'pointer' }}>
            <ListItemText primary={item.label} />
          </ListItem>
        ))}
      </List>
    </Box>
  );
};
