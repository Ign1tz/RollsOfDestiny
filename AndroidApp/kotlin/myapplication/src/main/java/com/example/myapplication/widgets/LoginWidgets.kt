package com.example.myapplication.widgets

import android.content.Context
import android.widget.TextView
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
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.android.volley.Request
import com.example.myapplication.connection.performHttpRequest
import com.example.myapplication.viewmodels.LoginViewModel
import org.json.JSONObject

@Composable
fun LoginBox (navController: NavController, viewModel: LoginViewModel) {
    val padding = 10.dp
    Column (
        modifier = Modifier,
        horizontalAlignment = Alignment.CenterHorizontally

    ) {
        var nameState by remember { mutableStateOf(TextFieldValue()) }
        var pwState by remember { mutableStateOf(TextFieldValue()) }

        TextField(value = nameState,
            onValueChange = { nameState = it},
            label = {Text("username")},
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp))

        TextField(value = pwState,
            onValueChange = { pwState = it},
            label = {Text("password")},
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp))

        Spacer(modifier = Modifier.padding(padding))
        RegisterButton()
        Spacer(Modifier.padding(padding))
        LoginButton(viewModel)
    }
}

fun makeRequest(context: Context) {
    val url = "https://www.google.com"
    val requestBody = JSONObject()

    performHttpRequest(
        context,
        url,
        Request.Method.GET,
        requestBody,
        onSuccess = { response ->
            println("Response: $response")
        },
        onError = { error ->
            println("Error: $error")
        }
    )
}

