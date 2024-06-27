import React, {useState} from "react";
import {Modal, TextField} from "@mui/material";
import Button from "@mui/material/Button";
import {authFetch, login} from "../auth"
import "../css/LoginSignup.css"

export default function Login() {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [isError, setIsError] = useState(false)

    function handleError() {
        setIsError(!isError)
    }

    function checkUsername() {
        if (username.length < 3 || username.length > 20) {
            return false
        }

        for (let character of username) {
            if (!"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_".includes(character)) {
                return false;
            }
        }
        return true
    }

    function checkPassword() {
        if (password.length < 6 || password.length > 50) {
            return false;
        }
        for (let character of password) {
            if (!"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_!$@&#+*-€".includes(character)) {
                return false;
            }
        }
        return true;
    }

    function submit() {
        if (checkUsername() && checkPassword()) {
            fetch("http://localhost:9090/login", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({username: username, password: password})
            }).then(r => {
                if (r.status === 200) {
                    return r.json()
                }else {
                    setIsError(true)
                    return
                }
            }).then(token => {
                console.log(token)
                if (token && token.token) {
                    login(token.token)
                    authFetch("http://localhost:9090/isLoggedIn").then(response => {
                        if (response.status === 200) {
                            window.location.href = "/"
                        }
                    })

                } /*else {
                        window.alert(token.errors + "\nPlease try again!")
                        setPassword("")
                        setUsername("")
                    }*/
            })
        } else {
            setIsError(true)
        }
    }


    return (
        <div className={"loginSignUpDivision"}>
            <Modal open={isError} onClose={handleError}>
                <div className="errorMenu">
                    <h2>Oops...something went wrong.</h2>
                    <div className="errorText">
                        {"Username or Password wrong. Please try again or reset your password."}
                    </div>
                    <Button variant="contained" color={"secondary"} onClick={handleError}>
                        I understand.
                    </Button>
                </div>
            </Modal>
            <h1>Login</h1>
            <h3>Username</h3>
            <TextField required id="filled-basic" label="Username" variant="filled"
                       value={username}
                       onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                           setUsername(event.target.value);
                       }}/>
            <h3>Password</h3>
            <TextField required id="filled-basic" label="Password" type="password" variant="filled"
                       value={password}
                       onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                           setPassword(event.target.value);
                       }}/>
            <br/>
            <br/>
            <Button variant="contained" onClick={submit}>Login</Button>
        </div>
    )
}
