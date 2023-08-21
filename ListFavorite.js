import React, { useState, useEffect } from 'react';
import { Box, List, ListItemSecondaryAction, IconButton, Button, Container, Typography, ListItem, ListItemAvatar, ListItemText, Avatar } from '@mui/material';
import { Delete } from '@mui/icons-material';
import UserNavbar from './components/UserNavbar';
import { Link } from "react-router-dom";
import axios from 'axios';

const ListFavorite = () => {
	const [dataList, setDataList] = useState([]);
	const BASE_URL = 'http://localhost:3001';

	useEffect(() => {
		handleFavoriteList();
	}, []);


	const handleFavoriteList = async () => {

		try {
			var userid = localStorage.getItem("userId")
			const resp = await axios.post(`${BASE_URL}/ListFavoriteModel`, {
				UserID: userid,
			});
			console.log(resp)
			setDataList(resp.data.list)

		} catch (error) {
			console.log('Error:', error);
		}
	};

	const handleDeleteFavorite = async (item) => {

		try {
			const resp = await axios.post(`${BASE_URL}/RemoveFromFavoriteModel`, {
				ID: item.id,
			});

			handleDeleteFavorite()
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
						Favorites
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
											<IconButton onClick={() => handleDeleteFavorite(data)} edge="end" aria-label="Delete">
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
						<Button variant="contained" color="primary" component={Link} to="/sepet">
							Go To Your Cart
						</Button>
					</Box>
				</Container>
			</Box>
		</div>
	);
};

export default ListFavorite;