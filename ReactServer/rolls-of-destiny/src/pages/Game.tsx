import { useState } from "react";
import { Button, Modal } from "@mui/material"; // Assuming you're using Material-UI components

import Grid from "../components/Grid";
import OpponentGrid from "../components/OpponentGrid";
import Dice from "react-dice-roll";
import SimpleBox from "../components/SimpleBox";
import { profile } from "../types/profileTypes";
import "../css/Game.css";
import background from "../images/game.jpg";
import testImage from "../images/1.png";

export default function Game() {
    const player1: profile = {
        username: "Lukas",
        rating: 3450913,
        picture: testImage,
        biography: "Player 1's bio"
    };

    const player2: profile = {
        username: "Moritz",
        rating: 1,
        picture: "/path/to/player2.jpg",
        biography: "Player 2's bio"
    };

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

    return (
        <div className="gameDivision"
             style={{
                 backgroundImage: `url(${background})`,
                 backgroundSize: 'cover',
                 backgroundPosition: 'center',
                 height: '100%',
                 width: '100%'
             }}>
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
