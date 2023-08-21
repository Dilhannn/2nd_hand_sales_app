import * as React from "react";
import { AppBar, Toolbar, Typography, Button, TextField } from "@mui/material";
import { Link } from "react-router-dom";
import logo from "./img/logo.jpeg";

const Navbar = ({ onTextFieldChange }) => {
  const handleInputChange = (e) => {
    const value = e.target.value;
    onTextFieldChange(value);
  };

  return (
    <AppBar position="static" color="default">
      <Toolbar sx={{ display: "flex", justifyContent: "space-between", backgroundColor: "turquoise" }}>
        <Typography variant="h6" component={Link} to="/" sx={{ color: "lightgreen" }}>
          <img src={logo} alt="Logo" style={{ width: "200px", height: "auto" }} />
        </Typography>
        <div>
          <TextField
            variant="outlined"
            placeholder="Ara..."
            size="small"
            onChange={handleInputChange}
            sx={{ marginRight: "10px" }}
          />
          <Button color="inherit" component={Link} to="/">
            ANASAYFA
          </Button>
          <Button color="inherit" component={Link} to="/giris">
            GİRİŞ YAP
          </Button>
          <Button color="inherit" component={Link} to="/kayıt">
            KAYIT OL
          </Button>
        </div>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
