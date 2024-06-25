import {TextField} from "@mui/material";
import "../css/Login.css"
import Button from "@mui/material/Button";
import React, {useState} from "react";


export default function SignUp() {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [email, setEmail] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")


    function submit() {
        const usernamePassed = checkUsername();
        const passwordPassed = checkPassword();
        const emailPassed = checkEmail();
        if (usernamePassed && passwordPassed && emailPassed) {
            fetch("http://localhost:9090/signup", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({username: username, email: email, password: password, confirmPassword: confirmPassword})
            })
                .then(res => {
                    if (res.status === 200) {
                        window.location.href = "/login";
                    }
                })
        } else {

        }
    }

    function checkUsername() {
        for (let character of username) {
            if (!"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_".includes(character)) {
                return false;
            }
        }
        return true;
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
        return password === confirmPassword;
    }

    function checkEmail() {
        const pattern = /^\w+@[a-zA-Z_]+?.[a-zA-Z]{2,3}$/;
        return pattern.test(email)
    }


    return (
        <div className={"loginSignUpDivision"}>
            <h1>SignUp</h1>
            <h3>Username</h3>
            <TextField required id="filled-basic" label="Username" variant="filled"
                       value={username}
                       onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                           setUsername(event.target.value);
                       }}/>
            <h3>Email Address</h3>
            <TextField required id="filled-basic" label="E-Mail" variant="filled"
                       value={email}
                       onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                           setEmail(event.target.value);
                       }}/>
            <h3>Password</h3>
            <TextField required id="filled-basic" label="Password" type="password" variant="filled"
                       value={password}
                       onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                           setPassword(event.target.value);
                       }}/>
            <h3>Confirm Password</h3>
            <TextField required id="filled-basic" label="ConfirmPassword" type="password" variant="filled"
                       value={confirmPassword}
                       onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                           setConfirmPassword(event.target.value);
                       }}/>
            <br/>
            <br/>
            <Button variant="contained" onClick={submit}>Submit</Button>
        </div>
    )
}

