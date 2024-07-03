package com.example.myapplication.localdb

import kotlinx.coroutines.flow.Flow

class Repository (private val dao: dao) {
    companion object {
        @Volatile
        private var instance:Repository?=null

        fun returnInstance (dao: dao) = instance ?: synchronized(this) {
            instance ?: Repository(dao).also { instance=it }
        }
    }

    fun returnInsert (user:User) = dao.Insert(user)
    fun returnUpdate (user:User) = dao.Update(user)
    fun returnDelete () = dao.Delete()
    fun getUser (): User = dao.getUser()
    fun isEmpty(): Boolean = dao.isEmpty()
}