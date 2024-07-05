package com.example.myapplication.screens

import android.annotation.SuppressLint
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import com.example.myapplication.widgets.ScreenTopBar



@Composable
fun RulesButton(navController: NavController) {
    Button(
        modifier = Modifier.size(300.dp, 50.dp),
        onClick = { navController.navigate(route = "rules")},
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Rules",
            color = Color.White,
            fontSize = 20.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
@SuppressLint("UnusedMaterial3ScaffoldPaddingParameter")
fun RulesScreen(navController: NavController) {
    Scaffold(
        topBar = { ScreenTopBar(navController = navController, title = "Rules") }
    ) { paddingValues ->
        Column(
            modifier = Modifier
                .padding(16.dp)
                .fillMaxWidth()
                .background(Color(0xFFF3E5F5), shape = RoundedCornerShape(10.dp))
                .padding(16.dp)
        ) {
            Text(
                text = "Rules",
                fontSize = 24.sp,
                fontWeight = FontWeight.Bold,
                color = Color(0xFF4A148C),
                modifier = Modifier.align(Alignment.CenterHorizontally)
            )
            Spacer(modifier = Modifier.height(16.dp))
            Text(
                text = "Roll your own Destiny.",
                fontSize = 16.sp,
                fontWeight = FontWeight.Medium,
                textAlign = TextAlign.Center,
                modifier = Modifier.align(Alignment.CenterHorizontally)
            )
            Spacer(modifier = Modifier.height(16.dp))
            RulesList()
        }
    }
}

@Composable
fun RulesList() {
    LazyColumn(
        modifier = Modifier
            .fillMaxWidth()
            .padding(8.dp)
    ) {
        item {
            val rulesText = buildAnnotatedString {
                withStyle(style = SpanStyle(fontWeight = FontWeight.Bold, color = Color(0xFF6A1B9A))) {
                    append("General Rules\n")
                }
                append("• Play with each other and have fun.\n")
                append("• Do your best.\n")
                append("• No harassing and racism.\n")
                append("• You play for yourself, do not cheat.\n\n")

                withStyle(style = SpanStyle(fontWeight = FontWeight.Bold, color = Color(0xFF6A1B9A))) {
                    append("Game Rules\n")
                }
                append("• You roll a die.\n")
                append("• Place die in a column of your grid.\n")
                append("• Opponent does the same.\n")
                append("• Points are calculated by values of your columns in your grid.\n")
                append("• If you have the same value 2 or even 3 times in your column, it is multiplied with each other.\n")
                append("• You can optionally play cards which give you certain effects.\n")
                append("• If you win, you gain rating points. If you lose, you lose rating points.\n")
                append("• You can play against a computer, against your friend or against a random enemy based on queueing.\n")
                append("• The game is over when one grid is filled up completely.\n")
                append("• Do your best and roll your destiny!")
            }

            Text(text = rulesText, color = Color(0xFF6A1B9A), lineHeight = 24.sp)
        }
    }
}
