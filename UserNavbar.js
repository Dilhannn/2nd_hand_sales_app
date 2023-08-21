import { AppBar, Toolbar, Typography, Button } from "@mui/material";
import { Link } from "react-router-dom";
import logo from "./img/logo.jpeg";

function cikis() {
  localStorage.clear();
}

const UserNavbar = () => {
  const username = localStorage.getItem("username");

  return (
    <AppBar position="static" color="default">
     <Toolbar sx={{ display: "flex", justifyContent: "space-between", backgroundColor: "turquoise" }}>
        <Typography variant="h6" component={Link} to="/" sx={{ color: "lightgreen" }}>
          <img src={logo} alt="Logo" style={{ width: "200px", height: "auto" }} />
        </Typography>
        <div>
          <Button color="inherit" component={Link} to="/">Anasayfa</Button> 
          <Button color="inherit" onClick={cikis} component={Link} to="/giris">Çıkış</Button>
          {username ? (
            <Button color="inherit" component={Link} to={`/user/${username}`}>Hesabım</Button>
          ) : (
            <Button color="inherit" component={Link} to="/giris">Giriş</Button>
          )}
          <Button color="inherit" component={Link} to="/sepet">Sepetim</Button>
          <Button color="inherit" component={Link} to="/favori">Favorilerim</Button>
        </div>
      </Toolbar>
    </AppBar>
  );
};

export default UserNavbar;
