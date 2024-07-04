package com.example.myapplication.screens

import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.NavController
import com.example.myapplication.viewmodels.Deck
import com.example.myapplication.viewmodels.DeckViewModel
import com.example.myapplication.widgets.DeckList
import com.example.myapplication.widgets.ScreenTopBar

@Composable
fun DeckScreen (deckViewModel: DeckViewModel, navController: NavController) {

    Scaffold (
        topBar = { ScreenTopBar(navController = navController, "Your Decks") },

    ){innerPadding ->
        DeckList(modifier = Modifier.padding(innerPadding), deckViewModel, deck = deckViewModel.getDeckList(), navController)
    }

}