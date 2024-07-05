package com.example.myapplication.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.material3.DrawerState
import androidx.compose.material3.DrawerValue
import androidx.compose.material3.rememberDrawerState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.example.myapplication.localdb.Repository
import com.example.myapplication.localdb.User
import com.example.myapplication.types.friends
import com.example.myapplication.types.friendsWraper
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

class HomeViewModel (val repository: Repository, val IPADDRESS: String) : ViewModel(), BasicViewModel {

    var friends = mutableStateOf<List<friends>?>(null)
    var addFriend = mutableStateOf("")

    var isFriendPlayClicked by mutableStateOf(false)
        private set
    fun toggleFriendClick() {
        isFriendPlayClicked = !isFriendPlayClicked
        isHostButtonClicked = false
        isJoinButtonClicked = false
    }

    var isHostButtonClicked by mutableStateOf(false)
    fun toggleHostButtonClicked() {
        isHostButtonClicked = !isHostButtonClicked
        isJoinButtonClicked = false
        isFriendPlayClicked = false
    }

    var isJoinButtonClicked by mutableStateOf(false)

    fun toggleJoinButtonClicked() {
        isJoinButtonClicked = !isJoinButtonClicked
        isHostButtonClicked = false
        isFriendPlayClicked = false
    }

    fun getUser(): User? {
        val user = repository.getUser()
        try {
            user.userName
            return user
        } catch (e: Exception) {
            return null
        }
    }

    var drawerState = DrawerValue.Closed

    fun getUsername(): String {
        val user = repository.getUser()
        try {
            return user.userName
        }catch (e : Exception){
            return ""
        }
    }

    fun openDrawer(drawerState: DrawerState): DrawerState{
        runBlocking{drawerState.open()}
        return drawerState
    }


    suspend private fun requestGetFriends() {
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
        val response: HttpResponse = newClient.get("http://$IPADDRESS:9090/getFriends") {
        }
        Log.d("userinfo", response.status.value.toString())
        if (response.status.value == 200) {
            val userInfo: friendsWraper = Json.decodeFromString(response.body())
            friends.value = userInfo.friends

            newClient.close()
        }
    }

    suspend private fun requestRemoveFriend(username: String){
        val user = repository.getUser()
        val client = HttpClient(CIO) {
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

        val responseText: HttpResponse = client.post("http://$IPADDRESS:9090/removeFriend") {
            contentType(ContentType.Application.Json)
            setBody("{\"username\":\"$username\"}")
        }
    }

    suspend private fun requestAddNewFriend(){
        val user = repository.getUser()
        val client = HttpClient(CIO) {
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

        val responseText: HttpResponse = client.post("http://$IPADDRESS:9090/addFriend") {
            contentType(ContentType.Application.Json)
            setBody("{\"username\":\"${addFriend.value}\"}")
        }
    }

    fun getFriends(){
        runBlocking {requestGetFriends() }
    }

    fun removeFriend(username: String){
        runBlocking { requestRemoveFriend(username) }
        runBlocking {requestGetFriends() }
    }

    fun addNewFriend(){

        runBlocking { requestAddNewFriend() }
        runBlocking {requestGetFriends() }
        addFriend.value = ""
    }
}