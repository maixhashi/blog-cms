import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { LayoutEditor } from '../components/LayoutEditor';

const LayoutEditorPage: React.FC = () => {
  const { layoutId } = useParams<{ layoutId: string }>();
  
  return (
    <div className="layout-editor-page">
      <h1>レイアウトエディタ</h1>
      <LayoutEditor layoutId={layoutId ? parseInt(layoutId) : undefined} />
    </div>
  );
};

export default LayoutEditorPage;