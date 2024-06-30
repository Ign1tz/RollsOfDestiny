import React from 'react';
import TopAppBar from "../bars/TopAppBar";
import { profile } from "../types/profileTypes";
import '../css/Friendlist.css';
import onlineImage from "../images/green.jpeg"
import offlineImage from "../images/red.jpeg"

export default function Friendlist({ loggedIn }: { loggedIn: boolean }) {

    const users: profile[] = [
        { username: "Bernd", rating: 839, profilePicture: "https://via.placeholder.com/100"},
        { username: "Anna", rating: 902, profilePicture: "https://via.placeholder.com/100"},
        { username: "Carlos", rating: 756, profilePicture: "https://via.placeholder.com/100"},
        { username: "Diana", rating: 820, profilePicture: "https://via.placeholder.com/100" },
        { username: "Edward", rating: 890, profilePicture: "https://via.placeholder.com/100"}
    ];

    const online = true;

    // TODO: remove from friendlist button


    return (
        <>
            <TopAppBar loggedIn={loggedIn} />
            <div className="friendlist-container">
                <h1 className="friendlist-title">Your Friends</h1>
                <ul className="friendlist">
                    {users.map((user, index) => (
                        <li key={index} className="friendlist-item">
                            <div className="friendlist-image-info">
                                <img src={user.profilePicture} alt={user.username} className="leaderboard-picture"/>
                                <div className="friendlist-info">
                                    <h2 className="friendlist-username">{user.username}</h2>
                                    <p className="friendlist-rating">Rating: {user.rating}</p>
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
