import React, { FC } from 'react';
import { Card, CardContent, CardMedia, Typography, Box, Chip, Button } from '@mui/material';

interface DefaultArticleCardProps {
  title?: string;
  excerpt?: string;
  imageUrl?: string;
  date?: string;
  author?: string;
  tags?: string[];
  primaryColor?: string;
  secondaryColor?: string;
  showImage?: boolean;
  showTags?: boolean;
}

export const DefaultArticleCard: FC<DefaultArticleCardProps> = ({
  title = '記事タイトル',
  excerpt = 'これは記事の抜粋です。ここに記事の概要が表示されます。',
  imageUrl = 'https://source.unsplash.com/random/300x200/?blog',
  date = '2023年1月1日',
  author = '著者名',
  tags = ['タグ1', 'タグ2'],
  primaryColor = '#2196f3',
  secondaryColor = '#bbdefb',
  showImage = true,
  showTags = true,
}) => {
  return (
    <Card sx={{ width: '100%', height: '100%', display: 'flex', flexDirection: 'column' }}>
      {showImage && (
        <CardMedia
          component="img"
          height="140"
          image={imageUrl}
          alt={title}
          sx={{ objectFit: 'cover' }}
        />
      )}
      <CardContent sx={{ flexGrow: 1 }}>
        <Typography gutterBottom variant="h5" component="div">
          {title}
        </Typography>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 1 }}>
          <Typography variant="body2" color="text.secondary">
            {date}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            {author}
          </Typography>
        </Box>
        <Typography variant="body2" color="text.secondary" paragraph>
          {excerpt}
        </Typography>
        {showTags && tags.length > 0 && (
          <Box sx={{ mt: 2, mb: 2 }}>
            {tags.map((tag, index) => (
              <Chip
                key={index}
                label={tag}
                size="small"
                sx={{ mr: 0.5, mb: 0.5, backgroundColor: secondaryColor, color: 'rgba(0, 0, 0, 0.7)' }}
              />
            ))}
          </Box>
        )}
        <Button 
          variant="contained" 
          size="small"
          sx={{ backgroundColor: primaryColor, '&:hover': { backgroundColor: primaryColor } }}
        >
          続きを読む
        </Button>
      </CardContent>
    </Card>
  );
};