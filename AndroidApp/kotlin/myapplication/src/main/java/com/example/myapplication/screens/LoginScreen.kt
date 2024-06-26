package com.example.myapplication.screens

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.NavController
import com.example.myapplication.viewmodels.LoginViewModel
import com.example.myapplication.widgets.BottomBar
import com.example.myapplication.widgets.HeaderTopBar
import com.example.myapplication.widgets.LoginBox


@Composable
fun LoginScreen (navController: NavController, loginViewModel: LoginViewModel) {
    Scaffold (
        topBar = {
            HeaderTopBar(navController, "X")
        },
        bottomBar = {
            BottomBar()
        }
    ) { innerPadding ->
        Column (modifier = Modifier.padding(innerPadding)){
            LoginBox(navController = navController)
        }

    }
}