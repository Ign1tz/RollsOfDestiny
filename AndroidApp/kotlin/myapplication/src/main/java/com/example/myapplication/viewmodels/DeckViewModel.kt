package com.example.myapplication.viewmodels

import androidx.lifecycle.ViewModel
import com.example.myapplication.localdb.Repository

class DeckViewModel(val repository: Repository) : ViewModel(), BasicViewModel {


    fun getDeckList(): List<Deck> {
        val deck1 = Deck("Turbodeck", true)
        val deck2 = Deck("Gigantodeck", false)
        val deck3 = Deck("Terradeck", false)
        return listOf(deck1, deck2, deck3)
    }

}

data class Deck(val name: String, val isActive: Boolean)