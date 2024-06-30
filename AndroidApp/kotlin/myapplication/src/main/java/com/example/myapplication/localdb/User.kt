package com.example.myapplication.localdb

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "user")
data class User (
    @PrimaryKey()
    var userName: String,
    var password: String,
    var jwtToken: String
)