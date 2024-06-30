import React, {useEffect, useState} from 'react';
import './App.css';
import Home from "./pages/Home";
import {BrowserRouter, Routes, Route, Navigate} from "react-router-dom"
import Profile from "./pages/Profile";
import {profile} from "./types/profileTypes";
import Game from "./pages/Game";
import Login from "./pages/Login";
import SignUp from "./pages/SignUp";
import Rules from "./pages/Rules";
import Settings from "./pages/Settings";
import testImage from "./soundtracks/testImage.png"
import Leaderboard from "./pages/Leaderboard";
import Friendlist from "./pages/Friendlist";

function App() {
    let p: profile = {username: "Bernd", picture: testImage, biography: "This is the bio", rating:839}
    const [loggedIn, setLoggedIn] = useState<boolean> (true)
    const [ingame, setIngame] = useState<boolean> (false)
    const [gameInfo, setGameInfo] = useState<string> ("")
    const [websocket, setWebsocket] = useState<WebSocket>()

    useEffect(() => {
        console.log("ws", websocket)
    }, [websocket]);


    return (
    <>
        <BrowserRouter>
            <Routes>
                <Route index element={<Home loggedIn={loggedIn} setLoggedIn={setLoggedIn} setGameInfo={setGameInfo} websocket={websocket} setWebsocket={setWebsocket}/>}/>
                <Route path={"/profile"} element={<Profile user={p}/>}/>
                <Route path="/leaderboard" element={<Leaderboard loggedIn={loggedIn}/>}/>
                <Route path="/friendlist" element={<Friendlist loggedIn={loggedIn}/>}/>
                <Route path={"/game" } element={loggedIn ? <Game/> : <Login/>}/>
                <Route path="/login" element={<Login/>}/>
                <Route path="/signup" element={<SignUp/>}/>
                <Route path="/rules" element={<Rules/>}/>
                <Route path="/settings" element={loggedIn ? <Settings profile={p}/> : <Login/>}/>
            </Routes>
        </BrowserRouter>
    </>
  );
}

export default App;
