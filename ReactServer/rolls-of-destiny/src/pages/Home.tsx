import "../css/Home.css"
import {useEffect, useState} from "react";
import TopAppBar from "../bars/TopAppBar";
import homeScreenButtonGroup from "../bars/buttonGroup";
import {Link} from "react-router-dom";
import Button from '@mui/material/Button';
import {ButtonGroup} from "@mui/material";

export default function Home() {
    const [loggedIn, setLoggedIn] = useState<boolean> (false)
    const [playOpened, setPlayOpened] = useState<boolean> (false)

    useEffect(() => {
        console.log(loggedIn)
    }, [])

    console.log("hello")
    const relocate = () => {
        window.location.href = "/profile";
        // setLoggedIn(!loggedIn)
    };
    return (
        <>
            <header>
                <TopAppBar/>
            </header>
            <div className="homepage">
                <div className="homeText">
                    <h1>Rolls of Destiny</h1>
                    <h3>A game made by</h3>
                    <p className={"contributor"}><Link to={"https://github.com/Ign1tz"}>Moritz Pertl</Link></p>
                    <p className={"contributor"}><Link to={"https://github.com/LukasBrezina"}>Lukas Brezina</Link></p>
                    <p className={"contributor"}><Link to={"https://github.com/Sweisser7"}>Simon Weisser</Link></p>
                </div>
                <div className="homeButtons">
                    <Button variant="contained" onClick={() => window.location.href = "/game"}> Play </Button>
                </div>
                <footer>
                </footer>
            </div>
        </>
    )
}