package com.example.myapplication.localdb

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "user")
data class User (
    @PrimaryKey(autoGenerate = true)
    val userId: Long=0,
    var userName: String,
    var password: String
)