import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import LayoutEditor from '../components/LayoutEditor';

const LayoutEditorPage: React.FC = () => {
  const { layoutId } = useParams<{ layoutId: string }>();
  
  if (!layoutId) {
    return <div>レイアウトIDが指定されていません</div>;
  }

  return (
    <div className="layout-editor-page">
      <h1>レイアウトエディタ</h1>
      <LayoutEditor layoutId={parseInt(layoutId)} />
    </div>
  );
};

export default LayoutEditorPage;
