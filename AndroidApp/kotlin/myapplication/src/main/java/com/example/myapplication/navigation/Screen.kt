package com.example.myapplication.navigation

sealed class Screen(val route: String) {
    object LoginScreen: Screen("login")
    object HomeScreen: Screen("home")
    object GameScreen: Screen("game")
    object DeckScreen: Screen("decks")
    object CardScreen: Screen("cards")
    object SettingScreen: Screen("setting")
    object ScoreBoard: Screen("scoreboard")

    object RulesScreen: Screen("rules")
}