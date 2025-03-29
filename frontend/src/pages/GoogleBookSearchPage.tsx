import { useState, FormEvent, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { 
  Container, 
  Typography, 
  TextField, 
  Button, 
  Grid, 
  Card, 
  CardContent, 
  CardMedia, 
  CardActions,
  Box,
  CircularProgress,
  Divider,
  Paper,
  IconButton,
  InputAdornment
} from '@mui/material'
import SearchIcon from '@mui/icons-material/Search'
import ArrowBackIcon from '@mui/icons-material/ArrowBack'
import AddIcon from '@mui/icons-material/Add'
import { useQueryGoogleBooks } from '../hooks/useQueryGoogleBooks'
import { useMutateBook } from '../hooks/useMutateBook'
import useStore from '../store'
import { GoogleBookVolume, GoogleBookSearchResponse } from '../types/models/googleBook'

export const GoogleBookSearchPage = () => {
  const navigate = useNavigate()
  const [searchEnabled, setSearchEnabled] = useState(false)
  const { 
    searchQuery, 
    updateSearchQuery, 
    setSearchResults, 
    searchResults, 
    selectedBook, 
    selectBook 
  } = useStore(state => state)
  const { importGoogleBookMutation } = useMutateBook()
  
  // 検索クエリフック
  const { data: searchData, isLoading, isError } = useQueryGoogleBooks(searchEnabled)
  
  // 検索結果を更新
  useEffect(() => {
    if (searchData && searchEnabled) {
      // 型アサーションを使用して、searchDataがGoogleBookSearchResponseであることを明示
      const responseData = searchData as GoogleBookSearchResponse;
      setSearchResults(responseData.items || [])
      setSearchEnabled(false)
    }
  }, [searchData, setSearchResults, searchEnabled])
  
  // 検索実行
  const handleSearch = (e: FormEvent) => {
    e.preventDefault()
    if (searchQuery.trim()) {
      setSearchEnabled(true)
    }
  }
  
  // 書籍インポート
  const handleImportBook = (googleBookId: string) => {
    importGoogleBookMutation.mutate(googleBookId)
  }
  
  // 書籍詳細表示
  const handleSelectBook = (book: GoogleBookVolume) => {
    selectBook(book)
  }
  
  // 書籍詳細をクリア
  const handleClearSelection = () => {
    selectBook(null)
  }
  
  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Button 
        startIcon={<ArrowBackIcon />} 
        onClick={() => navigate('/route-map')}
        sx={{ mb: 2 }}
      >
        戻る
      </Button>
      
      <Typography variant="h4" component="h1" gutterBottom>
        Google Books 検索
      </Typography>
      
      <Paper component="form" onSubmit={handleSearch} sx={{ p: 2, mb: 4 }}>
        <TextField
          fullWidth
          label="書籍を検索"
          variant="outlined"
          value={searchQuery}
          onChange={(e) => updateSearchQuery(e.target.value)}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton type="submit" edge="end">
                  <SearchIcon />
                </IconButton>
              </InputAdornment>
            ),
          }}
        />
      </Paper>
      
      {isLoading && (
        <Box display="flex" justifyContent="center" my={4}>
          <CircularProgress />
        </Box>
      )}
      
      {isError && (
        <Box textAlign="center" my={4}>
          <Typography color="error" gutterBottom>
            検索中にエラーが発生しました。もう一度お試しください。
          </Typography>
          <Button 
            variant="contained" 
            color="primary" 
            onClick={() => setSearchEnabled(true)}
            sx={{ mt: 2 }}
          >
            再試行
          </Button>
        </Box>
      )}
      
      <Grid container spacing={3}>
        {/* 検索結果リスト */}
        <Grid item xs={12} md={selectedBook ? 6 : 12}>
          {searchResults && searchResults.length > 0 ? (
            <>
              <Typography variant="h6" gutterBottom>
                検索結果: {searchResults.length}件
              </Typography>
              <Grid container spacing={2}>
                {searchResults.map((book, index) => (
                  <Grid item xs={12} sm={6} md={selectedBook ? 6 : 4} key={book.id || index}>
                    <Card 
                      sx={{ 
                        height: '100%', 
                        display: 'flex', 
                        flexDirection: 'column',
                        cursor: 'pointer',
                        bgcolor: selectedBook?.id === book.id ? 'action.selected' : 'background.paper'
                      }}
                      onClick={() => handleSelectBook(book)}
                    >
                      <CardMedia
                        component="img"
                        sx={{ 
                          height: 140, 
                          objectFit: 'contain',
                          p: 1,
                          bgcolor: 'grey.100'
                        }}
                        image={book.image_url || '/placeholder-book.png'}
                        alt={book.title || '書籍タイトル'}
                      />
                      <CardContent sx={{ flexGrow: 1 }}>
                        <Typography gutterBottom variant="h6" component="div" noWrap>
                          {book.title || 'タイトルなし'}
                        </Typography>
                        <Typography variant="body2" color="text.secondary" noWrap>
                          {book.authors?.join(', ') || '著者不明'}
                        </Typography>
                      </CardContent>
                      <CardActions>
                        <Button 
                          size="small" 
                          startIcon={<AddIcon />}
                          onClick={(e) => {
                            e.stopPropagation();
                            handleImportBook(book.id);
                          }}
                          disabled={importGoogleBookMutation.isLoading}
                        >
                          インポート
                        </Button>
                      </CardActions>
                    </Card>
                  </Grid>
                ))}
              </Grid>
            </>
          ) : searchEnabled && !isLoading ? (
            <Typography align="center">
              検索結果がありません。別のキーワードで試してください。
            </Typography>
          ) : null}
        </Grid>
        
        {/* 書籍詳細 */}
        {selectedBook && (
          <Grid item xs={12} md={6}>
            <Paper sx={{ p: 3, height: '100%' }}>
              <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
                <Typography variant="h5">書籍詳細</Typography>
                <IconButton onClick={handleClearSelection} size="small">
                  <ArrowBackIcon />
                </IconButton>
              </Box>
              
              <Divider sx={{ mb: 2 }} />
              
              <Box display="flex" mb={3}>
                <CardMedia
                  component="img"
                  sx={{ 
                    width: 128, 
                    height: 192, 
                    objectFit: 'contain',
                    mr: 2,
                    bgcolor: 'grey.100'
                  }}
                  image={selectedBook.image_url || '/placeholder-book.png'}
                  alt={selectedBook.title || '書籍タイトル'}
                />
                <Box>
                  <Typography variant="h6" gutterBottom>
                    {selectedBook.title || 'タイトルなし'}
                  </Typography>
                  <Typography variant="body1" gutterBottom>
                    著者: {selectedBook.authors?.join(', ') || '著者不明'}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    ISBN: {selectedBook.isbn || '不明'}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    出版日: {selectedBook.published_date || '不明'}
                  </Typography>
                </Box>
              </Box>
              
              <Typography variant="body1" paragraph>
                {selectedBook.description || '説明はありません。'}
              </Typography>
              
              <Button 
                variant="contained" 
                color="primary" 
                startIcon={<AddIcon />}
                onClick={() => handleImportBook(selectedBook.id)}
                disabled={importGoogleBookMutation.isLoading}
                fullWidth
              >
                {importGoogleBookMutation.isLoading ? '処理中...' : 'この書籍をインポート'}
              </Button>
            </Paper>
          </Grid>
        )}
      </Grid>
    </Container>
  )
}
