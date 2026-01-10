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
        if st.session_state["player"]["is_admin"]:
            collections.append({"name":"Criar Nova"})
        
        collection_selected = st.selectbox("Selecione uma coleção",
                                            options=collections,
                                            format_func=lambda x: x["name"])

        st.markdown("---")

        if not st.session_state["player"]["is_admin"]:
            show_collection_to_user(collection_selected)
            return
        
        if collection_selected["name"] == "Criar Nova":
            data = {"name": "", "description": ""}
            collection_inputs(data, mode='create')

        else:
            collection_inputs(collection_selected, mode='edit')


def show_collection_to_user(collection):
    st.markdown(f"### {collection['name']}")
    st.markdown(f"{collection['description']}")

def collection_inputs(collection, mode='create'):
    bt_name = st.text_input("Nome da coleção",
                            value=collection["name"],
                            key=f"collection_name_{mode}",
                            )
    
    bt_description = st.text_area("Descrição",
                                   value=collection["description"],
                                   key=f"collection_description_{mode}",
                                   )

    if st.button("Salvar Coleção", key=f'collection_save_bt_{mode}'):
        data = {
            "name": bt_name,
            "description": bt_description,
        }

        if mode == "create":
            resp = api.create_collection(**data)
        
            if "error" in resp:
                st.error(resp["error"])
            
            else:
                st.success("Coleção criada com sucesso!")
                time.sleep(1)
                st.rerun()
        
        elif mode == "edit":
            data["id"] = collection["id"]
            resp = api.put_collection(**data)
        
            if "error" in resp:
                st.error(resp["error"])
            
            else:
                st.success("Coleção atualizada com sucesso!")
                time.sleep(1)
                st.rerun()