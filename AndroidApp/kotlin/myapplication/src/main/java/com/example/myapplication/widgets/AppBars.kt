

package com.example.myapplication.widgets

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Menu
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.BottomAppBar
import androidx.compose.material3.CenterAlignedTopAppBar
import androidx.compose.material3.DrawerValue
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.rememberDrawerState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import androidx.navigation.compose.rememberNavController
import com.example.myapplication.navigation.Screen
import com.example.myapplication.viewmodels.LoginViewModel
import java.util.Objects


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun HomeTopBar (navController: NavController, loginViewModel: LoginViewModel) {

    Column {
        CenterAlignedTopAppBar(
            modifier = Modifier
                .fillMaxWidth(),
            title = {},
            navigationIcon = { LogOut(loginViewModel, navController = navController) },
        )
        TitleTopBar()
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ScreenTopBar (navController: NavController, title: String) {

    Column {
        CenterAlignedTopAppBar(
            modifier = Modifier
                .fillMaxWidth(),
            title = { Text(title) },
            navigationIcon = { TopButton(navController = navController)},
        )
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
                .fillMaxWidth()
                .padding(10.dp),
            horizontalArrangement = Arrangement.Absolute.SpaceBetween

        ) {
            Column {
                Text(text = "Lukas", color = Color.Black, fontFamily = FontFamily.Serif)
                Text(text = "Brezina", color = Color.Black, fontFamily = FontFamily.Serif)
            }
            Column {
                Text(text = "Moritz", color = Color.Black, fontFamily = FontFamily.Serif)
                Text(text = "Pertl", color = Color.Black, fontFamily = FontFamily.Serif)
            }
            Column {
                Text(text = "Simon", color = Color.Black, fontFamily = FontFamily.Serif)
                Text(text = "Weisser", color = Color.Black, fontFamily = FontFamily.Serif)
            }
        }
    }
}
