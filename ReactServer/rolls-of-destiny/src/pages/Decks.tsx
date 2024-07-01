import React, {useState} from "react";
import Card from '@mui/material/Card';
import "../css/Deck.css"
import {Modal, TextField} from "@mui/material";
import Button from "@mui/material/Button";
import TopAppBar from "../bars/TopAppBar";
import {profile} from "../types/profileTypes";
import {login} from "../auth";

export default function Decks() {

    type Deck = {
        name: string,
        numberOfCards: number,
        deckID: string
    }

    type CardType = {
        name: string,
        mana: number,
        image: string
    }

    const initialDeck: Deck = {
        name: "Sample Deck",
        numberOfCards: 50,
        deckID: "1"
    };

    const [createDeckButtonClicked, setCreateDeckButtonClicked] = useState(false)
    const [clickedDeck, setClickedDeck] = useState<Deck>(initialDeck);
    const [openDeckMenu, setOpenDeckMenu] = useState(false)
    const [showCards, setShowCards] = useState(false);
    const [newDeckName, setNewDeckName] = useState("");


    let decks: Deck[] = [
        { name: "Test", numberOfCards: 8, deckID: "1"},
        { name: "gdrgrdg", numberOfCards: 8, deckID: "2"},
        { name: "gdad3w", numberOfCards: 8, deckID: "3"},
        { name: "maurits", numberOfCards: 8, deckID: "4"},
        { name: "heyho", numberOfCards: 8, deckID: "5"},
        { name: "siuuuu", numberOfCards: 8, deckID: "6"},
    ]

    let cards: CardType[] = [
        {name: "Test", mana: 7, image: "Not here yet"},
        {name: "awdawd", mana: 8, image: "Not here yet"},
        {name: "siiuuu", mana: 7, image: "Not here yet"},
        {name: "123", mana: 5, image: "Not here yet"},
        {name: "hiiii", mana: 2, image: "Not here yet"}
    ]

    const clickEvent = (deck: Deck) => {
        setClickedDeck(deck)
        setOpenDeckMenu(true)
    }

    const closeDeckMenu = () => {
        setClickedDeck(initialDeck)
        setOpenDeckMenu(false)
    }

    const closeCreateDeckMenu = () => {
        setCreateDeckButtonClicked(false)
        setNewDeckName("")
    }


    function submitDeckCreation(deck: Deck) {
        console.log("submit new deck clicked")

        fetch("http://localhost:9090/createDeck", {
            method: "POST",
            headers:  {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({username: "TODO USERNAME", deck: deck})
        }).then(r => {
            if (r.status===200) {
                return r.json()
            } else {
                // Error handling
            }
        })
    }

    function deleteDeck(deck: Deck) {
        console.log("delete deck clicked")

        fetch("http://localhost:9090/deleteDeck", {
            method: "POST",
            headers:  {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({username: "TODO USERNAME", deckID: deck.deckID})
        }).then(r => {
            if (r.status===200) {
                return r.json()
            } else {
                // Error handling
            }
        })
    }

    return (
        <>
            <TopAppBar loggedIn={true}/>
            <h2>Your Decks</h2>
            <div className={"differentDecks"}>
                <Modal open={openDeckMenu} onClose={closeDeckMenu}>
                    <div className={"specificDeckMenu"}>
                        <div className={"deckMenuHeader"}>
                            <h3>{clickedDeck.name}</h3>
                            <Button variant={"contained"} color={"error"} onClick={closeDeckMenu}>Close</Button>
                        </div>
                    </div>
                </Modal>
                <Modal open={createDeckButtonClicked} onClose={closeCreateDeckMenu}>
                    <div className={"createDeckMenu"}>
                        <div className={"headerCreateDeckMenu"}>
                            <TextField required id="filled-basic" label="Deck Name" variant="filled"
                                       value={newDeckName}
                                       onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                                           setNewDeckName(event.target.value);
                                       }}/>
                            <Button variant={"contained"} color={"error"} onClick={closeCreateDeckMenu}>Exit</Button>
                        </div>
                        <div className={"chooseCardsMenu"}>
                            {cards.map((card) => (
                                <div className={"specificCardInCreatDeckMenu"}>
                                    <h3>{card.name}</h3>
                                </div>
                            ))}
                        </div>
                        <div className={"confirmButtonCreateDeckMenu"}>
                            <Button variant={"contained"} color={"success"} onClick={() => submitDeckCreation({name: newDeckName, numberOfCards: 1, deckID: "10"})}>Create Deck</Button>
                        </div>
                    </div>
                </Modal>
                {decks.map((deck, index) => (
                        <div className={"deckInstance"}>
                            <Card>
                                <h4>{deck.name}</h4>
                                <h5>Size: {deck.numberOfCards}</h5>
                                <h6>DeckID: {deck.deckID}</h6>
                                <Button variant="contained" color="secondary" onClick={() => clickEvent(deck)}>Edit
                                    Deck</Button>
                                <Button variant="contained" color="error" onClick={() => deleteDeck(deck)}>Delete
                                    Deck</Button>
                            </Card>
                        </div>
                    )
                )}
                <Button variant={"contained"} color={"inherit"} onClick={() => setCreateDeckButtonClicked(true)}>{
                    <h1>+</h1>}</Button>
            </div>
            <h2>Your Cards</h2>
            <div className={"allCards"}>
                {cards.map((card) => (
                    <h3>{card.name}</h3>
                ))}
            </div>
        </>
    )
}