import time

import pandas as pd

import streamlit as st

from clients import api

def show_card():
    st.markdown("## Cartas")

    collections = api.get_collections()
    collection_names = [i["name"] for i in collections]

    lista_cartas, create, edit = st.tabs(["Lista Completa", "Criar Carta", "Editar Carta"])

    with lista_cartas:
        collections_select_box = st.selectbox("Selecione uma coleção", collection_names+["Todas"])

        collection_selected = [i for i in collections if i["name"] == collections_select_box]

        params = {} if len(collection_selected) == 0 else {"collections_id": collection_selected[0]["id"]}

        cards = api.get_cards(**params)

        if len(cards) == 0:
            st.warning("Nenhuma carta encontrada. Aproveite para criar uma nova carta.")

        else:
            show_dataframe_cards(cards)

    with create:
        create_card(collections)

    with edit:
        edit_card(collections)

def show_dataframe_cards(cards):

    df = pd.DataFrame(cards)
    columns = ['collection', "number", "name", "description"]
    df = df[columns]
    df["collection"] = df["collection"].apply(lambda x: x["name"])
    df.sort_values(by=["collection", "number"], inplace=True)

    column_config = {
        "collection": st.column_config.TextColumn(
            "Coleção",
            help="Nome da coleção",
        ),
        "number": st.column_config.NumberColumn(
            "Número",
            help="Número da Carta na Coleção",
        ),
        "name": st.column_config.TextColumn(
            "Nome",
            help="Nome da Carta",
        ),
        "description": st.column_config.TextColumn(
            "Descrição",
            help="Descrição da Carta",
            width="large",
        )
    }

    st.dataframe(df, hide_index=True, column_config=column_config)


def create_card(collections):
    collection_names = [i["name"] for i in collections]

    card_collection_select_box = st.selectbox("Coleção", options=collection_names, key="create_card_collection")
    card_number_select_box = st.number_input("Número", min_value=1, max_value=500, value=1,)
    card_name_select_box = st.text_input("Nome", value="")
    card_description_select_box = st.text_area("Descrição", value="")

    collection_selected_card = [i for i in collections if i['name'] == card_collection_select_box][0]

    if st.button("Criar Carta", key="create_card"):
        data = {
            "collection": collection_selected_card,
            "collection_id": collection_selected_card["id"],
            "number": card_number_select_box,
            "name": card_name_select_box,
            "description": card_description_select_box,
        }

        resp = api.create_card(**data)

        if "error" in resp:
            st.error(resp["error"])
        else:
            st.success("Carta criada com sucesso!")
            time.sleep(1)
            st.rerun()

def edit_card(collections):
    collection_names = [i["name"] for i in collections]

    col1, col2 = st.columns(2)

    card_collection_select_box = col1.selectbox("Coleção", options=collection_names, key="create_card_collection_edit_card")
    collection_id=[i["id"] for i in collections if i['name'] == card_collection_select_box][0]

    cards_collection = api.get_cards(collection_id=collection_id)

    numbers = [i["number"] for i in cards_collection]
    card_number_select_box = col2.selectbox("Número", options=numbers, key="edit_card_number", disabled=len(numbers) == 0)
    
    if len(numbers) == 0:
        st.warning("Nenhuma carta encontrada nesta coleção.")
        return

    card = api.get_cards(collection_id=collection_id, number=card_number_select_box)[0]
    card_name_select_box = st.text_input("Nome", value=card["name"])
    card_description_select_box = st.text_area("Descrição", value=card["description"])

    card["name"] = card_name_select_box
    card["description"] = card_description_select_box

    col_edit, _, col_excl = st.columns([1, 2, 1])

    if col_edit.button("Salvar Carta"):
        resp = api.put_card(**card)
        if "error" in resp:
            st.error(resp["error"])
        else:
            st.success("Carta editada com sucesso!")
            time.sleep(1)
            st.rerun()

    if col_excl.button("Excluir Carta"):
        resp = api.delete_card(**card)
        if "error" in resp:
            st.error(resp["error"])
        else:
            st.success("Carta excluída com sucesso!")
            time.sleep(1)
            st.rerun()
