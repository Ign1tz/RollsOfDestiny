package com.example.myapplication.screens

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.NavController
import com.example.myapplication.viewmodels.SettingViewModel
import com.example.myapplication.widgets.BottomBar
import com.example.myapplication.widgets.HeaderTopBar
import com.example.myapplication.widgets.CenterSettings

@Composable
fun SettingScreen (navController: NavController, settingViewModel: SettingViewModel) {
    Scaffold (
        topBar = { HeaderTopBar(navController,"<") },

    ){
        innerPadding ->
        Column (modifier = Modifier.padding(innerPadding)){
            CenterSettings(settingViewModel, navController)
        }
    }
}