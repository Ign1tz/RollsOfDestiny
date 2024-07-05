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

@Serializable
data class decks(
    val decks: List<deck>
)

@Serializable
data class deck(
    val name: String,
    val deckid: String,
    val active: Boolean,
    val cards: List<deckCard>
)

@Serializable
data class singleDeck(
    val deck: deck
)

@Serializable
data class deckCard(
    val name: String,
    val count: String,
    val image: String
)


@Serializable
data class deckDetailsCards(
    val oldCards: List<String>,
    val newCards: List<String>
)