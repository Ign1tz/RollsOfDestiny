package com.example.myapplication.viewmodels

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import com.example.myapplication.localdb.Repository

class GameViewModel (val repository: Repository) : ViewModel(), BasicViewModel {

    val board = mutableStateListOf(
        mutableStateListOf(-1,-1,-1),
        mutableStateListOf(-1,-1,-1),
        mutableStateListOf(-1,-1,-1)
    )

    fun placeDie(column: Int) {
        for (row in 2 downTo 0) {
            if (board[row][column] == -1) {
                board[row][column] = 0
                break
            }
        }
    }

}