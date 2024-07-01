package com.example.myapplication.screens

import android.util.Log
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.NavController
import com.example.myapplication.R
import com.example.myapplication.connection.websocket.WebSocketClient
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.widgets.ProfileRow
import com.example.myapplication.widgets.PlayField
import java.net.URI


@Composable
fun GameScreen (navController: NavController, gameViewModel: GameViewModel) {

    val webSocketClient = gameViewModel.websocket()

    Scaffold (

        topBar = { ProfileRow(profileImage = R.drawable.caught, username = "Enemy", score = "50")},

        bottomBar = { ProfileRow(profileImage = R.drawable.xdd, username = "eziekel", score = "120") }



    ){innerPadding ->
        Column (modifier = Modifier.padding(innerPadding)) {
            PlayField(6,gameViewModel)

        }
    }

}