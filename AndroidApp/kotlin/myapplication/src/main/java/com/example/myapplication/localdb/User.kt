package com.example.myapplication.localdb

import android.provider.ContactsContract.Profile
import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "user")
data class User (
    @PrimaryKey()
    var userid: String,
    var password: String,
    var jwtToken: String,
    var userName: String,
    var email: String,
    var profilePicture: String,
    var rating: String
)