package com.example.myapplication.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import com.example.myapplication.localdb.Repository
import com.example.myapplication.localdb.User
import com.example.myapplication.types.token
import com.example.myapplication.types.userInfoMessage
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.auth.Auth
import io.ktor.client.plugins.auth.providers.BearerTokens
import io.ktor.client.plugins.auth.providers.bearer
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.get
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.client.statement.HttpResponse
import io.ktor.http.ContentType
import io.ktor.http.contentType
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.runBlocking
import kotlinx.serialization.json.Json

class SettingViewModel (val repository: Repository, val IPADDRESS: String): ViewModel(), BasicViewModel  {


    var username = mutableStateOf("")
    var oldPassword = mutableStateOf("")
    var newPassword = mutableStateOf("")
    var confirmNewPassword = mutableStateOf("")
    var user by mutableStateOf(repository.getUser())


    suspend private fun changeUsernameRequest(): Boolean {

        var user = repository.getUser()

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
                        BearerTokens(user.jwtToken, user.jwtToken)
                    }
                    sendWithoutRequest { true }
                }
            }

        }
        var newUsername = username.value
        Log.d("changeUsername", newUsername)
        val response: HttpResponse = newClient.post("http://$IPADDRESS:9090/changeUsername") {
            contentType(ContentType.Application.Json)
            setBody("{\"oldUsername\":\"${user.userName}\", \"newUsername\":\"${newUsername}\"}")
        }
        Log.d("userinfo", response.status.value.toString())
        if (response.status.value == 200) {
            user.userName = newUsername
            repository.returnUpdate(user)
            newClient.close()
            return true
        } else {
            return false
        }
        return false
    }

    suspend private fun changePasswordRequest(): Boolean {

        var user = repository.getUser()

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
                        BearerTokens(user.jwtToken, user.jwtToken)
                    }
                    sendWithoutRequest { true }
                }
            }

        }
        var password = newPassword.value
        val response: HttpResponse = newClient.post("http://$IPADDRESS:9090/changePassword") {
            contentType(ContentType.Application.Json)
            setBody("{\"oldPassword\":\"${oldPassword.value}\", \"newPassword\":\"${newPassword.value}\", \"confirmNewPassword\":\"${confirmNewPassword.value}\"}")
        }
        Log.d("userinfo", response.status.value.toString())
        if (response.status.value == 200) {
            user.password = password
            repository.returnUpdate(user)
            newClient.close()
            return true
        } else {
            return false
        }
        return false
    }

    fun changeUsername(): Boolean {
        var worked = false
        runBlocking { worked = changeUsernameRequest() }
        username.value = ""
        Log.d("test", worked.toString())
        return worked
    }
    fun changePassword(): Boolean {
        var worked = false
        runBlocking { worked = changePasswordRequest() }
        newPassword.value = ""
        oldPassword.value = ""
        confirmNewPassword.value = ""
        Log.d("test", worked.toString())
        return worked
    }

    fun deleteAccount(){
        runBlocking { requestDeleteAccount() }
        runBlocking {
            repository.returnDelete() }
    }

    suspend private fun requestDeleteAccount(){
        val user = repository.getUser()
        try {
            user.userName
        }catch (e : Exception){
            return
        }
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
                        BearerTokens(user.jwtToken, user.jwtToken)
                    }
                    sendWithoutRequest { true }
                }
            }

        }
        var password = newPassword.value
        val response: HttpResponse = newClient.post("http://$IPADDRESS:9090/deleteAccount") {
            contentType(ContentType.Application.Json)
            }

    }
}