import React, {useState} from "react";
import {TextField} from "@mui/material";
import Button from "@mui/material/Button";
import {login, authFetch, useAuth} from "../auth"
import {Link} from "react-router-dom";

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
            <p className={"link"}> Don't have an account yet?
                <br/>
                <Link to="http://localhost:3000/signup">Sign Up Here</Link>
            </p>
        </div>
    )
}
