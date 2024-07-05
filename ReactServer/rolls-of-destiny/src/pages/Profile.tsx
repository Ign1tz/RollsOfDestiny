import { profile } from "../types/profileTypes";
import Button from '@mui/material/Button';
import TopAppBar from "../bars/TopAppBar";
import "../css/Profile.css";
import {useEffect, useState} from "react";
import {authFetch} from "../auth";

export default function Profile() {

    const matches = ["Win", "Win", "Lose", "Lose", "Win", "Lose", "Win", "Lose"];
    const [friends, setFriends] = useState<profile[]>([
    ])
    useEffect(() => {
        authFetch("http://10.0.0.2:9090/getFriends").then(response => {
            console.log(response); return response.json()
        }).then(response => {
            console.log(response)
            setFriends(response.friends)
        })
    }, []);

    return (
        <>
            <TopAppBar loggedIn={true}></TopAppBar>
            <div className="profilePage">
                <div className="profileDiv">
                    <img src={"data:image/jpeg;base64," + JSON.parse(sessionStorage.getItem("userInfo") || "").profilePicture} alt={"profile picture"} />
                    <h1>{JSON.parse(sessionStorage.getItem("userInfo") || "").username}</h1>
                    <div className="list">
                        <h4>Rating: {JSON.parse(sessionStorage.getItem("userInfo") || "").rating}</h4>
                        <h4>Friends: {friends.length}</h4>
                    </div>
                    {false && <div className="lastMatches">
                        <h3>Last Matches:</h3>
                       <div className="matchesList">
                            {matches.map((result, index) => (
                                <div
                                    key={index}
                                    className={`match ${result.toLowerCase()}`}
                                >
                                    {result}
                                </div>
                            ))}
                        </div>
                    </div>}
                </div>
            </div>
        </>
    );
}
