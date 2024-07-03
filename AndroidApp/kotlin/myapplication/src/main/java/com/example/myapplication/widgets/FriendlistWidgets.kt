package com.example.myapplication.widgets


import android.content.Context
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.DrawerValue
import androidx.compose.material3.ModalDrawerSheet
import androidx.compose.material3.ModalNavigationDrawer
import androidx.compose.material3.rememberDrawerState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.material3.Text
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController


@Composable
fun FriendlistDrawer (navController: NavController) {



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

        }
        )
}


