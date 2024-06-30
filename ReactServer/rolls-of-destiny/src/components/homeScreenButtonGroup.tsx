import Button from "@mui/material/Button";
import {useEffect} from "react";
//import {ws} from "../pages/Game";

type ButtonGroupPatameter = {
    setPlayOpened: Function,
    playOpened: boolean
}



export default function HomeScreenButtonGroup({
                                                  setPlayOpened,
                                                  playOpened,
                                              }: ButtonGroupPatameter) {

    return (
        <>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => {sessionStorage.setItem("GameType", "bot")
                        window.location.href = "/game"
                    }}>
                Play Against Bot
            </Button>
            <br/>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => {sessionStorage.setItem("GameType", "")
                        window.location.href = "/game"
                    }}>
                Play Against Real Enemy
            </Button>
            <br/>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => setPlayOpened(!playOpened)}> Back </Button>
        </>
    )
}