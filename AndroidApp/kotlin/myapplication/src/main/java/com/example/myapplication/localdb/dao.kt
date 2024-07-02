package com.example.myapplication.localdb

import androidx.room.Dao
import androidx.room.Delete
import androidx.room.Insert
import androidx.room.Query
import androidx.room.Update
import kotlinx.coroutines.flow.Flow

@Dao
interface dao {
    @Insert
    fun Insert (user: User)

    @Query("Delete from user")
    fun Delete ()

    @Update
    fun Update (user: User)

    @Query("Select * from user")
    fun getUser (): User

    @Query("SELECT (SELECT COUNT(*) FROM user) == 0")
    fun isEmpty(): Boolean

}