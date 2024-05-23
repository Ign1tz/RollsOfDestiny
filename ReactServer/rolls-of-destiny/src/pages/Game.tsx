import { Button } from "@mui/material";
import Grid from "../components/Grid";
import Dice from "react-dice-roll";
import SimpleBox from "../components/SimpleBox";
import { profile } from "../types/profileTypes"; // Assuming profile type is defined in types.ts
import "../css/BasicCSS.css";

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

    return (
        <div className="gameDivision">
            <div className="header">
                <h1>Welcome to the Game!</h1>
                <Button variant="contained" onClick={() => window.location.href="/"}>
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
                            <Dice onRoll={(value) => console.log(value)} defaultValue={6} size={100}/>
                        </div>
                        <Grid/>
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
                            <Dice onRoll={(value) => console.log(value)} defaultValue={6} size={100}/>
                        </div>
                        <Grid/>
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
