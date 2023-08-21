import React, { useState, useEffect } from 'react';
import { Container, Box, List, ListItem, ListItemAvatar, CardActions,IconButton, ListItemText, Avatar, Grid, Card, CardMedia, CardContent, Typography } from '@mui/material';
import Navbar from './components/Navbar';
import UserNavbar from './components/UserNavbar';
import axios from 'axios';
import { ShoppingCart, Favorite } from '@mui/icons-material';

function App() {
  const [data, setData] = useState([]);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [formData, setFormData] = useState({
    url: '',
    description: '',
    price: '',
    tags: '',
  });
  const [dataList, setDataList] = useState([]);

  useEffect(() => {
    handleListAllPhotos();
  }, []);

  useEffect(() => {
    handleAddCartPhoto();
  }, []);

  const BASE_URL = 'http://localhost:3001';

  const fetchApiData = async () => {
    try {
      const response = await axios.get('http://localhost:3001/status');
      setData(response.data);
      if (localStorage.getItem("userId") != null && localStorage.getItem("username") != null) {
        setIsLoggedIn(true);
      } else {
        setIsLoggedIn(false);
      }
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchApiData();
  }, []);

  const handleListAllPhotos = async () => {
    try {
      const resp = await axios.post(`${BASE_URL}/ListAllPhotosModel`, {});
      console.log(resp)
      setDataList(resp.data.list);
    } catch (error) {
      console.log('Error:', error);
    }
  };

  const handleAddCartPhoto = async (item) => {
    try {
      var userid = localStorage.getItem("userId");
      const resp = await axios.post(`${BASE_URL}/AddToCartModel`, {
        userid: userid,
        photo: {
          Description: item.description,
          UserID: item.userid,
          URL: item.url,
          Tags: item.tags,
          Price: item.price,
        },
      });

      handleListAllPhotos();
    } catch (error) {
      console.log('Error:', error);
    }
  };

  const handleAddFavorite = async (item) => {
    try {
      var userid = localStorage.getItem("userId");
      const resp = await axios.post(`${BASE_URL}/AddToFavoritesModel`, {
        userid: userid,
        photo: {
          Description: item.description,
          UserID: item.userid,
          URL: item.url,
          Tags: item.tags,
          Price: item.price,
        },
      });

      handleListAllPhotos();
    } catch (error) {
      console.log('Error:', error);
    }
  };

  const handleTextFieldChange = async (value) => {
    try {
      const field = value.split(',').map(item => item.trim());
      const resp = await axios.post(`${BASE_URL}/SearchPhotosByTag`, {
        field,
      });
      console.log(resp);
      if (resp.data.photos) {
        setDataList(resp.data.photos);
      }
    } catch (error) {
      console.log('Error:', error);
    }
  };

  return (
    <div>
      {isLoggedIn ? (
        <UserNavbar />
      ) : (
        <Navbar onTextFieldChange={handleTextFieldChange} />
      )}
<Container>
  <Box sx={{ marginTop: 4 }}>
    {dataList !== undefined && dataList != null ? (
      <Grid container spacing={2}>
        {dataList.map((data) => (
          <Grid item xs={12} sm={6} md={3} key={data.id}>
            <Card>
              <CardMedia component="img" height="200" image={`data:image/png;base64, ${data.url}`} alt={data.description} />
              <CardContent>
                <Typography variant="body1" component="div">
                  {data.description}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Price: {data.price}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Tags: {data.tags}
                </Typography>
              </CardContent>
              <CardActions>
                <IconButton
                  onClick={() => {
                    if (localStorage.getItem("userId")) {
                      handleAddCartPhoto(data);
                    } else {
                      window.location.href = "/giris";
                    }
                  }}
                  aria-label="Add to Cart"
                >
                  <ShoppingCart />
                </IconButton>
                <IconButton
                  onClick={() => {
                    if (localStorage.getItem("userId")) {
                      handleAddFavorite(data);
                    } else {
                      window.location.href = "/giris";
                    }
                  }}
                  aria-label="Add to Favorites"
                >
                  <Favorite />
                </IconButton>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    ) : (
      <p>Veriler yükleniyor...</p>
    )}
  </Box>
</Container>
    </div>
  );
}

export default App;