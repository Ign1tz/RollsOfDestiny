package com.example.myapplication.types

import kotlinx.serialization.Serializable

@Serializable
data class message(
    val info: String,
    val message: String,
)

@Serializable
data class idMessageBody(
    val id: String
)

@Serializable
data class gameMessageBody(
    val ActivePlayer: ActivePlayer,
    val YourInfo: yourInfo,
    val EnemyInfo: enemyInfo,
    val gameid: String
)

@Serializable
data class gameInfo(
    val gameInfo: String
)

@Serializable
data class ActivePlayer(
    val active: Boolean,
    val roll: String
)

@Serializable
data class yourInfo(
    val WebsocketId: String,
    val Username: String,
    val userId: String,
    val LeftColumn: Column,
    val MiddleColumn: Column,
    val RightColumn: Column,
    val Score: Int
)

@Serializable
data class Column(
    val First: String,
    val Second: String,
    val Third: String,
    val IsFull: Boolean
)

@Serializable
data class enemyInfo(
    val Username: String,
    val WebsocketId: String,
    val LeftColumn: Column,
    val MiddleColumn: Column,
    val RightColumn: Column,
    val Score: Int
)

@Serializable
data class EndResults(
    val yourScore: Int,
    val enemyScore: Int,
    val youWon: String
)
@Serializable
data class EndResultsBody(
    val gameInfo: String,
    val endResults: String,
)

@Serializable
data class userInfoMessage(
    val username: String,
    val email: String,
    val profilePicture: String,
    val rating: String,
    val userid: String
)

