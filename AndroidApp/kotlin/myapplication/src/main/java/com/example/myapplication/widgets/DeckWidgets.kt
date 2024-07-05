package com.example.myapplication.widgets

import android.graphics.drawable.Icon
import android.util.Log
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Card
import androidx.compose.material3.Icon
import androidx.compose.material3.SegmentedButtonDefaults.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import com.example.myapplication.types.deck
import com.example.myapplication.viewmodels.DeckViewModel
import androidx.compose.material3.Icon
import androidx.compose.material3.TextField
import androidx.compose.ui.unit.sp

@Composable
fun DeckList (modifier: Modifier, deckViewModel: DeckViewModel, navController: NavController) {
    LazyColumn (modifier) {
        items(deckViewModel.decks.value!!) { deck ->
            DeckRow(deckViewModel,
                deck,
                navController,
                onItemClick = {
                    deckName ->
                    navController.navigate(route = "cards")
                }
            )
        }
        items(1){
            Card(modifier = Modifier
                .fillMaxWidth()
                .padding(5.dp)
                .clickable {  }) {

                TextField(value = deckViewModel.newDeckName.value,
                    onValueChange = { newName ->
                        deckViewModel.newDeckName.value = newName },
                    label = {Text("New Deck")},
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(16.dp))
                Button(
                    modifier = Modifier.size(180.dp,50.dp),
                    onClick = { deckViewModel.createNewDeck() },
                    colors = ButtonDefaults.buttonColors(
                        containerColor = Color.Black
                    )
                ) {
                    Text("Create new Deck!",
                        color = Color.White,
                        fontFamily = FontFamily.Serif
                    )
                }

            }
        }
    }
}

@Composable
fun DeckRow (deckViewModel: DeckViewModel,
             deck: deck,
             navController: NavController,
             onItemClick: (deck) -> Unit = {}) {
    Card(modifier = Modifier
        .fillMaxWidth()
        .padding(5.dp)
        .clickable { navController.navigate("deckDetails/${deck.deckid}") }) {
        Column (modifier = Modifier.padding(10.dp) ){
            Text(text = deck.name, fontFamily = FontFamily.Serif)
            Row (horizontalArrangement = Arrangement.SpaceBetween, modifier = Modifier.fillMaxWidth()){
                Button(
                    onClick = { deckViewModel.setActive(deck) },
                    colors = ButtonDefaults.buttonColors(
                        containerColor = Color.Black
                    )
                ) {
                    Text(
                        if (deck.active) "Active" else "Activate",
                        color = Color.White,
                        fontFamily = FontFamily.Serif
                    )
                }
                Button(
                    onClick = { deckViewModel.removeDeck(deck) },
                    colors = ButtonDefaults.buttonColors(
                        containerColor = Color.Black
                    )
                ) {
                    Text(
                        "Delete",
                        color = Color.White,
                        fontFamily = FontFamily.Serif
                    )
                }
            }
        }
    }
}