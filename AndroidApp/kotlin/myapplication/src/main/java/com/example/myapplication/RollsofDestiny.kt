package com.example.myapplication

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.example.myapplication.navigation.Navigation
import com.example.myapplication.ui.theme.RollsofDestinyAppTheme

class RollsofDestiny : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            RollsofDestinyAppTheme {
                Navigation()
            }
        }
    }
}