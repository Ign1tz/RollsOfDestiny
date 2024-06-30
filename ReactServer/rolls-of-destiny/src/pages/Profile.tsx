import { profile } from "../types/profileTypes";
import Button from '@mui/material/Button';
import TopAppBar from "../bars/TopAppBar";
import "../css/Profile.css";

export default function Profile({ user }: { user: profile }) {

    const matches = ["Win", "Win", "Lose", "Lose", "Win", "Lose", "Win", "Lose"];

    return (
        <>
            <TopAppBar loggedIn={true}></TopAppBar>
            <div className="profilePage">
                <div className="profileDiv">
                    <img src={user.profilePicture} alt={"profile picture"} />
                    <h1>{user.username}</h1>
                    <div className="list">
                        <h4>Rating: {user.rating}</h4>
                        <h4>Friends: 5</h4>
                    </div>
                    <div className="lastMatches">
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
                    </div>
                </div>
            </div>
        </>
    );
}
