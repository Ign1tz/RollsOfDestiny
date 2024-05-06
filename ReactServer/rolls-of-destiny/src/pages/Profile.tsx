import {profile} from "../types/profileTypes";
import Button from '@mui/material/Button'

// mua.com for styling

export default function Profile({user}: { user: profile }) {
    const relocate = () => {
        window.location.href = "/"
    }
    return (
        <>
            <h1> {user.username} </h1>
            <Button variant="outlined" onClick={relocate}> Back button</Button>
        </>
    )
}