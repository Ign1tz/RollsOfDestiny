import Tooltip from "@mui/material/Tooltip";
import IconButton from "@mui/material/IconButton";
import Avatar from "@mui/material/Avatar";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import * as React from "react";

type ProfileButtonType = {
    anchorElUser: null | HTMLElement,
    handleOpenUserMenu: Function,
    handleCloseUserMenu: Function,
    settings: string[],
    relocate: Function
}


export default function ProfileButton({
                                          anchorElUser,
                                          handleOpenUserMenu,
                                          handleCloseUserMenu,
                                          settings,
                                          relocate
                                      }: ProfileButtonType) {
    return (
        <Box sx={{flexGrow: 0}}>
            <Tooltip title="Profile">
                <IconButton onClick={(event) => handleOpenUserMenu(event)} sx={{p: 0}}>
                    <Avatar alt="Remy Sharp" src={"data:image/jpeg;base64," + sessionStorage.getItem("profilePicture")}/>
                </IconButton>
            </Tooltip>
            <Menu
                sx={{mt: '45px'}}
                id="menu-appbar"
                anchorEl={anchorElUser}
                anchorOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
                open={Boolean(anchorElUser)}
                onClose={(event) => handleCloseUserMenu(event)}
            >
                {settings.map((setting) => (
                    <MenuItem key={setting}
                              onClick={() => relocate(setting)}>
                        <Typography textAlign="center">{setting}</Typography>
                    </MenuItem>
                ))}
            </Menu>
        </Box>
    )
}