import React, { FC } from 'react';
import ReactMarkdown from 'react-markdown';
import { Box, Typography } from '@mui/material';
import remarkGfm from 'remark-gfm';
import rehypeRaw from 'rehype-raw';
import rehypeSanitize from 'rehype-sanitize';

interface ArticlePreviewProps {
  title: string;
  markdown: string;
}

export const ArticlePreview: FC<ArticlePreviewProps> = ({ title, markdown }) => {
  return (
    <Box sx={{ height: '100%', overflow: 'auto' }}>
      <Typography variant="h4" component="h1" gutterBottom>
        {title || '(タイトルなし)'}
      </Typography>
      
      <Box className="markdown-preview">
        <ReactMarkdown
          remarkPlugins={[remarkGfm]}
          rehypePlugins={[rehypeRaw, rehypeSanitize]}
        >
          {markdown || '*プレビューはここに表示されます*'}
        </ReactMarkdown>
      </Box>
    </Box>
  );
};
