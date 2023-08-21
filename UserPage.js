import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Typography, Container, Box, Card, IconButton, TextField, CardMedia, CardContent, CardActions, Button, FormControl, Input, FormHelperText } from '@mui/material';
import { Clear, Favorite, Edit, Delete } from '@mui/icons-material';
import UserNavbar from './components/UserNavbar';
import axios from 'axios';
import ReactDOM from 'react-dom';




const Popup = ({ children }) => {

  const popupContainerStyle = {
    position: 'fixed',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
    zIndex: 9999
  };

  const popupContentStyle = {
    backgroundColor: '#fff',
    padding: '20px',
    borderRadius: '5px',
    boxShadow: '0 2px 10px rgba(0, 0, 0, 0.2)'
  };

  return ReactDOM.createPortal(
    <div style={popupContainerStyle}>
      <div style={popupContentStyle}>
        {children}
      </div>
    </div>,
    document.getElementById('popup-root')
  );
};

const UserPage = () => {
  const navigate = useNavigate();

  const fetchApiData = async () => {

    try {
      if (localStorage.getItem("userId") == null || localStorage.getItem("username") == null) {
        navigate(`/`);
      }
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchApiData();
  }, []);

  const { username } = useParams();
  const [selectedFile, setSelectedFile] = useState(null);
  const [selectedFileEdit, setSelectedFileEdit] = useState(null);
  const [base64Data, setBase64Data] = useState('');
  const [base64DataEdit, setBase64DataEdit] = useState('');
  const [error, setError] = useState('');
  const [selectedItem, setSelectedItem] = useState(null);

  const handleListItemClick = (item) => {
    console.log(item)
    setSelectedItem(item);
    openPopup();
  };
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  const openPopup = () => {
    setIsPopupOpen(true);
  };

  const closePopup = () => {
    setIsPopupOpen(false);
  };

  const [formData, setFormData] = useState({
    url: '',
    description: '',
    price: '',
    tags: '',
  });

  const [dataList, setDataList] = useState([]);

  useEffect(() => {
    handleListUserPhotos();
  }, []);

  const convertToBase64 = (file) => {
    const reader = new FileReader();

    reader.onloadend = () => {
      const base64String = reader.result.split(',')[1];
      setBase64Data(base64String);
    };

    reader.readAsDataURL(file);
  };

  const convertToBase64Edit = (file) => {
    const reader = new FileReader();

    reader.onloadend = () => {
      const base64String = reader.result.split(',')[1];
      setBase64DataEdit(base64String);
    };

    reader.readAsDataURL(file);
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setSelectedFile(file);
      setError('');
      convertToBase64(file);
    } else {
      setSelectedFile(null);
      setBase64Data('');
      setError('Dosya seçiniz');
    }
  };

  const handleFileChangeEdit = (e) => {
    const file = e.target.files[0];
    if (file) {
      setSelectedFileEdit(file);
      setError('');
      convertToBase64Edit(file);
    } else {
      setSelectedFileEdit(null);
      setBase64DataEdit('');
      setError('Dosya seçiniz');
    }
  };

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleChangeEdit = (e) => {
    setSelectedItem({
      ...selectedItem,
      [e.target.name]: e.target.value,
    });
  };

  const BASE_URL = 'http://localhost:3001';

  const handleCreatePhoto = async () => {
    try {
      const userid = localStorage.getItem('userId');
      const resp = await axios.post(`${BASE_URL}/CreatePhotoModel`, {
        Description: formData.description,
        UserID: userid,
        URL: base64Data,
        Tags: formData.tags,
        Price: formData.price,
      });

      handleListUserPhotos();
    } catch (error) {
      console.log('Error:', error);
    }
  };

  const handleEditPhoto = async () => {
    try {
      const userid = localStorage.getItem('userId');
      const resp = await axios.post(`${BASE_URL}/UpdatePhotoModel`, {
        ID: selectedItem.id,
        Description: selectedItem.description,
        UserID: userid,
        URL: base64DataEdit,
        Tags: selectedItem.tags,
        Price: selectedItem.price,
      });

      handleListUserPhotos();
      closePopup();
    } catch (error) {
      console.log('Error:', error);
    }
  };

  const handleDeletePhoto = async (item) => {
    try {
      const resp = await axios.post(`${BASE_URL}/DeletePhotoModel`, {
        ID: item.id,
      });

      handleListUserPhotos();
    } catch (error) {
      console.log('Error:', error);
    }
  };
  const handleListUserPhotos = async () => {
    try {
      const userid = localStorage.getItem('userId');
      const resp = await axios.post(`${BASE_URL}/ListPhotoModel`, {
        UserID: userid,
      });
      setDataList(resp.data.list);
    } catch (error) {
      console.log('Error:', error);
    }
  };

  // const handleAddCartPhoto = async (item) => {

  //   try {
  //     var userid = localStorage.getItem("userId")
  //     const resp = await axios.post(`${BASE_URL}/AddToCartModel`, {
  //      userid:userid, 
  //     photo:
  //     { Description: item.description,
  //       UserID: item.userid,
  //       URL: item.url,
  //       Tags: item.tags,
  //       Price: item.price,} 
  //     });

  //     handleListUserPhotos()
  //   } catch (error) {
  //     console.log('Error:', error);
  //   }
  // };

  return (
    <div>
      <UserNavbar />
      <Container maxWidth="md" sx={{ marginTop: '2rem' }}>
  <Typography variant="h4" component="h2" gutterBottom>
    Welcome, {username}!
  </Typography>
  <Box sx={{ marginTop: '2rem' }}>
    <Typography variant="h5" component="h3" gutterBottom>
      Create a new photo
    </Typography>
    <FormControl error={!!error} fullWidth>
      <Input type="file" onChange={handleFileChange} />
      {selectedFile && (
        <img style={{ maxWidth: "300px" }} src={URL.createObjectURL(selectedFile)} alt="Selected File" />
      )}
      {error && <FormHelperText>{error}</FormHelperText>}
    </FormControl>
    <TextField
      name="tags"
      label="Tags"
      placeholder="Enter tags"
      fullWidth
      required
      color="secondary"
      value={formData.tags}
      onChange={handleChange}
      sx={{ marginTop: '1rem' }}
    />
    <TextField
      name="price"
      label="Price"
      placeholder="Enter price"
      fullWidth
      required
      color="secondary"
      value={formData.price}
      onChange={handleChange}
      sx={{ marginTop: '1rem' }}
    />
    <TextField
      name="description"
      label="Description"
      placeholder="Enter description"
      fullWidth
      required
      color="secondary"
      value={formData.description}
      onChange={handleChange}
      sx={{ marginTop: '1rem' }}
    />
    <Button variant="contained" onClick={handleCreatePhoto} sx={{ marginTop: '1rem' ,backgroundColor: "turquoise" }}>
      Create Photo
    </Button>
  </Box>
  {/* Display user-specific content here */}
</Container>

      <div id="popup-root">
        {isPopupOpen && (
          <Popup item={selectedItem}>
            <IconButton onClick={closePopup} edge="end" aria-label="Add to Favorites">
              <Clear />
            </IconButton>
            <div>
              <Typography variant="h5" component="h3" gutterBottom>
                Create a new photo
              </Typography>
              <FormControl error={!!error}>
                <Input type="file" onChange={handleFileChangeEdit} />
                {selectedFileEdit ? (
                  <img style={{ maxWidth: "300px" }} src={URL.createObjectURL(selectedFileEdit)} alt="Seçilen Dosya" />
                ) : (
                  <span></span>
                )}
                {error && <FormHelperText>{error}</FormHelperText>}
              </FormControl>
              <Box mt={2}>
                <TextField
                  name="tags"
                  label="Tags"
                  placeholder="Enter your name"
                  fullWidth
                  required
                  color="secondary"
                  value={selectedItem.tags}
                  onChange={handleChangeEdit}
                />
              </Box>
              <Box mt={2}>
                <TextField
                  name="price"
                  label="Price"
                  placeholder="Enter your name"
                  fullWidth
                  required
                  color="secondary"
                  value={selectedItem.price}
                  onChange={handleChangeEdit}
                />
              </Box>
              <Box mt={2}>
                <TextField
                  name="description"
                  label="Description"
                  placeholder="Enter your name"
                  fullWidth
                  required
                  color="secondary"
                  value={selectedItem.description}
                  onChange={handleChangeEdit}
                />
              </Box>
              <Box mt={2}>
                <Button variant="contained" onClick={handleEditPhoto}>
                  Update Photo
                </Button>
              </Box>
            </div>
          </Popup>
        )}
      </div>
      <Container>
        <Box  sx={{ marginTop: 4 }} display="flex" flexWrap="wrap" justifyContent="flex-start">
          {dataList !== undefined && dataList != null ? (
            dataList.map((data, index) => (
              <Box key={data.id} width="25%" p={2}>
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
                    <IconButton edge="end" aria-label="Add to Favorites">
                      <Favorite />
                    </IconButton>
                    <IconButton onClick={() => handleListItemClick(data)} edge="end" aria-label="Add to Favorites">
                      <Edit />
                    </IconButton>
                    <IconButton onClick={() => handleDeletePhoto(data)} edge="end" aria-label="Add to Favorites">
                      <Delete />
                    </IconButton>
                  </CardActions>
                </Card>
              </Box>
            ))
          ) : (
            <p>Veriler yükleniyor...</p>
          )}
        </Box>
      </Container>

    </div>
  );
};

export default UserPage;
