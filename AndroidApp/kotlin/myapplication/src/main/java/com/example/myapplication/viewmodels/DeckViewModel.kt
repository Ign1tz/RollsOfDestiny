package com.example.myapplication.viewmodels

import android.util.Log
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import com.example.myapplication.localdb.Repository
import com.example.myapplication.localdb.User
import com.example.myapplication.types.deck
import com.example.myapplication.types.decks
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

class DeckViewModel(val repository: Repository, val IPADDRESS: String) : ViewModel(),
    BasicViewModel {

    val decks = mutableStateOf<List<deck>?>(null)
    val newDeckName = mutableStateOf("")

    fun getDecks() {
        runBlocking { requestGetDecks() }
    }

    suspend private fun requestGetDecks() {
        val newUser = repository.getUser()
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
                        BearerTokens(newUser.jwtToken, newUser.jwtToken)
                    }
                    sendWithoutRequest { true }
                }
            }

        }
        val response: HttpResponse = newClient.get("http://$IPADDRESS:9090/getDecks") {
        }
        Log.d("userinfo", response.status.value.toString())
        if (response.status.value == 200) {
            val userInfo: decks = Json.decodeFromString(response.body())
            decks.value = userInfo.decks
            newClient.close()
        }
    }

    fun setActive(deck: deck) {
        runBlocking {
            requestSetActive(deck)
            getDecks()
        }
    }

    suspend private fun requestSetActive(deck: deck) {

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
        val response: HttpResponse = newClient.post("http://$IPADDRESS:9090/setActiveDeck") {
            contentType(ContentType.Application.Json)
            setBody("{\"name\":\"${deck.name}\", \"deckid\":\"${deck.deckid}\"}")
        }
    }

    fun removeDeck(deck: deck) {
        runBlocking {
            requestRemoveDeck(deck)
            getDecks()
        }
    }

    suspend private fun requestRemoveDeck(deck: deck) {

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
        val response: HttpResponse = newClient.post("http://$IPADDRESS:9090/removeDeck") {
            contentType(ContentType.Application.Json)
            setBody("{\"name\":\"${deck.name}\", \"deckid\":\"${deck.deckid}\"}")
        }
    }

    fun createNewDeck() {
        runBlocking {
            requestCreateNewDeck()
            getDecks()
            newDeckName.value = ""
        }
    }

    suspend fun requestCreateNewDeck() {

        if (newDeckName.value.replace(" ", "") != "") {
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
            val response: HttpResponse = newClient.post("http://$IPADDRESS:9090/createDeck") {
                contentType(ContentType.Application.Json)
                setBody("{\"name\":\"${newDeckName.value}\"}")
            }
        }

    }
}



