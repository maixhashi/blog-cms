import React, { FC, useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { MarkdownEditor } from '../components/MarkdownEditor';
import { ArticlePreview } from '../components/ArticlePreview';
import { fetchArticle, createArticle, updateArticle } from '../api/articles';
import { Button, TextField, Box, Stack, Paper, useTheme, useMediaQuery } from '@mui/material';

export const ArticleEditorPage: FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [tags, setTags] = useState<string[]>([]);
  const [isPublished, setIsPublished] = useState(false);
  const [isSaving, setIsSaving] = useState(false);
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));

  useEffect(() => {
    if (id) {
      // 既存の記事を編集する場合
      const loadArticle = async () => {
        try {
          const article = await fetchArticle(id);
          setTitle(article.title);
          setContent(article.content);
          setTags(article.tags);
          setIsPublished(article.published);
        } catch (error) {
          console.error('記事の読み込みに失敗しました', error);
        }
      };
      loadArticle();
    }
  }, [id]);

  const handleSave = async (publish: boolean = false) => {
    setIsSaving(true);
    try {
      const articleData = {
        title,
        content,
        tags,
        published: publish || isPublished
      };

      if (id) {
        await updateArticle(id, articleData);
      } else {
        const newArticle = await createArticle(articleData);
        navigate(`/editor/${newArticle.id}`);
      }
      
      setIsPublished(publish || isPublished);
    } catch (error) {
      console.error('記事の保存に失敗しました', error);
    } finally {
      setIsSaving(false);
    }
  };

  return (
    <div className="container">
      <h1>{id ? '記事を編集' : '新規記事作成'}</h1>
      
      <Box mb={3}>
        <TextField
          fullWidth
          label="タイトル"
          value={title}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setTitle(e.target.value)}
          variant="outlined"
        />
      </Box>
      
      <Stack 
        direction={isMobile ? 'column' : 'row'} 
        spacing={2} 
        sx={{ width: '100%' }}
      >
        <Paper elevation={3} sx={{ p: 2, height: '600px', overflow: 'auto', flex: 1 }}>
          <MarkdownEditor
            value={content}
            onChange={setContent}
          />
        </Paper>
        <Paper elevation={3} sx={{ p: 2, height: '600px', overflow: 'auto', flex: 1 }}>
          <ArticlePreview markdown={content} title={title} />
        </Paper>
      </Stack>
      
      <Box mt={3} display="flex" justifyContent="space-between">
        <Button
          variant="outlined"
          onClick={() => navigate(-1)}
          disabled={isSaving}
        >
          キャンセル
        </Button>
        <div>
          <Button
            variant="outlined"
            onClick={() => handleSave(false)}
            disabled={isSaving}
            sx={{ mr: 1 }}
          >
            下書き保存
          </Button>
          <Button
            variant="contained"
            color="primary"
            onClick={() => handleSave(true)}
            disabled={isSaving}
          >
            {isPublished ? '更新して公開' : '公開する'}
          </Button>
        </div>
      </Box>
    </div>
  );
};