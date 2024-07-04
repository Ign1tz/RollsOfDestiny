package com.example.myapplication.navigation

import androidx.compose.runtime.Composable
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.myapplication.screens.GameScreen
import com.example.myapplication.screens.HomeScreen
import com.example.myapplication.screens.LoginScreen
import com.example.myapplication.screens.ScoreBoardScreen
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.screens.SettingScreen
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.viewmodels.Injector
import com.example.myapplication.viewmodels.LoginViewModel
import com.example.myapplication.viewmodels.ScoreboardViewModel
import com.example.myapplication.viewmodels.SettingViewModel

@Composable
fun Navigation() {
    val navController = rememberNavController()

    val loginViewModel: LoginViewModel = viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val homeViewModel: HomeViewModel = viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val gameViewModel: GameViewModel = viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val settingViewModel: SettingViewModel = viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))

    val scoreboardViewModel: ScoreboardViewModel = viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))


    NavHost(navController = navController,
        startDestination = if (loginViewModel.checkAlreadyLoggedIn()) {Screen.HomeScreen.route} else {Screen.LoginScreen.route}) {
        composable(route = Screen.LoginScreen.route) {
            LoginScreen(navController = navController, loginViewModel = loginViewModel)
        }

        composable(route = Screen.HomeScreen.route) {
            HomeScreen(navController = navController, homeViewModel = homeViewModel, loginViewModel, gameViewModel, scoreboardViewModel)
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
            gameViewModel.resetAllValues()
            GameScreen(navController = navController, gameViewModel = gameViewModel)
        }
        composable(route = "game/bot") {
            gameViewModel.GameType.value = "bot"
            gameViewModel.resetAllValues()
            GameScreen(navController = navController, gameViewModel = gameViewModel)
        }

        composable(route = Screen.ScoreBoard.route) {
            ScoreBoardScreen(navController = navController, scoreboardViewModel = scoreboardViewModel)
        }
    }
}