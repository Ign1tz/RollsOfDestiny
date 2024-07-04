package com.example.myapplication.screens

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import com.example.myapplication.widgets.BottomBar
import com.example.myapplication.widgets.HomeTopBar
import com.example.myapplication.widgets.HomeScreenButtons
import androidx.compose.ui.Modifier
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.viewmodels.LoginViewModel


@Composable
fun HomeScreen (navController: NavController, homeViewModel: HomeViewModel, loginViewModel: LoginViewModel, gameViewModel: GameViewModel) {
    Scaffold (
        topBar = {
            HomeTopBar(navController,loginViewModel)
        },
        bottomBar = {
            BottomBar()
        }
    ){ innerPadding ->
        Column (modifier = Modifier.padding(innerPadding)) {
            HomeScreenButtons(navController = navController, loginViewModel = loginViewModel, gameViewModel = gameViewModel )

        }
    }
}