import time

import streamlit as st
from clients import api

def show_cenario():
    
    st.markdown("## Cenários")

    collections = api.get_collections()
    if len(collections) == 0:
        st.warning("Nenhuma coleção encontrada para criar cenários")
        return

    cenarios = api.get_cenarios()
    if len(cenarios) > 0:

        if st.session_state["player"]["is_admin"]:
            cenarios.append({"name":"Criar Novo"})
        
        cenario_selected = st.selectbox("Selecione um cenário",
                                          options=cenarios,
                                          format_func=lambda x: x["name"])

        if cenario_selected["name"] == "Criar Novo":
            cenario = {"name":"", "description":"", "order":1, "collection":collections[0]}
            cenario_inputs(cenario, collections, mode='create')
            return

        st.markdown("---")

        if st.session_state["player"]["is_admin"]:
            cenario_inputs(cenario_selected, collections, mode='edit')

        else:
            show_cenario_to_user(cenario_selected)

    elif st.session_state["player"]["is_admin"]:
        st.warning("Nenhum cenário encontrado. Aproveite para criar um novo")
        cenario = {"name":"", "description":"", "order":1, "collection":collections[0]}
        cenario_inputs(cenario, collections, mode='create')


def show_cenario_to_user(cenario):
    st.markdown(f"### {cenario['name']}")
    st.markdown(f"**Coleção:** {cenario['collection']['name']} | **Etapa:** {cenario['order']}")
    st.markdown(f"#### Descrição: {cenario["description"]}")

def cenario_inputs(cenario, collections, mode='create'):

    bt_name = st.text_input("Nome do cenário", value=cenario["name"], key=f"cenario_name_{mode}")
    bt_description = st.text_area("Descrição", value=cenario["description"], key=f"cenario_description_{mode}")
    bt_order = st.number_input("Ordem", min_value=1, max_value=15, value=max([1, int(cenario["order"])]), key=f"cenario_order_{mode}")
    bt_collection = st.selectbox("Coleção", options=collections, format_func=lambda x: x["name"], key=f"cenario_collection_{mode}")

    col1, _, col2 = st.columns([1, 2, 1])
    if col1.button("Salvar Cenário", key=f'cenario_save_bt_{mode}'):

        data = {            
            "name": bt_name,
            "description": bt_description,
            "order": bt_order,
            "collection": bt_collection,
            "collection_id": bt_collection["id"],
        }

        if mode == 'create':
            resp = api.create_cenario(**data)
        else:
            data['id'] = cenario["id"]
            resp = api.put_cenarios(**data)
        
        if "error" in resp:
            st.error(resp["error"])
        
        else:
            if mode == 'create':
                st.success("Cenário criado com sucesso!")
            else:
                st.success("Cenário atualizado com sucesso!")
            
            time.sleep(1)
            st.rerun()

    if col2.button("Deletar Cenário", key=f'cenario_delete_bt_{mode}'):
        resp = api.delete_cenario(**cenario)
        if "error" in resp:
            st.error(resp["error"])
        else:
            st.success("Cenário deletado com sucesso!")
            time.sleep(1)
            st.rerun()