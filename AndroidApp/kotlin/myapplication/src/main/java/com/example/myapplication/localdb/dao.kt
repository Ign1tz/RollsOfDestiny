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

    @Delete
    fun Delete (user: User)

    @Update
    fun Update (user: User)

    @Query("Select * from user where userId=:id")
    fun getUserById (id: String): Flow<User>

}