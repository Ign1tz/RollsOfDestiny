package com.example.myapplication.viewmodels

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.myapplication.localdb.Repository
import com.example.myapplication.localdb.User
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.request.get
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.client.statement.HttpResponse
import kotlinx.coroutines.launch

class LoginViewModel (val repository: Repository) : ViewModel(), BasicViewModel {

    private val ipAddress = "192.168.0.181"
    suspend private fun testHttp (userName: String, password: String) {
        val client = HttpClient(CIO)


        //Log.d("HttpTest", System.getenv("LOCAL_IP"))

        val response: HttpResponse = client.post("http://" + ipAddress + ":9090/login") {
            setBody("{\"username\":\"" + userName +"\", \"password\":\"" + password+"\"}")
        }

        Log.d("HttpTest", response.status.toString())

        client.close()

    }

    fun login (userName: String, password: String) {
        viewModelScope.launch { testHttp(userName, password) }
        Log.d("HttpTest", "login")

    }
}