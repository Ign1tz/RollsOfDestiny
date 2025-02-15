package com.example.myapplication.widgets

import android.graphics.BitmapFactory
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
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.asImageBitmap
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.core.content.ContextCompat.getDrawable
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.myapplication.R
import com.example.myapplication.viewmodels.GameViewModel
import com.google.accompanist.drawablepainter.rememberDrawablePainter
import com.example.myapplication.types.Column
import com.example.myapplication.types.card
import com.example.myapplication.viewmodels.Card
import kotlinx.coroutines.delay
import kotlin.io.encoding.Base64
import kotlin.io.encoding.ExperimentalEncodingApi

@Composable
fun PlayField(viewModel: GameViewModel) {

    Column(verticalArrangement = Arrangement.SpaceBetween) {
        EnemyCardField(viewModel)
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
        Row(
            modifier = Modifier
                .fillMaxWidth()
                .padding(5.dp),
            horizontalArrangement = Arrangement.Start
        ) {
            Box() {

                if (viewModel.hasRolled.value && viewModel.gameInfo?.ActivePlayer?.active ?: false) {
                    Die(gameViewModel = viewModel)
                } else {
                    DefaultDie(gameViewModel = viewModel)
                }
            }
            CardField(gameViewModel = viewModel)
        }

    }
}

@Composable
fun CardField(gameViewModel: GameViewModel) {
    LazyRow(modifier = Modifier) {
        items(gameViewModel.gameInfo?.YourInfo?.deck?.inHand ?: listOf()) { card ->
            OwnSingleCard(card, gameViewModel)
        }
    }
}

@Composable
fun OwnSingleCard(card: card, gameViewModel: GameViewModel) {
    Image(
        painter = painterResource(id = gameViewModel.getCardImageById(card.name)),
        contentDescription = "",
        modifier = Modifier.clickable {
            if (gameViewModel.isActive.value) {
                gameViewModel.WebSocketClient!!.sendMessage("{\"type\":\"playCard\", \"messageBody\":\"${card.cardid}\", \"gameId\":\"${gameViewModel.gameInfo?.gameid}\", \"userid\":\"${gameViewModel.getUser()?.userid}\"}")
            }
        })
}

@Composable
fun EnemyCardField(gameViewModel: GameViewModel) {
    LazyRow {

        items(gameViewModel.gameInfo?.EnemyInfo?.deck?.inHand ?: 0) {
            EnemyCard()
        }

    }
}

@Composable
fun EnemyCard() {
    Image(
        painter = painterResource(id = R.drawable.cardback), contentDescription = ""
    )
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
            column(if (info != null) {
                info.YourInfo.LeftColumn
            } else null, { handleColumnClick(viewModel, info?.YourInfo?.LeftColumn, 0) })
            column(if (info != null) {
                info.YourInfo.MiddleColumn
            } else null, { handleColumnClick(viewModel, info?.YourInfo?.MiddleColumn, 1) })
            column(if (info != null) {
                info.YourInfo.RightColumn
            } else null, { handleColumnClick(viewModel, info?.YourInfo?.RightColumn, 2) })
        }
    }
}

fun handleColumnClick(viewModel: GameViewModel, column: Column?, key: Int) {
    Log.d("isFull?", column?.IsFull.toString())
    if (column?.IsFull ?: true && !viewModel.hasRolled.value && viewModel.pickedColumn.value) {
        return
    }
    if (viewModel.isActive.value) {
        viewModel.WebSocketClient!!.sendMessage("{\"type\":\"${viewModel.GameType.value}PickColumn\", \"messageBody\":\"${key.toString()}\", \"gameId\":\"${viewModel.gameInfo?.gameid}\", \"userid\":\"${viewModel.getUser()?.userid}\"}")
    }
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
            .background(Color.Transparent)) {
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
            Row(
                Modifier
                    .height(170.dp * 0.33333f)
                    .fillMaxWidth()
                    .border(2.dp, Color.Black),
                horizontalArrangement = Arrangement.Center
            ) {
                gridBox(third)
            }
            Row(
                Modifier
                    .height(170.dp * 0.33333f)
                    .fillMaxWidth()
                    .border(2.dp, Color.Black),
                horizontalArrangement = Arrangement.Center
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
                gridBox(first)
            }
        }
    }
}

@OptIn(ExperimentalEncodingApi::class)
@Composable
fun ProfileRow(profileImage: Int?,
    gameViewModel: GameViewModel, username: String, score: Int, mana: String
) {
    var user = gameViewModel.getUser()
    Box(
        modifier = Modifier
            .fillMaxWidth()
            .height(60.dp)
            .background(Color.Gray, shape = RoundedCornerShape(2.dp))
            .padding(horizontal = 16.dp, vertical = 4.dp),
    ) {
        Row(
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.SpaceBetween,
            modifier = Modifier.fillMaxWidth()
        ) {
            if (profileImage != null){

                Image(
                    painter = painterResource(id = profileImage) ,
                    contentDescription = user?.userName,
                    modifier = Modifier
                        .size(45.dp)
                        .padding(1.dp)
                        .clip(CircleShape),
                    contentScale = ContentScale.Crop
                )
            }else{

                Image(
                    bitmap = BitmapFactory.decodeByteArray(Base64.decode(user!!.profilePicture, 0), 0, Base64.decode(user!!.profilePicture, 0).size).asImageBitmap(),
                    contentDescription = user?.userName,
                    modifier = Modifier
                        .size(45.dp)
                        .padding(1.dp)
                        .clip(CircleShape),
                    contentScale = ContentScale.Crop
                )
            }
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
            Spacer(modifier = Modifier.weight(1f))
            Column {
                Text(
                    text = "Mana:",
                    fontSize = 10.sp,
                    color = Color.White,
                    fontFamily = FontFamily.Serif
                )
                Text(
                    text = mana,
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
                .clickable {
                    gameViewModel.hasRolled.value = true

                    if (gameViewModel.isActive.value && !gameViewModel.roll.value) {
                        gameViewModel.WebSocketClient!!.sendMessage("{\"type\":\"rolled\", \"messageBody\":\"\", \"gameId\":\"${gameViewModel.gameInfo?.gameid}\", \"userid\":\"${gameViewModel.getUser()?.userid}\"}")
                    }
                },
            contentScale = ContentScale.FillBounds
        )
    }
}

fun handleDieRoll(gameViewModel: GameViewModel) {
    gameViewModel.roll.value = true
}

@Composable
fun Die(gameViewModel: GameViewModel) {
    Box(
        modifier = Modifier.clickable(onClick = {
            Log.d("roll", gameViewModel.roll.value.toString())
            handleDieRoll(gameViewModel)
        })
    ) {

        val diceSize = 90.dp
        if (gameViewModel.roll.value || true) {
            Log.d("game", gameViewModel.gameInfo?.ActivePlayer?.roll.toString())
            when (gameViewModel.gameInfo?.ActivePlayer?.roll) {
                "1" -> Image(
                    painter = rememberDrawablePainter(
                        drawable = getDrawable(
                            LocalContext.current, R.drawable.dice1
                        ),

                        ),
                    contentDescription = null,
                    contentScale = ContentScale.FillBounds,
                    modifier = Modifier.size(diceSize)
                )

                "2" -> Image(
                    painter = rememberDrawablePainter(
                        drawable = getDrawable(
                            LocalContext.current, R.drawable.dice2
                        )
                    ),
                    contentDescription = null,
                    contentScale = ContentScale.Fit,
                    modifier = Modifier.size(diceSize)
                )

                "3" -> Image(
                    painter = rememberDrawablePainter(
                        drawable = getDrawable(
                            LocalContext.current, R.drawable.dice3
                        )
                    ),
                    contentDescription = null,
                    contentScale = ContentScale.Fit,
                    modifier = Modifier.size(diceSize)
                )

                "4" -> Image(
                    painter = rememberDrawablePainter(
                        drawable = getDrawable(
                            LocalContext.current, R.drawable.dice4
                        )
                    ),
                    contentDescription = null,
                    contentScale = ContentScale.Fit,
                    modifier = Modifier.size(diceSize)
                )

                "5" -> Image(
                    painter = rememberDrawablePainter(
                        drawable = getDrawable(
                            LocalContext.current, R.drawable.dice5
                        )
                    ),
                    contentDescription = null,
                    contentScale = ContentScale.Fit,
                    modifier = Modifier.size(diceSize)
                )

                "6" -> Image(
                    painter = rememberDrawablePainter(
                        drawable = getDrawable(
                            LocalContext.current, R.drawable.dice6
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



