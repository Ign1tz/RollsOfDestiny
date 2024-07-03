
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
    Score: number
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
    Score: number
}

type endResults = {
    yourScore: number,
    enemyScore: number,
    youWon: string
}

export type {messageBody, column , yourInfo, enemyInfo, endResults }