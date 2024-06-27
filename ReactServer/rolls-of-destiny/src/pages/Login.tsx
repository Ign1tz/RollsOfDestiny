import React, {useState} from "react";
import {TextField} from "@mui/material";
import Button from "@mui/material/Button";
import {login} from "../auth"

export default function Login() {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    function submit() {
        if (username && password) {
            fetch("http://localhost:9090/login", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({username: username, password: password})
            }).then(r => {

                return r.json()
            }).then(token => {
                console.log(token)
                if (token.token) {
                    login(token.token)
                    sessionStorage.setItem("username", username)
                    window.location.href = "/"
                }
            })
        }
    }


    return (
        <div className={"loginSignUpDivision"}>
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
