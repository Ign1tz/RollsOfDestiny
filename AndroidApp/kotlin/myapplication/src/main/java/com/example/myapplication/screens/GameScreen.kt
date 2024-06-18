package com.example.myapplication.screens

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import com.example.myapplication.widgets.Dice
import com.example.myapplication.widgets.PlayField

@Composable
fun GameScreen (navController: NavController) {
    Column {
        PlayField(3)

    }
}