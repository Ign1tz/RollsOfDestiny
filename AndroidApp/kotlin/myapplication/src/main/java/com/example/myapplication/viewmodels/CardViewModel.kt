package com.example.myapplication.viewmodels

import android.util.Log
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import com.example.myapplication.R
import com.example.myapplication.localdb.Repository
import com.example.myapplication.types.deck
import com.example.myapplication.types.deckDetailsCards
import com.example.myapplication.types.singleDeck
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

class CardViewModel(val repository: Repository, val IPADDRESS: String, var deckid: String) :
    ViewModel(), BasicViewModel {

    val deck = mutableStateOf<deck?>(null)
    val cards = mutableStateOf<List<String>?>(null)

    fun getDeck() {
        runBlocking {
            requestGetDeck()
            getYourCards()
        }
    }

    suspend private fun requestGetDeck() {
        Log.d("deckId", deckid)
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
        val response: HttpResponse =
            newClient.get("http://$IPADDRESS:9090/getDeck?deckid=" + deckid) {
            }
        Log.d("userinfo", response.status.value.toString())
        if (response.status.value == 200) {
            val userInfo: singleDeck = Json.decodeFromString(response.body())
            deck.value = userInfo.deck
            newClient.close()
        }
    }

    fun getYourCards() {
        runBlocking { requestGetYourCards() }
    }

    suspend private fun requestGetYourCards() {
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
        val response: HttpResponse = newClient.get("http://$IPADDRESS:9090/getNewCards") {
        }
        Log.d("userinfo", response.status.value.toString())
        if (response.status.value == 200) {
            val userInfo: deckDetailsCards = Json.decodeFromString(response.body())
            cards.value = userInfo.oldCards
            Log.d("deckinfo", cards.value.toString())
            newClient.close()
        }
    }

    fun getCardImageById(cardName: String): Int {
        return when (cardName) {
            "Destroy Column" -> R.drawable.destroy_column_app
            "Double Mana" -> R.drawable.double_mana_app
            "Roll Again" -> R.drawable.roll_again_app
            "Flip Clockwise" -> R.drawable.rotate_grid_app
            else -> R.drawable.double_mana_app
        }
    }

    fun removeFromDeck(cardname: String) {
        runBlocking {
            requestRemoveFromDeck(cardname)
            getDeck()
        }
    }

    suspend fun requestRemoveFromDeck(cardname: String) {

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
        val response: HttpResponse = newClient.post("http://$IPADDRESS:9090/removeCardFromDeck") {
            contentType(ContentType.Application.Json)
            setBody("{\"name\":\"${cardname}\", \"deckid\":\"${deck.value?.deckid}\"}")
        }
    }

    fun addToDeck(cardname: String) {
        runBlocking {
            requestAddToDeck(cardname)
            getDeck()
        }
    }

    suspend fun requestAddToDeck(cardname: String) {

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
        val response: HttpResponse = newClient.post("http://$IPADDRESS:9090/addCardToDeck") {
            contentType(ContentType.Application.Json)
            setBody("{\"name\":\"${cardname}\", \"deckid\":\"${deck.value?.deckid}\"}")
        }
    }
}