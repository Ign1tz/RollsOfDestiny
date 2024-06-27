import "../css/Home.css"
import {useState} from "react";
import TopAppBar from "../bars/TopAppBar";
import {Link} from "react-router-dom";
import Button from '@mui/material/Button';
import HomeScreenButtonGroup from "../components/homeScreenButtonGroup";
import {authFetch} from "../auth";
import {profile} from "../types/profileTypes";

export default function Home({loggedIn, setLoggedIn, setGameInfo, websocket, setWebsocket}: { loggedIn: boolean, setLoggedIn: Function, setGameInfo: Function, websocket: WebSocket|undefined, setWebsocket: Function }) {
    const [playOpened, setPlayOpened] = useState<boolean>(false)
    const [connected, setConnected] = useState(false)
    const [websoketId, setWebsoketId] = useState("")

    const relocate = () => {
        window.location.href = "/profile";
        setLoggedIn(!loggedIn)
    };

    useEffect(() => {
        if (localStorage.getItem("access_token")) {
            console.log("Access token", localStorage.getItem("access_token"));
            authFetch("http://localhost:9090/userInfo?username=" + sessionStorage.getItem("username")).then(r => {

                return r.json()
            }).then(response => {
                let profile: profile = response
                sessionStorage.setItem("userInfo", JSON.stringify(profile))
            })
        }
    }, []);

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
                <TopAppBar loggedIn={loggedIn}/>
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
                    {visibleButtons()}
                </div>
            </div>
            <footer style={{textAlign: "center", fontSize: "x-small"}}>
                Copyright
            </footer>
        </>
    )
}