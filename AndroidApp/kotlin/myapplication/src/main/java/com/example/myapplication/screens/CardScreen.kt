package com.example.myapplication.screens

import android.util.Log
import androidx.compose.foundation.Image
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import com.example.myapplication.viewmodels.CardViewModel
import com.example.myapplication.widgets.DeckList
import com.example.myapplication.widgets.DeckRow
import com.example.myapplication.widgets.ScreenTopBar

@Composable
fun CardScreen(navController: NavController, cardViewModel: CardViewModel) {
    cardViewModel.getDeck()
    Scaffold(
        topBar = {
            ScreenTopBar(
                navController = navController,
                cardViewModel.deck.value?.name ?: " "
            )
        },

        ) { innerPadding ->
        LazyColumn(Modifier.padding(innerPadding)) {
            items(cardViewModel.cards.value!!) { cardname ->
                var inDeck = cardViewModel.deck.value!!.cards.any { it.name == cardname }
                Log.d("inDeck", inDeck.toString())
                Column {
                    Image(
                        painter = painterResource(id = cardViewModel.getCardImageById(cardname)),
                        contentDescription = "",
                        modifier = Modifier
                            .border(3.dp, if (inDeck) Color.Green else Color.Red)
                            .clickable {if (inDeck) cardViewModel.removeFromDeck(cardname) else cardViewModel.addToDeck(cardname)},
                    )
                    Text(text = if (inDeck) "Remove." else "Add.")

                }
            }

        }
    }
}