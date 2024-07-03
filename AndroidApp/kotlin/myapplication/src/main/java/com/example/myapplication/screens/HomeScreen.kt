package com.example.myapplication.screens

import android.graphics.drawable.Icon
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.DrawerValue
import androidx.compose.material3.ModalDrawerSheet
import androidx.compose.material3.ModalNavigationDrawer
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.rememberDrawerState
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import com.example.myapplication.widgets.BottomBar
import com.example.myapplication.widgets.HeaderTopBar
import com.example.myapplication.widgets.HomeScreenButtons
import androidx.compose.ui.Modifier
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.widgets.FriendlistDrawer
import com.example.myapplication.widgets.UpButton


@Composable
fun HomeScreen (navController: NavController, homeViewModel: HomeViewModel) {


    ModalNavigationDrawer (
        drawerContent = {
            ModalDrawerSheet (
                content = {
                    Text(text = "Friends")
                },
                modifier = Modifier
                    .fillMaxHeight()
                    .fillMaxWidth(0.6f)
            )
        },
        content = {
            Scaffold (
                topBar = {
                    HeaderTopBar(navController = navController, icon = "<")
                },
                bottomBar = {
                    BottomBar()
                }
            ){ innerPadding ->
                Column (modifier = Modifier.padding(innerPadding)) {
                    HomeScreenButtons(navController = navController)

                }
            }
        }
    )

}