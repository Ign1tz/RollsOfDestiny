package com.example.myapplication.localdb

import android.content.Context
import androidx.room.Database
import androidx.room.Room
import androidx.room.RoomDatabase

@Database(entities = [User::class], version = 2, exportSchema = false)
abstract class userDatabase: RoomDatabase() {
    abstract fun dao():dao

    companion object {
        @Volatile
        private var instance: userDatabase? = null

        fun getDatabase(context: Context): userDatabase {
            return instance ?: synchronized(this) {
                Room.databaseBuilder(context, userDatabase::class.java, "user_db")
                    .fallbackToDestructiveMigration()
                    .allowMainThreadQueries()
                    .build()
                    .also { instance = it}
            }
        }
    }
}