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
    <Box sx={{ backgroundColor, color: textColor, py: 6 }}>
      <Container maxWidth="lg">
        <Grid container spacing={4}>
          <Grid item xs={12} md={4}>
            <Typography variant="h6" gutterBottom>
              {title}
            </Typography>
            <Typography variant="body2" paragraph>
              {description}
            </Typography>
          </Grid>
          
          {links.map((section, index) => (
            <Grid item xs={6} md={2} key={index}>
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
        
        <Box mt={5} pt={3} borderTop={`1px solid ${textColor}40`}>
          <Typography variant="body2" align="center">
            {copyright}
          </Typography>
        </Box>
      </Container>
    </Box>
  );
};
