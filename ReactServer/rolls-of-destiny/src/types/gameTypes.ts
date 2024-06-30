type GameInfo = {
    messageType: string,
    messageBody: messageBody,
    gameId: string
}

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
    username: string,
    websocketId: string,
    userId: string,
    left: column,
    middle: column,
    right: column
}

type column = {
    first: string,
    second: string,
    third: string,
    isFull: boolean
}

type enemyInfo = {
    username: string,
    websocketId: string,
    left: column,
    middle: column,
    right: column
}

export type { GameInfo , messageBody }