package com.example.myapplication.navigation

sealed class Screen(val route: String) {
    object LoginScreen: Screen("login")
    object HomeScreen: Screen("home")
    object GameScreen: Screen("game")
    object SettingScreen: Screen("setting")

    object ScoreBoard: Screen("scoreboard")
}