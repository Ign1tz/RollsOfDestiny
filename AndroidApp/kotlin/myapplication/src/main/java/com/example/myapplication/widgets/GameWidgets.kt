package com.example.myapplication.widgets

import android.util.Log
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.core.content.ContextCompat.getDrawable
import coil.compose.rememberAsyncImagePainter
import coil.compose.rememberImagePainter
import com.example.myapplication.R
import com.example.myapplication.viewmodels.GameViewModel
import com.google.accompanist.drawablepainter.rememberDrawablePainter
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.myapplication.localdb.Repository

@Composable
fun PlayField (int: Int, viewModel: GameViewModel) {

    Column {
        Column (
            modifier = Modifier
                .fillMaxWidth()
                .fillMaxHeight(0.5f),
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.Bottom
        ) {
            EnemyField()
        }
        HorizontalDivider(
            color = Color.Black,
            thickness = 1.dp,
            modifier = Modifier.fillMaxWidth())
        Column (
            modifier = Modifier
                .fillMaxWidth(),
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.Top
        ){
            OwnField(viewModel)
        }
        Column (
            modifier = Modifier
                .fillMaxWidth()
                .padding(5.dp),
            horizontalAlignment = Alignment.Start
        ) {
            Box() {
                var originStateOfDie by remember { mutableStateOf(false) }

                if (originStateOfDie) {
                    Dice(int = int)
                } else {
                    DefaultDie(onClick = {originStateOfDie = true})
                }
            }

        }

    }
}

@Composable
fun OwnField (viewModel: GameViewModel) {
    Box() {
        val image: Painter = painterResource(id = R.drawable.grid_image)
        Image(
            painter = image,
            contentDescription = null,
            modifier = Modifier
                .width(170.dp)
                .height(170.dp),
            contentScale = ContentScale.Crop
        )
        Row() {
            repeat(3) { index ->
                Box(
                    modifier = Modifier
                        .height(170.dp)
                        .width(57.dp)
                        .clickable { }
                        .background(Color.Transparent)
                ) {
                    Column () {
                        repeat(3) {index ->
                            Box(modifier = Modifier
                                .fillMaxSize()
                                .background(Color.Transparent)
                            ){

                            }
                        }
                    }
                }
            }
        }
    }
}

@Composable
fun addingDice (int: Int) {

    val image1: Painter = painterResource(id = R.drawable.appdie1)
    val image2: Painter = painterResource(id = R.drawable.appdie2)
    val image3: Painter = painterResource(id = R.drawable.appdie3)
    val image4: Painter = painterResource(id = R.drawable.appdie4)
    val image5: Painter = painterResource(id = R.drawable.appdie5)
    val image6: Painter = painterResource(id = R.drawable.appdie6)

    when (int) {
        1 -> Image(painter = image1, contentDescription = "die 1")
        2 -> Image(painter = image2, contentDescription = "die 2")
        3 -> Image(painter = image3, contentDescription = "die 3")
        4 -> Image(painter = image4, contentDescription = "die 4")
        5 -> Image(painter = image5, contentDescription = "die 5")
        6 -> Image(painter = image6, contentDescription = "die 6")
    }
}

@Composable
fun EnemyField () {
    Box() {
        val image: Painter = painterResource(id = R.drawable.grid_image)
        Image(
            painter = image,
            contentDescription = null,
            modifier = Modifier
                .width(170.dp)
                .height(170.dp),
            contentScale = ContentScale.Crop
        )
    }
}

@Composable
fun ProfileRow(
    profileImage: Int,
    username: String,
    score: String
) {
    Box(
        modifier = Modifier
            .fillMaxWidth()
            .height(80.dp)
            .background(Color.Gray, shape = RoundedCornerShape(2.dp))
            .padding(horizontal = 16.dp, vertical = 8.dp),
    ) {
        Row(
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.SpaceBetween,
            modifier = Modifier.fillMaxWidth()
        ) {
            Image(
                painter = painterResource(id = profileImage),
                contentDescription = null,
                modifier = Modifier
                    .size(55.dp)
                    .padding(1.dp),
                contentScale = ContentScale.Crop
            )
            Spacer(modifier = Modifier.width(16.dp))
            Column {
                Text(
                    text = "Username",
                    fontSize = 10.sp,
                    color = Color.White,
                    fontFamily = FontFamily.Serif
                )
                Text(
                    text = username,
                    fontSize = 20.sp,
                    color = Color.White,
                    fontFamily = FontFamily.Serif
                )
            }
            Spacer(modifier = Modifier.weight(1f))
            Column {
                Text(
                    text = "Score:",
                    fontSize = 10.sp,
                    color = Color.White,
                    fontFamily = FontFamily.Serif
                )
                Text(
                    text = score,
                    fontSize = 20.sp,
                    color = Color.White,
                    fontFamily = FontFamily.Serif
                )
            }
        }
    }
}

@Composable
fun DefaultDie (onClick: () -> Unit) {
    Box (modifier = Modifier.clickable(onClick = onClick)) {
        val image: Painter = painterResource(id = R.drawable.die_picture)
        Image(painter = image,
            contentDescription = null,
            modifier = Modifier
                .width(90.dp)
                .height(90.dp),
            contentScale = ContentScale.FillBounds
        )
    }
}

@Composable
fun Dice (int: Int) {
    Box (
    ) {

        val diceSize = 90.dp
        when (int) {
            1 -> Image(
                painter = rememberDrawablePainter(
                    drawable = getDrawable(
                        LocalContext.current,
                        R.drawable.dice1
                    ),

                ),
                contentDescription = null,
                contentScale = ContentScale.FillBounds,
                modifier = Modifier.size(diceSize)
            )
            2 -> Image(
                painter = rememberDrawablePainter(
                    drawable = getDrawable(
                        LocalContext.current,
                        R.drawable.dice2
                    )
                ),
                contentDescription = null,
                contentScale = ContentScale.Fit,
                modifier = Modifier.size(diceSize)
            )
            3 -> Image(
                painter = rememberDrawablePainter(
                    drawable = getDrawable(
                        LocalContext.current,
                        R.drawable.dice3
                    )
                ),
                contentDescription = null,
                contentScale = ContentScale.Fit,
                modifier = Modifier.size(diceSize)
            )
            4 -> Image(
                painter = rememberDrawablePainter(
                    drawable = getDrawable(
                        LocalContext.current,
                        R.drawable.dice4
                    )
                ),
                contentDescription = null,
                contentScale = ContentScale.Fit,
                modifier = Modifier.size(diceSize)
            )
            5 -> Image(
                painter = rememberDrawablePainter(
                    drawable = getDrawable(
                        LocalContext.current,
                        R.drawable.dice5
                    )
                ),
                contentDescription = null,
                contentScale = ContentScale.Fit,
                modifier = Modifier.size(diceSize)
            )
            6 -> Image(
                painter = rememberDrawablePainter(
                    drawable = getDrawable(
                        LocalContext.current,
                        R.drawable.dice6
                    )
                ),
                contentDescription = null,
                contentScale = ContentScale.Fit,
                modifier = Modifier.size(diceSize)
            )
        }
    }
}

