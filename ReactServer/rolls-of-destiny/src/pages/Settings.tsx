import TopAppBar from "../bars/TopAppBar";
import {Modal, TextField} from "@mui/material";
import React, {useState} from "react";
import Button from "@mui/material/Button";
import "../css/Settings.css"

export default function Settings() {

    const [oldUsername, setOldUsername] = useState("");
    const [newUsername, setNewUsername] = useState("");
    const [confirmNewUsername, setConfirmNewUsername] = useState("");
    const [oldPassword, setOldPassword] = useState("");
    const [newPassword, setNewPassword ] = useState("");
    const [confirmNewPassword, setConfirmNewPassword] = useState("");
    const [isUsernameError, setIsUsernameError] = useState(false)
    const [isPasswordError, setIsPasswordError] = useState(false)

    let errorMessage: string = "Nothing typed.";

    function checkPasswordChange() {

        // return false if oldPassword is incorrect
        // if (oldPassword != database entry password ...)

        if (oldPassword === newPassword || newPassword !== confirmNewPassword)  {
            return false;
        }
        if (newPassword.length < 6 || newPassword.length > 50) {
            return false
        }

        for (let character of newPassword) {
            if (!"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_!$@&#+*-€".includes(character)) {
                return false;
            }
        }

        return true;

    }

    function checkUsernameChange() {
        if (oldUsername === newUsername || newUsername !== confirmNewUsername || oldUsername === "" || newUsername === "" || confirmNewUsername === "") {
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
        setIsUsernameError(false)
        setIsPasswordError(false)
        errorMessage = "Nothing typed."
    }

    function submitNewUsername() {
        if (checkUsernameChange()) {
            fetch("http://localhost:9090/changeUsername", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({oldUsername: oldUsername, newUsername: newUsername})
            }).then(r => {

                return r.json()
            })
        } else {
            setIsUsernameError(true)
            errorMessage = "Wrong old username or new username and confirmation for it are not the same or invalid username."
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

                return r.json()
            })
        } else {
            setIsPasswordError(true)
            errorMessage = "New Password not valid or Old Password false or New Password != Confirmation"
        }
    }

    return (
        <>
            <TopAppBar loggedIn={true}></TopAppBar>
            <div className="settings" style={{textAlign: "center", justifyContent: "center", alignItems: "center"}}>
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
                    <div className="volume">
                        <h2>Change Volume here:</h2>
                        <h3>Volume Bar here</h3>
                        <Button variant={"contained"}>Submit Volume Change</Button>
                    </div>
                    <div className="username">
                        <h2>Change Username here</h2>
                        <h3>Old Username</h3>
                        <TextField id="filled-basic" label="Username" variant="filled"
                                   value={oldUsername}
                                   onChange={(event) => setOldUsername(event.target.value)}/>
                        <h3>New Username</h3>
                        <TextField id="filled-basic" label="Password" type="password" variant="filled"
                                   value={newUsername}
                                   onChange={(event) => setNewUsername(event.target.value)}/>
                        <h3>Confirm New Username</h3>
                        <TextField id="filled-basic" label="Password" type="password" variant="filled"
                                   value={confirmNewUsername}
                                   onChange={(event) => setConfirmNewUsername(event.target.value)}/>
                        <Button variant={"contained"} onClick={submitNewUsername}>Submit Username Change</Button>
                    </div>
                    <div className="password">
                        <h2>Change Password here</h2>
                        <h3>Old Password</h3>
                        <TextField id="filled-basic" label="Username" variant="filled"
                                   value={oldPassword}
                                   onChange={(event) => setOldPassword(event.target.value)}/>
                        <h3>New Password</h3>
                        <TextField id="filled-basic" label="Password" type="password" variant="filled"
                                   value={newPassword}
                                   onChange={(event) => setNewPassword(event.target.value)}/>
                        <h3>Confirm New Password</h3>
                        <TextField id="filled-basic" label="Password" type="password" variant="filled"
                                   value={confirmNewPassword}
                                   onChange={(event) => setConfirmNewPassword(event.target.value)}/>
                        <Button variant={"contained"} onClick={submitPasswordChange}>Submit Password Change</Button>
                    </div>
                </div>
            </div>
        </>
    )
}