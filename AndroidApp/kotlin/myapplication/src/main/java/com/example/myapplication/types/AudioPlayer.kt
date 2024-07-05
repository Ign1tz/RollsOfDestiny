package com.example.myapplication.types

import android.media.MediaPlayer
import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateOf
import androidx.compose.ui.platform.LocalContext
import com.example.myapplication.R

class AudioPlayer {

    var audioStarted = false
    val audioPlayer = mutableStateOf<MediaPlayer?>(null)

    companion object {

        @Volatile
        private var instance: AudioPlayer? = null

        fun getInstance() =
            instance ?: synchronized(this) {
                instance ?: AudioPlayer().also { instance = it }
            }
    }




    @Composable
    fun startAudio(): MediaPlayer? {
        if (!audioStarted) {
            val audio = MediaPlayer.create(LocalContext.current, R.raw.background_music);
            audio.setLooping(true);
            audioStarted = true
            return audio
        }
        return audioPlayer.value
    }
}