package com.example.myapplication.widgets


import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.DrawerValue
import androidx.compose.material3.ModalDrawerSheet
import androidx.compose.material3.ModalNavigationDrawer
import androidx.compose.material3.rememberDrawerState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.material3.Text
import androidx.compose.ui.unit.dp


@Composable
fun FriendlistDrawer () {

    val drawerState = rememberDrawerState(DrawerValue.Closed)

    ModalNavigationDrawer ( drawerContent = { /*TODO*/ }) {

    }
}

@Composable
fun FriendlistDrawerSheet () {

    ModalDrawerSheet (
        content = {
            Text(text = "Friends")
        }
    )
}

