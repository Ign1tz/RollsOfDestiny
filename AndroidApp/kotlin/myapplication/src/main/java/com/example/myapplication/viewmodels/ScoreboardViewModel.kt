package com.example.myapplication.viewmodels

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import com.example.myapplication.R
import com.example.myapplication.localdb.Repository
import com.example.myapplication.types.scoreboardPlayer



class ScoreboardViewModel(val repository: Repository) : ViewModel(), BasicViewModel {

    val players = mutableStateOf(listOf<scoreboardPlayer>())

    fun getPlayers() {
        val createdPlayers = listOf(
            scoreboardPlayer(username = "Hubert", rating = 1000, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Alice", rating = 1200, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Bob", rating = 900, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Carol", rating = 1100, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Dave", rating = 950, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Eve", rating = 1300, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Frank", rating = 800, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Grace", rating = 1150, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Heidi", rating = 1050, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Ivan", rating = 1250, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Judy", rating = 1100, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Mallory", rating = 900, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Niaj", rating = 850, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Olivia", rating = 1200, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Peggy", rating = 950, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Trent", rating = 1000, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Victor", rating = 1400, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Wendy", rating = 900, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Xander", rating = 1050, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Yvonne", rating = 1150, profilePicture = "placeholder"),
            scoreboardPlayer(username = "Zara", rating = 1100, profilePicture = "placeholder")
        )

        players.value = createdPlayers
    }

}