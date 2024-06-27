import TopAppBar from "../bars/TopAppBar";
import {TextField} from "@mui/material";
import React, {useState} from "react";
import Button from "@mui/material/Button";

export default function Settings() {

    const [oldUsername, setOldUsername] = useState("");
    const [newUsername, setNewUsername] = useState("");
    const [confirmNewUsername, setConfirmNewUsername] = useState("");
    const [oldPassword, setOldPassword] = useState("");
    const [newPassword, setNewPassword ] = useState("");
    const [confirmNewPassword, setConfirmNewPassword] = useState("");

    function checkPasswordChange() {
        if (oldPassword === newPassword || newPassword !== confirmNewPassword)  {
            return false;
        }
        if (newPassword.length < 6 || newPassword.length > 50) {
            return false
        }

        return true;

    }

    function checkUsernameChange() {
        if (oldUsername === newUsername || newUsername !== confirmNewUsername) {
            return false;
        }

        for (let character of newUsername) {
            if (!"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_".includes(character)) {
                return false;
            }
        }
        return true;

    }

    function submitNewUsername() {
        if (checkUsernameChange()) {
            // handle fetch logic
        } else {
            // return error like in Login
        }
    }

    function submitPasswordChange() {
        if (checkPasswordChange()) {
            // handle fetch logic
        } else {
            // return error like in Login
        }
    }

    return (
        <>
            <TopAppBar loggedIn={true}></TopAppBar>
            <div className="settings" style={{textAlign: "center", justifyContent: "center", alignItems: "center"}}>
                <h2>Change Volume here:</h2>
                <h3>Volume Bar here </h3>
                <Button variant={"contained"}>Submit Volume Change</Button>
                <br/>
                <br/>
                <h2>Change Username here</h2>
                <h3>Old Username</h3>
                <TextField id="filled-basic" label="Username" variant="filled"
                           value={oldUsername}
                           onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                               setOldUsername(event.target.value);
                           }}/>
                <h3>New Username</h3>
                <TextField id="filled-basic" label="Password" type="password" variant="filled"
                           value={newUsername}
                           onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                               setNewUsername(event.target.value);
                           }}/>
                <h3>Confirm New Username</h3>
                <TextField id="filled-basic" label="Password" type="password" variant="filled"
                           value={confirmNewUsername}
                           onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                               setConfirmNewUsername(event.target.value);
                           }}/>
                <br/>
                <br/>
                <Button variant={"contained"}>Submit Username Change</Button>
                <br/>
                <br/>
                <h2>Change Password here</h2>
                <h3>Old Password</h3>
                <TextField id="filled-basic" label="Username" variant="filled"
                           value={oldPassword}
                           onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                               setOldPassword(event.target.value);
                           }}/>
                <h3>New Password</h3>
                <TextField id="filled-basic" label="Password" type="password" variant="filled"
                           value={newPassword}
                           onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                               setNewPassword(event.target.value);
                           }}/>
                <h3>Confirm New Password</h3>
                <TextField id="filled-basic" label="Password" type="password" variant="filled"
                           value={confirmNewPassword}
                           onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                               setConfirmNewPassword(event.target.value);
                           }}/>
                <br/>
                <br/>
                <Button variant={"contained"}>Submit Password Change</Button>

            </div>
        </>
    )
}