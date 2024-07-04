package com.example.myapplication.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import com.example.myapplication.R
import com.example.myapplication.localdb.Repository
import com.example.myapplication.types.friendsWraper
import com.example.myapplication.types.scoreboardPlayer
import com.example.myapplication.types.topTenPlayers
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.auth.Auth
import io.ktor.client.plugins.auth.providers.BearerTokens
import io.ktor.client.plugins.auth.providers.bearer
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.get
import io.ktor.client.statement.HttpResponse
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.runBlocking
import kotlinx.serialization.json.Json


class ScoreboardViewModel(val repository: Repository, val IPADDRESS: String) : ViewModel(), BasicViewModel {

    val players = mutableStateOf(listOf<scoreboardPlayer>())
    var user by mutableStateOf(repository.getUser())

    fun getPlayers() {



        var createdPlayers: List<scoreboardPlayer> = listOf()
        runBlocking { createdPlayers = requestGetPlayers() }


        players.value = createdPlayers
    }

    suspend private fun requestGetPlayers(): List<scoreboardPlayer>{
        var createdPlayers: List<scoreboardPlayer> = listOf()
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
        val response: HttpResponse = newClient.get("http://$IPADDRESS:9090/getTopTen") {
        }
        Log.d("userinfo", response.status.value.toString())
        if (response.status.value == 200) {
            val userInfo: topTenPlayers = Json.decodeFromString(response.body())
            createdPlayers = userInfo.topTenPlayers
            newClient.close()
        }
        return createdPlayers
    }

}