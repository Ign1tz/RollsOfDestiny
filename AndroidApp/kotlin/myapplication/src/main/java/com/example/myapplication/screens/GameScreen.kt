package com.example.myapplication.screens

import android.util.Log
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import com.example.myapplication.R
import com.example.myapplication.connection.websocket.WebSocketClient
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.widgets.ProfileRow
import com.example.myapplication.widgets.PlayField
import java.net.URI


@Composable
fun GameScreen (navController: NavController, gameViewModel: GameViewModel) {
    gameViewModel.websocket()
    if (gameViewModel.endResults != null){
        Scaffold (

        ){innerPadding ->
            Column (modifier = Modifier
                .padding(innerPadding)
                .background(Color.White)) {
                Box(){
                    Text(
                        text = "Game finished",
                        fontSize = 10.sp,
                        color = Color.Black,
                        fontFamily = FontFamily.Serif
                    )
                }
                Box(){
                    Text(
                        text = gameViewModel.endResults!!.youWon,
                        fontSize = 10.sp,
                        color = Color.Black,
                        fontFamily = FontFamily.Serif
                    )
                }
                Box(){
                    Text(
                        text = "Score",
                        fontSize = 10.sp,
                        color = Color.Black,
                        fontFamily = FontFamily.Serif
                    )
                }
                Row {
                    Box(){
                        Text(
                            text = gameViewModel.endResults!!.yourScore.toString(),
                            fontSize = 10.sp,
                            color = Color.Black,
                            fontFamily = FontFamily.Serif
                        )
                    }
                    Box(){
                        Text(
                            text = "to",
                            fontSize = 10.sp,
                            color = Color.Black,
                            fontFamily = FontFamily.Serif
                        )
                    }
                    Box(){
                        Text(
                            text = gameViewModel.endResults!!.enemyScore.toString(),
                            fontSize = 10.sp,
                            color = Color.Black,
                            fontFamily = FontFamily.Serif
                        )
                    }
                }
                Button(onClick = { gameViewModel.WebSocketClient!!.close(); navController.navigate("home"); gameViewModel.resetAllValues(); }) {

                }

            }
        }
    }else{
        Scaffold (

            topBar = { ProfileRow(profileImage = R.drawable.caught, username = gameViewModel.gameInfo?.EnemyInfo?.Username ?: "", score = gameViewModel.gameInfo?.EnemyInfo?.Score ?: 0, mana = gameViewModel.gameInfo?.EnemyInfo?.mana?: "0")},

            bottomBar = { ProfileRow(profileImage = R.drawable.xdd, username = gameViewModel.user.userName, score =  gameViewModel.gameInfo?.YourInfo?.Score ?: 0, mana = gameViewModel.gameInfo?.YourInfo?.mana?: "0") }



        ){innerPadding ->
            Column (modifier = Modifier
                .padding(innerPadding)
                .background(Color.White)) {
                PlayField(gameViewModel)

            }
        }
    }


}