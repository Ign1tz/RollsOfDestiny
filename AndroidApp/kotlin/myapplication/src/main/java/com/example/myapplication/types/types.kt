package com.example.myapplication.types

import kotlinx.serialization.Serializable


@Serializable
data class scoreboardPlayer(
    val username: String,
    val rating: Int,
    val profilePicture: String
)

@Serializable
data class topTenPlayers(
    val topTenPlayers: List<scoreboardPlayer>
)