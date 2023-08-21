//api.js
import axios from 'axios';

const BASE_URL = 'http://localhost:3001'; // Backend URL

//API call
export const signIn = async (credentials) => {
  try {
    const response = await axios.post(`${BASE_URL}/signin`, credentials);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const signUp = async (userData) => {
  try {
    const response = await axios.post(`${BASE_URL}/signup`, userData);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

