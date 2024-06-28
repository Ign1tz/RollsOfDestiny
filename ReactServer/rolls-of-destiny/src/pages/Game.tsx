import { Button, Modal } from "@mui/material"; // Assuming you're using Material-UI components

import Grid from "../components/Grid";
import OpponentGrid from "../components/OpponentGrid";
import Dice from "react-dice-roll";
import SimpleBox from "../components/SimpleBox";
import {profile} from "../types/profileTypes";
import "../css/Game.css";
import background from "../images/game.jpg";
import testImage from "../images/1.png";
import {useEffect, useState} from "react";


export default function Game() {
    localStorage.setItem("gameInfo", "")
    console.log(localStorage.getItem("gameInfo"))
    let gameInfo = localStorage.getItem("gameInfo")
    if (!gameInfo) {
        gameInfo = '{"gameid": "", "YourInfo": { "WebsocketId": "", "Username": "Host"}, "EnemyInfo": { "WebsocketId":"", "Username": ""}}'
    }
    const [websocket, setWebsocket] = useState<WebSocket>(new WebSocket('ws://localhost:8080/ws'))
    const [websocketId, setWebsocketId] = useState("")
    const [connected, setConnected] = useState(false)
    const [gameInfoJson, setGameInfoJson] = useState(JSON.parse(gameInfo))
    const [gameId, setGameId] = useState("")
    const [player1Score, setPlayer1Score] = useState(0);
    const [player2Score, setPlayer2Score] = useState(0);
    const [diceRoll, setDiceRoll] = useState<number | null>(null);
    const [disableRoll, setDisableRoll] = useState(false);
    const [canPlace, setCanPlace] = useState(true);
    const [isPaused, setIsPaused] = useState(false);
    const [confirmSurrender, setConfirmSurrender] = useState(false);

    const [rollValue, setRollValue] = useState< 1 | 2 | 3 | 4 | 5 | 6 | undefined >(undefined);

    const handleRoll = (player: 'player1' | 'player2', value: number) => {
        setDiceRoll(value);
        setDisableRoll(true);
    };

    const togglePause = () => {
        setIsPaused(!isPaused);
    };

    const handleQuit = () => {
        window.location.href = "/";
    };

    const toggleSurrender = () => {
        setConfirmSurrender(!confirmSurrender)
    };


    let player1: profile = {
        username: gameInfoJson.EnemyInfo.Username,
        rating: 3450913,
        picture: testImage,
        biography: "Player 1's bio"
    };

    let player2: profile = {
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
            } else if (e.data.includes("id:")) {
                console.log(e.data.split(":")[e.data.split(":").length - 1])
                setWebsocketId(e.data.split(":")[e.data.split(":").length - 1])
            } else if (e.data.includes("{")) {
                console.log(e.data)
                localStorage.setItem("gameInfo", e.data)
                setGameInfoJson(JSON.parse(e.data))
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


    return (<div className="gameDivision">
            <div className="header">
                <Button variant="contained" onClick={togglePause}>
                    Pause
                </Button>
            </div>
            <div className="content">
                <Modal open={confirmSurrender} onClose={toggleSurrender}>
                    <div className="confirmSurrenderMenu">
                        <Button variant="contained" onClick={() => {toggleSurrender(); togglePause()}}>
                            Cancel
                        </Button>
                        <Button variant="contained" onClick={handleQuit}>
                            Confirm Surrender
                        </Button>
                    </div>
                </Modal>
                <Modal open={isPaused} onClose={togglePause}>
                    <div className="pauseMenu">
                        <h2>Pause Menu</h2>
                        <Button variant="contained" onClick={togglePause}>
                            Continue playing
                        </Button>
                        <Button variant="contained" onClick={() => console.log("Go to Settings")}>
                            Settings
                        </Button>
                        <Button variant="contained" onClick={() => {toggleSurrender(); togglePause()}}>
                        Surrender
                        </Button>
                    </div>
                </Modal>
                <div className="playerSection">
                    <div className="playerInfoOpp">
                        <div className="score">
                            <p>Score: <span id="player1Score">{player1Score}</span></p>
                        </div>
                        <div className="playerInfoUsernameRating">
                            <h2>{player1.username + " (Opponent)"}</h2>
                            <p>Rating: {player1.rating}</p>
                        </div>
                        <img src={player1.picture} alt={player1.username}/>
                    </div>
                    <div className="playerActions">
                        <div className="playerCards">
                            <h3>Deck </h3>
                            <SimpleBox diceValue={null}/>
                        </div>
                        <div className="grid">
                            <OpponentGrid diceRoll={diceRoll}/>
                        </div>
                        <div className="diceWrapper">
                            <Dice defaultValue={6} size={100} cheatValue={undefined} disabled={true}/>
                        </div>
                    </div>
                </div>
                <div className="divider"></div>
                <div className="playerSection">
                    <div className="playerActions">
                        <div className="diceWrapper">
                            <Dice onRoll={(value) => handleRoll('player2', value)} defaultValue={6} size={100}
                                  cheatValue={rollValue} disabled={disableRoll}/>
                        </div>
                        <div className="grid">
                            <Grid canPlace={canPlace} setCanPlace={setCanPlace} diceRoll={diceRoll}/>
                        </div>
                        <div className="playerCards">
                            <h3>Deck</h3>
                            <SimpleBox diceValue={null}/>
                        </div>
                    </div>
                    <div className="playerInfo">
                        <img src={player2.picture} alt={player2.username}/>
                        <div className="playerInfoUsernameRating">
                            <h2>{player2.username + " (You)"}</h2>
                            <p>Rating: {player2.rating}</p>
                        </div>
                        <div className="score">
                            <p>Score: <span id="player2Score">{player2Score}</span></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
