import time
import pandas as pd

import streamlit as st

from clients import api

def show_deck():

    st.markdown("## Decks")
    
    lista_decks, details, create = st.tabs(["Lista de Decks", "Detalhes", "Criar"])

    decks = api.get_decks()
    cards = api.get_cards()

    if len(cards) == 0:
        st.error("Ainda não há cartas cadastradas")
        return

    with lista_decks:
        if len(decks) == 0:
            st.warning("Ainda não há decks criados")

        else:        
            df = pd.DataFrame(decks)
            df["owner"] = df["owner"].apply(lambda owner: owner['nick'])
            df = df[['name', 'description', "owner"]].rename(columns={"name": "Nome", "description": "Descrição", "owner": "Criador"})
            st.dataframe(df, hide_index=True)

    with details:

        if len(decks) == 0:
            st.warning("Ainda não há decks criados")

        else:
            deck_selected = st.selectbox("Selecione o Deck",
                                         options=decks,
                                         format_func=lambda deck: f"{deck['name']} - {deck['owner'].get('nick', 'Desconhecido')}",
                                         key="deck_selected_edit")

            st.markdown("---")

            edit_togle = st.toggle("Habilitar Edição", disabled=deck_selected["owner_id"]!=st.session_state['player']['id'])

            deck_name = st.text_input("Nome", value=deck_selected["name"], key="deck_name_edit", disabled=not edit_togle)
            deck_description = st.text_area("Descrição", value=deck_selected["description"], key="deck_description_edit", disabled=not edit_togle)
            cards_selected = st.multiselect("Selecione as cartas",
                                            default=deck_selected["cards"],
                                            options=cards,
                                            format_func=lambda card: f"{card['name']} ({card['collection']['name']})",
                                            key="cards_selected_deck_edit",
                                            disabled=not edit_togle,
                                            )

            if edit_togle:
                col1, _, col2 = st.columns([1, 2, 1])
                if col1.button("Salvar Deck"):
                    resp = api.update_deck(id=deck_selected['id'], owner_id=deck_selected['owner_id'], name=deck_name, description=deck_description, cards=cards_selected)
                    if "error" in resp:
                        st.error(resp["error"])
                    else:
                        st.success("Deck atualizado com sucesso!")
                    time.sleep(2)
                    st.rerun()

                if col2.button("Excluir Deck"):
                    resp = api.delete_deck(id=deck_selected['id'])
                    if 'error' in resp:
                        st.error(resp["error"])
                    else:
                        st.success("Deck excluído com sucesso!")
                        time.sleep(1)
                        st.rerun()

    with create:
        deck_name = st.text_input("Nome", key="deck_name")
        deck_description = st.text_area("Descrição", key="deck_description")
        cards_selected = st.multiselect("Selecione as cartas",
                                        options=cards,
                                        format_func=lambda card: f"{card['name']} ({card['collection']['name']})",
                                        key="cards_selected_deck")

        if st.button("Criar Deck"):
            api.create_deck(owner_id=st.session_state['player']['id'], name=deck_name, description=deck_description, cards=cards_selected)
            st.success("Deck criado com sucesso!")
            time.sleep(1)
            st.rerun()