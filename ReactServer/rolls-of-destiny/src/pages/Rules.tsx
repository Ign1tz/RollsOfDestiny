import React from 'react';
import TopAppBar from "../bars/TopAppBar";
import "../css/Rules.css"

export default function Rules({ loggedIn }: { loggedIn: boolean }) {
    return (
        <>
            <TopAppBar loggedIn={loggedIn} />
            <div className="rules-container">
                <div className="rules-header">
                    <h1>Rules</h1>
                </div>
                <div className="rules-content">
                    <p style={{textAlign:"center"}}>Roll your own Destiny.</p>
                    <ul>
                        <li>General Rules</li>
                        <ul>
                            <li>Play with each other and have fun.</li>
                            <li>Do your best.</li>
                            <li>No harassing and racism.</li>
                            <li>You play for yourself, do not cheat.</li>
                        </ul>
                        <li>Game Rules</li>
                        <ul>
                            <li>You roll a die.</li>
                            <li>Place die in a column of your grid.</li>
                            <li>Opponent does the same.</li>
                            <li>Points are caculated by values of your columns in your grid.</li>
                            <li>If you have the same value 2 or even 3 times in your column, it is multiplied with each
                                other.
                            </li>
                            <li>You can optionally play cards which give you certain effects.</li>
                            <li>If you win, you gain rating points. If you lose, you lose rating points.</li>
                            <li>You can play against a computer, against your friend or against a random enemy based on
                                queueing.
                            </li>
                            <li>The game is over when one grid is filled up completely.</li>
                            <li>Do your best and roll your destiny!</li>
                        </ul>


                    </ul>
                </div>
            </div>
        </>
    );
}
