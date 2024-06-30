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
import {endResults, enemyInfo, messageBody, yourInfo} from "../types/gameTypes";




export default function Game() {
    localStorage.setItem("gameInfo", "")
    let startGameInfo = localStorage.getItem("gameInfo")
    if (!startGameInfo) {
        startGameInfo = '{"gameid": "", "YourInfo": { "WebsocketId": "", "Username": "Host"}, "EnemyInfo": { "WebsocketId":"", "Username": ""}}'
    }
    const [websocket, setWebsocket] = useState<WebSocket>(new WebSocket('ws://localhost:8080/ws'))
    const [websocketId, setWebsocketId] = useState("")
    const [connected, setConnected] = useState(false)
    const [gameInfoJson, setGameInfoJson] = useState(JSON.parse(startGameInfo))
    const [gameId, setGameId] = useState("")
    const [gameInfo, setGameInfo] = useState<messageBody>({} as messageBody)
    const [rolled, setRolled] = useState(false)
    const [placed, setPlaced] = useState(false)
    const [player1Score, setPlayer1Score] = useState(0);
    const [player2Score, setPlayer2Score] = useState(0);
    const [diceRoll, setDiceRoll] = useState<number | null>(null);
    const [disableRoll, setDisableRoll] = useState(false);
    const [canPlace, setCanPlace] = useState(true);
    const [isPaused, setIsPaused] = useState(false);
    const [confirmSurrender, setConfirmSurrender] = useState(false);
    const [yourInfo, setYourInfo] = useState<yourInfo | null>( null )
    const [enemyInfo, setEnemyInfo] = useState<enemyInfo | null>( null )
    const [gameEnded, setGameEnded] = useState(false)
    const [endResults, setEndResults] = useState<endResults | null>()
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
            //websocket.send("test")
        }
    }, [connected])

    useEffect(() => {
        if (websocketId !== "") {
            console.log("queuing websocket")
            if (sessionStorage.getItem("GameType") === "bot"){
                startBot()
            }else {
                queueForGame()
            }
        }
    }, [websocketId])

    if (websocket) {
        websocket.onmessage = (e) => {
            console.log(e.data)
            let message = JSON.parse(e.data)
            console.log("got a message")
            console.log(message)
            if (message.info == "connected") {
                console.log("connected")
                setConnected(true)
                websocket.send(JSON.stringify({type: "id"}))
            } else if (message.info == "id") {
                console.log("id:", message.message.id)
                setWebsocketId(message.message.id)
            } else if (message.info == "gameInfo") {
                console.log(message.message.gameInfo)
                localStorage.setItem("gameInfo", message.message.gameInfo)
                setGameInfoJson(message.message.gameInfo)
                setGameId(message.message.gameInfo.gameid)
                console.log("setGameInfo")
                setGameInfo(message.message.gameInfo)
            } else if (message.info == "gameEnded") {
                console.log(message.message.gameInfo)
                localStorage.setItem("gameInfo", message.message.gameInfo)
                setGameInfoJson(message.message.gameInfo)
                setGameId(message.message.gameInfo.gameid)
                console.log("endResults", message.message.endResults)
                setGameInfo(message.message.gameInfo)
                setGameEnded(true)
                setEndResults(message.message.endResults)
            }
        }
    }

    const handleColumnClick = (key: number) => {
        console.log("handleColumnClicK", connected)
        if (websocket && connected && gameInfo){
            console.log(gameId)
            setPlaced(true)
            websocket.send(JSON.stringify({type: sessionStorage.getItem("GameType") + "PickColumn", messageBody: key.toString(), gameId: gameId}))
        }
    };

    async function queueForGame() {
        let userinfo = sessionStorage.getItem("userInfo")
        if (userinfo) {
            console.log("test")
            const response = await fetch("http://localhost:8080/queue", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({
                    userid: JSON.parse(userinfo).userid, websocketconnectionid: websocketId
                })
            });
        }else {
            window.location.href = "/login"
        }
    }

    async function startBot() {
        let userinfo = sessionStorage.getItem("userInfo")
        if (userinfo) {
            console.log("test")
            const response = await fetch("http://localhost:8080/startBot", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({
                    userid: JSON.parse(userinfo).userid, websocketconnectionid: websocketId
                })
            });
        }else {
            window.location.href = "/login"
        }
    }

    const parseRoll = (roll: string) : 1 | 2 | 3 | 4 | 5 | 6 | undefined => {
        switch (roll) {
            case "1":
                return 1
            case "2":
                return 2
            case "3":
                return 3
            case "4":
                return 4
            case "5":
                return 5
            case "6":
                return 6
        }
    }

    useEffect(() => {
        if (gameInfo.ActivePlayer){
            console.log(gameInfo.ActivePlayer.active)
            if (gameInfo.ActivePlayer.active) {
                console.log("setFalse")
                setRolled(false)
                setPlaced(false)
            }else{
                setRolled(true)
                setPlaced(true)
            }
            setYourInfo(gameInfo.YourInfo)
            setEnemyInfo(gameInfo.EnemyInfo)
        }
    }, [gameInfo]);

    useEffect(() => { console.log(rolled)}, [rolled])

    return (
        <div className="gameDivision" style={{
            backgroundImage: `url(${background})`,
            backgroundSize: 'cover',
            backgroundPosition: 'center',
            height: '100%',
            width: '100%'
        }} >
            <div className="header">
                <Button variant="contained" onClick={togglePause}>
                    Pause
                </Button>
            </div>
            <div className="content">
                <Modal open={gameEnded} >
                    <div className="pauseMenu">
                        <h2>Game finished</h2>
                        <h2>{endResults?.youWon}</h2>
                        <h2>Score</h2>
                        <h2>{endResults?.yourScore} to {endResults?.enemyScore}</h2>
                        <Button variant="contained" onClick={() => window.location.href = "/"}>
                            Menu
                        </Button>
                    </div>
                </Modal>
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
                            <p>Score: <span id="player1Score">{enemyInfo ? enemyInfo?.Score : 0}</span></p>
                        </div>
                        <img src={player1.picture} alt={player1.username}/>
                    </div>
                    <div className="playerActions">
                        <div className="playerCards">
                            <h3>Deck </h3>
                            <SimpleBox diceValue={null}/>
                        </div>
                        <div className="grid">
                            <OpponentGrid grid={enemyInfo ? enemyInfo: null}/>
                        </div>
                        <div className="diceWrapper">
                            <Dice onRoll={(value) => console.log(value)} defaultValue={6} size={100}
                                  cheatValue={gameInfo.ActivePlayer ? parseRoll(gameInfo?.ActivePlayer.roll): undefined} disabled = {true}/>
                        </div>
                    </div>
                </div>
                <div className="divider"></div>
                <div className="playerSection">
                    <div className="playerActions">
                        <div className="diceWrapper">
                            <Dice onRoll={(value) => {
                                console.log(value); setRolled(true)
                            }} defaultValue={6} size={100}
                                  cheatValue={gameInfo.ActivePlayer ? parseRoll(gameInfo?.ActivePlayer.roll): undefined} disabled={(gameInfo.ActivePlayer ? !gameInfo?.ActivePlayer.active : true) || rolled}/>
                        </div>
                        <div className="grid">
                            <Grid handleColumnClick={handleColumnClick} active={rolled && !placed} grid={yourInfo ? yourInfo: null}/>
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
                            <p>Score: <span id="player2Score">{yourInfo ? yourInfo?.Score : 0}</span></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
