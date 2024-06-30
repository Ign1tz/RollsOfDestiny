import {useState} from "react";

export default function Decks() {

    const [createDeck, setCreateDeck] = useState(false)
    const [clickedDeck, setClickedDeck] = useState(false)

    type deck = {
        name: string,
        numberOfCards: number,
        deckID: number
    }

    let decks: deck[] = [
        { name: "Test", numberOfCards: 8, deckID: 1},
        { name: "gdrgrdg", numberOfCards: 8, deckID: 2},
        { name: "gdad3w", numberOfCards: 8, deckID: 3},
        { name: "maurits", numberOfCards: 8, deckID: 4},
        { name: "heyho", numberOfCards: 8, deckID: 5},
        { name: "siuuuu", numberOfCards: 8, deckID: 6},
    ]




    return (
        <></>
    )
}