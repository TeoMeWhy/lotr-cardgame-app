import time

import pandas as pd

import streamlit as st

from clients import api

def show_card():
    st.markdown("## Cartas")

    collections = api.get_collections()
    collection_names = [i["name"] for i in collections]

    if len(collections) == 0:
        st.warning("Ainda não há coleções criadas. Crie uma coleção para poder criar cartas.")
        return

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
    columns = ['collection', "number", "name","type", "description"]
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
        "type": st.column_config.TextColumn(
            "Tipo",
            help="Tipo da Carta",
        ),
        "description": st.column_config.TextColumn(
            "Descrição",
            help="Descrição da Carta",
            width="large",
        )
    }

    st.dataframe(df, hide_index=True, column_config=column_config)


def create_card(collections):
    
    data = {
        "collection": {},
        "collection_id": "",
        "number": 1,
        "name": "",
        "description": "",
        "type": "",
        "cost": 0,
        "willpower": 0,
        "attack": 0,
        "defense": 0,
        "sphere_type": "",
        "hit_points": 0,
    }

    data = collect_card_info(data, collections)

    if st.button("Criar Carta", key="create_card"):
        resp = api.create_card(**data)

        if "error" in resp:
            st.error(resp["error"])
        else:
            st.success("Carta criada com sucesso!")
            time.sleep(1)
            st.rerun()


def edit_card(collections):

    cards = api.get_cards()

    card = st.selectbox("Selecione a carta para editar",
                        options=cards,
                        format_func=lambda x: f'{x["name"]} ({x["collection"]["name"]} - {x["number"]})',
                        key="edit_card_selectbox"
                        )

    card.update(collect_card_info(card, collections, mode='edit'))

    col_edit, _, col_excl = st.columns([1, 2, 1])

    if col_edit.button("Salvar Carta"):

        st.write(card)

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


def collect_card_info(data, collections, mode='create'):

    c1, c2 = st.columns(2)
    card_collection_select_box = c1.selectbox("Coleção",
                                              options=collections,
                                              format_func=lambda x: x["name"],
                                              key=f"create_card_collection_{mode}",
                                              )
    
    card_number_select_box = c2.number_input("Número", min_value=1, max_value=500, value=data["number"],key=f"create_card_number_{mode}")

    card_name_select_box = st.text_input("Nome", value=data["name"], key=f"create_card_name_{mode}")
    card_description_select_box = st.text_area("Descrição", value=data["description"], key=f"create_card_description_{mode}")

    type_cards_list = ["Herói", "Aliado", "Complemento", "Evento"]
    type_cards_list.sort()
    card_type_select_box = st.selectbox("Tipo", options=type_cards_list, key=f"create_card_type_{mode}", index=type_cards_list.index(data["type"]) if data['type']!="" else 0)

    c1, c2 = st.columns(2)
    card_cost = c1.number_input("Custo", min_value=0, max_value=20, value=data["cost"], key=f"create_card_cost_{mode}")
    sphere_types = ['Liderança', 'Conhecimento', 'Espírito', 'Tática']
    sphere_types.sort()
    card_sphereType = c2.selectbox("Tipo de Esfera", options=sphere_types ,key=f"create_card_sphereType_{mode}", index=sphere_types.index(data["sphere_type"]) if data['sphere_type']!="" else 0)


    c1,c2,c3 = st.columns(3)
    card_willpower = c1.number_input("Força de Vontade", min_value=0, max_value=20, value=data["willpower"],key=f"create_card_willpower_{mode}")
    card_attack = c2.number_input("Força de Ataque", min_value=0, max_value=20, value=data["attack"],key=f"create_card_attack_{mode}")
    card_defense = c3.number_input("Força de Defesa", min_value=0, max_value=20, value=data["defense"],key=f"create_card_defense_{mode}")

    card_hitPoints = st.number_input("Pontos de Vida", min_value=0, max_value=20, value=data["hit_points"],key=f"create_card_hitPoints_{mode}")

    data = {
        "collection": card_collection_select_box,
        "collection_id": card_collection_select_box["id"],
        "number": card_number_select_box,
        "name": card_name_select_box,
        "description": card_description_select_box,
        "type": card_type_select_box,
        "cost": card_cost,
        "willpower": card_willpower,
        "attack": card_attack,
        "defense": card_defense,
        "sphere_type": card_sphereType,
        "hit_points": card_hitPoints,
    }

    return data