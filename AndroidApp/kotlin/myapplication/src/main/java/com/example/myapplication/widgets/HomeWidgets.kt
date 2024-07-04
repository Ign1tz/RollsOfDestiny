package com.example.myapplication.widgets

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.viewmodels.LoginViewModel

@Composable
fun HomeScreenButtons (navController: NavController, loginViewModel: LoginViewModel, gameViewModel: GameViewModel) {
    val padding = 10.dp
    Column (
        modifier = Modifier.fillMaxWidth(),
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Spacer(Modifier.padding(20.dp))
        QuickPlayButton(navController, gameViewModel)
        Spacer(Modifier.padding(padding))
        FriendPlayButton()
        Spacer(Modifier.padding(padding))
        RankedPlayButton(navController, gameViewModel)
        Spacer(Modifier.padding(padding))
        ScoreboardButton()
        Spacer(Modifier.padding(padding))
        SettingsButton()
        Spacer(Modifier.padding(padding))
        DecksButton(navController = navController)
    }
}