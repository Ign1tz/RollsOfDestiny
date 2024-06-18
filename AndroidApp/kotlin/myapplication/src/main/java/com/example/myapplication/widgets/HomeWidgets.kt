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

@Composable
fun HomeScreenButtons (navController: NavController) {
    val padding = 10.dp
    Column (
        modifier = Modifier.fillMaxWidth(),
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Spacer(Modifier.padding(20.dp))
        QuickPlayButton(navController)
        Spacer(Modifier.padding(padding))
        FriendPlayButton()
        Spacer(Modifier.padding(padding))
        RankedPlayButton()
        Spacer(Modifier.padding(padding))
        ScoreboardButton()
        Spacer(Modifier.padding(padding))
        SettingsButton()
    }
}