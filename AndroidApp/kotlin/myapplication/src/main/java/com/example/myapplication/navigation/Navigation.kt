package com.example.myapplication.navigation

import android.util.Log
import androidx.compose.runtime.Composable
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavType
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import androidx.navigation.navArgument
import com.example.myapplication.screens.CardScreen
import com.example.myapplication.screens.DeckScreen
import com.example.myapplication.screens.GameScreen
import com.example.myapplication.screens.HomeScreen
import com.example.myapplication.screens.LoginScreen
import com.example.myapplication.screens.RulesScreen
import com.example.myapplication.viewmodels.CardViewModel
import com.example.myapplication.viewmodels.DeckViewModel
import com.example.myapplication.screens.ScoreBoardScreen
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.screens.SettingScreen
import com.example.myapplication.types.AudioPlayer
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.viewmodels.Injector
import com.example.myapplication.viewmodels.LoginViewModel
import com.example.myapplication.viewmodels.ScoreboardViewModel
import com.example.myapplication.viewmodels.SettingViewModel
import kotlinx.coroutines.delay
import kotlinx.coroutines.runBlocking

@Composable
fun Navigation() {
    val navController = rememberNavController()

    val loginViewModel: LoginViewModel =
        viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val homeViewModel: HomeViewModel =
        viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val gameViewModel: GameViewModel =
        viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val deckViewModel: DeckViewModel =
        viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val cardViewModel: CardViewModel =
        viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val settingViewModel: SettingViewModel =
        viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val scoreboardViewModel: ScoreboardViewModel =
        viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))



    NavHost(
        navController = navController,
        startDestination = if (loginViewModel.checkAlreadyLoggedIn()) {
            Screen.HomeScreen.route
        } else {
            Screen.LoginScreen.route
        }
    ) {
        composable(route = Screen.LoginScreen.route) {
            LoginScreen(navController = navController, loginViewModel = loginViewModel)
        }

        composable(route = Screen.HomeScreen.route) {
            HomeScreen(
                navController = navController,
                homeViewModel = homeViewModel,
                loginViewModel,
                gameViewModel,
                scoreboardViewModel
            )
        }

        composable(route = Screen.SettingScreen.route) {

            settingViewModel.newPassword.value = ""
            settingViewModel.oldPassword.value = ""
            settingViewModel.confirmNewPassword.value = ""
            settingViewModel.username.value = ""
            SettingScreen(navController = navController, settingViewModel = settingViewModel)
        }
        composable(route = Screen.GameScreen.route) {
            gameViewModel.GameType.value = ""
            gameViewModel.endResults = null
            gameViewModel.gameInfo = null
            gameViewModel.connected.value = false
            GameScreen(navController = navController, gameViewModel = gameViewModel)
        }

        composable(route = "game/bot") {

            gameViewModel.GameType.value = "bot"
            gameViewModel.endResults = null
            gameViewModel.gameInfo = null
            gameViewModel.connected.value = false
            GameScreen(navController = navController, gameViewModel = gameViewModel)
        }

        composable(route = Screen.RulesScreen.route) {
            RulesScreen(navController = navController)
        }

        composable(route = "game/friend/{friendId}",
            arguments = listOf(navArgument(name = "friendId") { type = NavType.StringType })
        )
        { backStackEntry ->
            gameViewModel.GameType.value = "Friend"
            gameViewModel.gameInfo = null
            gameViewModel.endResults = null
            gameViewModel.connected.value = false
            gameViewModel.FriendId.value = backStackEntry.arguments?.getString("friendId") ?: ""
            GameScreen(navController = navController, gameViewModel = gameViewModel)
        }

        composable(route = Screen.ScoreBoard.route) {
            ScoreBoardScreen(
                navController = navController,
                scoreboardViewModel = scoreboardViewModel
            )
        }

        composable(route = "decks") {
            DeckScreen(navController = navController, deckViewModel = deckViewModel)
        }
        composable(route = "deckDetails/{deckid}",
            arguments = listOf(navArgument(name = "deckid") { type = NavType.StringType })
        )
        { backStackEntry ->

            cardViewModel.deckid = backStackEntry.arguments?.getString("deckid") ?: ""
            CardScreen(navController = navController, cardViewModel = cardViewModel)
        }
    }
}