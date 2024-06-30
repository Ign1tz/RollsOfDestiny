package com.example.myapplication.widgets

import android.util.Log
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.material3.Divider
import androidx.compose.material3.HorizontalDivider
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
import androidx.compose.ui.unit.dp
import androidx.core.content.ContextCompat.getDrawable
import com.example.myapplication.R
import com.google.accompanist.drawablepainter.rememberDrawablePainter

@Composable
fun PlayField (int: Int) {

    Column {
        Column (
            modifier = Modifier
                .fillMaxWidth()
                .fillMaxHeight(0.5f),
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.Bottom
        ) {
            NewGrid()
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
            NewGrid()
        }
        Column (
            modifier = Modifier
                .fillMaxWidth(),
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
fun NewGrid () {
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
                        .clickable {
                            Log.d("grid", "$index")
                        }
                        .background(Color.Transparent)
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

