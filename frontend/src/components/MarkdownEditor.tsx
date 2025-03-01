import React, { FC, useState, useEffect, useRef } from 'react';
import { getShortcutConfig } from '../api/shortcuts';
import { Box, Button, Typography } from '@mui/material';
import CodeMirror from '@uiw/react-codemirror';
import { markdown } from '@codemirror/lang-markdown';
import { EditorView } from '@codemirror/view';

interface MarkdownEditorProps {
  value: string;
  onChange: (value: string) => void;
}

export const MarkdownEditor: FC<MarkdownEditorProps> = ({ value, onChange }) => {
  const [shortcuts, setShortcuts] = useState<Record<string, string>>({
    'Ctrl-B': '**${selection}**',
    'Ctrl-I': '*${selection}*',
    'Ctrl-K': '[${selection}](url)',
    'Ctrl-H': '# ${selection}',
    'Ctrl-L': '* ${selection}',
  });
  
  const editorRef = useRef<EditorView | null>(null);

  useEffect(() => {
    const loadShortcuts = async () => {
      try {
        const config = await getShortcutConfig();
        if (config && config.shortcuts) {
          setShortcuts(config.shortcuts);
        }
      } catch (error) {
        console.error('ショートカット設定の読み込みに失敗しました', error);
      }
    };
    loadShortcuts();
  }, []);

  const handleKeyDown = (e: React.KeyboardEvent) => {
    const key = `${e.ctrlKey ? 'Ctrl-' : ''}${e.key.toUpperCase()}`;
    if (shortcuts[key] && editorRef.current) {
      e.preventDefault();
      const editor = editorRef.current;
      const selection = editor.state.sliceDoc(
        editor.state.selection.main.from,
        editor.state.selection.main.to
      );
      
      const template = shortcuts[key].replace('${selection}', selection || '');
      
      editor.dispatch({
        changes: {
          from: editor.state.selection.main.from,
          to: editor.state.selection.main.to,
          insert: template
        }
      });
    }
  };

  const shortcutButtons = [
    { key: 'Ctrl-B', label: '太字 (Ctrl+B)', icon: 'B' },
    { key: 'Ctrl-I', label: '斜体 (Ctrl+I)', icon: 'I' },
    { key: 'Ctrl-K', label: 'リンク (Ctrl+K)', icon: '🔗' },
    { key: 'Ctrl-H', label: '見出し (Ctrl+H)', icon: 'H' },
    { key: 'Ctrl-L', label: 'リスト (Ctrl+L)', icon: '•' },
  ];

  const handleShortcutButtonClick = (shortcutKey: string) => {
    if (!editorRef.current) return;
    
    const editor = editorRef.current;
    const selection = editor.state.sliceDoc(
      editor.state.selection.main.from,
      editor.state.selection.main.to
    );
    
    const template = shortcuts[shortcutKey].replace('${selection}', selection || '');
    
    editor.dispatch({
      changes: {
        from: editor.state.selection.main.from,
        to: editor.state.selection.main.to,
        insert: template
      }
    });
    editor.focus();
  };

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
      <Box sx={{ mb: 1, display: 'flex', flexWrap: 'wrap', gap: 1 }}>
        {shortcutButtons.map((button) => (
          <Button
            key={button.key}
            variant="outlined"
            size="small"
            onClick={() => handleShortcutButtonClick(button.key)}
            title={button.label}
          >
            {button.icon}
          </Button>
        ))}
      </Box>
      
      <Box sx={{ flexGrow: 1 }}>
        <CodeMirror
          value={value}
          onChange={onChange}
          extensions={[markdown()]}
          onKeyDown={handleKeyDown}
          height="100%"
          theme="light"
          basicSetup={{
            lineNumbers: true,
            foldGutter: true,
            highlightActiveLine: true,
          }}
          onCreateEditor={(view) => {
            editorRef.current = view;
          }}
        />
      </Box>
      
      <Box sx={{ mt: 1 }}>
        <Typography variant="caption" color="text.secondary">
          ヒント: Ctrl+B (太字), Ctrl+I (斜体), Ctrl+K (リンク), Ctrl+H (見出し), Ctrl+L (リスト)
        </Typography>
      </Box>
    </Box>
  );
};