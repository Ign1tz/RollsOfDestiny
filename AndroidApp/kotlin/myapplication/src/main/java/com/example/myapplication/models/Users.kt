package com.example.myapplication.models

import androidx.room.Entity
import com.example.myapplication.localdb.User
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder


data class User(
    val id: Long,
    var userName: String,
    var password: String
)
object UserRepository {


    fun addUser (user: User) {

    }

    fun setPassword (password: String): String {
        val passwordEncoder = BCryptPasswordEncoder()
        return passwordEncoder.encode(password)
    }

}