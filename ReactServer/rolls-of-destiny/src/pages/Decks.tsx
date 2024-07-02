import React, {useEffect, useState} from "react";
import Card from '@mui/material/Card';
import "../css/Deck.css"
import {Modal, TextField} from "@mui/material";
import Button from "@mui/material/Button";
import TopAppBar from "../bars/TopAppBar";
import destroyColumnCard from "../cards/destroy_column.png"
import doubleManaCard from "../cards/double_mana.png"
import rollAgainCard from "../cards/roll_again.png"
import rotateGridCard from "../cards/rotate_grid.png"
import {authFetch} from "../auth";

export default function Decks() {

    type Deck = {
        name: string,
        //numberOfCards: number,
        deckid: string,
        cards: CardType[],
        active: boolean
    };

    type CardType = {
        name: string,
        mana: number,
        image: string,
        count: number
    };

    const initialDeck: Deck = {
        name: "Sample Deck",
        //numberOfCards: 50,
        deckid: "1",
        cards: [],
        active: false
    };

    const [createDeckButtonClicked, setCreateDeckButtonClicked] = useState(false)
    const [clickedDeck, setClickedDeck] = useState<Deck>(initialDeck);
    const [openDeckMenu, setOpenDeckMenu] = useState(false)
    const [showCards, setShowCards] = useState(false);
    const [newDeckName, setNewDeckName] = useState("");
    const [cardsForNewDeck, setCardsForNewDeck] = useState<CardType[]>([])
    const [errorMessage, setErrorMessage] = useState("");
    const [isError, setIsError] = useState(false)

    const [decks, setDecks] = useState<Deck[]>([])

    let cards: CardType[] = [
        {
            name: "Destroy Column",
            mana: 7,
            image: destroyColumnCard,
            count: 0
        },
        {name: "Double Mana", mana: 8, image: doubleManaCard, count: 0},
        {name: "Roll again", mana: 7, image: rollAgainCard, count: 0},
        {name: "Rotate Grid", mana: 5, image: rotateGridCard, count: 0}
    ];

    /*let decks: Deck[] = [
        {
            name: "Test", /!*numberOfCards: 8,*!/ deckID: "1", cards: cards, active: true
        },
        {
            name: "gdrgrdg", /!*numberOfCards: 8,*!/ deckID: "2", cards: [], active: false
        },
        {
            name: "gdad3w", /!*numberOfCards: 8,*!/ deckID: "3", cards: [], active: false
        },
        {
            name: "maurits", /!*numberOfCards: 8,*!/ deckID: "4", cards: [], active: false
        },
        {
            name: "heyho", /!*numberOfCards: 8,*!/ deckID: "5", cards: [], active: false
        },
        {
            name: "siuuuu", /!*numberOfCards: 8,*!/ deckID: "6", cards: [], active: false
        }
    ];*/

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

    const addCardToDeck = (card: CardType) => {
        cardsForNewDeck.push(card)
        console.log(cardsForNewDeck)
        setCardsForNewDeck(cardsForNewDeck)
    }

    const handleError = () => {
        setIsError(false)
        setErrorMessage("")
    }

    useEffect(() => {
        console.log(cards[0].image)
        console.log(cards[1].image)
        console.log(cards[2].image)
        console.log(cards[3].image)
        authFetch("http://localhost:9090/getDecks").then(response => {

            return response.json()
        }).then(r => {
                console.log(r.decks[0].deckid)
                setDecks(r.decks.reverse())
            }
        )
    }, []);


    function setActiveDeck(deck: Deck) {
        console.log(deck.deckid)
        if (!deck.active) {
            authFetch("http://localhost:9090/setActiveDeck", {
                method: "POST",
                headers: {
                    'Accept': 'application/json, text/plain',
                    'Content-Type': 'application/json;charset=UTF-8'
                },
                body: JSON.stringify({name: deck.name, deckid: deck.deckid})
            }).then( () => {
                window.location.reload()
            })
        } else {
            setErrorMessage("This deck is already activated.")
            setIsError(true)
        }
    }


    function submitDeckCreation(deck: Deck) {
        console.log("submit new deck clicked")
        if (newDeckName != "") {
            decks.push(deck)
        }
        console.log(decks)
        setCardsForNewDeck([])
        setNewDeckName("")
        setCreateDeckButtonClicked(false)

        authFetch("http://localhost:9090/createDeck", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({name: deck.name})
        }).then( () => {
            window.location.reload()
        })
    }

    function deleteDeck(deck: Deck) {
        console.log("delete deck clicked")

        authFetch("http://localhost:9090/removeDeck", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({name: deck.name, deckid: deck.deckid})
        }).then( () => {
            window.location.reload()
        })
    }

    return (
        <div className={"deckPage"}>
            <TopAppBar loggedIn={true}/>
            <Modal open={isError} onClose={handleError}>
                <div className="errorMenu">
                    <h2>Oops...something went wrong.</h2>
                    <div className="errorText">
                        {errorMessage}
                    </div>
                    <Button variant="contained" color={"secondary"} onClick={handleError}>
                        I understand.
                    </Button>
                </div>
            </Modal>
            <Modal open={openDeckMenu} onClose={closeDeckMenu}>
                <div className={"specificDeckMenu"}>
                    <div className={"deckMenuHeader"}>
                        <h3>{clickedDeck.name}</h3>
                        <Button variant={"contained"} color={"error"} onClick={closeDeckMenu}>Close</Button>
                    </div>
                    <div className={"specificDeckCards"}>
                        {clickedDeck.cards.map((card) => (
                            <div className={"specificCardInCreatDeckMenu"}>
                                <h3>{card.name}</h3>
                                <img id="cardImages" src={card.image} alt={"card image"}/>
                                <h5>In Deck: {cardsForNewDeck.filter(c => c.name === card.name).length + 1}</h5>
                            </div>
                        ))}
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
                                <img id="cardImages" src={card.image} alt={"card image"}/>
                                <Button onClick={() => addCardToDeck(card)} variant={"contained"}
                                        color={"secondary"} style={{marginTop: "20px"}}>Add to Deck</Button>
                            </div>
                        ))}
                    </div>
                    <div className={"confirmButtonCreateDeckMenu"}>
                        <Button variant={"contained"} color={"success"} onClick={() => submitDeckCreation({
                            name: newDeckName,
                            deckid: "10",
                            cards: cardsForNewDeck,
                            active: false
                        })}>Create Deck</Button>
                    </div>
                </div>
            </Modal>
            <div className={"titleWithDecksAndCards"}>
                <h2 style={{textAlign: "center"}}>Your Decks</h2>
                <div className={"differentDecks"}>
                    {decks.map((deck, index) => (
                            <div className={"deckInstance"}>
                                <Card style={{backgroundColor: "lightgray"}}>
                                    <h4>{deck.name}</h4>
                                    <div className={"infosForDeck"}>
                                        <h5>Currently no Infos</h5>
                                    </div>
                                    <div className={"buttonForDeck"}>
                                        <Button variant="contained" color="secondary" onClick={() => clickEvent(deck)}>Edit
                                            Deck</Button>
                                        <Button variant="contained" color="error" onClick={() => deleteDeck(deck)}>Delete
                                            Deck</Button>
                                        <Button variant={"contained"} color={deck.active ? "info" : "success"}
                                                onClick={() => setActiveDeck(deck)}> {deck.active ? "Currently Using" : "Activate"}</Button>
                                    </div>
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
                        <div className={"individualCardOwned"}>
                            <h3>{card.name}</h3>
                            <img src={card.image} alt={"card image"}/>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}