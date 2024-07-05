package com.example.myapplication.widgets

import android.content.Context
import android.util.Log
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
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
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import androidx.navigation.compose.rememberNavController
import com.example.myapplication.navigation.Screen
import com.example.myapplication.viewmodels.GameViewModel
import com.example.myapplication.viewmodels.HomeViewModel
import com.example.myapplication.viewmodels.LoginViewModel
import com.example.myapplication.viewmodels.SettingViewModel
import com.example.myapplication.viewmodels.ScoreboardViewModel
import android.content.ClipData
import android.content.ClipboardManager
import android.widget.Toast
import androidx.compose.foundation.clickable

import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.request.get
import io.ktor.client.statement.HttpResponse
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.launch


//@Composable
//fun LoginButton (viewModel: LoginViewModel) {
//
//}
@Composable
fun LoginButton(viewModel: LoginViewModel) {

}


//@Composable
//fun RegisterButton (viewModel: LoginViewModel) {
//    Button(
//        modifier = Modifier.size(130.dp,50.dp),
//        onClick = {  },
//        colors = ButtonDefaults.buttonColors(
//            containerColor = Color.Black
//        )
//    ) {
//        Text("Register",
//            color = Color.White,
//            fontSize = 20.sp,
//            fontFamily = FontFamily.Serif
//        )
//    }
//}


@Composable
fun TopButton (navController: NavController) {
    Button(
        modifier = Modifier
            .size(50.dp, 50.dp),
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        ),
        shape = RoundedCornerShape(10.dp),
        onClick = { navController.popBackStack() }
    ) {
        Text("<",
            color = Color.White,
            fontSize = 25.sp,
            fontFamily = FontFamily.Serif,
            textAlign = TextAlign.Center,
            modifier = Modifier.align(Alignment.CenterVertically)
        )
    }
}

//@Composable
//fun UpButton () {
//    Button(
//        modifier = Modifier
//            .size(50.dp,50.dp),
//        colors = ButtonDefaults.buttonColors(
//            containerColor = Color.Black),
//        shape = RoundedCornerShape(10.dp),
//        onClick = {}
//    ) {
//        Text("<",
//            color = Color.White,
//            fontSize = 25.sp,
//            fontFamily = FontFamily.Serif,
//            textAlign = TextAlign.Center,
//            modifier = Modifier.align(Alignment.CenterVertically)
//        )
//    }
//}

@Composable
fun QuickPlayButton(navController: NavController, gameViewModel: GameViewModel) {

    Button(
        modifier = Modifier.size(300.dp, 50.dp),
        onClick = { navController.navigate(route = "game/bot"); gameViewModel.GameType.value = "bot" },
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
fun FriendPlayButton(homeViewModel: HomeViewModel, navController: NavController) {

    val isFriendPlayClicked by homeViewModel::isFriendPlayClicked
    val isHostButtonClicked by homeViewModel::isHostButtonClicked
    val isJoinButtonClicked by homeViewModel::isJoinButtonClicked


    Button(
        modifier = Modifier.size(300.dp, 50.dp),
        onClick = { homeViewModel.toggleFriendClick() },
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
    // extra buttons when Play vs Friend is clicked
    if (isFriendPlayClicked) {
        Button(
            modifier = Modifier.size(300.dp, 50.dp),
            onClick = {homeViewModel.toggleHostButtonClicked()},
            colors = ButtonDefaults.buttonColors(
                containerColor = Color.Gray
            )
        ) {
            Text(
                "Host",
                color = Color.White,
                fontSize = 20.sp,
                fontFamily = FontFamily.Serif
            )
        }
        Button(
            modifier = Modifier.size(300.dp, 50.dp),
            onClick = {homeViewModel.toggleJoinButtonClicked()},
            colors = ButtonDefaults.buttonColors(
                containerColor = Color.Gray
            )
        ) {
            Text(
                "Join",
                color = Color.White,
                fontSize = 20.sp,
                fontFamily = FontFamily.Serif
            )
        }
    }
    if (isJoinButtonClicked) {


        TextField(value = homeViewModel.friendId.value,
            onValueChange = { newName ->
                homeViewModel.friendId.value = newName },
            label = {Text("Enter User ID")},
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp))
            Button(
                modifier = Modifier.size(130.dp,50.dp),
                onClick = { navController.navigate("game/friend/" + homeViewModel.friendId.value) },
                colors = ButtonDefaults.buttonColors(
                    containerColor = Color.Black
                )
            ) {
                Text("Play",
                    color = Color.Green,
                    fontSize = 25.sp,
                    fontFamily = FontFamily.Serif
                )
            }
    }
    if (isHostButtonClicked) {
        UserIDTextWithCopyClickboard(homeViewModel)
        Button(
            modifier = Modifier.size(130.dp,50.dp),
            onClick = { navController.navigate("game/friend/ ") },
            colors = ButtonDefaults.buttonColors(
                containerColor = Color.Black
            )
        ) {
            Text("Play",
                color = Color.Green,
                fontSize = 25.sp,
                fontFamily = FontFamily.Serif
            )
        }
    }
}

@Composable
fun UserIDTextWithCopyClickboard(homeViewModel: HomeViewModel) {
    val context = LocalContext.current
    homeViewModel.getUser()?.userid?.let { userId ->
        Text(
            text = userId,
            color = Color.Black,
            fontSize = 14.sp,
            fontFamily = FontFamily.Serif,
            modifier = Modifier.clickable {
                copyToClipboard(context, userId)
                Toast.makeText(context, "Copied to clipboard", Toast.LENGTH_SHORT).show()
            }
        )
    }
}

fun copyToClipboard(context: Context, text: String) {
    val clipboard = context.getSystemService(Context.CLIPBOARD_SERVICE) as ClipboardManager
    val clip = ClipData.newPlainText("UserId", text)
    clipboard.setPrimaryClip(clip)
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
fun ScoreboardButton(navController: NavController, scoreboardViewModel: ScoreboardViewModel) {
    Button(
        modifier = Modifier.size(300.dp, 50.dp),
        onClick = { navController.navigate(route = "scoreboard") },
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
        modifier = Modifier.size(100.dp,35.dp),
        onClick = { navController.navigate("login"); viewModel.repository.returnDelete()},
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text("LogOut",
            color = Color.White,
            fontSize = 13.sp,
            fontFamily = FontFamily.Serif
        )
    }
}

@Composable
fun DecksButton (navController: NavController) {
    Button(
        modifier = Modifier.size(300.dp,50.dp),
        onClick = {navController.navigate(route = "decks")},
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Black
        )
    ) {
        Text("Card Decks",
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
@Composable
fun DeleteAccount(navController: NavController, settingViewModel: SettingViewModel) {
    Button(
        modifier = Modifier.size(230.dp, 40.dp),
        onClick = {
            settingViewModel.deleteAccount()
            navController.navigate("login")

        },
        colors = ButtonDefaults.buttonColors(
            containerColor = Color.Red
        )
    ) {
        Text(
            "Delete Account",
            color = Color.White,
            fontSize = 18.sp,
            fontFamily = FontFamily.Serif
        )
    }
}