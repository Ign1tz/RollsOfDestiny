import Button from "@mui/material/Button";

export default function HomeScreenButtonGroup({setPlayOpened, playOpened}: {setPlayOpened: Function, playOpened: boolean}) {
    return (
        <>
            <Button className="buttonInHomeScreenGroup"  color = "secondary" variant="contained" onClick = {() => window.location.href = "/game"}>
                Play Against Bot
            </Button>
            <br/>
            <Button className="buttonInHomeScreenGroup" color = "secondary"  variant="contained" onClick = {() => window.location.href = "/game"}>
                Play Against Real Enemy
            </Button>
            <br/>
            <Button className="buttonInHomeScreenGroup"  color = "secondary" variant="contained" onClick = {() => setPlayOpened(!playOpened)}> Back </Button>
        </>
    )
}