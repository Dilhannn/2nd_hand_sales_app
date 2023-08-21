import React, { useState, useEffect } from 'react';
import { AppBar, Toolbar, Typography, Container, CustomButton, Box, List, ListItemSecondaryAction, IconButton, TextField, ListItem, ListItemAvatar, ListItemText, Avatar, Button, FormControl, Input, FormHelperText } from '@mui/material';
import { ShoppingCart, Clear, Favorite, Edit, Delete } from '@mui/icons-material';
import axios from 'axios';
import UserNavbar from './components/UserNavbar';

const ListCart = () => {
	const [dataList, setDataList] = useState([]);
	const BASE_URL = 'http://localhost:3001';

	useEffect(() => {
		handleCartList();
	}, []);


	const handleCartList = async () => {

		try {
			var userid = localStorage.getItem("userId")
			const resp = await axios.post(`${BASE_URL}/ListCartModel`, {
				UserID: userid,
			});
			console.log(resp)
			setDataList(resp.data.list)

		} catch (error) {
			console.log('Error:', error);
		}
	};

	const handleDeleteCart = async (item) => {

		try {
			const resp = await axios.post(`${BASE_URL}/RemoveFromCartModel`, {
				ID: item.id,
			});

			handleCartList()
		} catch (error) {
			console.log('Error:', error);
		}
	};



	return (

		<div>
			<UserNavbar />
			<Box sx={{ display: 'flex', marginTop: "30px", justifyContent: 'center' }}>
				<Container maxWidth="md">
					<Typography variant="h5" component="h3" gutterBottom>
						My Cart
					</Typography>
					<Box>
						{dataList !== undefined && dataList != null ? (
							<List>
								{dataList.map((data) => (
									<ListItem key={data.photo.id}>
										<ListItemAvatar>
											<Avatar alt={data.photo.description} src={`data:image/png;base64, ${data.photo.url}`} />
										</ListItemAvatar>
										<ListItemText primary={data.photo.description} secondary={data.photo.tags} />
										<ListItemSecondaryAction>
											<IconButton onClick={() => handleDeleteCart(data)} edge="end" aria-label="Delete">
												<Delete />
											</IconButton>
										</ListItemSecondaryAction>
									</ListItem>
								))}
							</List>
						) : (
							<p>Veriler yükleniyor...</p>
						)}
					</Box>
					<Box sx={{ display: 'flex', justifyContent: 'flex-end', marginTop: '2rem' }}>

						<Box sx={{ display: 'flex', justifyContent: 'flex-end', marginTop: '2rem' }}>
							<Button variant="contained" color="primary">
								Go Payment
							</Button>
						</Box>
					</Box>
				</Container>
			</Box>
		</div>
	);
};

export default ListCart;