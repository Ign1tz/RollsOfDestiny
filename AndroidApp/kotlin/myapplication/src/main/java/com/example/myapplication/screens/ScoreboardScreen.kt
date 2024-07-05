package com.example.myapplication.screens

import android.graphics.BitmapFactory
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.asImageBitmap
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import coil.compose.rememberAsyncImagePainter
import coil.compose.rememberImagePainter
import com.example.myapplication.types.scoreboardPlayer
import com.example.myapplication.viewmodels.ScoreboardViewModel
import com.example.myapplication.widgets.BottomBar
import com.example.myapplication.widgets.HeaderTopBar
import kotlin.io.encoding.Base64
import kotlin.io.encoding.ExperimentalEncodingApi

@OptIn(ExperimentalEncodingApi::class)
@Composable
fun DisplayPlayer(player: scoreboardPlayer, place: Int) {
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .padding(horizontal = 10.dp, vertical = 5.dp)
            .background(color = Color(0xFFb5b5b5), shape = RoundedCornerShape(5.dp))
            .padding(10.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.SpaceBetween
    ) {
        Image(
            bitmap = BitmapFactory.decodeByteArray(Base64.decode(player.profilePicture, 0), 0, Base64.decode(player.profilePicture, 0).size).asImageBitmap(),
            contentDescription = player.username,
            modifier = Modifier
                .size(50.dp)
                .clip(CircleShape),
            contentScale = ContentScale.Crop
        )

        Spacer(modifier = Modifier.width(10.dp))
        Column(
            modifier = Modifier.weight(1f)
        ) {
            Text(
                text = player.username,
                style = TextStyle(fontSize = 20.sp)
            )
            Text(
                text = "Rating: ${player.rating}",
                style = TextStyle(fontSize = 14.sp, color = Color.Gray)
            )
        }
        Text(
            text = place.toString(),
            style = TextStyle(fontSize = 16.sp, color = Color.Gray),
            modifier = Modifier.align(Alignment.CenterVertically)
        )
    }
}


@Composable
fun ScoreBoardScreen(navController: NavController, scoreboardViewModel: ScoreboardViewModel) {

    scoreboardViewModel.getPlayers()

    Scaffold(
        topBar = {
            HeaderTopBar(navController, "<")
        },
        bottomBar = {
            BottomBar()
        }
    ) { innerPadding ->
        Column(
            modifier = Modifier
                .padding(innerPadding)
                .fillMaxSize()
                .background(Color.White)
        ) {
            Text(
                text = "Top Ten",
                style = TextStyle(fontSize = 32.sp, fontWeight = FontWeight.Bold),
                modifier = Modifier.padding(20.dp)
            )
            LazyColumn {
                itemsIndexed(scoreboardViewModel.players.value) { index, player ->
                    DisplayPlayer(player, index + 1)
                }
            }
        }
    }
}
