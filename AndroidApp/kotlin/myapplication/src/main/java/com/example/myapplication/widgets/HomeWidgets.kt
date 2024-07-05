package com.example.myapplication.widgets

import android.util.Log
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import com.example.myapplication.screens.RulesButton
import com.example.myapplication.screens.RulesScreen
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.viewmodels.LoginViewModel
import com.example.myapplication.viewmodels.ScoreboardViewModel

@Composable
fun HomeScreenButtons(
    navController: NavController,
    loginViewModel: LoginViewModel,
    gameViewModel: GameViewModel,
    scoreboardViewModel: ScoreboardViewModel,
    homeViewModel: HomeViewModel
) {
    val padding = 10.dp

    Column(
        modifier = Modifier.fillMaxWidth().verticalScroll(rememberScrollState()),
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Spacer(Modifier.padding(20.dp))
        QuickPlayButton(navController, gameViewModel)
        Spacer(Modifier.padding(padding))
        FriendPlayButton(homeViewModel, navController)
        Spacer(Modifier.padding(padding))
        RankedPlayButton(navController, gameViewModel)
        Spacer(Modifier.padding(padding))
        ScoreboardButton(navController, scoreboardViewModel = scoreboardViewModel)
        Spacer(Modifier.padding(padding))
        SettingsButton(navController)
        Spacer(Modifier.padding(padding))
        DecksButton(navController = navController)
        Spacer(Modifier.padding(padding))
        RulesButton(navController)
    }
}