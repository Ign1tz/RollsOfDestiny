import React, {useEffect, useState} from 'react';
import './App.css';
import Home from "./pages/Home";
import {BrowserRouter, Routes, Route, Navigate} from "react-router-dom"
import Profile from "./pages/Profile";
import {profile} from "./types/profileTypes";
import Game from "./pages/Game";
import Login from "./pages/Login";
import SignUp from "./pages/SignUp";
import {authFetch} from "./auth";

function App() {
    let p: profile = {username: "Bernd", biography: "H", profilePicture:" H", rating:839}
    const [loggedIn, setLoggedIn] = useState<boolean> (false)


    useEffect(() => {
        if (localStorage.getItem("access_token")) {
            authFetch("http://localhost:9090/isLoggedIn").then(response => {
            })
        }
    }, []);

    return (
    <>
        <BrowserRouter>
            <Routes>
                <Route index element={<Home loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>}/>
                <Route path={"/profile"} element={<Profile user={p}/>}/>
                <Route path={"/game" } element={<Game/>}/>
                <Route path="/login" element={<Login/>}/>
                <Route path="/signup" element={<SignUp/>}/>
            </Routes>
        </BrowserRouter>
    </>
  );
}

export default App;
