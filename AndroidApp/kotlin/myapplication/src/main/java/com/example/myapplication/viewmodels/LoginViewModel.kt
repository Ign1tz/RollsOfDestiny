package com.example.myapplication.viewmodels

import android.util.Log
import androidx.compose.ui.platform.LocalUriHandler
import androidx.lifecycle.ViewModel
import com.example.myapplication.localdb.Repository
import com.example.myapplication.localdb.User
import com.example.myapplication.types.token
import com.example.myapplication.types.userInfoMessage
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.get
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.client.statement.HttpResponse
import io.ktor.client.utils.EmptyContent.headers
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.json.Json
import io.ktor.http.ContentType
import io.ktor.http.HttpHeaders
import io.ktor.http.contentType
import io.ktor.http.headers
import kotlinx.coroutines.runBlocking
import io.ktor.client.plugins.auth.*
import io.ktor.client.plugins.auth.providers.BearerTokens
import io.ktor.client.plugins.auth.providers.basic
import io.ktor.client.plugins.auth.providers.bearer
import okhttp3.Credentials.basic
import kotlinx.coroutines.withContext

class LoginViewModel(val repository: Repository) : ViewModel(), BasicViewModel {


    val IPADDRESS = "10.0.0.2"


    suspend private fun loginRequest(userName: String, password: String): Boolean {
        var client = HttpClient(CIO) {
            install(ContentNegotiation) {
                json(Json {
                    prettyPrint = true
                    isLenient = true
                    ignoreUnknownKeys =
                        true // Useful if the JSON has more fields than the data class
                })
            }

        }

        val responseText: HttpResponse = client.post("http://$IPADDRESS:9090/login") {
            contentType(ContentType.Application.Json)
            setBody("{\"username\":\"$userName\", \"password\":\"$password\"}")
        }

        if (responseText.status.value != 200) {
            return false
        }
        val token: token = Json.decodeFromString(responseText.body())
        val newClient = HttpClient(CIO) {
            install(ContentNegotiation) {
                json(Json {
                    prettyPrint = true
                    isLenient = true
                    ignoreUnknownKeys =
                        true // Useful if the JSON has more fields than the data class
                })
            }
            install(Auth) {
                bearer {
                    loadTokens {
                        BearerTokens(token.token, token.token)
                    }
                    sendWithoutRequest { true }
                }
            }

        }
        val response: HttpResponse = newClient.get("http://$IPADDRESS:9090/userInfo") {
        }
        Log.d("userinfo", response.status.value.toString())
        if (response.status.value == 200) {
            val userInfo: userInfoMessage = Json.decodeFromString(response.body())
            repository.returnDelete()
            repository.returnInsert(
                User(
                    userName = userInfo.username,
                    password = password,
                    jwtToken = token.token,
                    userid = userInfo.userid,
                    email = userInfo.email,
                    profilePicture = userInfo.profilePicture,
                    rating = userInfo.rating
                )
            )
            client.close()
            return true
        } else {
            return false
        }
        return false
    }

    fun login(userName: String, password: String): Boolean {
        var worked = false

        runBlocking { worked = loginRequest(userName, password) }
        Log.d("test", worked.toString())
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