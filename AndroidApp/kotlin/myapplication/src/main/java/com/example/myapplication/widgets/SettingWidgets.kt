package com.example.myapplication.widgets

import android.util.Log
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import com.example.myapplication.viewmodels.SettingViewModel

@Composable
fun CenterSettings (viewModel: SettingViewModel, navController: NavController) {
    val padding = 10.dp
    Spacer(modifier = Modifier.padding(padding))
    Column (modifier = Modifier
        .fillMaxWidth(),
        horizontalAlignment = Alignment.CenterHorizontally) {

        OutlinedTextField(
            value = viewModel.username.value,
            onValueChange = { viewModel.username.value = it;
                            Log.d("change", viewModel.username.value + " " + it)},
            label = { Text("New Username") },
            placeholder = { Text("Enter new username") },
            modifier = Modifier
                .fillMaxWidth()
                .padding(padding)
        )

        UsernameConfirmation(navController = navController, settingViewModel = viewModel)


        OutlinedTextField(
            value = viewModel.oldPassword.value,
            onValueChange = { viewModel.oldPassword.value = it },
            label = { Text("Old Password") },
            placeholder = { Text("Enter old password") },
            visualTransformation = PasswordVisualTransformation(),
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
            modifier = Modifier
                .fillMaxWidth()
                .padding(padding)
        )


        OutlinedTextField(
            value = viewModel.newPassword.value,
            onValueChange = { viewModel.newPassword.value = it },
            label = { Text("New Password") },
            placeholder = { Text("Enter new password") },
            visualTransformation = PasswordVisualTransformation(),
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
            modifier = Modifier
                .fillMaxWidth()
                .padding(padding)
        )


        OutlinedTextField(
            value = viewModel.confirmNewPassword.value,
            onValueChange = { viewModel.confirmNewPassword.value = it },
            label = { Text("Confirm New Password") },
            placeholder = { Text("Re-enter new password") },
            visualTransformation = PasswordVisualTransformation(),
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
            modifier = Modifier
                .fillMaxWidth()
                .padding(padding)
        )

        PasswordConfirmation(navController = navController, settingViewModel = viewModel)
        DeleteAccount(navController = navController, settingViewModel = viewModel)
    }
}