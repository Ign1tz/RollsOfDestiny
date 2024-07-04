package com.example.myapplication.widgets

import android.content.Context
import android.widget.TextView
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalUriHandler
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.input.TextFieldValue
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.android.volley.Request
import com.example.myapplication.navigation.Screen
import com.example.myapplication.viewmodels.LoginViewModel
import org.json.JSONObject
import java.net.URI

@Composable
fun LoginBox (navController: NavController, viewModel: LoginViewModel) {
    val padding = 10.dp
    Column (
        modifier = Modifier,
        horizontalAlignment = Alignment.CenterHorizontally

    ) {
        var nameState by remember { mutableStateOf("") }
        var pwState by remember { mutableStateOf("") }

        TextField(value = nameState,
            onValueChange = { newName ->
                nameState = newName },
            label = {Text("username")},
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp))

        TextField(value = pwState,
            onValueChange = { newPassword -> pwState = newPassword},
            label = {Text("password")},
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp))

        Spacer(modifier = Modifier.padding(padding))
        Button(
            modifier = Modifier.size(130.dp,50.dp),
            onClick = { if (viewModel.login(nameState, pwState)) {navController.navigate(Screen.HomeScreen.route)} },
            colors = ButtonDefaults.buttonColors(
                containerColor = Color.Black
            )
        ) {
            Text("Login",
                color = Color.White,
                fontSize = 25.sp,
                fontFamily = FontFamily.Serif
            )
        }
        Spacer(Modifier.padding(padding))
        val uriHandler = LocalUriHandler.current
        Button(
                modifier = Modifier.size(150.dp,50.dp),
        onClick = { uriHandler.openUri("http://${viewModel.IPADDRESS}:3000/signup")},
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
        ) {
        Text("SignUp",
            color = Color.White,
            fontSize = 25.sp,
            fontFamily = FontFamily.Serif
        )
    }
    }
}


