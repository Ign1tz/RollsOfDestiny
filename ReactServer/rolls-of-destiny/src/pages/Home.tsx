import "../css/Home.css"
import {useEffect, useState} from "react";

export default function Home() {
    const [loggedIn, setLoggedIn] = useState<boolean> (false)

    useEffect(() => {
        console.log(loggedIn)
    }, [])
    console.log("hello")
    const relocate = () => {
        // window.location.href = "/profile";
        setLoggedIn(!loggedIn)
    };
    return (
        <div className="homepage">
            <div className="homeText">
                <h1>Rolls of Destiny</h1>
                <h3>A game made by</h3>
                <p className={"contributor"}><a href={"https://github.com/Ign1tz"}>Moritz Pertl</a></p>
                <p className={"contributor"}><a href={"https://github.com/LukasBrezina"}>Lukas Brezina</a></p>
                <p className={"contributor"}><a href={"https://github.com/Sweisser7"}>Simon Weisser</a></p>
            </div>
            <div className="profile">
                <button onClick={relocate}>Profile</button>
            </div>
            <div className="homeButtons">
                <button>Play</button>
                <button>Settings</button>
            </div>
        </div>
    )
}