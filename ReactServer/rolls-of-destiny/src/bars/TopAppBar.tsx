import * as React from 'react';
import {useEffect, useState} from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import MenuIcon from '@mui/icons-material/Menu';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import MenuItem from '@mui/material/MenuItem';
import AdbIcon from '@mui/icons-material/Adb';
import ProfileButton from "../components/ProfileButton";
import LoginSignUpButton from "../components/LoginSignUpButton";
import "../css/ExtraTopAppBar.css"

const pages = ["Home", "Landing Page", 'Friends', 'Leaderboard', 'Rules'];
const pagesLogOut = ['Leaderboard', 'Rules']
const settings = ['Profile', 'Decks', 'Settings', 'Logout'];

export default function TopAppBar({loggedIn}: { loggedIn: boolean }) {
    const [anchorElNav, setAnchorElNav] = React.useState<null | HTMLElement>(null);
    const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);
    const [username, setUsername] = useState()

    const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorElNav(event.currentTarget);
    };
    const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorElUser(event.currentTarget);
    };

    const handleCloseNavMenu = () => {
        setAnchorElNav(null);
    };

    const handleCloseUserMenu = () => {
        setAnchorElUser(null);
    };

    const relocate = (page: string) => {
        console.log(loggedIn);
        switch (page) {
            case "Profile":
                window.location.href = loggedIn ? "/profile" : "/login";
                break;
            case "Settings":
                window.location.href = "/settings";
                break;
            case "Leaderboard":
                window.location.href = "/leaderboard";
                break;
            case "Friends":
                window.location.href = "/friendlist";
                break;
            case "Home":
                window.location.href = "/";
                break;
            case "Rules":
                window.location.href = "/rules";
                break;
            case "Decks":
                window.location.href = "/decks";
                break;
            case "Landing Page":
                window.location.href = "/landingpage"
                break;
            case "Logout":
                localStorage.clear()
                sessionStorage.clear()
                window.location.href = "/"
        }
    }

    useEffect(() => {
        let username = sessionStorage.getItem("userInfo")
        if (username != "" && username != null) {
            setUsername(JSON.parse(username).username || "")

        }
    }, []);

    return (
        <>
            <AppBar position="static" color={"secondary"}>
                <Container maxWidth="xl">
                    <Toolbar disableGutters>
                        <AdbIcon sx={{display: {xs: 'none', md: 'flex'}, mr: 1}}/>
                        <Typography
                            variant="h6"
                            noWrap
                            component="a"
                            sx={{
                                mr: 2,
                                display: {xs: 'none', md: 'flex'},
                                fontFamily: 'monospace',
                                fontWeight: 700,
                                letterSpacing: '.3rem',
                                color: 'inherit',
                                textDecoration: 'none',
                            }}
                        >
                            Rolls of Destiny
                        </Typography>

                        <Box sx={{flexGrow: 1, display: {xs: 'flex', md: 'none'}}}>
                            <IconButton
                                size="large"
                                aria-label="account of current user"
                                aria-controls="menu-appbar"
                                aria-haspopup="true"
                                onClick={handleOpenNavMenu}
                                color="inherit"
                            >
                                <MenuIcon/>
                                <h2>{username}</h2>
                            </IconButton>
                            <Menu
                                id="menu-appbar"
                                anchorEl={anchorElNav}
                                anchorOrigin={{
                                    vertical: 'bottom',
                                    horizontal: 'left',
                                }}
                                keepMounted
                                transformOrigin={{
                                    vertical: 'top',
                                    horizontal: 'left',
                                }}
                                open={Boolean(anchorElNav)}
                                onClose={handleCloseNavMenu}
                                sx={{
                                    display: {xs: 'block', md: 'none'},
                                }}
                            >
                                {(loggedIn ? pages : pagesLogOut).map((page) => (
                                    <MenuItem key={page} onClick={() => relocate(page)}>
                                        <Typography textAlign="center">{page}</Typography>
                                    </MenuItem>
                                ))}
                            </Menu>
                        </Box>
                        <AdbIcon sx={{display: {xs: 'flex', md: 'none'}, mr: 1}}/>
                        <Typography
                            variant="h5"
                            noWrap
                            component="a"
                            href="#app-bar-with-responsive-menu"
                            sx={{
                                mr: 2,
                                display: {xs: 'flex', md: 'none'},
                                flexGrow: 1,
                                fontFamily: 'monospace',
                                fontWeight: 700,
                                letterSpacing: '.3rem',
                                color: 'inherit',
                                textDecoration: 'none',
                            }}
                        >
                            Rolls of Destiny
                        </Typography>

                        <Box sx={{flexGrow: 1, display: {xs: 'none', md: 'flex'}}}>
                            {(loggedIn ? pages : pagesLogOut).map((page) => (
                                <Button
                                    key={page}
                                    onClick={() => relocate(page)}
                                    sx={{my: 2, color: 'white', display: 'block'}}
                                >
                                    {page}
                                </Button>
                            ))}
                        </Box>

                        <h3 style={{marginRight: "10px"}}>{username}</h3>
                        {loggedIn ? <ProfileButton anchorElUser={anchorElUser} handleOpenUserMenu={handleOpenUserMenu}
                                                   handleCloseUserMenu={handleCloseUserMenu} settings={settings}
                                                   relocate={relocate}/> : <LoginSignUpButton></LoginSignUpButton>}
                    </Toolbar>
                </Container>
            </AppBar>
        </>
    );
}