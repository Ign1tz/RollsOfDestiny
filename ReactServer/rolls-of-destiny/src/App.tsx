import React from 'react';
import './App.css';
import Home from "./pages/Home";
import {BrowserRouter, Routes, Route, Navigate} from "react-router-dom"
import Profile from "./pages/Profile";
import {profile} from "./types/profileTypes";
import Game from "./pages/Game";

function App() {
    let p: profile = {username: "Bernd", biography: "H", picture:" H", rating:839}
    return (
    <>
        <BrowserRouter>
            <Routes>
                <Route index element={<Home/>}/>
                <Route path="/profile" element={<Profile user={p}/>}/>
                <Route path="/game" element={<Game/>}/>
            </Routes>
        </BrowserRouter>
    </>
  );
}

export default App;
