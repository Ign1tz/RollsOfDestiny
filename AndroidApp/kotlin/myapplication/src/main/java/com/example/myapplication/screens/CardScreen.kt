package com.example.myapplication.screens

import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import com.example.myapplication.viewmodels.CardViewModel

@Composable
fun CardScreen (navController: NavController, cardViewModel: CardViewModel) {
    
    Text(text = "Ich bin der Cardscreen")
}