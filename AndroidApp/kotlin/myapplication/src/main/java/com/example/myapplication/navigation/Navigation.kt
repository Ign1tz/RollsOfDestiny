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
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.viewmodels.Injector
import com.example.myapplication.viewmodels.LoginViewModel

@Composable
fun Navigation() {
    val navController = rememberNavController()

    val loginViewModel: LoginViewModel = viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))
    val homeViewModel: HomeViewModel = viewModel(factory = Injector.provideModelFactory(context = LocalContext.current))


    NavHost(navController = navController,
        startDestination = Screen.GameScreen.route) {
        composable(route = Screen.LoginScreen.route) {
            LoginScreen(navController = navController, loginViewModel = loginViewModel)
        }

        composable(route = Screen.HomeScreen.route) {
            HomeScreen(navController = navController, homeViewModel = homeViewModel)
        }

        composable(route = Screen.GameScreen.route) {
            GameScreen(navController = navController)
        }
    }
}