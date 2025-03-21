import React, { FC, useState, useEffect } from 'react';
import { 
  Box, 
  Typography, 
  TextField, 
  Button, 
  Accordion, 
  AccordionSummary, 
  AccordionDetails,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Chip,
  IconButton,
  Grid,
} from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import AddIcon from '@mui/icons-material/Add';
import DeleteIcon from '@mui/icons-material/Delete';
import { DefaultComponent } from './default-components';

interface ComponentPropertyEditorProps {
  component: DefaultComponent;
  currentProps: any;
  onPropsChange: (newProps: any) => void;
}

export const ComponentPropertyEditor: FC<ComponentPropertyEditorProps> = ({
  component,
  currentProps,
  onPropsChange,
}) => {
  const [props, setProps] = useState<any>(currentProps || component.defaultProps);

  useEffect(() => {
    setProps(currentProps || component.defaultProps);
  }, [component, currentProps]);

  const handleTextChange = (key: string, value: string) => {
    const newProps = { ...props, [key]: value };
    setProps(newProps);
    onPropsChange(newProps);
  };

  const handleColorChange = (key: string, value: string) => {
    const newProps = { ...props, [key]: value };
    setProps(newProps);
    onPropsChange(newProps);
  };

  const handleNumberChange = (key: string, value: string) => {
    const numberValue = value === '' ? '' : Number(value);
    const newProps = { ...props, [key]: numberValue };
    setProps(newProps);
    onPropsChange(newProps);
  };

  const handleArrayItemChange = (key: string, index: number, field: string, value: string) => {
    const newArray = [...props[key]];
    newArray[index] = { ...newArray[index], [field]: value };
    
    const newProps = { ...props, [key]: newArray };
    setProps(newProps);
    onPropsChange(newProps);
  };

  const handleAddArrayItem = (key: string, template: any) => {
    const newArray = [...(props[key] || []), { ...template }];
    const newProps = { ...props, [key]: newArray };
    setProps(newProps);
    onPropsChange(newProps);
  };

  const handleRemoveArrayItem = (key: string, index: number) => {
    const newArray = [...props[key]];
    newArray.splice(index, 1);
    
    const newProps = { ...props, [key]: newArray };
    setProps(newProps);
    onPropsChange(newProps);
  };

  // プロパティの型に基づいて適切な入力フィールドを生成
  const renderPropertyField = (key: string, value: any) => {
    // 値の型に基づいて適切なエディタを表示
    if (typeof value === 'string') {
      // 色の場合はカラーピッカーを表示
      if (key.toLowerCase().includes('color')) {
        return (
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
            <TextField
              fullWidth
              label={key}
              value={value}
              onChange={(e) => handleColorChange(key, e.target.value)}
              margin="normal"
            />
            <input
              type="color"
              value={value}
              onChange={(e) => handleColorChange(key, e.target.value)}
              style={{ width: 40, height: 40, padding: 0, border: 'none' }}
            />
          </Box>
        );
      }
      
      // 通常のテキストフィールド
      return (
        <TextField
          fullWidth
          label={key}
          value={value}
          onChange={(e) => handleTextChange(key, e.target.value)}
          margin="normal"
        />
      );
    }
    
    // 数値の場合
    if (typeof value === 'number') {
      return (
        <TextField
          fullWidth
          label={key}
          type="number"
          value={value}
          onChange={(e) => handleNumberChange(key, e.target.value)}
          margin="normal"
        />
      );
    }
    
    // 配列の場合（リンクやアイテムなど）
    if (Array.isArray(value)) {
      return (
        <Box sx={{ mt: 2, mb: 2 }}>
          <Typography variant="subtitle1" sx={{ mb: 1 }}>
            {key}
          </Typography>
          
          {value.map((item, index) => {
            if (typeof item === 'object') {
              // オブジェクトの配列（リンクなど）
              return (
                <Box 
                  key={index} 
                  sx={{ 
                    display: 'flex', 
                    alignItems: 'center', 
                    mb: 1,
                    p: 1,
                    border: '1px solid #e0e0e0',
                    borderRadius: 1,
                  }}
                >
                  <Grid container spacing={2} sx={{ flexGrow: 1 }}>
                    {Object.entries(item).map(([itemKey, itemValue]) => (
                      <Grid item xs={12} sm={6} key={itemKey}>
                        <TextField
                          fullWidth
                          size="small"
                          label={itemKey}
                          value={itemValue as string}
                          onChange={(e) => handleArrayItemChange(key, index, itemKey, e.target.value)}
                        />
                      </Grid>
                    ))}
                  </Grid>
                  <IconButton 
                    color="error" 
                    onClick={() => handleRemoveArrayItem(key, index)}
                    sx={{ ml: 1 }}
                  >
                    <DeleteIcon />
                  </IconButton>
                </Box>
              );
            }
            
            // プリミティブ値の配列
            return (
              <Box 
                key={index} 
                sx={{ 
                  display: 'flex', 
                  alignItems: 'center', 
                  mb: 1 
                }}
              >
                <TextField
                  fullWidth
                  size="small"
                  value={item}
                  onChange={(e) => {
                    const newArray = [...value];
                    newArray[index] = e.target.value;
                    const newProps = { ...props, [key]: newArray };
                    setProps(newProps);
                    onPropsChange(newProps);
                  }}
                />
                <IconButton 
                  color="error" 
                  onClick={() => handleRemoveArrayItem(key, index)}
                  sx={{ ml: 1 }}
                >
                  <DeleteIcon />
                </IconButton>
              </Box>
            );
          })}
          
          <Button
            startIcon={<AddIcon />}
            onClick={() => {
              // 配列の最初の要素をテンプレートとして使用
              const template = value.length > 0 
                ? (typeof value[0] === 'object' ? { ...value[0] } : '')
                : (key === 'links' ? { label: '', url: '' } : '');
              
              handleAddArrayItem(key, template);
            }}
            variant="outlined"
            size="small"
            sx={{ mt: 1 }}
          >
            追加
          </Button>
        </Box>
      );
    }
    
    // その他のタイプ（未対応）
    return (
      <Typography color="error">
        未対応のプロパティタイプ: {typeof value}
      </Typography>
    );
  };

  return (
    <Box sx={{ mt: 2 }}>
      <Typography variant="h6" gutterBottom>
        {component.name} のプロパティ
      </Typography>
      
      <Accordion defaultExpanded>
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Typography>基本設定</Typography>
        </AccordionSummary>
        <AccordionDetails>
          {Object.entries(props).map(([key, value]) => (
            <Box key={key}>
              {renderPropertyField(key, value)}
            </Box>
          ))}
        </AccordionDetails>
      </Accordion>
      
      <Box sx={{ mt: 2, display: 'flex', justifyContent: 'flex-end' }}>
        <Button 
          variant="outlined" 
          onClick={() => {
            setProps(component.defaultProps);
            onPropsChange(component.defaultProps);
          }}
          sx={{ mr: 1 }}
        >
          デフォルトに戻す
        </Button>
      </Box>
    </Box>
  );
};
