
type messageBody = {
    ActivePlayer: activePlayer,
    YourInfo: yourInfo,
    EnemyInfo: enemyInfo
}

type activePlayer = {
    active: boolean,
    roll: string
}

type yourInfo = {
    Username: string,
    WebsocketId: string,
    userId: string,
    LeftColumn: column,
    MiddleColumn: column,
    RightColumn: column,
    Score: number,
    deck: deck
}

type column = {
    First: string,
    Second: string,
    Third: string,
    IsFull: boolean
}

type enemyInfo = {
    Username: string,
    websocketId: string,
    LeftColumn: column,
    MiddleColumn: column,
    RightColumn: column,
    Score: number,
    deck: enemyDeck
}

type endResults = {
    yourScore: number,
    enemyScore: number,
    youWon: string
}

type card = {
    name: string,
    cost: number,
    picture: string,
    effect: string
}

type deck = {
    cardsLeft: number,
    inHand: card[]
}

type enemyDeck = {
    cardsLeft: number,
    inHand: number
}

export type {messageBody, column , yourInfo, enemyInfo, endResults }