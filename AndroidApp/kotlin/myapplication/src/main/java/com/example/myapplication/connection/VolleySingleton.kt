package com.example.myapplication.connection

import android.content.Context
import com.android.volley.RequestQueue
import com.android.volley.Response
import com.android.volley.toolbox.JsonObjectRequest
import com.android.volley.toolbox.Volley
import org.json.JSONObject

class VolleySingleton private constructor(context: Context) {
    companion object {
        @Volatile
        private var INSTANCE: VolleySingleton? = null
        fun getInstance(context: Context) =
            INSTANCE ?: synchronized(this) {
                INSTANCE ?: VolleySingleton(context).also {
                    INSTANCE = it
                }
            }

    }

    val requestQueue: RequestQueue by lazy {
        Volley.newRequestQueue(context.applicationContext)
    }
}

fun performHttpRequest(
    context: Context,
    url: String,
    method: Int,
    requestBody: JSONObject?,
    onSuccess: (response: JSONObject) -> Unit,
    onError: (error: String) -> Unit
) {
    val requestQueue = VolleySingleton.getInstance(context).requestQueue

    val jsonObjectRequest = JsonObjectRequest(
        method, url, requestBody,
        { response ->
            onSuccess(response)
        },
        { error ->
            onError(error.toString())
        }
    )

    requestQueue.add(jsonObjectRequest)
}