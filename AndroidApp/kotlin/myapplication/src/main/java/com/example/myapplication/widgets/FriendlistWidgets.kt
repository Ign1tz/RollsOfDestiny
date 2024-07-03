package com.example.myapplication.widgets


import android.content.Context
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.DrawerValue
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalDrawerSheet
import androidx.compose.material3.ModalNavigationDrawer
import androidx.compose.material3.NavigationDrawerItem
import androidx.compose.material3.rememberDrawerState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.myapplication.localdb.User


@Composable
fun FriendList() {
    Column(
        modifier = Modifier.padding(16.dp).fillMaxHeight(0.9f)) {

        var friendState by remember { mutableStateOf("") }

        Text(text = "Friends", style = MaterialTheme.typography.headlineMedium, fontFamily = FontFamily.Serif)
        Text("Search for a Friend:", style = MaterialTheme.typography.bodyMedium, fontFamily = FontFamily.Serif)
        TextField(value = friendState,
            onValueChange = { searchedFriend ->
                friendState = searchedFriend },
            label = {Text("username", fontFamily = FontFamily.Serif)},
            modifier = Modifier
                .padding(16.dp)
                .fillMaxWidth(0.6f))
        HorizontalDivider(modifier = Modifier
            .fillMaxWidth(0.6f))
        Spacer(modifier = Modifier.height(8.dp))
        val friends = listOf("Alice", "Bob", "Charlie", "David")
        friends.forEach { friend ->
            Text(text = friend, fontFamily = FontFamily.Serif)
            Spacer(modifier = Modifier.height(4.dp))
        }


    }
    Column (modifier = Modifier.padding(16.dp)){
        HorizontalDivider(modifier = Modifier
            .fillMaxWidth(0.6f))
        Text("username", fontFamily = FontFamily.Serif)
    }
}
