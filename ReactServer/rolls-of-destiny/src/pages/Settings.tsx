import TopAppBar from "../bars/TopAppBar";
import {Modal, TextField} from "@mui/material";
import React, {useEffect, useState} from "react";
import Button from "@mui/material/Button";
import "../css/Settings.css"
import {profile} from "../types/profileTypes";
import VolumeSlider from "../components/VolumeSlider";
import {authFetch} from "../auth";

export default function Settings({profile}: { profile: profile }) {

    const [newUsername, setNewUsername] = useState("");
    const [oldPassword, setOldPassword] = useState("");
    const [newPassword, setNewPassword] = useState("");
    const [confirmNewPassword, setConfirmNewPassword] = useState("");
    const [isUsernameError, setIsUsernameError] = useState(false);
    const [isPasswordError, setIsPasswordError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [volume, setVolume] = useState<number>(30)
    const [imageString, setImageString] = useState("")
    const [image, setImage] = useState(<img alt={""}/>)
    const [username, setUsername] = useState("")

    useEffect(() => {
        authFetch("http://localhost:9090/userInfo").then(r => {

            return r.json()
        }).then(response => {
            sessionStorage.setItem("userInfo", JSON.stringify(response))
            sessionStorage.setItem("profilePicture", response.profilePicture)
        })
        setVolume(Number(sessionStorage.getItem("volume")) || 30)

        let imageString = sessionStorage.getItem("userInfo")
        if (imageString) {
            imageString = JSON.parse(imageString).profilePicture
            setImage(<img className={"img"} src={"data:image/jpeg;base64," + imageString}
                          alt={"Something went wrong"}/>)
        }
        let userinfo = sessionStorage.getItem("userInfo")
        if (userinfo) {
            let username = JSON.parse(userinfo).username
            setUsername(username)
        }
    }, []);


    function checkPasswordChange() {
        // return false if oldPassword is incorrect
        // if (oldPassword != database entry password ...)

        if (oldPassword === newPassword || newPassword !== confirmNewPassword) {
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
            authFetch("http://localhost:9090/changeUsername", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({oldUsername: profile.username, newUsername: newUsername})
            }).then(r => {
                if (r.status === 200) {
                    authFetch("http://localhost:9090/userInfo").then(r => {
                        return r.json()
                    }).then(response => {
                        sessionStorage.setItem("userInfo", JSON.stringify(response))
                        window.location.reload()
                    })
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
            authFetch("http://localhost:9090/changePassword", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({
                    oldPassword: oldPassword,
                    newPassword: newPassword,
                    confirmNewPassword: confirmNewPassword
                })
            }).then(r => {
                if (r.status === 200) {
                    window.location.reload()
                }
                setErrorMessage("Something went wrong while trying to save your password. Please try again.");
            });
        } else {
            setIsPasswordError(true);
            if (oldPassword !== "" && newPassword !== "" && confirmNewPassword !== "") {
                setErrorMessage("New Password may not be valid. Old Password may be false. Maybe Password and Confirm Password are not the same.");
            } else {
                setErrorMessage("Please fill out every field if you want to change your password.");
            }
        }
    }

    function submitProfilePicture() {
        if (imageString != "") {
            authFetch("http://localhost:9090/changeProfilePicture", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({profilePicture: imageString})
            }).then(r => {
                if (r.status === 200) {
                }
                setErrorMessage("Something went wrong while trying to save your password. Please try again.");
            });
        }
    }

    function deleteAccount() {
        authFetch("http://localhost:9090/deleteAccount", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            }
        }).then(r => {
            localStorage.clear()
            sessionStorage.clear()
            window.location.href = "/"
        });
    }

    function handleImage(e: any) {
        let reader = new FileReader()
        reader.onloadend = function () {
            if (typeof reader.result === "string") {
                let convertedImg = reader.result.split(',')[1]
                setImageString(convertedImg)
                let image = <img className={"img"} src={"data:image/jpeg;base64," + convertedImg}
                                 alt={"Something went wrong"}></img>
                setImage(image)

            }
        }
        reader.readAsDataURL(e.target.files[0])

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
                        {image}
                        <input className={"profile_input"} type={"file"} accept={"image/png, image/gif, image/jpeg"}
                               name={"file"}
                               onChange={handleImage}/>
                        <Button variant={"contained"} color="secondary" onClick={submitProfilePicture}>Submit Profile
                            Picture</Button>
                    </div>
                    <div className="volume">
                        <h2 id={"h2Text"}>Volume</h2>
                        <VolumeSlider volume={volume} setVolume={setVolume}/>
                        <h5>Attention: it is required to "slide" the bar, not click it.</h5>
                        <Button style={{marginTop: "30px"}} variant={"contained"} color="secondary"
                                onClick={deleteAccount}>Delete Account</Button>
                    </div>
                    <div className="username">
                        <h2 id={"h2Text"}>Username</h2>
                        <h3 id={"h3Text"}>New Username</h3>
                        <TextField id="filled-basic" label={username} variant="filled"
                                   value={newUsername}
                                   onChange={(event) => setNewUsername(event.target.value)}/>
                        <br/>
                        <br/>
                        <Button variant={"contained"} color="secondary" onClick={submitNewUsername}>Submit Username
                            Change</Button>
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
                        <Button variant={"contained"} color="secondary" onClick={submitPasswordChange}>Submit Password
                            Change</Button>
                    </div>
                </div>
            </div>
        </>
    );
}
