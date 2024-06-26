package com.example.myapplication.viewmodels

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.myapplication.localdb.Repository

class LoginViewModel (val repository: Repository) : ViewModel(), BasicViewModel {


}