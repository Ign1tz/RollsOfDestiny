import {Button} from "@mui/material";
import Grid from "../components/Grid";
import Dice from "react-dice-roll";
import SimpleBox from "../components/SimpleBox";
import {profile} from "../types/profileTypes";
import "../css/Game.css";
import {useEffect, useState} from "react";


/*export const ws = new WebSocket('http://localhost:8080/ws');

ws.onopen = () => {
    console.log('WebSocket connected')
    ws.send(JSON.stringify({purpose:"login", UserId:"testuser1", Username:"testuser1"}))
};
ws.onclose = () => console.log('WebSocket disconnected');
*/
export default function Game() {
    console.log(localStorage.getItem("gameInfo"))
    let gameInfo = localStorage.getItem("gameInfo")
    if (!gameInfo) {
        gameInfo = '{"gameid": "", "YourInfo": { "WebsocketId": "", "Username": "Host"}, "EnemyInfo": { "WebsocketId":"", "Username": ""}}'
    }
    const [websocket, setWebsocket] = useState<WebSocket>(new WebSocket('http://localhost:8080/ws'))
    const [websocketId, setWebsocketId] = useState("")
    const [connected, setConnected] = useState(false)
    const [gameInfoJson, setGameInfoJson] = useState(JSON.parse(gameInfo))
    const [gameId, setGameId] = useState("")
    const player1: profile = {
        username: gameInfoJson.EnemyInfo.Username,
        rating: 3450913,
        picture: "/path/to/player1.jpg",
        biography: "Player 1's bio"
    };

    const player2: profile = {
        username: gameInfoJson.YourInfo.Username,
        rating: 1,
        picture: "/path/to/player2.jpg",
        biography: "Player 2's bio"
    };




    useEffect(() => {
        console.log("Starting websocket")
        //setWebsocket(prevWebsocket => ([...prevWebsocket, ...new WebSocket('http://localhost:8080/ws')]))
        localStorage.setItem("gameInfo", '{"gameid": "", "YourInfo": { "WebsocketId": "", "Username": "Host"}, "EnemyInfo": { "WebsocketId":"", "Username": ""}}')
    }, [])

    useEffect(() => {
        console.log(websocket)
        if (connected && websocket) {
            console.log("test")
            websocket.send("test")
        }
    }, [connected])

    useEffect(() => {
        console.log("queuing websocket")
        if (websocketId !== "") {
            queueForGame()
        }
    }, [websocketId])

    if (websocket) {
        websocket.onmessage = (e) => {
            console.log("go a message")
            console.log("message: " + e.data)
            if (e.data == "connected") {
                setConnected(true)
                websocket.send("id")
            } else if (e.data.includes("id:[::1]:")) {
                console.log(e.data.split(":")[e.data.split(":").length - 1])
                setWebsocketId(e.data.split(":")[e.data.split(":").length - 1])
            } else if (e.data.includes("{")) {
                console.log(e.data)
                localStorage.setItem("gameInfo", e.data)
            }
        }
    }

    async function queueForGame() {
        console.log("test")
        const response = await fetch("http://localhost:8080/queue", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({
                userid: "testasdasfasasdasd".split('').sort(function () {
                    return 0.5 - Math.random()
                }).join(''), websocketconnectionid: websocketId
            })
        });
    }


    return (
        <div className="gameDivision">
            <div className="header">
                <h1>Welcome to the Game!</h1>
                <Button variant="contained" onClick={() => window.location.href = "/"}>
                    Back
                </Button>
            </div>
            <div className="content">
                <div className="playerSection">
                    <div className="playerInfo">
                        <img src={player1.picture} alt={player1.username}/>
                        <div>
                            <h2>{player1.username}</h2>
                            <p>Rating: {player1.rating}</p>
                            <p>Score: <span id="player1Score">0</span></p>
                        </div>
                    </div>
                    <div className="playerActions">
                        <div className="diceWrapper">
                            <Dice onRoll={(value) => console.log(value)} defaultValue={6} size={100}
                                  cheatValue={undefined}/>
                        </div>
                        <Grid websocket={websocket}/>
                        <div className="playerCards">
                            {/* Placeholder for player's cards */}
                            <SimpleBox/>
                            <SimpleBox/>
                            <SimpleBox/>
                        </div>
                    </div>
                </div>
                <div className="playerSection">
                    <div className="playerActions">
                        <div className="diceWrapper">
                            <Dice onRoll={(value) => console.log(value)} defaultValue={6} size={100}
                                  cheatValue={undefined}/>
                        </div>
                        <Grid websocket={websocket}/>
                        <div className="playerCards">
                            <SimpleBox/>
                            <SimpleBox/>
                            <SimpleBox/>
                        </div>
                    </div>
                    <div className="playerInfo">
                        <img src={player2.picture} alt={player2.username}/>
                        <div>
                            <h2>{player2.username}</h2>
                            <p>Rating: {player2.rating}</p>
                            <p>Score: <span id="player2Score">0</span></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
