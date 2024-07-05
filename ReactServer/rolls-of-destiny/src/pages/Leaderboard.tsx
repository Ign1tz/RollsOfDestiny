import React, {useEffect, useState} from 'react';
import TopAppBar from "../bars/TopAppBar";
import { profile } from "../types/profileTypes";
import '../css/Leaderboard.css';
import testImage from "../soundtracks/testImage.png";
import {authFetch} from "../auth";

export default function Leaderboard({ loggedIn }: { loggedIn: boolean }) {

    const [users, setUsers] = useState<profile[]>([])




    useEffect(() => {
        authFetch("http://10.0.0.2:9090/getTopTen").then((response) => response.json()).then(r =>
            setUsers(r.topTenPlayers.sort((a:profile, b:profile) => b.rating - a.rating))
        )
    }, []);



    return (
        <>
            <TopAppBar loggedIn={loggedIn} />
            <div className="leaderboard-container">
                <h1 className="leaderboard-title">Leaderboard</h1>
                <ul className="leaderboard">
                    {users.map((user, index) => (
                        <li key={index} className="leaderboard-item">
                                <div className={"someItemsFromLeaderbord"}>
                                    <img src={"data:image/jpeg;base64," + user.profilePicture} alt={user.username} className="leaderboard-picture" />
                                    <div className="leaderboard-info">
                                        <h2 className="leaderboard-username">{user.username}</h2>
                                        <p className="leaderboard-rating">Rating: {user.rating}</p>
                                    </div>
                                </div>
                                <h4>{index+1}</h4>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    );
}
