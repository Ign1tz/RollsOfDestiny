package com.example.myapplication.widgets

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Card
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import com.example.myapplication.viewmodels.Deck
import com.example.myapplication.viewmodels.DeckViewModel

@Composable
fun DeckList (modifier: Modifier, deckViewModel: DeckViewModel, deck: List<Deck>) {
    LazyColumn (modifier) {
        items(deckViewModel.getDeckList()) { deck -> DeckRow(deckViewModel, deck) }
    }
}

@Composable
fun DeckRow (deckViewModel: DeckViewModel, deck: Deck) {
    Card(modifier = Modifier
        .fillMaxWidth()
        .padding(5.dp)
        .clickable { /*TODO*/ }) {
        Column (modifier = Modifier.padding(10.dp) ){
            Text(text = deck.name, fontFamily = FontFamily.Serif)
            Row (horizontalArrangement = Arrangement.SpaceBetween, modifier = Modifier.fillMaxWidth()){
                Button(
                    onClick = {/*TODO*/ },
                    colors = ButtonDefaults.buttonColors(
                        containerColor = Color.Black
                    )
                ) {
                    Text(
                        "Activate",
                        color = Color.White,
                        fontFamily = FontFamily.Serif
                    )
                }
                Button(
                    onClick = {/*TODO*/ },
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