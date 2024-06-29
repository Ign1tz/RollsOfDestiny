package com.example.myapplication.viewmodels

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.myapplication.localdb.Repository
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.request.get
import io.ktor.client.statement.HttpResponse
import kotlinx.coroutines.launch

class LoginViewModel (val repository: Repository) : ViewModel(), BasicViewModel {

    suspend fun testHttp () {
        val client = HttpClient(CIO)

        val response: HttpResponse = client.get("https://ktor.io/")

        Log.d("HttpTest", response.status.toString())

        client.close()

    }

    fun testenWirZusammen () {
        viewModelScope.launch { testHttp() }
    }
}