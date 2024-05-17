import { Button } from "@mui/material";
import Grid from "../components/Grid";

export default function Game() {
    return (
        <>
            <h1> Welcome to the Game!</h1>
            <Button variant="contained" onClick={() => window.location.href="/"}>Back</Button>
            <Grid/>
        </>
    )
}