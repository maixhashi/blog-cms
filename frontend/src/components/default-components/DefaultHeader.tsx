import React, { FC } from 'react';
import { Box, Typography, AppBar, Toolbar, Button } from '@mui/material';

interface DefaultHeaderProps {
  title?: string;
  links?: Array<{ label: string; url: string }>;
  backgroundColor?: string;
  textColor?: string;
}

export const DefaultHeader: FC<DefaultHeaderProps> = ({
  title = 'ブログタイトル',
  links = [
    { label: 'ホーム', url: '/' },
    { label: '記事一覧', url: '/articles' },
    { label: 'カテゴリー', url: '/categories' },
    { label: '問い合わせ', url: '/contact' },
  ],
  backgroundColor = '#2196f3',
  textColor = '#ffffff',
}) => {
  return (
    <AppBar position="static" sx={{ backgroundColor, width: '100%', height: '100%', boxSizing: 'border-box' }}>
      <Toolbar sx={{ height: '100%' }}>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1, color: textColor }}>
          {title}
        </Typography>
        <Box>
          {links.map((link, index) => (
            <Button key={index} color="inherit" sx={{ color: textColor }}>
              {link.label}
            </Button>
          ))}
        </Box>
      </Toolbar>
    </AppBar>
  );
};