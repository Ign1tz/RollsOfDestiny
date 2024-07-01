package com.example.myapplication.viewmodels

import android.util.Log
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.myapplication.connection.websocket.WebSocketClient
import com.example.myapplication.localdb.Repository
import com.example.myapplication.types.EndResults
import com.example.myapplication.types.EndResultsBody
import com.example.myapplication.types.gameMessageBody
import com.example.myapplication.types.idMessageBody
import com.example.myapplication.types.message
import io.ktor.client.call.body
import kotlinx.coroutines.launch
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.json.Json
import java.net.URI

class GameViewModel(val repository: Repository) : ViewModel(), BasicViewModel {

    var connected = false
    var WebSocketClient: WebSocketClient? = null
    var WebsocketId  = ""
    var gameInfo: gameMessageBody? = null
    var endResults: EndResults? = null
    var isActive = false
    var hasRolled = false
    var pickedColumn = false

    private val IPADDRESS = "10.0.0.2"

    val board = mutableStateListOf(
        mutableStateListOf(-1, -1, -1),
        mutableStateListOf(-1, -1, -1),
        mutableStateListOf(-1, -1, -1)
    )

    fun placeDie(column: Int) {
        for (row in 2 downTo 0) {
            if (board[row][column] == -1) {
                board[row][column] = 0
                break
            }
        }
    }

    fun websocket(): WebSocketClient? {
        if (connected){
            return null
        }
        val serverUri = URI("http://10.0.0.2:8080/ws")
        Log.d("websocket", "starting")
        val webSocketClient = WebSocketClient(serverUri) { message ->
            // display incoming message in ListView
            viewModelScope.launch {
                run {

                    Log.d("websocket", message)
                    val msg: message = Json.decodeFromString(message)
                    if (msg.info == "connected"){
                        connected = true
                    }else if (msg.info == "id"){
                        val idBody: idMessageBody = Json.decodeFromString(msg.message)
                        WebsocketId = idBody.id
                    } else if (msg.info == "gameInfo"){
                        gameInfo = Json.decodeFromString(msg.message)
                        isActive = gameInfo!!.ActivePlayer.active
                        if (isActive){
                            hasRolled = false
                            pickedColumn = false
                        }
                    } else if (msg.info == "gameEnded"){
                        val body: EndResultsBody = Json.decodeFromString(msg.message)
                        gameInfo = Json.decodeFromString(body.gameInfo)
                        endResults = Json.decodeFromString(body.endResults)
                    }
                }
            }
        }
        // connect to websocket server
        if(!connected) {
            webSocketClient.connectBlocking()
        }
        webSocketClient.sendMessage("{\"type\": \"id\"}")
        WebSocketClient = webSocketClient
        return webSocketClient
    }
}