import time
import streamlit as st

from clients import api

def show_collection():
    st.markdown("## Coleções")

    collections = api.get_collections()

    if len(collections) == 0:
        st.warning("Nenhuma coleção encontrada. Aproveite para criar uma agora!")
        data = {"name": "", "description": ""}
        collection_inputs(data, mode='create')

    else:
        collections.append({"name":"Criar Nova"})
        collection_selected = st.selectbox("Selecione uma coleção",
                                            options=collections,
                                            format_func=lambda x: x["name"])

        st.markdown("---")

        if collection_selected["name"] == "Criar Nova":
            data = {"name": "", "description": ""}
            collection_inputs(data, mode='create')

        else:
            collection_inputs(collection_selected, mode='view')


def collection_inputs(collection, mode='create'):
    bt_name = st.text_input("Nome da coleção",
                            value=collection["name"],
                            key=f"collection_name_{mode}",
                            disabled=mode!='create',
                            )
    
    bt_description = st.text_area("Descrição",
                                   value=collection["description"],
                                   key=f"collection_description_{mode}",
                                   disabled=mode!='create',
                                   )

    if st.button("Salvar Coleção", key=f'collection_save_bt_{mode}', disabled=mode!='create'):
        data = {
            "name": bt_name,
            "description": bt_description,
        }

        resp = api.create_collection(**data)
        
        if "error" in resp:
            st.error(resp["error"])
        
        else:
            st.success("Coleção criada com sucesso!")
            time.sleep(1)
            st.rerun()