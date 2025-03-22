import React, { FC } from 'react';
import { Box, Typography, Container, Grid, Link } from '@mui/material';

interface DefaultFooterProps {
  title?: string;
  description?: string;
  links?: Array<{ section: string; items: Array<{ label: string; url: string }> }>;
  copyright?: string;
  backgroundColor?: string;
  textColor?: string;
}

export const DefaultFooter: FC<DefaultFooterProps> = ({
  title = 'ブログタイトル',
  description = 'ブログの説明文をここに入力します。',
  links = [
    {
      section: 'カテゴリー',
      items: [
        { label: 'テクノロジー', url: '/category/tech' },
        { label: 'ライフスタイル', url: '/category/lifestyle' },
        { label: '旅行', url: '/category/travel' },
      ],
    },
    {
      section: 'リンク',
      items: [
        { label: 'ホーム', url: '/' },
        { label: '記事一覧', url: '/articles' },
        { label: '問い合わせ', url: '/contact' },
      ],
    },
  ],
  copyright = '© 2023 ブログタイトル. All rights reserved.',
  backgroundColor = '#333333',
  textColor = '#ffffff',
}) => {
  return (
    <Box 
      sx={{ 
        backgroundColor, 
        color: textColor, 
        width: '100%', 
        height: '100%', 
        boxSizing: 'border-box',
        overflow: 'auto',
        display: 'flex',
        flexDirection: 'column',
        position: 'relative', // 追加: 絶対位置指定のための相対位置
      }}
      className="react-draggable-handle" // 追加: ドラッグハンドルのクラス
    >
      <Box sx={{ flex: 1, p: 2 }}>
        <Grid container spacing={2}>
          <Grid item xs={12} md={4}>
            <Typography variant="h6" gutterBottom>
              {title}
            </Typography>
            <Typography variant="body2" paragraph>
              {description}
            </Typography>
          </Grid>
          
          {links.map((section, index) => (
            <Grid item xs={6} md={4} key={index}>
              <Typography variant="h6" gutterBottom>
                {section.section}
              </Typography>
              <Box>
                {section.items.map((item, idx) => (
                  <Link
                    key={idx}
                    href={item.url}
                    color="inherit"
                    sx={{ display: 'block', mb: 1, textDecoration: 'none', '&:hover': { textDecoration: 'underline' } }}
                  >
                    {item.label}
                  </Link>
                ))}
              </Box>
            </Grid>
          ))}
        </Grid>
      </Box>
      
      <Box sx={{ 
        p: 2, 
        borderTop: `1px solid ${textColor}40`,
        mt: 'auto'
      }}>
        <Typography variant="body2" align="center">
          {copyright}
        </Typography>
      </Box>
      
      {/* 追加: ドラッグハンドルの視覚的インジケーター（オプション） */}
      <Box
        sx={{
          position: 'absolute',
          top: 0,
          right: 0,
          width: '20px',
          height: '20px',
          backgroundColor: 'rgba(255, 255, 255, 0.3)',
          cursor: 'move',
          borderBottomLeftRadius: '4px',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          fontSize: '12px',
          color: 'rgba(255, 255, 255, 0.7)',
        }}
      >
        ⋮⋮
      </Box>
    </Box>
  );
};