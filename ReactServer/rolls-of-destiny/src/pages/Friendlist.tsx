import React from 'react';
import TopAppBar from "../bars/TopAppBar";
import { profile } from "../types/profileTypes";
import '../css/Friendlist.css';
import onlineImage from "../images/green.jpeg"
import offlineImage from "../images/red.jpeg"

export default function Friendlist({ loggedIn }: { loggedIn: boolean }) {

    const users: profile[] = [
        { username: "Bernd", rating: 839, picture: "https://via.placeholder.com/100", biography: "Bio for Bernd" },
        { username: "Anna", rating: 902, picture: "https://via.placeholder.com/100", biography: "Bio for Anna" },
        { username: "Carlos", rating: 756, picture: "https://via.placeholder.com/100", biography: "Bio for Carlos" },
        { username: "Diana", rating: 820, picture: "https://via.placeholder.com/100", biography: "Bio for Diana" },
        { username: "Edward", rating: 890, picture: "https://via.placeholder.com/100", biography: "Bio for Edward" }
    ];

    const online = true;


    return (
        <>
            <TopAppBar loggedIn={loggedIn} />
            <div className="friendlist-container">
                <h1 className="friendlist-title">Your Friends</h1>
                <ul className="friendlist">
                    {users.map((user, index) => (
                        <li key={index} className="friendlist-item">
                            <div className="friendlist-image-info">
                                <img src={user.picture} alt={user.username} className="leaderboard-picture"/>
                                <div className="friendlist-info">
                                    <h2 className="friendlist-username">{user.username}</h2>
                                    <p className="friendlist-rating">Rating: {user.rating}</p>
                                    <p className="friendlist-biography">{user.biography}</p>
                                </div>
                            </div>
                            <img src={online ? onlineImage : offlineImage} className="online-status"
                                 alt={"online/offline status"}/>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    );
}
