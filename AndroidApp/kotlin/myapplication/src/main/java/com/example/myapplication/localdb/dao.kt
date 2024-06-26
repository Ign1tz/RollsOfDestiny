package com.example.myapplication.localdb

import androidx.room.Dao
import androidx.room.Delete
import androidx.room.Insert
import androidx.room.Update

@Dao
interface dao {
    @Insert
    fun Insert (user: User)

    @Delete
    fun Delete (user: User)

    @Update
    fun Update (user: User)

}