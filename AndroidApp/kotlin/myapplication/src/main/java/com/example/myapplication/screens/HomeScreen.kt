package com.example.myapplication.screens

import android.graphics.drawable.Icon
import android.util.Log
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.waterfall
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Menu
import androidx.compose.material3.CenterAlignedTopAppBar
import androidx.compose.material3.DrawerValue
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalDrawerSheet
import androidx.compose.material3.ModalNavigationDrawer
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.rememberDrawerState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.rememberCoroutineScope
import androidx.navigation.NavController
import com.example.myapplication.widgets.BottomBar
import com.example.myapplication.widgets.HomeScreenButtons
import androidx.compose.ui.Modifier
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.viewmodels.LoginViewModel
import com.example.myapplication.viewmodels.ScoreboardViewModel
import com.example.myapplication.widgets.FriendList
import com.example.myapplication.widgets.LogOut
import com.example.myapplication.widgets.TitleTopBar
import kotlinx.coroutines.launch


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun HomeScreen (navController: NavController, homeViewModel: HomeViewModel, loginViewModel: LoginViewModel, gameViewModel: GameViewModel, scoreboardViewModel: ScoreboardViewModel) {

    val drawerState = rememberDrawerState(DrawerValue.Closed)
    val scope = rememberCoroutineScope()
    homeViewModel.getFriends()
    ModalNavigationDrawer(

        drawerState = drawerState,
        drawerContent = {
            ModalDrawerSheet {
                FriendList(modifier = Modifier, homeViewModel)
            }
        },
        content = {
            Scaffold (
                topBar = {
                    Column {
                        CenterAlignedTopAppBar(
                            modifier = Modifier
                                .fillMaxWidth(),
                            title = {},
                            navigationIcon = { LogOut(loginViewModel, navController = navController) },
                            actions = {
                                IconButton(onClick = { scope.launch {   drawerState.open()} }) {
                                    Icon(
                                        imageVector = Icons.Filled.Menu,
                                        contentDescription = "Friendlist"
                                    )
                                }
                            }
                        )
                        TitleTopBar()
                        Log.d("bug", "before")
                    }
                },
                bottomBar = {
                    BottomBar()
                }
            ){ innerPadding ->
                Column (modifier = Modifier.padding(innerPadding)) {
                    HomeScreenButtons(navController = navController, loginViewModel, gameViewModel, scoreboardViewModel, homeViewModel)
                }
            }
        })
}