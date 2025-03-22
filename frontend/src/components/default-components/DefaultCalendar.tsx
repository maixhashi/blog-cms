import React, { FC, useState } from 'react';
import { Box, Typography, Grid, Paper } from '@mui/material';

interface DefaultCalendarProps {
  initialDate?: Date;
  highlightedDates?: Date[];
  primaryColor?: string;
  secondaryColor?: string;
}

export const DefaultCalendar: FC<DefaultCalendarProps> = ({
  initialDate = new Date(),
  highlightedDates = [],
  primaryColor = '#2196f3',
  secondaryColor = '#bbdefb',
}) => {
  const [currentDate, setCurrentDate] = useState(initialDate);
  
  // 月の最初の日と最後の日を取得
  const firstDayOfMonth = new Date(currentDate.getFullYear(), currentDate.getMonth(), 1);
  const lastDayOfMonth = new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 0);
  
  // 月の日数を取得
  const daysInMonth = lastDayOfMonth.getDate();
  
  // 月の最初の日の曜日を取得（0: 日曜日, 1: 月曜日, ...）
  const firstDayOfWeek = firstDayOfMonth.getDay();
  
  // カレンダーの日付を生成
  const days = [];
  for (let i = 0; i < firstDayOfWeek; i++) {
    days.push(null); // 月の最初の日より前の空白
  }
  
  for (let i = 1; i <= daysInMonth; i++) {
    days.push(i);
  }
  
  // 前月・次月に移動する関数
  const prevMonth = () => {
    setCurrentDate(new Date(currentDate.getFullYear(), currentDate.getMonth() - 1, 1));
  };
  
  const nextMonth = () => {
    setCurrentDate(new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 1));
  };
  
  // 日付がハイライトされているかチェック
  const isHighlighted = (day: number) => {
    if (!day) return false;
    
    return highlightedDates.some(date => 
      date.getDate() === day && 
      date.getMonth() === currentDate.getMonth() && 
      date.getFullYear() === currentDate.getFullYear()
    );
  };
  
  return (
    <Paper elevation={3} sx={{ 
      p: 2, 
      width: '100%', 
      height: '100%', 
      boxSizing: 'border-box',
      overflow: 'auto'
    }}>
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
        <Typography 
          variant="button" 
          onClick={prevMonth} 
          sx={{ cursor: 'pointer', color: primaryColor }}
        >
          ＜
        </Typography>
        <Typography variant="h6">
          {currentDate.getFullYear()}年{currentDate.getMonth() + 1}月
        </Typography>
        <Typography 
          variant="button" 
          onClick={nextMonth} 
          sx={{ cursor: 'pointer', color: primaryColor }}
        >
          ＞
        </Typography>
      </Box>
      
      <Grid container spacing={1}>
        {['日', '月', '火', '水', '木', '金', '土'].map((day, index) => (
          <Grid item key={index} xs={1.7}>
            <Box 
              textAlign="center" 
              fontWeight="bold"
              color={index === 0 ? 'error.main' : index === 6 ? 'primary.main' : 'text.primary'}
            >
              {day}
            </Box>
          </Grid>
        ))}
        
        {days.map((day, index) => (
          <Grid item key={index} xs={1.7}>
            {day && (
              <Box 
                textAlign="center" 
                p={1}
                borderRadius={1}
                bgcolor={isHighlighted(day) ? secondaryColor : 'transparent'}
                sx={{ 
                  cursor: 'pointer',
                  '&:hover': { bgcolor: primaryColor, color: 'white' }
                }}
              >
                {day}
              </Box>
            )}
          </Grid>
        ))}
      </Grid>
    </Paper>
  );
};