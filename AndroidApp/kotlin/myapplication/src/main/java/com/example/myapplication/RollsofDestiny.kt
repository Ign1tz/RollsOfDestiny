package com.example.myapplication

import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.example.myapplication.navigation.Navigation
import com.example.myapplication.ui.theme.RollsofDestinyAppTheme

class RollsofDestiny : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            RollsofDestinyAppTheme(darkTheme = false) {
                Navigation()
            }
        }
    }
    override fun onStart() {
        super.onStart()
        Log.i("MainActivity", "onStart called.")
    }

    override fun onResume() {
        super.onResume()
        Log.i("MainActivity", "onResume called.")
    }

    override fun onPause() {
        super.onPause()
        Log.i("MainActivity", "onPause called.")
    }

    override fun onStop() {
        super.onStop()
        Log.i("MainActivity", "onStop called.")
    }

    override fun onRestart() {
        super.onRestart()
        Log.i("MainActivity", "onRestart called.")
    }

    override fun onDestroy() {
        super.onDestroy()
        Log.i("MainActivity", "onDestroy called.")
    }
}