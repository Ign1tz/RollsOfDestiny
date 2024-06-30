import React from 'react';
import TopAppBar from "../bars/TopAppBar";
import { profile } from "../types/profileTypes";
import '../css/Leaderboard.css';

export default function Leaderboard({ loggedIn }: { loggedIn: boolean }) {

    const users: profile[] = [
        { username: "Bernd", rating: 839, profilePicture: "https://via.placeholder.com/100"},
        { username: "Anna", rating: 902, profilePicture: "https://via.placeholder.com/100"},
        { username: "Carlos", rating: 756, profilePicture: "https://via.placeholder.com/100" },
        { username: "Diana", rating: 820, profilePicture: "https://via.placeholder.com/100"},
        { username: "Edward", rating: 890, profilePicture: "https://via.placeholder.com/100"}
    ];

    return (
        <>
            <TopAppBar loggedIn={loggedIn} />
            <div className="leaderboard-container">
                <h1 className="leaderboard-title">Leaderboard</h1>
                <ul className="leaderboard">
                    {users.map((user, index) => (
                        <li key={index} className="leaderboard-item">
                            <img src={user.profilePicture} alt={user.username} className="leaderboard-picture" />
                            <div className="leaderboard-info">
                                <h2 className="leaderboard-username">{user.username}</h2>
                                <p className="leaderboard-rating">Rating: {user.rating}</p>
                            </div>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    );
}
