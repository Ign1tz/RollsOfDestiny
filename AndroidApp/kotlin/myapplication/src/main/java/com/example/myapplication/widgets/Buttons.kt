package com.example.myapplication.widgets

import android.content.Context
import android.util.Log
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalUriHandler
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import androidx.navigation.compose.rememberNavController
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.viewmodels.LoginViewModel
import com.example.myapplication.viewmodels.SettingViewModel
import com.example.myapplication.viewmodels.ScoreboardViewModel
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.request.get
import io.ktor.client.statement.HttpResponse
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.launch


@Composable
fun LoginButton(viewModel: LoginViewModel) {

}


@Composable
fun RegisterButton() {
    Button(
        modifier = Modifier.size(130.dp, 50.dp),
        onClick = { },
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Register",
            color = Color.White,
            fontSize = 20.sp,
            fontFamily = FontFamily.Serif
        )
    }
}


@Composable
fun TopButton(navController: NavController, icon: String) {
    Button(
        modifier = Modifier
            .size(50.dp, 50.dp),
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        ),
        shape = RoundedCornerShape(10.dp),
        onClick = { navController.popBackStack() }
    ) {
        Text(
            icon,
            color = Color.White,
            fontSize = 25.sp,
            fontFamily = FontFamily.Serif,
            textAlign = TextAlign.Center,
            modifier = Modifier.align(Alignment.CenterVertically)
        )
    }
}

@Composable
fun UpButton() {
    Button(
        modifier = Modifier
            .size(50.dp, 50.dp),
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        ),
        shape = RoundedCornerShape(10.dp),
        onClick = {}
    ) {
        Text(
            "<",
            color = Color.White,
            fontSize = 25.sp,
            fontFamily = FontFamily.Serif,
            textAlign = TextAlign.Center,
            modifier = Modifier.align(Alignment.CenterVertically)
        )
    }
}

@Composable
fun QuickPlayButton(navController: NavController, gameViewModel: GameViewModel) {
    gameViewModel.GameType.value = "bot"
    Button(
        modifier = Modifier.size(300.dp, 50.dp),
        onClick = { navController.navigate(route = "game/bot") },
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Quickplay",
            color = Color.White,
            fontSize = 20.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
fun FriendPlayButton() {
    Button(
        modifier = Modifier.size(300.dp, 50.dp),
        onClick = {},
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Play vs Friend",
            color = Color.White,
            fontSize = 20.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
fun RankedPlayButton(navController: NavController, gameViewModel: GameViewModel) {
    gameViewModel.GameType.value = ""
    Button(
        modifier = Modifier.size(300.dp, 50.dp),
        onClick = { navController.navigate(route = "game") },
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Ranked",
            color = Color.White,
            fontSize = 20.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
fun ScoreboardButton (navController: NavController, scoreboardViewModel: ScoreboardViewModel) {
    Button(
        modifier = Modifier.size(300.dp,50.dp),
        onClick = {navController.navigate(route = "scoreboard")},
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Scoreboard",
            color = Color.White,
            fontSize = 20.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
fun SettingsButton(navController: NavController) {
    Button(
        modifier = Modifier.size(300.dp, 50.dp),
        onClick = { navController.navigate(route = "setting") },
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Settings",
            color = Color.White,
            fontSize = 20.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
fun LogOut(viewModel: LoginViewModel, navController: NavController) {
    Button(
        modifier = Modifier.size(300.dp,50.dp),
        onClick = {viewModel.repository.returnDelete(); navController.navigate("login")},
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "LogOut",
            color = Color.White,
            fontSize = 20.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
fun PasswordConfirmation(navController: NavController, settingViewModel: SettingViewModel) {
    Button(
        modifier = Modifier.size(230.dp, 40.dp),
        onClick = {
            if (settingViewModel.changePassword()) {
                navController.navigate(route = "home")
            } else {
                settingViewModel.oldPassword.value = ""
                settingViewModel.newPassword.value = ""
                settingViewModel.confirmNewPassword.value = ""
            }
        },
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Confirm Password",
            color = Color.White,
            fontSize = 18.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
fun UsernameConfirmation(navController: NavController, settingViewModel: SettingViewModel) {
    Button(
        modifier = Modifier.size(230.dp, 40.dp),
        onClick = {
            if (settingViewModel.changeUsername()) {
                navController.navigate(route = "home")
            } else {
                settingViewModel.username.value = ""
            }
        },
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text(
            "Confirm Username",
            color = Color.White,
            fontSize = 18.sp,
            fontFamily = FontFamily.Serif
        )
    }
}