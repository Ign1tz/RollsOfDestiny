package com.example.myapplication.screens

import android.media.MediaPlayer
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
import androidx.compose.runtime.DisposableEffect
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.sp
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.compose.LifecycleEventEffect
import androidx.navigation.NavController
import com.example.myapplication.R
import com.example.myapplication.connection.websocket.WebSocketClient
import com.example.myapplication.types.AudioPlayer
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.widgets.ProfileRow
import com.example.myapplication.widgets.PlayField
import java.net.URI


@Composable
fun GameScreen (navController: NavController, gameViewModel: GameViewModel) {
    gameViewModel.websocket()
    var audio = AudioPlayer.getInstance().startAudio()
    if (audio != null){
        try {
            audio.start()
        }catch (e: Exception){
            AudioPlayer.getInstance().audioPlayer.value = null
            AudioPlayer.getInstance().audioStarted = false
            audio = AudioPlayer.getInstance().startAudio()
            audio?.start()
        }
    }
    AudioPlayer.getInstance().audioPlayer.value = audio
    DisposableEffect(Unit) {
        onDispose {
            audio?.release()
        }
    }
    LifecycleEventEffect(Lifecycle.Event.ON_STOP){
        Log.d("stop", "sssssssssssssssssssssssssssssssssssssssssssssssssssssssssss")
        audio?.pause()
    }
    LifecycleEventEffect(Lifecycle.Event.ON_RESUME){
        Log.d("stop", "sssssssssssssssssssssssssssssssssssssssssssssssssssssssssss")
        audio?.start()
    }

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

            topBar = { ProfileRow(profileImage = R.drawable.caught, gameViewModel, username = gameViewModel.gameInfo?.EnemyInfo?.Username ?: "", score = gameViewModel.gameInfo?.EnemyInfo?.Score ?: 0, mana = gameViewModel.gameInfo?.EnemyInfo?.mana?: "0")},

            bottomBar = { ProfileRow(profileImage = null, gameViewModel, username = gameViewModel.getUser()!!.userName, score =  gameViewModel.gameInfo?.YourInfo?.Score ?: 0, mana = gameViewModel.gameInfo?.YourInfo?.mana?: "0") }



        ){innerPadding ->
            Column (modifier = Modifier
                .padding(innerPadding)
                .background(Color.White)) {
                PlayField(gameViewModel)

            }
        }
    }


}