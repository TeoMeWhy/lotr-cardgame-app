import pandas as pd
import datetime

import streamlit as st

from clients import api
from campaign import show_campaign
from cards import show_card
from cenario import show_cenario
from collection import show_collection
from decks import show_deck
from players import show_player

st.set_page_config(page_title="Senhor dos Anéis")

def main():
    st.markdown("""
                
    # Senhor dos Anéis
    #### Ficha de Campanha

    Por aqui, você poderá criar suas campanhas, players, e gerenciar seus avanços. A ideia é facilitar o acompanhamento do progresso de campanha.            
                
    """)


    players = api.get_players()
    decks = api.get_decks()
    cards = api.get_cards()
    cenarios = api.get_cenarios()
    collections = api.get_collections()


    tab_campaign, tab_players, tab_decks, tab_cards, tab_cenarios, tab_collections = st.tabs(["Campanhas", "Players", "Decks", "Cartas", "Cenários", "Coleções"])

    with tab_campaign:
        
        error = False
        if len(players) == 0:
            error = True
            st.error("Não há jogadores cadastrados")

        if len(decks) == 0:
            error = True
            st.error("Não há baralhos cadastrados")

        if len(cards) == 0:
            error = True
            st.error("Não há cartas cadastradas")

        if len(cenarios) == 0:
            error = True
            st.error("Não há cenarios cadastrados")

        if len(collections) == 0:
            error = True
            st.error("Nao há coleções cadastradas")

        if not error:
            show_campaign()

    with tab_players:
        show_player()

    with tab_decks:
        show_deck()

    with tab_cards:
        show_card()

    with tab_cenarios:
        if len(collections) == 0:
            st.error("Nao há coleções cadastradas")
        else:
            show_cenario()

    with tab_collections:
        show_collection()

if __name__ == "__main__":
    main()