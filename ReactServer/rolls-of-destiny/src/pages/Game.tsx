import {useState} from "react";
import {Button} from "@mui/material";
import Grid from "../components/Grid";
import OpponentGrid from "../components/OpponentGrid"
import Dice from "react-dice-roll";
import SimpleBox from "../components/SimpleBox";
import {profile} from "../types/profileTypes";
import "../css/Game.css";
import background from "../images/game.jpg"


export default function Game() {
    const player1: profile = {
        username: "Lukas",
        rating: 3450913,
        picture: "/path/to/player1.jpg",
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
    const [disableRoll, setDisableRoll] = useState(false) // TODO: change to true
    const [canPlace, setCanPlace] = useState(true) // TODO: change to false

    const handleRoll = (player: 'player1' | 'player2', value: number) => {
        setDiceRoll(value);
        setDisableRoll(true);
        if (player === 'player1') {
            setPlayer1Score(player1Score + value);
        } else {
            setPlayer2Score(player2Score + value);
        }
    };

    return (
        <div className="gameDivision"
             style={{
                 backgroundImage: `url(${background})`,
                 backgroundSize: 'cover',
                 backgroundPosition: 'center',
                 height: '100vh',
                 width: '100%'
             }}>
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
                            <p>Score: <span id="player1Score">{player1Score}</span></p>
                        </div>
                    </div>
                    <div className="playerActions">
                        <div className="diceWrapper">
                            <Dice defaultValue={6} size={100} cheatValue={undefined} disabled={true}/>
                        </div>
                        <OpponentGrid diceRoll={diceRoll}/>
                        <div className="playerCards">
                            <SimpleBox diceValue={null}/>
                            <SimpleBox diceValue={null}/>
                            <SimpleBox diceValue={null}/>
                        </div>
                    </div>
                </div>
                <div className="playerSection">
                    <div className="playerActions">
                        <div className="diceWrapper">
                            <Dice onRoll={(value) => handleRoll('player2', value)} defaultValue={6} size={100}
                                  cheatValue={undefined} disabled={disableRoll}/>
                        </div>
                        <Grid canPlace={canPlace} setCanPlace={setCanPlace} diceRoll={diceRoll}/>
                        <div className="playerCards">
                            <SimpleBox diceValue={null}/>
                            <SimpleBox diceValue={null}/>
                            <SimpleBox diceValue={null}/>
                        </div>
                    </div>
                    <div className="playerInfo">
                        <img src={player2.picture} alt={player2.username}/>
                        <div>
                            <h2>{player2.username}</h2>
                            <p>Rating: {player2.rating}</p>
                            <p>Score: <span id="player2Score">{player2Score}</span></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
