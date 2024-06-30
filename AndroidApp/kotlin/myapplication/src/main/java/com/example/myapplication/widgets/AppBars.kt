

package com.example.myapplication.widgets

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material3.BottomAppBar
import androidx.compose.material3.CenterAlignedTopAppBar
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import androidx.navigation.compose.rememberNavController
import java.util.Objects


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun HeaderTopBar (navController: NavController, icon: String) {

    Column {
        CenterAlignedTopAppBar(
            modifier = Modifier
                .fillMaxWidth(),
            title = {},
            navigationIcon = { TopButton(navController, icon) },
        )
        TitleTopBar()
    }
}

@Composable
fun TitleTopBar () {
    Box(
        modifier = Modifier
            .fillMaxWidth(),
        contentAlignment = Alignment.BottomCenter,
    ) {
        Column (
            horizontalAlignment = Alignment.CenterHorizontally,
        ) {
            Text(
                text = "Rolls of",
                style = TextStyle(
                    color = Color.Black,
                    fontSize = 50.sp,
                    fontFamily = FontFamily.Serif
                )
            )
            Text(
                text = "Destiny",
                style = TextStyle(
                    color = Color.Black,
                    fontSize = 50.sp,
                    fontFamily = FontFamily.Serif
                )
            )
        }
    }
}
@Composable
fun BottomBar () {
    BottomAppBar {
        Row (
            modifier = Modifier
                .fillMaxWidth(),
            horizontalArrangement = Arrangement.Center
        ) {
            Text(text = "Brezina Lukas  ", color = Color.Black, fontFamily = FontFamily.Serif)
            Text(text = "Pertl Moritz  ", color = Color.Black, fontFamily = FontFamily.Serif)
            Text(text = "Weisser Simon", color = Color.Black, fontFamily = FontFamily.Serif)
        }
    }
}
