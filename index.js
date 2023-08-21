import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import reportWebVitals from './reportWebVitals';
import App from './App';
import Markalar from './Markalar';
import LoginPage from './components/LoginPage';
import SignUpPage from './components/SignUpPage';
import UserPage from './UserPage';
import { createBrowserRouter, RouterProvider, Route } from 'react-router-dom';
import ListCart from './ListCart';
import ListFavorite from './ListFavorite';

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
  },
  {
    path: '/markalar',
    element: <Markalar />,
  },
  {
    path: '/user/:username',
    element: <UserPage />,
  },
  {
    path: '/giris',
    element: <LoginPage />,
  },
  {
    path: '/kayıt',
    element: <SignUpPage />,
  },
  {
    path: '/sepet',
    element: <ListCart />,
  },
  {
    path: '/favori',
    element: <ListFavorite/>,
  },
]);

ReactDOM.render(
  <RouterProvider router={router}>
    <Route />
  </RouterProvider>,
  document.getElementById('root')
);

reportWebVitals();
