import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Box, Button, Typography, Breadcrumbs, Link } from '@mui/material';
import { LayoutEditor } from '../components/LayoutEditor';

const LayoutEditorPage: React.FC = () => {
  const { layoutId } = useParams<{ layoutId: string }>();
  const navigate = useNavigate();
  const [layoutName, setLayoutName] = useState<string>('');
  
  // レイアウト名を取得
  useEffect(() => {
    if (layoutId) {
      // ローカルストレージからレイアウト情報を取得
      const savedLayouts = JSON.parse(localStorage.getItem('savedLayouts') || '[]');
      const layout = savedLayouts.find((l: any) => l.id === parseInt(layoutId));
      
      if (layout) {
        setLayoutName(layout.name);
      } else {
        // レイアウトが見つからない場合
        setLayoutName('不明なレイアウト');
      }
    } else {
      setLayoutName('新規レイアウト');
    }
  }, [layoutId]);
  
  return (
    <Box className="layout-editor-page" sx={{ p: 2 }}>
      {/* パンくずリスト */}
      <Breadcrumbs sx={{ mb: 2 }}>
        <Link 
          color="inherit" 
          href="#" 
          onClick={(e) => {
            e.preventDefault();
            navigate('/layout-manager');
          }}
        >
          レイアウト管理
        </Link>
        <Typography color="text.primary">
          {layoutId ? `${layoutName} の編集` : '新規レイアウト作成'}
        </Typography>
      </Breadcrumbs>
      
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">
          {layoutId ? `レイアウト編集: ${layoutName}` : '新規レイアウト作成'}
        </Typography>
        <Button 
          variant="outlined" 
          onClick={() => navigate('/layout-manager')}
        >
          レイアウト一覧に戻る
        </Button>
      </Box>
      
      <Box sx={{ bgcolor: 'background.paper', borderRadius: 1, boxShadow: 1 }}>
        <LayoutEditor layoutId={layoutId ? parseInt(layoutId) : undefined} />
      </Box>
    </Box>
  );
};

export default LayoutEditorPage;