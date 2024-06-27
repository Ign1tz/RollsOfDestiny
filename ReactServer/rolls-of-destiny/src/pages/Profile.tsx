import {profile} from "../types/profileTypes";
import Button from '@mui/material/Button'
import TopAppBar from "../bars/TopAppBar";
import "../css/Profile.css";


export default function Profile({user}: { user: profile }) {

    return (
        <>
            <TopAppBar loggedIn={true}></TopAppBar>
            <div className = "profilePage">
                <div className = "profileDiv">
                    <img src={user.picture} alt={"profile picture"}/>
                    <h1>{user.username}</h1>
                    <div className = "list">
                        <h4>Rating: {user.rating}</h4>
                        <h4>Friends: 5</h4>
                    </div>
                    <h4>{user.biography}</h4>
                    <div className = "lastMatches">
                        <h3>Last Matches:</h3>
                        <h4>Coming soon...</h4>
                    </div>
                </div>
            </div>
        </>
    )
}