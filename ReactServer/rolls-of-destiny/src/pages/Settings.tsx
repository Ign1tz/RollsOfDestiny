import TopAppBar from "../bars/TopAppBar";
import {Modal, TextField} from "@mui/material";
import React, {useState} from "react";
import Button from "@mui/material/Button";
import "../css/Settings.css"
import {profile} from "../types/profileTypes";
import VolumeSlider from "../components/VolumeSlider";

export default function Settings({profile}: {profile:profile }) {

    const [newUsername, setNewUsername] = useState("");
    const [oldPassword, setOldPassword] = useState("");
    const [newPassword, setNewPassword ] = useState("");
    const [confirmNewPassword, setConfirmNewPassword] = useState("");
    const [isUsernameError, setIsUsernameError] = useState(false);
    const [isPasswordError, setIsPasswordError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    function checkPasswordChange() {
        // return false if oldPassword is incorrect
        // if (oldPassword != database entry password ...)

        if (oldPassword === newPassword || newPassword !== confirmNewPassword)  {
            return false;
        }
        if (newPassword.length < 6 || newPassword.length > 50) {
            return false;
        }

        for (let character of newPassword) {
            if (!"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_!$@&#+*-€".includes(character)) {
                return false;
            }
        }

        return true;
    }

    function checkUsernameChange() {
        if (profile.username === newUsername || newUsername === "") {
            return false;
        }

        for (let character of newUsername) {
            if (!"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_".includes(character)) {
                return false;
            }
        }
        return true;
    }

    function closeError() {
        setIsUsernameError(false);
        setIsPasswordError(false);
        setErrorMessage("");
    }

    function submitNewUsername() {
        if (checkUsernameChange()) {
            fetch("http://localhost:9090/changeUsername", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({oldUsername: profile.username, newUsername: newUsername})
            }).then(r => {
                if (r.status === 404) {
                    return r.json();
                }
                setErrorMessage("Username already taken.");
            });
        } else {
            setIsUsernameError(true);
            setErrorMessage("New username not valid.");
        }
    }

    function submitPasswordChange() {
        if (checkPasswordChange()) {
            fetch("http://localhost:9090/changePassword", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({oldPassword: oldPassword, newPassword: newPassword})
            }).then(r => {
                if (r.status === 404) {
                    return r.json();
                }
                setErrorMessage("Something went wrong while trying to save your password. Please try again.");
            });
        } else {
            setIsPasswordError(true);
            if (oldPassword !== "" && newPassword !=="" && confirmNewPassword!=="") {
                setErrorMessage("New Password may not be valid. Old Password may be false. Maybe Password and Confirm Password are not the same.");
            } else {
                setErrorMessage("Please fill out every field if you want to change your password.");
            }
        }
    }

    return (
        <>
            <TopAppBar loggedIn={true}></TopAppBar>
            <div className="settings" style={{textAlign: "center", justifyContent: "center", alignItems: "center"}}>
                <h1>Settings</h1>
                <Modal open={isUsernameError || isPasswordError} onClose={closeError}>
                    <div className="errorMenu">
                        <h2>Oops...something went wrong.</h2>
                        <div className="errorText">
                            {errorMessage}
                        </div>
                        <Button variant="contained" color={"secondary"} onClick={closeError}>
                            I understand.
                        </Button>
                    </div>
                </Modal>
                <div className="grid-container">
                    <div className="profilePicture">
                        <h2 id={"h2Text"}>Profile Picture</h2>
                        <h3 id={"h3Text"}>Your current Profile Picture:</h3>
                        <img src={profile.picture} alt={"profile picture current"}/>
                        <h3 id={"h3Text"}>Upload a new picture</h3>
                        <h4>Coming soon...</h4>
                    </div>
                    <div className="volume">
                        <h2 id={"h2Text"}>Volume</h2>
                        <VolumeSlider/>
                        <h5>Attention: it is required to "slide" the bar, not click it.</h5>
                    </div>
                    <div className="username">
                        <h2 id={"h2Text"}>Username</h2>
                        <h3 id={"h3Text"}>New Username</h3>
                        <TextField id="filled-basic" label={profile.username} type="password" variant="filled"
                                   value={newUsername}
                                   onChange={(event) => setNewUsername(event.target.value)}/>
                        <br/>
                        <br/>
                        <Button variant={"contained"} color = "secondary" onClick={submitNewUsername}>Submit Username Change</Button>
                    </div>
                    <div className="password">
                        <h2 id={"h2Text"}>Password</h2>
                        <h3 id={"h3Text"}>Old Password</h3>
                        <TextField id="filled-basic" label="Password" type={"password"} variant="filled"
                                   value={oldPassword}
                                   onChange={(event) => setOldPassword(event.target.value)}/>
                        <h3 id={"h3Text"}>New Password</h3>
                        <TextField id="filled-basic" label="Password" type="password" variant="filled"
                                   value={newPassword}
                                   onChange={(event) => setNewPassword(event.target.value)}/>
                        <h3 id={"h3Text"}>Confirm New Password</h3>
                        <TextField id="filled-basic" label="Password" type="password" variant="filled"
                                   value={confirmNewPassword}
                                   onChange={(event) => setConfirmNewPassword(event.target.value)}/>
                        <br/>
                        <br/>
                        <Button variant={"contained"} color = "secondary" onClick={submitPasswordChange}>Submit Password Change</Button>
                    </div>
                </div>
            </div>
        </>
    );
}
