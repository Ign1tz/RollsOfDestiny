package com.example.myapplication.viewmodels

import android.app.Application
import android.media.MediaPlayer
import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.myapplication.R
import com.example.myapplication.connection.websocket.WebSocketClient
import com.example.myapplication.localdb.Repository
import com.example.myapplication.localdb.User
import com.example.myapplication.types.EndResults
import com.example.myapplication.types.gameMessageBody
import com.example.myapplication.types.idMessageBody
import com.example.myapplication.types.message
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.client.statement.HttpResponse
import io.ktor.http.ContentType
import io.ktor.http.contentType
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.jsonObject
import java.net.URI

class GameViewModel(val repository: Repository, val IPADDRESS: String) : ViewModel(),
    BasicViewModel {

    var connected = mutableStateOf(false)
    var WebSocketClient: WebSocketClient? = null
    var WebsocketId = mutableStateOf("")
    var gameInfo: gameMessageBody? by mutableStateOf(null)
    var endResults: EndResults? by mutableStateOf(null)
    var isActive = mutableStateOf(false)
    var hasRolled = mutableStateOf(false)
    var pickedColumn = mutableStateOf(false)
    var roll = mutableStateOf(false)
    var started = mutableStateOf(false)
    var GameType = mutableStateOf("")
    var FriendId = mutableStateOf("")

    fun getUser(): User? {
        val user = repository.getUser()
        try {
            user.userName
            return user
        } catch (e: Exception) {
            return null
        }
    }

    fun resetAllValues() {
        connected.value = false
        WebSocketClient = null
        WebsocketId.value = ""
        gameInfo = null
        endResults = null
        isActive.value = false
        hasRolled.value = false
        pickedColumn.value = false
        roll.value = false
        started.value = false
        FriendId.value = ""
    }


    fun websocket(): WebSocketClient? {
       // resetAllValues()
        if (connected.value && started.value) {
            return null
        }
        Log.d("bot?2", GameType.value)
        Log.d("websocket", started.value.toString())
        started.value = true
        Log.d("bot?3", GameType.value)
        val serverUri = URI("http://$IPADDRESS:8080/ws")
        Log.d("bot?4", GameType.value)
        val gameType = GameType.value
        val webSocketClient = WebSocketClient(serverUri) { message ->
            // display incoming message in ListView
            Log.d("websocket", message)
            val jsonObject = Json.parseToJsonElement(message).jsonObject
            val msg = message(
                info = jsonObject["info"]!!.toString().replace("\"", ""),
                message = jsonObject["message"]!!.toString()
            )
            if (msg.info == "connected") {
                connected.value = true
            } else if (msg.info == "id") {
                val idBody: idMessageBody = Json.decodeFromString(msg.message)
                WebsocketId.value = idBody.id
                runBlocking {
                    Log.d("GameType", gameType)
                    if (gameType == "bot") {
                        botRequest()
                    } else if (gameType == "Friend") {
                        queueForGameWithFriendsRequest()
                    } else {
                        queueRequest()
                    }
                }
            } else if (msg.info == "gameInfo") {
                val gameInfoJson = Json.parseToJsonElement(msg.message).jsonObject
                gameInfo =
                    Json.decodeFromString<gameMessageBody>(gameInfoJson["gameInfo"]!!.toString())
                isActive.value = gameInfo?.ActivePlayer?.active ?: false
                if (isActive.value) {
                    hasRolled.value = false
                    pickedColumn.value = false
                }
            } else if (msg.info == "gameEnded") {

                val gameInfoJson = Json.parseToJsonElement(msg.message).jsonObject
                gameInfo =
                    Json.decodeFromString<gameMessageBody>(gameInfoJson["gameInfo"]!!.toString())
                endResults =
                    Json.decodeFromString<EndResults>(gameInfoJson["endResults"]!!.toString())
            }

        }
        // connect to websocket server
        if (!connected.value && started.value) {
            webSocketClient.connectBlocking()
        }
        webSocketClient.sendMessage("{\"type\": \"id\", \"message\": \"\", \"gameId\": \"\", \"userid\": \"\"}")
        WebSocketClient = webSocketClient
        return webSocketClient
    }

    suspend private fun queueRequest(): Boolean {
        Log.d("websocket", "ququeing")

        val userInfo = repository.getUser()

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
            val responseText: HttpResponse = client.post("http://$IPADDRESS:8080/queue") {
                contentType(ContentType.Application.Json)
                setBody("{\"userid\":\"${userInfo.userid}\", \"websocketconnectionid\":\"${WebsocketId.value}\", \"username\":\"${userInfo.userName}\"}")
            }

            if (responseText.status.value != 200) {
                return false
            }

            client.close()
            return true
        } catch (e: Exception) {
            Log.d("HttpTest", "Received error: ${e.message}")
        }
        return false
    }

    fun remove(){
        viewModelScope.launch { removeQueueRequest() }
    }

    suspend private fun removeQueueRequest(): Boolean {
        Log.d("websocket", "ququeing")

        val userInfo = repository.getUser()

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
            val responseText: HttpResponse = client.post("http://$IPADDRESS:8080/deleteQueue") {
                contentType(ContentType.Application.Json)
                setBody("{\"userid\":\"${userInfo.userid}\", \"websocketconnectionid\":\"${WebsocketId.value}\", \"username\":\"${userInfo.userName}\"}")
            }

            if (responseText.status.value != 200) {
                return false
            }

            client.close()
            return true
        } catch (e: Exception) {
            Log.d("HttpTest", "Received error: ${e.message}")
        }
        return false
    }


    suspend private fun queueForGameWithFriendsRequest(): Boolean {
        Log.d("friendQueue", "Friend ququeing")

        val userInfo = repository.getUser()
        Log.d("websocket", userInfo.userName)

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
            val responseText: HttpResponse =
                client.post("http://$IPADDRESS:8080/queueFroGameWithFriend") {
                    contentType(ContentType.Application.Json)
                    setBody("{\"userid\":\"${userInfo.userid}\", \"websocketconnectionid\":\"${WebsocketId.value}\", \"username\":\"${userInfo.userName}\", \"FriendId\":\"${FriendId.value.replace(" ","")}\"}")
                }

            if (responseText.status.value != 200) {
                return false
            }

            client.close()
            return true
        } catch (e: Exception) {
            Log.d("HttpTest", "Received error: ${e.message}")
        }
        return false
    }

    suspend private fun botRequest(): Boolean {
        Log.d("websocket", "botting")

        val userInfo = repository.getUser()

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
            val responseText: HttpResponse = client.post("http://$IPADDRESS:8080/startBot") {
                contentType(ContentType.Application.Json)
                setBody("{\"userid\":\"${userInfo.userid}\", \"websocketconnectionid\":\"${WebsocketId.value}\"}")
            }

            if (responseText.status.value != 200) {
                return false
            }

            client.close()
            return true
        } catch (e: Exception) {
            Log.d("HttpTest", "Received error: ${e.message}")
        }
        return false
    }

    fun delayRoll() {
        viewModelScope.launch {
            delay(3000)
            hasRolled.value = true
        }
    }

    fun getCardList(): List<Card> {
        val cards = listOf("Destroy Column", "Double Mana", "Roll Again", "Flip Clockwise")
        return cards.map { name -> Card(name, getCardImageById(name)) }
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


}

data class Card(val cardName: String, val cardImageId: Int)
