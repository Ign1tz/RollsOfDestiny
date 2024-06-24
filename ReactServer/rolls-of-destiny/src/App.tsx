import React, {useState} from 'react';
import './App.css';
import Home from "./pages/Home";
import {BrowserRouter, Routes, Route, Navigate} from "react-router-dom"
import Profile from "./pages/Profile";
import {profile} from "./types/profileTypes";
import Game from "./pages/Game";
import Login from "./pages/Login";
import Leaderboard from "./pages/Leaderboard";
import Friendlist from "./pages/Friendlist";

function App() {
    let p: profile = {username: "Bernd", biography: "H", picture:" H", rating:839}
    const [loggedIn, setLoggedIn] = useState<boolean> (false)
    return (
    <>
        <BrowserRouter>
            <Routes>
                <Route index element={<Home loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>}/>
                <Route path={"/profile"} element={<Profile user={p}/>}/>
                <Route path={"/game" } element={<Game/>}/>
                <Route path="/login" element={<Login loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>}/>
                <Route path="/leaderboard" element={<Leaderboard loggedIn={loggedIn}/>}/>
                <Route path="/friendlist" element={<Friendlist loggedIn={loggedIn}/>}/>
            </Routes>
        </BrowserRouter>
    </>
  );
}

export default App;
