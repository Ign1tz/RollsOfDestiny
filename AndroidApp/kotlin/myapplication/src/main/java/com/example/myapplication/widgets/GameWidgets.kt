package com.example.myapplication.widgets

import android.util.Log
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
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
import androidx.compose.ui.layout.VerticalAlignmentLine
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.core.content.ContextCompat.getDrawable
import com.example.myapplication.R
import com.example.myapplication.viewmodels.GameViewModel
import com.google.accompanist.drawablepainter.rememberDrawablePainter
import com.example.myapplication.types.Column
import kotlinx.coroutines.delay

@Composable
fun PlayField(viewModel: GameViewModel) {

    Column(verticalArrangement = Arrangement.SpaceBetween) {
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .fillMaxHeight(0.4f)
                .padding(20.dp),
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.Bottom
        ) {
            EnemyField(viewModel)
        }
        Box(
            Modifier
                .fillMaxWidth()
                .height(1.dp)
                .background(Color.Red)
        )
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .padding(20.dp),
            horizontalAlignment = Alignment.CenterHorizontally,
        ) {
            OwnField(viewModel)
        }
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .padding(5.dp),
            horizontalAlignment = Alignment.Start
        ) {
            Box() {

                if (viewModel.hasRolled.value && viewModel.gameInfo?.ActivePlayer?.active ?: false) {
                    Die(gameViewModel = viewModel)
                } else {
                    DefaultDie(gameViewModel = viewModel)
                }
            }

        }

    }
}

@Composable
fun OwnField(viewModel: GameViewModel) {
    val info = viewModel.gameInfo
    Box() {
        //val image: Painter = painterResource(id = R.drawable.grid_image)
        /*Image(
            painter = image,
            contentDescription = null,
            modifier = Modifier
                .width(170.dp)
                .height(170.dp),
            contentScale = ContentScale.Crop
        )*/
        Row() {
            column(
                if (info != null) {
                    info.YourInfo.LeftColumn
                } else null,
                { handleColumnClick(viewModel, info?.YourInfo?.LeftColumn, 0) }
            )
            column(
                if (info != null) {
                    info.YourInfo.MiddleColumn
                } else null,
                { handleColumnClick(viewModel, info?.YourInfo?.MiddleColumn, 1) }
            )
            column(
                if (info != null) {
                    info.YourInfo.RightColumn
                } else null,
                { handleColumnClick(viewModel, info?.YourInfo?.RightColumn, 2) }
            )
        }
    }
}

fun handleColumnClick(viewModel: GameViewModel, column: Column?, key: Int){
    Log.d("isFull?", column?.IsFull.toString())
    if (column?.IsFull ?: true){
        return
    }
    viewModel.WebSocketClient!!.sendMessage("{\"type\":\"${viewModel.GameType.value}PickColumn\", \"messageBody\":\"${key.toString()}\", \"gameId\":\"${viewModel.gameInfo?.gameid}\"}")
}

@Composable
fun column(column: Column?, onClick: () -> Unit = {}) {

    Log.d("columnInfo", column.toString())
    var first = column?.First
    var second = column?.Second
    var third = column?.Third
    Log.d("columnInfoSeperate", first.toString() + second.toString() + third.toString())

    if (first == null) {
        first = "0"
    }
    if (second == null) {
        second = "0"
    }
    if (third == null) {
        third = "0"
    }
    Box( //woks as one column
        modifier = Modifier
            .height(170.dp)
            .width(57.dp)
            .clickable { onClick() }
            .background(Color.Transparent)
    ) {
        Column(
            Modifier.fillMaxHeight(), verticalArrangement = Arrangement.SpaceBetween
        ) {
            Row(
                Modifier
                    .height(170.dp * 0.33333f)
                    .fillMaxWidth()
                    .border(2.dp, Color.Black),
                horizontalArrangement = Arrangement.Center
            ) {
                gridBox(first)
            }
            Row(
                Modifier
                    .height(170.dp * 0.33333f)
                    .fillMaxWidth()
                    .border(2.dp, Color.Black),
                horizontalArrangement = Arrangement.Center,

            ) {
                gridBox(second)
            }
            Row(
                Modifier
                    .height(170.dp * 0.33333f)
                    .fillMaxWidth()
                    .border(2.dp, Color.Black),
                horizontalArrangement = Arrangement.Center
            ) {
                gridBox(third)
            }
        }
    }
}


@Composable
fun gridBox(value: String) {
    Box(
        modifier = Modifier
            .background(Color.Transparent)
            .fillMaxWidth()
            .padding(start = 3.dp, top = 3.dp, end = 0.dp, bottom = 0.dp)
    ) {
        getDice(int = value.toInt())
    }
}

@Composable
fun getDice(int: Int) {

    val image1: Painter = painterResource(id = R.drawable.appdie1)
    val image2: Painter = painterResource(id = R.drawable.appdie2)
    val image3: Painter = painterResource(id = R.drawable.appdie3)
    val image4: Painter = painterResource(id = R.drawable.appdie4)
    val image5: Painter = painterResource(id = R.drawable.appdie5)
    val image6: Painter = painterResource(id = R.drawable.appdie6)

    when (int) {
        0 -> return
        1 -> return Image(
            painter = image1,
            contentDescription = "die 1",
            Modifier.size(50.dp),
            alignment = Alignment.Center
        )

        2 -> return Image(
            painter = image2,
            contentDescription = "die 2",
            Modifier.size(50.dp),
            alignment = Alignment.Center
        )

        3 -> return Image(
            painter = image3,
            contentDescription = "die 3",
            Modifier.size(50.dp),
            alignment = Alignment.Center
        )

        4 -> return Image(
            painter = image4,
            contentDescription = "die 4",
            Modifier.size(50.dp),
            alignment = Alignment.Center
        )

        5 -> return Image(
            painter = image5,
            contentDescription = "die 5",
            Modifier.size(50.dp),
            alignment = Alignment.Center
        )

        6 -> return Image(
            painter = image6,
            contentDescription = "die 6",
            Modifier.size(50.dp),
            alignment = Alignment.Center
        )
    }
}

@Composable
fun EnemyField(viewModel: GameViewModel) {
    val info = viewModel.gameInfo
    Box() {
        /*val image: Painter = painterResource(id = R.drawable.grid_image)
        Image(
            painter = image,
            contentDescription = null,
            modifier = Modifier
                .width(170.dp)
                .height(170.dp),
            contentScale = ContentScale.Crop
        )*/
        Row() {
            enemyColumn(
                if (info != null) {
                    info.EnemyInfo.LeftColumn
                } else null
            )
            enemyColumn(
                if (info != null) {
                    info.EnemyInfo.MiddleColumn
                } else null
            )
            enemyColumn(
                if (info != null) {
                    info.EnemyInfo.RightColumn
                } else null
            )
        }
    }
}

@Composable
fun enemyColumn(column: Column?) {

    Log.d("columnInfo", column.toString())
    var first = column?.First
    var second = column?.Second
    var third = column?.Third

    if (first == null) {
        first = "0"
    }
    if (second == null) {
        second = "0"
    }
    if (third == null) {
        third = "0"
    }
    Box( //woks as one column
        modifier = Modifier
            .height(170.dp)
            .width(57.dp)
            .background(Color.Transparent)
    ) {
        Column(Modifier.fillMaxHeight(), verticalArrangement = Arrangement.SpaceBetween) {
            Row(Modifier
                .height(170.dp * 0.33333f)
                .fillMaxWidth()
                .border(2.dp, Color.Black), horizontalArrangement = Arrangement.Center) {
                gridBox(third)
            }
            Row(Modifier
                .height(170.dp * 0.33333f)
                .fillMaxWidth()
                .border(2.dp, Color.Black), horizontalArrangement = Arrangement.Center) {
                gridBox(second)
            }
            Row(Modifier
                .height(170.dp * 0.33333f)
                .fillMaxWidth()
                .border(2.dp, Color.Black), horizontalArrangement = Arrangement.Center) {
                gridBox(first)
            }
        }
    }
}

@Composable
fun ProfileRow(
    profileImage: Int,
    username: String,
    score: Int
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
                    text = score.toString(),
                    fontSize = 20.sp,
                    color = Color.White,
                    fontFamily = FontFamily.Serif
                )
            }
        }
    }
}

@Composable
fun DefaultDie(gameViewModel: GameViewModel) {
    Log.d("die", "hhhhhhhhhahsdhfahdf")
    Box() {
        val image: Painter = painterResource(id = R.drawable.die_picture)
        Image(
            painter = image,
            contentDescription = null,
            modifier = Modifier
                .width(90.dp)
                .height(90.dp)
                .clickable { gameViewModel.hasRolled.value = true },
            contentScale = ContentScale.FillBounds
        )
    }
}

fun handleDieRoll(gameViewModel: GameViewModel) {
    Log.d("game", gameViewModel.roll.value.toString())
    gameViewModel.roll.value = true
    Log.d("game", gameViewModel.roll.value.toString())
}

@Composable
fun Die(gameViewModel: GameViewModel) {
    Box(
        modifier = Modifier.clickable(onClick = { handleDieRoll(gameViewModel) })
    ) {

        val diceSize = 90.dp
        if (gameViewModel.roll.value || true) {
            Log.d("game", gameViewModel.gameInfo?.ActivePlayer?.roll.toString())
            when (gameViewModel.gameInfo?.ActivePlayer?.roll) {
                "1" -> Image(
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

                "2" -> Image(
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

                "3" -> Image(
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

                "4" -> Image(
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

                "5" -> Image(
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

                "6" -> Image(
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
            gameViewModel.delayRoll()
        }
    }
}



