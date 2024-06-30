import "../css/Home.css"
import {useState} from "react";
import TopAppBar from "../bars/TopAppBar";
import {Link} from "react-router-dom";
import Button from '@mui/material/Button';
import HomeScreenButtonGroup from "../components/homeScreenButtonGroup";
import background from "../images/game.jpg";

export default function Home({loggedIn, setLoggedIn, setGameInfo, websocket, setWebsocket}: { loggedIn: boolean, setLoggedIn: Function, setGameInfo: Function, websocket: WebSocket|undefined, setWebsocket: Function }) {
    const [playOpened, setPlayOpened] = useState<boolean>(false)
    const [connected, setConnected] = useState(false)
    const [websoketId, setWebsoketId] = useState("")

    const relocate = () => {
        window.location.href = "/profile";
        setLoggedIn(!loggedIn)
    };

    function visibleButtons() {
        if (playOpened) {
            return (
                <>
                    <HomeScreenButtonGroup setPlayOpened={setPlayOpened}
                                           playOpened={playOpened}
                                           connected={connected}
                                           websocket={websocket}
                                           setWebsocket={setWebsocket}
                                           setConnected={setConnected}
                                           websocketId={websoketId}
                                           setWebsocketId={setWebsoketId}
                    setGameInfo={setGameInfo}/>
                </>
            )
        } else {
            return (
                <Button variant="contained" color="secondary" onClick={() => setPlayOpened(!playOpened)}> Play </Button>
            )
        }
    }

    return (
        <>
            <header>
                <TopAppBar loggedIn={loggedIn} />
            </header>
            <div className="homepage">
                <div className="textAndButtons">
                    <div className="homeText">
                        <h1>Rolls of Destiny</h1>
                        <h3>A game made by</h3>
                        <p className={"contributor"}><Link to={"https://github.com/Ign1tz"}>Moritz Pertl</Link></p>
                        <p className={"contributor"}><Link to={"https://github.com/LukasBrezina"}>Lukas Brezina</Link>
                        </p>
                        <p className={"contributor"}><Link to={"https://github.com/Sweisser7"}>Simon Weisser</Link></p>
                    </div>
                    <div className="homeButtons">
                        {visibleButtons()}
                    </div>
                </div>
            </div>
        </>
    )
}