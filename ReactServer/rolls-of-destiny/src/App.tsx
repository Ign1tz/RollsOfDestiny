import React, {useEffect, useState} from 'react';
import './App.css';
import Home from "./pages/Home";
import {BrowserRouter, Route, Routes} from "react-router-dom"
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
import {authFetch} from "./auth";
import Decks from "./pages/Decks";
import LandingPage from "./pages/LandingPage";

function App() {
    let p: profile = {username: "Bernd", profilePicture: "testImage", rating:839}
    const [loggedIn, setLoggedIn] = useState<boolean> (false)
    const [ingame, setIngame] = useState<boolean> (false)
    const [gameInfo, setGameInfo] = useState<string> ("")
    const [websocket, setWebsocket] = useState<WebSocket>()
    const [fetched, setFetched] = useState<boolean> (false)

    useEffect(() => {
        const tempLoggedIn = sessionStorage.getItem("loggedIn")
        if (localStorage.getItem("access_token") && tempLoggedIn !== "true") {
            authFetch("http://10.0.0.2:9090/isLoggedIn").then(response => {
                setLoggedIn(response.status === 200)
                sessionStorage.setItem("loggedIn", "true")
                setFetched(true)
                console.log("userInfo")
                if (response.status === 200) {
                    authFetch("http://10.0.0.2:9090/userInfo").then(r => {

                        return r.json()
                    }).then(response => {
                        sessionStorage.setItem("userInfo", JSON.stringify(response))

                    })
                }
            })
        } else {
            setLoggedIn(tempLoggedIn === "true")
            setFetched(true)
        }
    }, []);

    return (
    <>
        { fetched && (typeof loggedIn).toString() !== "undefined" &&
            <BrowserRouter>
                <Routes>
                    <Route index element={ loggedIn ? <Home loggedIn={loggedIn || false} setLoggedIn={setLoggedIn}/> : <LandingPage loggedIn={loggedIn || false}/>}/>
                    <Route path={"/profile"} element={<Profile/>}/>
                    <Route path="/leaderboard" element={<Leaderboard loggedIn={loggedIn || false}/>}/>
                    <Route path="/friendlist" element={<Friendlist loggedIn={loggedIn || false}/>}/>
                    <Route path={"/game" } element={loggedIn ? <Game/> : <Login/>}/>
                    <Route path="/login" element={!loggedIn ? <Login/> : <Home loggedIn={loggedIn || false} setLoggedIn={setLoggedIn}/>}/>
                    <Route path="/signup" element={!loggedIn ? <SignUp/> : <Home loggedIn={loggedIn || false} setLoggedIn={setLoggedIn}/>}/>
                    <Route path="/rules" element={<Rules loggedIn={loggedIn || false}/>}/>
                    <Route path="/settings" element={loggedIn ? <Settings profile={p}/> : <Login/>}/>
                    <Route path="/decks" element={loggedIn ? <Decks/> : <Login/>}/>
                    <Route path="/landingpage" element={<LandingPage loggedIn={loggedIn || false}/>}/>
                </Routes>
            </BrowserRouter>
        }
    </>
  );
}

export default App;
