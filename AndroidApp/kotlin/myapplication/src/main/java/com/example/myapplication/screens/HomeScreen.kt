package com.example.myapplication.screens

import android.graphics.drawable.Icon
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import com.example.myapplication.widgets.BottomBar
import com.example.myapplication.widgets.HeaderTopBar
import com.example.myapplication.widgets.HomeScreenButtons
import androidx.compose.ui.Modifier
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.widgets.FriendlistDrawer
import com.example.myapplication.widgets.UpButton


@Composable
fun HomeScreen (navController: NavController, homeViewModel: HomeViewModel) {
    Scaffold (
        topBar = {
            FriendlistDrawer(navController = navController)
        },
        bottomBar = {
            BottomBar()
        }
    ){ innerPadding ->
        Column (modifier = Modifier.padding(innerPadding)) {
            HomeScreenButtons(navController = navController)

        }
    }
}