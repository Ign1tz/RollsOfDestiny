package com.example.myapplication.viewmodels

import android.util.Log
import androidx.lifecycle.ViewModel
import com.example.myapplication.localdb.Repository
import com.example.myapplication.localdb.User
import com.example.myapplication.types.token
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.client.statement.HttpResponse
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.json.Json
import io.ktor.http.ContentType
import io.ktor.http.contentType
import kotlinx.coroutines.runBlocking

class LoginViewModel(val repository: Repository) : ViewModel(), BasicViewModel {


    private val IPADDRESS = "192.168.0.181"


    suspend private fun loginRequest(userName: String, password: String): Boolean {
        val client = HttpClient(CIO) {
            install(ContentNegotiation) {
                json(Json {
                    prettyPrint = true
                    isLenient = true
                    ignoreUnknownKeys =
                        true // Useful if the JSON has more fields than the data class
                })
            }
        }

        try {
            val responseText: HttpResponse = client.post("http://$IPADDRESS:9090/login") {
                contentType(ContentType.Application.Json)
                setBody("{\"username\":\"$userName\", \"password\":\"$password\"}")
            }

            if (responseText.status.value != 200) {
                return false
            }
            Log.d("HttpTest", "Response text: $responseText")

            val token: token = Json.decodeFromString(responseText.body())
            Log.d("HttpTest", "Received token: ${token.token}")
            repository.returnDelete(
                User(
                    userName = userName,
                    password = password,
                    jwtToken = token.token
                )
            )
            repository.returnInsert(
                User(
                    userName = userName,
                    password = password,
                    jwtToken = token.token
                )
            )
            client.close()
            return true
        } catch (e: Exception) {
            Log.d("HttpTest", "Received error: ${e.message}")
        }
        return false
    }

    fun login(userName: String, password: String): Boolean {
        var worked = false

        runBlocking { worked = loginRequest(userName, password) }
        Log.d("HttpTest", worked.toString())

        return worked
    }

    fun checkAlreadyLoggedIn(): Boolean {
        var worked = false
        runBlocking {
            val isEmpty = repository.isEmpty()
            if (!isEmpty) {
                val newUser = repository.getUser()
                worked = loginRequest(newUser.userName, newUser.password)
            }
        }

        return worked
    }
}