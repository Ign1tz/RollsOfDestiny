import React, {useState} from "react";
import {TextField} from "@mui/material";
import Button from "@mui/material/Button";

export default function Login() {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    function submit() {
        if (username && password) {
            fetch("http://localhost:9090/signup", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({username: username, password: password})
            })
                .then(res => {
                    if (res.status === 200) {
                        window.location.href = "/home";
                    }
                })
        } else {

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
