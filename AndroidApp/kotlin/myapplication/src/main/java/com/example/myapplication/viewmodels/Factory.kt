package com.example.myapplication.viewmodels

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.example.myapplication.localdb.Repository

class Factory (private val repository: Repository):ViewModelProvider.Factory{
    @Override
    override fun <T:ViewModel> create(model: Class<T>):T=when (model) {
        LoginViewModel::class.java -> LoginViewModel(repository = repository)
        HomeViewModel::class.java -> HomeViewModel(repository = repository)
        GameViewModel::class.java -> GameViewModel(repository = repository)
        DeckViewModel::class.java -> DeckViewModel(repository = repository)
        else -> throw IllegalArgumentException("Oh oh!")
    } as T
}