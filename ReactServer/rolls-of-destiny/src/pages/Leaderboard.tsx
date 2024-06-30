import React from 'react';
import TopAppBar from "../bars/TopAppBar";
import { profile } from "../types/profileTypes";
import '../css/Leaderboard.css';

export default function Leaderboard({ loggedIn }: { loggedIn: boolean }) {

    const users: profile[] = [
        { username: "Bernd", rating: 839, picture: "https://via.placeholder.com/100", biography: "Bio for Bernd" },
        { username: "Anna", rating: 902, picture: "https://via.placeholder.com/100", biography: "Bio for Anna" },
        { username: "Carlos", rating: 756, picture: "https://via.placeholder.com/100", biography: "Bio for Carlos" },
        { username: "Diana", rating: 820, picture: "https://via.placeholder.com/100", biography: "Bio for Diana" },
        { username: "Edward", rating: 890, picture: "https://via.placeholder.com/100", biography: "Bio for Edward" }
    ];

    return (
        <>
            <TopAppBar loggedIn={loggedIn} />
            <div className="leaderboard-container">
                <h1 className="leaderboard-title">Leaderboard</h1>
                <ul className="leaderboard">
                    {users.map((user, index) => (
                        <li key={index} className="leaderboard-item">
                            <img src={user.picture} alt={user.username} className="leaderboard-picture" />
                            <div className="leaderboard-info">
                                <h2 className="leaderboard-username">{user.username}</h2>
                                <p className="leaderboard-rating">Rating: {user.rating}</p>
                                <p className="leaderboard-biography">{user.biography}</p>
                            </div>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    );
}
