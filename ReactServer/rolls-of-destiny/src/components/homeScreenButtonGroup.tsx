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

    /*async function startWebsocket() {
        console.log("Starting websocket")
        setWebsocket(new WebSocket('http://localhost:8080/ws'))
        localStorage.setItem("gameInfo", '{"gameid": "", "YourInfo": { "WebsocketId": "", "Username": "Host"}, "EnemyInfo": { "WebsocketId":"", "Username": ""}}')
    }

    useEffect(() => {
        console.log("queuing websocket")
        if (websocketId !== "") {
            queueForGame()
        }
    }, [websocketId])

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
        window.location.href = "/game"
    }

    useEffect(() => {
        if (connected && websocket) {
            websocket.send("test")
        }
    }, [connected])

    if (websocket) {
        console.log(websocket)
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
    }*/

    return (
        <>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => window.location.href = "/game"}>
                Play Against Bot
            </Button>
            <br/>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => window.location.href = "/game"}>
                Play Against Real Enemy
            </Button>
            <br/>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => setPlayOpened(!playOpened)}> Back </Button>
        </>
    )
}