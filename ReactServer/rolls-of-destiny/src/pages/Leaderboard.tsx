import React from 'react';
import TopAppBar from "../bars/TopAppBar";
import { profile } from "../types/profileTypes";
import '../css/Leaderboard.css';
import testImage from "../soundtracks/testImage.png";

export default function Leaderboard({ loggedIn }: { loggedIn: boolean }) {

    let users: profile[] = [
        { username: "Bernd", rating: 839, profilePicture: testImage},
        { username: "Anna", rating: 902, profilePicture: testImage},
        { username: "Carlos", rating: 756, profilePicture: "https://via.placeholder.com/100"},
        { username: "Diana", rating: 820, profilePicture: testImage},
        { username: "Edward", rating: 890, profilePicture: "https://via.placeholder.com/100"}
    ];

    users = users.sort((a, b) => b.rating - a.rating);


    return (
        <>
            <TopAppBar loggedIn={loggedIn} />
            <div className="leaderboard-container">
                <h1 className="leaderboard-title">Leaderboard</h1>
                <ul className="leaderboard">
                    {users.map((user, index) => (
                        <li key={index} className="leaderboard-item">
                                <div className={"someItemsFromLeaderbord"}>
                                    <img src={user.profilePicture} alt={user.username} className="leaderboard-picture" />
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
