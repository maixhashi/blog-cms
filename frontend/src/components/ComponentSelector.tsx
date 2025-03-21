import React, { FC, useState } from 'react';
import { 
  Box, 
  Typography, 
  Grid, 
  Card, 
  CardContent, 
  CardMedia, 
  Button, 
  Dialog, 
  DialogTitle, 
  DialogContent, 
  DialogActions 
} from '@mui/material';
import { defaultComponents, DefaultComponent } from './default-components';

interface ComponentSelectorProps {
  onSelectComponent: (component: DefaultComponent) => void;
}

export const ComponentSelector: FC<ComponentSelectorProps> = ({ onSelectComponent }) => {
  const [open, setOpen] = useState(false);
  const [selectedComponent, setSelectedComponent] = useState<DefaultComponent | null>(null);

  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleSelect = (component: DefaultComponent) => {
    setSelectedComponent(component);
  };

  const handleConfirm = () => {
    if (selectedComponent) {
      onSelectComponent(selectedComponent);
      handleClose();
    }
  };

  return (
    <>
      <Button variant="contained" color="primary" onClick={handleOpen}>
        コンポーネントを追加
      </Button>

      <Dialog open={open} onClose={handleClose} maxWidth="md" fullWidth>
        <DialogTitle>コンポーネントを選択</DialogTitle>
        <DialogContent>
          <Grid container spacing={3}>
            {defaultComponents.map((component) => (
              <Grid item xs={12} sm={6} md={4} key={component.id}>
                <Card 
                  sx={{ 
                    cursor: 'pointer',
                    border: selectedComponent?.id === component.id ? '2px solid #2196f3' : 'none',
                  }}
                  onClick={() => handleSelect(component)}
                >
                  <CardMedia
                    component="img"
                    height="140"
                    image={component.thumbnail}
                    alt={component.name}
                  />
                  <CardContent>
                    <Typography gutterBottom variant="h6" component="div">
                      {component.name}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      {component.description}
                    </Typography>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>キャンセル</Button>
          <Button 
            onClick={handleConfirm} 
            variant="contained" 
            disabled={!selectedComponent}
          >
            選択
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};
