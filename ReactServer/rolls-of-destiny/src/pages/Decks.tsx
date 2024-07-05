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
import flipClockwiseCard from "../cards/rotate_grid.png"
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

    const [hasMessage, setHasMessage] = useState(false)
    const [showNewCard, setShowNewCard] = useState(false)

    const [decks, setDecks] = useState<Deck[]>([])
    const [oldCards, setOldCards] = useState<string[]>([])
    const [newCards, setNewCards] = useState<string[]>([])

    let cards: CardType[] = [
        {
            name: "Destroy Column",
            mana: 7,
            image: destroyColumnCard,
            count: 0
        },
        {name: "Double Mana", mana: 8, image: doubleManaCard, count: 0},
        {name: "Roll Again", mana: 7, image: rollAgainCard, count: 0},
        {name: "Flip Clockwise", mana: 5, image: rotateGridCard, count: 0}
    ];

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

    const openNewCardModal = () => {
        setShowNewCard(true)
        // fetch logic for getting new card
    }

    const closeNewCardModal = () => {
        authFetch("http://localhost:9090/aknowledgeNewCard", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({name: newCards[0]})
        }).then(() =>
            window.location.reload()
        )
        setShowNewCard(false)
    }


    const addCardToDeck = (name: string) => {
        setCardsForNewDeck(cardsForNewDeck)

        authFetch("http://localhost:9090/addCardToDeck", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({name: name, deckId: clickedDeck.deckid})
        }).then(() =>
            window.location.reload()
        )


    }

    const removeCardFromDeck = (cardname: string) => {
        setCardsForNewDeck(cardsForNewDeck)

        authFetch("http://localhost:9090/removeCardFromDeck", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({name: cardname, deckId: clickedDeck.deckid})
        }).then(() =>
            window.location.reload())
    }

    const handleError = () => {
        setIsError(false)
        setErrorMessage("")
    }

    useEffect(() => {
        authFetch("http://localhost:9090/getDecks").then(response => {

            return response.json()
        }).then(r => {
                setDecks(r.decks.reverse())
            }
        )

        authFetch("http://localhost:9090/getNewCards").then(response => {
            return response.json()
        }).then(r => {
            setOldCards(r.oldCards)
            setHasMessage(r.newCards.length > 0)
            setNewCards(r.newCards)
            console.log(r)
        })
    }, []);


    function setActiveDeck(deck: Deck) {
        console.log(deck.deckid)
        authFetch("http://localhost:9090/setActiveDeck", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({name: deck.name, deckid: deck.deckid})
        }).then(() => {
            window.location.reload()
        })

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
        }).then(() => {
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
        }).then(() => {
            window.location.reload()
        })
    }

    function getCardPicture(card: string) {
        switch (card) {
            case "Roll Again":
                return rollAgainCard
            case "Double Mana":
                return doubleManaCard
            case "Destroy Column":
                return destroyColumnCard
            case "Flip Clockwise":
                return flipClockwiseCard
        }
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
                    <div className={"chooseCardsMenu"}>
                        {oldCards.map((cardName) => (
                            <div className={"specificCardInCreatDeckMenu"}>
                                <h3>{cardName}</h3>
                                <img id="cardImages" src={getCardPicture(cardName)} alt={"card image"}/>
                                <Button onClick={() => {
                                    !clickedDeck.cards.find(e => e.name === cardName) ? addCardToDeck(cardName) : removeCardFromDeck(cardName)
                                }} variant={"contained"}
                                        color={"secondary"}
                                        style={{marginTop: "20px"}}>{clickedDeck.cards.find(e => e.name === cardName) ? "Delete" : "Add to Deck"}</Button>
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
                        <Button variant={"contained"} color={"success"} onClick={() => submitDeckCreation({
                            name: newDeckName,
                            deckid: "10",
                            cards: cardsForNewDeck,
                            active: false
                        })}>Create Deck</Button>
                    </div>
                </div>
            </Modal>

            <Modal open={showNewCard} onClose={closeNewCardModal}>
                <div className="errorMenu">
                    <h2>You collected a new card! </h2>
                    <div className="errorText">
                        <div className={"individualCardOwned"}>
                            <h3>{newCards[0]}</h3>
                            <img src={getCardPicture(newCards[0])} alt={"card image"}/>
                        </div>
                        <Button variant={"contained"} color={"success"} onClick={closeNewCardModal}>WOW THIS IS THE BEST
                            THING EVER</Button>
                    </div>
                </div>
            </Modal>
            <div className={"titleWithDecksAndCards"}>
                <div className={"deckScreenHeader"} style={hasMessage ? {width: "1200px"} : {width: "135px"}}>
                    <h2>Your Decks</h2>
                    {hasMessage && (
                        <Button className="blinking-button" variant={"contained"} color={"warning"}
                                onClick={openNewCardModal}>New
                            Message</Button>
                    )}

                </div>

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
                                        <Button variant={"contained"} color={!deck.active ? "info" : "success"}
                                                onClick={() => setActiveDeck(deck)}> {!deck.active ? "Activate" : "Active"}</Button>
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
                    {oldCards.map((cardname) => (
                        <div className={"individualCardOwned"}>
                            <h3>{cardname}</h3>
                            <img src={getCardPicture(cardname)} alt={"card image"}/>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}