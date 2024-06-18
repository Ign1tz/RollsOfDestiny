package com.example.myapplication.widgets

import android.content.Context
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.input.TextFieldValue
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController

@Composable
fun LoginBox (navController: NavController) {
    val padding = 10.dp
    Column (
        modifier = Modifier,
        horizontalAlignment = Alignment.CenterHorizontally

    ) {
        StartScreenTextFields(initialText = "username")
        StartScreenTextFields(initialText = "password")
        Spacer(modifier = Modifier.padding(padding))
        RegisterButton(navController)
        Spacer(Modifier.padding(padding))
        LoginButton()
    }
}

@Composable
fun StartScreenTextFields (initialText: String) {
    var textState by remember { mutableStateOf(TextFieldValue()) }

    TextField(value = textState,
        onValueChange = { textState = it},
        label = {Text(initialText)},
        modifier = Modifier
            .fillMaxWidth()
            .padding(16.dp))
}