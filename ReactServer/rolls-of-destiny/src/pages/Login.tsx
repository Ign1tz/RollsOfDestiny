import React, {useState} from "react";
import {Modal, TextField} from "@mui/material";
import Button from "@mui/material/Button";
import {login} from "../auth"
import "../css/LoginSignup.css"
import {Link} from "react-router-dom";

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
            fetch("http://10.0.0.2:9090/login", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({username: username, password: password})
            }).then(r => {
                if (r.status === 200) {
                    r.json().then(token => {
                        if (token.token) {
                            login(token.token)
                            sessionStorage.setItem("username", username)
                            window.location.href = "/"
                        }
                    })

                    return
                } else {
                    setPassword("")
                    setIsError(true)
                    return
                }
            })
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
            <p className={"link"}> Don't have an account yet?
                <br/>
                <Link to="http://10.0.0.2:3000/signup">Sign Up Here</Link>
            </p>
        </div>
    )
}
