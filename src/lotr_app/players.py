import time

import pandas as pd

import streamlit as st

from clients import api

def show_player():

    st.markdown("## Players")

    players = api.get_players()

    lista_jogadores, create, edit = st.tabs(["Lista de Jogadores", "Criar Jogador", "Editar Jogador"])

    with lista_jogadores:
        if len(players) == 0:
            st.warning("Ainda não há jogadores criados")

        else:
            df = pd.DataFrame(players)
            st.dataframe(df, hide_index=True)

    with create:
        player_name = st.text_input("Nome", key="player_name")
        
        if st.button("Criar Jogador"):
            resp = api.create_player(name=player_name)
            
            if "error" in resp:
                st.error(f"Erro ao criar jogador: {resp['error']}")
            
            else:
                st.success("Jogador criado com sucesso!")
                time.sleep(2)
                st.rerun()

    with edit:

        player_selected = st.selectbox("Selecione o Jogador",
                                       options=players,
                                       format_func=lambda player: player['name'],
                                       key="player_selected_edit")

        st.markdown("---")

        player_name = st.text_input("Nome", value=player_selected["name"], key="player_name_edit")
        player_selected["name"] = player_name

        col1, _, col2 = st.columns([1, 2, 1])
        if col1.button("Salvar Jogador"):
            resp = api.update_player(**player_selected)
            if "error" in resp:
                st.error(resp["error"])
            else:
                st.success("Jogador atualizado com sucesso!")
            time.sleep(2)
            st.rerun()

        if col2.button("Excluir Jogador"):
            resp = api.delete_player(**player_selected)
            if "error" in resp:
                st.error(resp["error"])
            else:
                st.success("Jogador excluído com sucesso!")
            time.sleep(2)
            st.rerun()