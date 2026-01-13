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

st.set_page_config(page_title="Senhor dos Anéis", page_icon='🌿')

def main():
    st.markdown("""
                
    # Senhor dos Anéis
    #### Progresso da Campanha

    App destinado a entusiastas do jogo "O Senhor dos Anéis: cardgame" para gerenciar campanhas.
                
    No lugar de impreimir suas fichas de campanhas, utilize esse app para acompanhar seu progresso e testes diferentes decks!

    """)

    if not st.user.is_logged_in:
        st.warning("Por favor, faça login para acessar o app.")
        if st.button("Login com Google"):
            st.login()
        return
    
    else:
        cols = st.columns([8, 2])
        cols[0].success(f"Bem-vindo, {st.user.name}!")
        cols[1].button("Logout", on_click=st.logout)


    players = api.get_players()
    player = [i for i in players if i["email"] == st.user.email]
    if len(player) == 0:
        
        st.info("Novo jogador! Criando perfil... só precisamos de seu nick")
        nick_input = st.text_input("Nick", key="nick_input")
        
        if len([i for i in players if i["nick"] == nick_input]) > 0:
            st.error("Nick já está em uso, por favor escolha outro.")
            return
        
        if st.button("Finalizar!"):
            st.session_state["player"] = api.create_player(name=st.user.name, email=st.user.email, nick=nick_input)
            st.rerun()
        
        st.stop()
    else:
        st.session_state["player"] = player[0]

    if st.session_state["player"]["is_admin"]:
        show_admin()
        return

    show_user()    
    

def show_user():
    campaign_tab, decks_tab, cards_tab, cenarios_tab, collections_tab = st.tabs(["Campanha", "Decks", "Cartas", "Cenários", "Coleções"])
    with campaign_tab:
        show_campaign()

    with decks_tab:
        show_deck()

    with cards_tab:
        show_card()

    with cenarios_tab:
        show_cenario()

    with collections_tab:
        show_collection()


def show_admin():

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