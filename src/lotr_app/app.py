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

st.markdown("""
            
# Senhor dos Anéis
#### Ficha de Campanha

Por aqui, você poderá criar suas campanhas, players, e gerenciar seus avanços. A ideia é facilitar o acompanhamento do progresso de campanha.            
            
""")

campaign, players, decks, cards, cenarios, collections = st.tabs(["Campanhas", "Players", "Decks", "Cartas", "Cenários", "Coleções"])

with campaign:
    show_campaign()

with players:
    show_player()

with decks:
    show_deck()

with cards:
    show_card()

with cenarios:
    show_cenario()

with collections:
    show_collection()