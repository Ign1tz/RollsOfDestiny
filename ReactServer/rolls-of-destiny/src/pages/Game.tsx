import {Button, Modal} from "@mui/material"; // Assuming you're using Material-UI components
import Grid from "../components/Grid";
import OpponentGrid from "../components/OpponentGrid";
import Dice from "react-dice-roll";
import SimpleBox from "../components/SimpleBox";
import {profile} from "../types/profileTypes";
import "../css/Game.css";
import React, {useEffect, useState} from "react";
import background_music from "../soundtracks/background_music.mp3";
import diceSound from "../soundtracks/dice-95077.mp3"
import ReactAudioPlayer from "react-audio-player";
import background from "../images/game.jpg";
import cardBack from "../images/cardBack.jpg"
import {activePlayer, card, endResults, enemyInfo, messageBody, yourInfo} from "../types/gameTypes";
import VolumeSlider from "../components/VolumeSlider";
import destroyColumnCard from "../cards/destroy_column.png"
import doubleManaCard from "../cards/double_mana.png"
import rollAgainCard from "../cards/roll_again.png"
import flipClockwiseCard from "../cards/rotate_grid.png"


export default function Game() {
    sessionStorage.setItem("gameInfo", "")
    console.log(sessionStorage.getItem("gameInfo"))
    let startGameInfo = sessionStorage.getItem("gameInfo")
    if (!startGameInfo) {
        startGameInfo = '{"gameid": "", "YourInfo": { "WebsocketId": "", "Username": "Host"}, "EnemyInfo": { "WebsocketId":"", "Username": ""}}'
    }
    const [websocket, setWebsocket] = useState<WebSocket>(new WebSocket('ws://localhost:8080/ws'))
    const [websocketId, setWebsocketId] = useState("")
    const [connected, setConnected] = useState(false)
    const [userID, setUserID] = useState("")
    const [gameId, setGameId] = useState("")
    const [gameInfo, setGameInfo] = useState<messageBody>({} as messageBody)
    const [rolled, setRolled] = useState(false)
    const [placed, setPlaced] = useState(false)
    const [player1Score, setPlayer1Score] = useState(0);
    const [isPaused, setIsPaused] = useState(false);
    const [confirmSurrender, setConfirmSurrender] = useState(false);
    const [yourInfo, setYourInfo] = useState<yourInfo | null>(null)
    const [enemyInfo, setEnemyInfo] = useState<enemyInfo | null>(null)
    const [gameEnded, setGameEnded] = useState(false)
    const [endResults, setEndResults] = useState<endResults | null>()
    const [volume, setVolume] = useState<number>(30);
    const [reroll, setReroll] = useState(false)

    window.onbeforeunload = () =>{
        let userinfo = sessionStorage.getItem("userInfo")
        if (userinfo) {
            console.log("test")
            const response = fetch("http://localhost:8080/deleteQueue", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({
                    userid: JSON.parse(userinfo).userid,
                    websocketconnectionid: websocketId,
                    username: JSON.parse(userinfo).username
                })
            });
        } else {
            window.location.href = "/login"
        }
        websocket.close()
    }

    const togglePause = () => {
        setIsPaused(!isPaused);
    };

    const handleQuit = () => {
        websocket.send(JSON.stringify({type: "surrender", message: "", gameId: gameId, userid: userID}))
        window.location.href = "/";
    };

    const toggleSurrender = () => {
        setConfirmSurrender(!confirmSurrender)
    };

    const [player1, setPlayer1] = useState<profile>({
        username: "",
        rating: 3450913,
        profilePicture: "testImage"
    })

    const [player2, setPlayer2] = useState<profile>({
        username: "default",
        rating: 0,
        profilePicture: "/path/to/player1.jpg"
    })

    useEffect(() => {
        console.log("Starting websocket")
        //setWebsocket(prevWebsocket => ([...prevWebsocket, ...new WebSocket('http://localhost:8080/ws')]))
        sessionStorage.setItem("gameInfo", '{"gameid": "", "YourInfo": { "WebsocketId": "", "Username": "Host"}, "EnemyInfo": { "WebsocketId":"", "Username": ""}}')
        let user = sessionStorage.getItem("userInfo")
        if (user) {
            setPlayer2(JSON.parse(user))
            setUserID(JSON.parse(user).userid)
        }

        setVolume(Number(sessionStorage.getItem("volume")) || 30)

    }, []);


    useEffect(() => {
        console.log(volume)
    }, [volume]);


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
            if (sessionStorage.getItem("GameType") === "bot") {
                startBot()
            } else if (sessionStorage.getItem("GameType") === "Friend") {
                queueForGameWithFriend()
            } else {
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
                websocket.send(JSON.stringify({type: "id", message: "", gameId: "", userid: ""}))
            } else if (message.info == "id") {
                console.log("id:", message.message.id)
                setWebsocketId(message.message.id)
            } else if (message.info == "gameInfo") {
                console.log(message.message.gameInfo)
                sessionStorage.setItem("gameInfo", message.message.gameInfo)
                setGameId(message.message.gameInfo.gameid)
                console.log("setGameInfo")
                setGameInfo(message.message.gameInfo)
            } else if (message.info == "gameEnded") {
                console.log(message.message.gameInfo)
                sessionStorage.setItem("gameInfo", message.message.gameInfo)
                setGameId(message.message.gameInfo.gameid)
                console.log("endResults", message.message.endResults)
                setGameInfo(message.message.gameInfo)
                setGameEnded(true)
                setEndResults(message.message.endResults)
            } else if (message.info === "rollAgain") {
                let newActivePlayer: activePlayer = {active: message.message.active, roll: message.message.roll}
                setGameInfo(prev => ({...prev, ActivePlayer: newActivePlayer}))
                setRolled(false)
            }
        }
    }

    const handleColumnClick = (key: number) => {
        console.log("handleColumnClicK", connected)
        if (websocket && connected && gameInfo) {
            console.log(gameId)
            setPlaced(true)
            websocket.send(JSON.stringify({
                type: sessionStorage.getItem("GameType") + "PickColumn",
                messageBody: key.toString(),
                gameId: gameId, userid: userID
            }))
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
                    userid: JSON.parse(userinfo).userid,
                    websocketconnectionid: websocketId,
                    username: JSON.parse(userinfo).username
                })
            });
        } else {
            window.location.href = "/login"
        }
    }

    async function queueForGameWithFriend() {
        let userinfo = sessionStorage.getItem("userInfo")
        console.log(JSON.parse(userinfo || ""))
        if (userinfo) {
            console.log("test")
            const response = await fetch("http://localhost:8080/queueFroGameWithFriend", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({
                    userid: JSON.parse(userinfo).userid,
                    websocketconnectionid: websocketId,
                    username: JSON.parse(userinfo).username,
                    FriendId: sessionStorage.getItem("FriendId") || ""
                })
            });
        } else {
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
        } else {
            window.location.href = "/login"
        }
    }

    const parseRoll = (roll: string): 1 | 2 | 3 | 4 | 5 | 6 | undefined => {
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
        console.log(gameInfo)
        if (gameInfo.ActivePlayer) {
            console.log(gameInfo.ActivePlayer.active)
            if (gameInfo.ActivePlayer.active) {
                console.log("setting", placed,  reroll)
                setRolled(reroll)
                setReroll(false)
                setPlaced(false)
            } else {
                setRolled(true)
                setPlaced(true)
            }
            setYourInfo(gameInfo.YourInfo)
            setEnemyInfo(gameInfo.EnemyInfo)
            console.log(gameInfo.EnemyInfo)
            setPlayer1(prevState => ({...prevState, username: gameInfo.EnemyInfo.Username}))
        }
    }, [gameInfo]);

    useEffect(() => {
        console.log(rolled)
    }, [rolled])

    // testing purposes

    type CardType = {
        name: string,
        mana: number,
        image: string,
        description: string
    };

    let cards: CardType[] = [
        {
            name: "Destroy Column",
            mana: 7,
            image: destroyColumnCard,
            description: "Destroy a column from your opponent."
        },
        {name: "Double Mana", mana: 8, image: doubleManaCard, description: "You get double mana."}
    ];

    function playCard(card: card) {
        setReroll(card.name !== "Roll Again")
        websocket.send(JSON.stringify({type: "playCard", messageBody: card.cardid, gameId: gameId, userid: userID}))
    }

    function getCardPicture(card: card) {
        switch (card.name) {
            case "Roll Again":
                return rollAgainCard
            case "Double Mana":
                return  doubleManaCard
            case "Destroy Column":
                return  destroyColumnCard
            case "Flip Clockwise":
                return flipClockwiseCard
        }
    }

    function playSoundEffect() {
        var snd = new Audio(diceSound); // buffers automatically when created
        snd.play();
    }

    return (
        <>
            <ReactAudioPlayer
                src={background_music}
                autoPlay={true}
                loop={true}
                volume={volume / 100}
            />
            <div className="gameDivision" style={{
                backgroundImage: `url(${background})`,
                backgroundSize: 'cover',
                backgroundPosition: 'center',
                height: '100vh',
                width: '100%'
            }}>
                <div className="content">
                    <Modal open={gameEnded}>
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
                            <Button variant="contained" onClick={() => {
                                toggleSurrender();
                                togglePause()
                            }}>
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
                            <VolumeSlider volume={volume} setVolume={setVolume}/>
                            <Button variant="contained" onClick={togglePause}>
                                Continue playing
                            </Button>
                            <Button variant="contained" onClick={() => {
                                toggleSurrender();
                                togglePause()
                            }}>
                                Surrender
                            </Button>
                        </div>
                    </Modal>
                    <div className={"opponentAndCards"}>
                        <div className={"opponentCards"}>
                            {(typeof enemyInfo).toString() != null &&
                                Array.from({length: enemyInfo?.deck.inHand || 0}, (_, k) => (
                                    <div className={"specificOpponentCard"}>
                                        <img src={cardBack} alt={"card image"}/>
                                    </div>
                                ))
                            }
                        </div>
                        <div className="playerInfoOpp">
                            <div className="score">
                                <p>Score: <span id="player1Score">{enemyInfo ? enemyInfo?.Score : 0}</span></p>
                                <h3>Mana : {enemyInfo?.mana || "0"}</h3>
                            </div>
                            <div className="playerInfoUsernameRating">
                                <h2>{player1.username}</h2>
                                <p>Rating: {player1.rating}</p>
                            </div>
                            <img src={player1.profilePicture} alt={player1.username}/>
                        </div>
                    </div>
                    <div className="playerSection">

                        <div className="playerActions">
                            <div className="playerCards">
                                <h3>Deck </h3>
                                <SimpleBox diceValue={null}/>
                            </div>
                            <div className="grid">
                                <OpponentGrid grid={enemyInfo ? enemyInfo : null}/>
                            </div>
                            <div className="diceWrapper">
                                <Dice onRoll={(value) => console.log(value)} defaultValue={6} size={100}
                                      cheatValue={gameInfo.ActivePlayer ? parseRoll(gameInfo?.ActivePlayer.roll) : undefined}
                                      disabled={true}/>
                            </div>
                        </div>
                    </div>
                    <div className={"dividerAndPauseButton"}>
                        <div className="divider"></div>
                        <Button variant="contained" color="secondary" onClick={togglePause}
                                style={{marginLeft: "10px", marginRight: "10px"}}>
                            Pause
                        </Button>
                        <div className="divider"></div>
                    </div>
                    <div className="playerSection">
                        <div className="playerActions">
                            <div className="diceWrapper">
                                <Dice sound={diceSound} onRoll={(value) => {
                                    console.log(value);
                                    setRolled(true)
                                    websocket.send(JSON.stringify({type: "rolled", message: "", gameId: gameId, userid: userID}))
                                }} defaultValue={6} size={100}
                                      cheatValue={gameInfo.ActivePlayer ? parseRoll(gameInfo?.ActivePlayer.roll) : undefined}
                                      disabled={(gameInfo.ActivePlayer ? !gameInfo?.ActivePlayer.active : true) || rolled}/>
                            </div>
                            <div className="grid">
                                <Grid handleColumnClick={handleColumnClick} active={rolled && !placed}
                                      grid={yourInfo ? yourInfo : null}/>
                            </div>
                            <div className="playerCards">
                                <h3>Deck</h3>
                                <SimpleBox diceValue={null}/>
                            </div>
                        </div>

                    </div>
                    <div className={"playerAndCards"}>
                        <div className="playerInfo">
                            <img src={"data:image/jpeg;base64," + sessionStorage.getItem("profilePicture")}
                                 alt={player2.username}/>
                            <div className="playerInfoUsernameRating">
                                <h2>{player2.username}</h2>
                                <p>Rating: {player2.rating}</p>
                            </div>
                            <div className="score">
                                <h3>Mana : {yourInfo?.mana || "0"}</h3>
                                <p>Score: <span id="player2Score">{yourInfo ? yourInfo?.Score : 0}</span></p>
                            </div>
                        </div>
                        <div className={"playerOwnCards"}>
                            {(typeof yourInfo).toString() != null && yourInfo?.deck.inHand.map((card) => (
                                <div className={"specificPlayerCard"} onClick={() => {
                                    if ((gameInfo.ActivePlayer ? gameInfo?.ActivePlayer.active : false)) {
                                        playCard(card)
                                    }
                                }}>
                                    <img src={getCardPicture(card)} alt={"card image"}/>
                                </div>
                            ))}
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}
