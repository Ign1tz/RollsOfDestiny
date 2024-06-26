import {profile} from "../types/profileTypes";
import Button from '@mui/material/Button'
import TopAppBar from "../bars/TopAppBar";

export default function Profile({user}: { user: profile }) {

    return (
        <>
            <TopAppBar loggedIn={true}></TopAppBar>
        </>
    )
}