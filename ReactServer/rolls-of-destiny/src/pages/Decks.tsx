import {useState} from "react";
import Card from '@mui/material/Card';
import "../css/Deck.css"
import {Modal} from "@mui/material";
import Button from "@mui/material/Button";
import TopAppBar from "../bars/TopAppBar";

export default function Decks() {

    type Deck = {
        name: string,
        numberOfCards: number,
        deckID: number
    }

    type Card = {
        name: string,
        mana: number,
        image: string
    }

    const initialDeck: Deck = {
        name: "Sample Deck",
        numberOfCards: 50,
        deckID: 1
    };

    const [createDeck, setCreateDeck] = useState(false)
    const [clickedDeck, setClickedDeck] = useState<Deck>(initialDeck);
    const [openDeckMenu, setOpenDeckMenu] = useState(false)
    const [showCards, setShowCards] = useState(false);


    let decks: Deck[] = [
        { name: "Test", numberOfCards: 8, deckID: 1},
        { name: "gdrgrdg", numberOfCards: 8, deckID: 2},
        { name: "gdad3w", numberOfCards: 8, deckID: 3},
        { name: "maurits", numberOfCards: 8, deckID: 4},
        { name: "heyho", numberOfCards: 8, deckID: 5},
        { name: "siuuuu", numberOfCards: 8, deckID: 6},
    ]

    let cards: Card[] = [
        {name: "Test", mana: 7, image: "Not here yet"}
    ]

    const clickEvent = (deck: Deck) => {
        setClickedDeck(deck)
        setOpenDeckMenu(true)
    }

    const closeDeckMenu = () => {
        setClickedDeck(initialDeck)
        setOpenDeckMenu(false)
    }

    function deleteDeck(deck: Deck) {
        console.log("delete deck clicked")
        // TODO: delete logic
    }

    return (
        <>
            <TopAppBar loggedIn={true}/>
            <div className={"differentDecks"}>
                <Modal open={openDeckMenu} onClose={closeDeckMenu}>
                    <div className={"specificDeckMenu"}>
                        <div className={"deckMenuHeader"}>
                            <h3>{clickedDeck.name}</h3>
                            <Button variant={"contained"} color={"error"} onClick={closeDeckMenu}>Close</Button>
                        </div>
                    </div>
                </Modal>
                {decks.map((deck, index) => (
                    <div className={"deckInstance"}>
                        <Card>
                            <h4>{deck.name}</h4>
                            <h5>Size: {deck.numberOfCards}</h5>
                            <h6>DeckID: {deck.deckID}</h6>
                            <Button variant="contained" color="secondary" onClick={() => clickEvent(deck)}>Edit Deck</Button>
                            <Button variant="contained" color="error" onClick={() => deleteDeck(deck)}>Delete Deck</Button>
                        </Card>
                    </div>
                    )
                )}
                <Button variant={"contained"} color={"inherit"}>{<h1>+</h1>}</Button>
            </div>
        </>
    )
}