package com.example.myapplication.widgets

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonColors
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import com.example.myapplication.viewmodels.GameViewModel

@Composable
fun EndScreenRow (gameViewModel: GameViewModel, navController: NavController) {
    Scaffold (

    ){innerPadding ->
        Column (modifier = Modifier
            .padding(innerPadding)
            .fillMaxSize()
            .background(Color.White),
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.SpaceBetween) {
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
                    fontSize = 15.sp,
                    color = Color.Black,
                    fontFamily = FontFamily.Serif
                )
            }
            Box(){
                Text(
                    text = "Score",
                    fontSize = 15.sp,
                    color = Color.Black,
                    fontFamily = FontFamily.Serif
                )
            }
            Row {
                Box(){
                    Text(
                        text = gameViewModel.endResults!!.yourScore.toString(),
                        fontSize = 15.sp,
                        color = Color.Black,
                        fontFamily = FontFamily.Serif
                    )
                }
                Box(){
                    Text(
                        text = "to",
                        fontSize = 15.sp,
                        color = Color.Black,
                        fontFamily = FontFamily.Serif
                    )
                }
                Box(){
                    Text(
                        text = gameViewModel.endResults!!.enemyScore.toString(),
                        fontSize = 15.sp,
                        color = Color.Black,
                        fontFamily = FontFamily.Serif
                    )
                }
            }
            Button(onClick = { gameViewModel.WebSocketClient!!.close(); navController.navigate("home"); gameViewModel.resetAllValues(); },
                colors = ButtonDefaults.buttonColors(containerColor = Color.Black),
                modifier = Modifier.size(100.dp,35.dp)
            ) {
                Text(text = "Home")
            }

        }
    }
}