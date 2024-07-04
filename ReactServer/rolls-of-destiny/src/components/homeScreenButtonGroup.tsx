import {Modal, TextField} from "@mui/material";
import Button from "@mui/material/Button";
import React, {useState} from "react";
//import {ws} from "../pages/Game";
import "../css/something.css"


export default function HomeScreenButtonGroup() {


    const [playFriend, setPlayFriend] = useState(false)
    const [host, setHost] = useState(false)
    const [join, setJoin] = useState(false);
    const [id, setId] = useState("");

    return (
        <>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => {
                        sessionStorage.setItem("GameType", "bot")
                        window.location.href = "/game"
                    }}>
                Play Against Bot
            </Button>
            <br/>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => {
                        sessionStorage.setItem("GameType", "")
                        window.location.href = "/game"
                    }}>
                Play Against Real Enemy
            </Button>
            <br/>
            <Button className="buttonInHomeScreenGroup" color="secondary" variant="contained"
                    onClick={() => setPlayFriend(true)}> Play Against Friends </Button>
            <Modal open={playFriend} onClose={() => setPlayFriend(false)}>
                <div className={"errorMenu"}>
                    {!join && !host &&
                        <>
                            <Button variant="contained" color={"secondary"} onClick={() => setHost(true)}>
                                Host
                            </Button>
                            <Button variant="contained" color={"secondary"} style={{marginLeft: "20px"}}
                                    onClick={() => setJoin(true)}>
                                Join
                            </Button>
                        </>
                    }
                    <div className={"hostAndJoinModal"}>
                    {host &&
                        <>
                            <h4 style={{color: "white"}}>Test</h4>
                            <Button variant={"contained"} color={"success"}>Play</Button>
                        </>
                    }
                    {join &&
                        <>
                            <TextField required id="filled-basic" label="Username" variant="filled"
                                       value={id}
                                       onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                                           setId(event.target.value);
                                       }}/>
                            <Button style={{marginTop: "10px"}} variant={"contained"} color={"success"}>Play</Button>
                        </>
                    }
                    </div>
                </div>
            </Modal>
        </>
    )
}