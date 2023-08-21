//LoginPage.js
import React, { useState } from 'react';
import { signIn } from './api';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import { Link } from "react-router-dom";
import { useNavigate } from 'react-router-dom';
import Navbar from './Navbar';


const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    setError(null);
    setLoading(true);

    try {
      const credentials = {
        username,
        password,
      };

      const response = await signIn(credentials);

    // Handle successful sign in
    console.log(response);

    setUsername('');
    setPassword('');

    // Navigate to the user-specific page with the username as a parameter
    localStorage.setItem('userId', response.user);
    localStorage.setItem('username', username);
    navigate(`/user/${username}`);
  } catch (error) {
      setError('Giriş yapılamadı. Lütfen kullanıcı adı ve şifrenizi kontrol edin.');
    }

    setLoading(false);
  };
  

  return (
    <div><Navbar></Navbar>
    <Container maxWidth="sm">
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          height: '100vh',
        }}
      >
        <Typography path variant="h4" gutterBottom>
          Giriş Yap
        </Typography>
        {error && <Typography color="error">{error}</Typography>}
        <Box
          component="form"
          onSubmit={handleSubmit}
          sx={{
            display: 'flex',
            flexDirection: 'column',
            width: '100%',
            maxWidth: '400px',
          }}
        >
          <TextField
            label="Kullanıcı Adı"
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            margin="normal"
            required
          />
          <TextField
            label="Şifre"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            margin="normal"
            required
          />
          <Button type="submit" variant="contained" disabled={loading} sx={{ mt: 2 }}>
            Giriş Yap
          </Button>
        </Box>
      </Box>
    </Container>
    </div>
  );
};

export default LoginPage;
