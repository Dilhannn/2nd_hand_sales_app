import React, { useState } from 'react';
import { signUp } from './api';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import Navbar from './Navbar';

const SignUpPage = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    setErrorMessage('');
    setSuccess(false);
    setLoading(true);

    try {
      if (password !== confirmPassword) {
        throw new Error('Şifreler eşleşmiyor.');
      }

      const userData = {
        username,
        email,
        password,
      };

      const response = await signUp(userData);

      setSuccess(true);
      setErrorMessage('');

      setUsername('');
      setEmail('');
      setPassword('');
      setConfirmPassword('');
    } catch (error) {
      setErrorMessage(error.message);
      setSuccess(false);
    }

    setLoading(false);
  };

  return (
    <div> <Navbar></Navbar>
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
        <Typography variant="h4" gutterBottom>
          Kayıt Ol
        </Typography>
        {success && (
          <Typography color="green" sx={{ mb: 2 }}>
            Kayıt başarıyla tamamlandı!
          </Typography>
        )}
        {errorMessage && (
          <Typography color="error" sx={{ mb: 2 }}>
            {errorMessage}
          </Typography>
        )}
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
            label="E-posta"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
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
          <TextField
            label="Şifre Tekrarı"
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            margin="normal"
            required
          />
          <Button
            type="submit"
            variant="contained"
            disabled={loading}
            sx={{ mt: 2}}
          >
            Kayıt Ol
          </Button>
        </Box>
      </Box>
    </Container>
    </div>
  );
};

export default SignUpPage;
