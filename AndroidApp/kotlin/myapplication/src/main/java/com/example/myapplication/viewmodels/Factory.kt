package com.example.myapplication.viewmodels

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.example.myapplication.localdb.Repository

class Factory (private val repository: Repository):ViewModelProvider.Factory{

    private val IPADDRESS = "menews.site"
    @Override
    override fun <T:ViewModel> create(model: Class<T>):T=when (model) {
        LoginViewModel::class.java -> LoginViewModel(repository = repository, IPADDRESS)
        HomeViewModel::class.java -> HomeViewModel(repository = repository, IPADDRESS)
        SettingViewModel::class.java -> SettingViewModel(repository = repository, IPADDRESS)
        GameViewModel::class.java -> GameViewModel(repository = repository, IPADDRESS)
        ScoreboardViewModel::class.java -> ScoreboardViewModel(repository = repository, IPADDRESS)
        else -> throw IllegalArgumentException("Oh oh!")
    } as T
}