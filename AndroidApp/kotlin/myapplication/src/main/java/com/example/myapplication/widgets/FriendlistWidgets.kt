package com.example.myapplication.widgets


import android.content.Context
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column

import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Delete
import androidx.compose.material.icons.filled.PlayArrow
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Card
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.graphics.Color

import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.example.myapplication.viewmodels.HomeViewModel


@Composable
fun FriendList(modifier: Modifier, homeViewModel: HomeViewModel) {
    Column(
        modifier = Modifier
            .padding(16.dp)
            .fillMaxHeight(0.9f)) {



        Text("Search for a Friend:", style = MaterialTheme.typography.bodyMedium, fontFamily = FontFamily.Serif)
        TextField(value = homeViewModel.addFriend.value,
            onValueChange = { searchedFriend ->
                homeViewModel.addFriend.value = searchedFriend },
            label = {Text("username", fontFamily = FontFamily.Serif)},
            modifier = Modifier
                .padding(5.dp)
                .fillMaxWidth(0.6f))
        Button(
            modifier = Modifier.size(90.dp,30.dp),
            onClick = { homeViewModel.addNewFriend() },
            colors = ButtonDefaults.buttonColors(
                containerColor = Color.Black
            )
        ) {
            Text("Add",
                color = Color.White,
                fontSize = 13.sp,
                fontFamily = FontFamily.Serif
            )
        }
        Spacer(modifier = Modifier.height(8.dp))
        Text(text = "Friends", style = MaterialTheme.typography.headlineMedium, fontFamily = FontFamily.Serif)
        HorizontalDivider(modifier = Modifier
            .fillMaxWidth(0.6f))
        Spacer(modifier = Modifier.height(8.dp))

        if (homeViewModel.friends.value == null || homeViewModel.friends.value!!.isEmpty()) {
            Column {
                Text(text = "No friends yet. :(")
            }
        } else {
            LazyColumn (modifier = modifier){
                items(homeViewModel.friends.value!!) { friend ->
                    FriendField(friend.username, homeViewModel)
                }
            }
        }

    }
    Column (modifier = Modifier.padding(16.dp)){
        HorizontalDivider(modifier = Modifier
            .fillMaxWidth(0.6f))
        Text("username", fontFamily = FontFamily.Serif)
    }
}

@Composable
fun FriendField (
    friend: String, homeViewModel: HomeViewModel) {
    Card (modifier = Modifier.fillMaxWidth(0.6f)) {
        Row (horizontalArrangement = Arrangement.End){
            Text(text = friend, fontFamily = FontFamily.Serif, modifier = Modifier.padding(8.dp))
            Spacer(Modifier.weight(1f))
            Icon(modifier = Modifier.clickable { homeViewModel.removeFriend(friend) }.padding(8.dp),
                imageVector = Icons.Filled.Delete,
                contentDescription = "Delete from friendlist")
        }
        Spacer(modifier = Modifier.height(8.dp))
    }
}