package com.example.myapplication.viewmodels

import android.content.Context
import com.example.myapplication.localdb.userDatabase
import com.example.myapplication.localdb.Repository

object Injector {
    private fun getRepository(context: Context): Repository {
        return Repository.returnInstance(userDatabase.getDatabase(context.applicationContext).dao())
    }

    fun provideModelFactory(context: Context): Factory {
        val repository = getRepository(context)
        return Factory(repository = repository)
    }
}