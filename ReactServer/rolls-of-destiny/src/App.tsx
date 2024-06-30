import React, {useEffect, useState} from 'react';
import './App.css';
import Home from "./pages/Home";
import {BrowserRouter, Route, Routes} from "react-router-dom"
import Profile from "./pages/Profile";
import {profile} from "./types/profileTypes";
import Game from "./pages/Game";
import Login from "./pages/Login";
import SignUp from "./pages/SignUp";
import Leaderboard from "./pages/Leaderboard";
import Friendlist from "./pages/Friendlist";
import {authFetch} from "./auth";

function App() {
    let p: profile = {username: "Bernd", profilePicture: " H", rating: 839}
    const [loggedIn, setLoggedIn] = useState<boolean>(false)


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
                <Route index element={<Home loggedIn={loggedIn} setLoggedIn={setLoggedIn} setGameInfo={setGameInfo} websocket={websocket} setWebsocket={setWebsocket}/>}/>
                <Route path={"/profile"} element={<Profile user={p}/>}/>
                <Route path={"/game" } element={<Game/>}/>
                <Route path="/leaderboard" element={<Leaderboard loggedIn={loggedIn}/>}/>
                <Route path="/friendlist" element={<Friendlist loggedIn={loggedIn}/>}/>
                <Route path="/login" element={<Login/>}/>
                <Route path="/signup" element={<SignUp/>}/>
            </Routes>
        </BrowserRouter>
    </>
  );
}

export default App;
