import time

import streamlit as st
from clients import api


def show_campaign():
    st.markdown("## Campanhas")

    campaigns = api.get_campaigns()

    if len(campaigns) == 0:
        st.warning("Nenhuma campanha encontrada. Aproveite para criar uma nova!")

    campanha = st.selectbox("Selecione uma campanha", campaigns+[{"name":"Criar Nova Campanha"}], format_func=lambda x: x["name"], key="campaign_select")

    if campanha["name"] == "Criar Nova Campanha":
        create_player_to_campaign()

    else:
        campanha["cenario_campaigns"].sort(key=lambda x: x['cenario'].get('order', 0))

        geral, cenarios = st.tabs(["Geral", "Cenários"])
        with geral:
            ficha_base(campaign=campanha)
        with cenarios:
            summary_cenario(campaign=campanha)


def create_player_to_campaign():
    players = api.get_players()
    players.append({"name":"Sem Jogador"})

    heroes = api.get_cards()
    heroes = [i for i in heroes if i["type"]=="Herói"]

    name_campaign = st.text_input("Nome da Campanha", key="name_campaign_create")

    st.markdown("---\n### Jogadores\n")

    players_to_play = [prepare_player(i, players, heroes) for i in range(1, 5)]
    players_to_play = [i for i in players_to_play if i is not None]

    if st.button("Criar Campanha"):

        players_to_campaign = []
        for p in players_to_play:
            resp = api.create_player_to_play(**p)
            if "error" in resp:
                st.error(f"Erro ao criar jogador: {resp['error']}")
                return
            else:
                players_to_campaign.append(resp)


        campaign_data = {
            "name": name_campaign,
            "players": players_to_campaign,
        }
        api.create_campaign(**campaign_data)
        st.success("Campanha criada com sucesso!")
        time.sleep(2)
        st.rerun()

        
def prepare_player(number, players, cards):

    decks = api.get_decks()

    st.markdown(f"##### Jogador {number}")
    
    col1, col2, col3 = st.columns(3)

    player = col1.selectbox("Nome",
                          players,
                          index=len(players)-1,
                          format_func=lambda x: x["name"],
                          key=f"player_{number}_name")
    
    heroes = col2.multiselect("Heróis",
                                cards,
                                format_func=lambda x: f'{x["name"]} ({x["collection"]["name"]})',
                                key=f"player_{number}_hero",
                                max_selections=3)

    deck = col3.selectbox("Deck",
                           decks,
                           format_func=lambda x: x["name"],
                           key=f"player_{number}_deck")
    
    st.markdown("---")

    if player is None or player == {"name": "Sem Jogador"}:
        return None

    player2play = {
        "player_id": player['id'],
        "player": player,
        "heroes": heroes,
        "deck_id": deck['id'],
        "deck": deck,
    }

    return player2play


def delete_campaign(campaign):
    resp = api.delete_campaign(**campaign)
    if "error" in resp:
        st.error(f"Erro ao remover campanha: {resp['error']}")
    else:
        st.success("Campanha removida com sucesso!")
        return True


def ficha_base(campaign):

    st.markdown("#### Registro de Campanha - Ficha Base")

    enable_edit = st.toggle("Habilitar Edição")

    all_heroes = api.get_cards()
    all_heroes = [i for i in all_heroes if i["type"]=="Herói"]

    cols = st.columns(len(campaign["players"]))

    players = []
    for i in range(len(cols)):

        with cols[i]:
            with st.container(border=True):

                player = campaign["players"][i]

                deck = player['deck']

                st.markdown(f"###### Jogador {i+1} - {player['player']['name']}")
                
                st.markdown(f"###### Deck {deck['name']}")
                st.markdown(f"{deck['description']}")

                player['heroes'] = st.multiselect(label="",
                                                  options=all_heroes,
                                                  default=player['heroes'],
                                                  max_selections=3,
                                                  label_visibility='collapsed',
                                                  format_func=lambda x: f'{x["name"]} ({x["collection"]["name"]})',
                                                  disabled=not enable_edit)
                
                players.append(player)
    
    campaign["players"] = players

    st.markdown("---")

    st.markdown("##### Heróis Derrotados")
    
    col1, col2 = st.columns([2,1])

    with col1.container(border=True):
        defeated_heroes = st.multiselect("Marque todos heróis derrotados na campanha",
                                options=all_heroes,
                                default=campaign["defeated_heroes"],
                                format_func=lambda x: f'{x["name"]} ({x["collection"]["name"]})',
                                disabled=not enable_edit)

    with col2.container(border=True):
        st.markdown(f"###### Penalidade de ameaça:\n## {len(defeated_heroes)}")

    st.markdown("---")
    st.markdown("##### Notas")

    notes = st.text_area("Notas da campanha",
                        value=campaign["notes"],
                        disabled=not enable_edit,
                        key="campaign_notes",
                        )

    st.markdown("---")

    if enable_edit:
        campaign["defeated_heroes"] = defeated_heroes
        campaign["notes"] = notes

        col1, _, col2 = st.columns([1,2,1])
        if col1.button("Salvar Campanha"):

            for p in campaign["players"]:
                resp = api.put_player_to_play(**p)
                if 'error' in resp:
                    st.error(f"Erro ao salvar jogador {p['player']['name']}: {resp['error']}")

            resp = api.put_campaign(**campaign)
            if "error" in resp:
                st.error(f"Erro ao salvar notas: {resp['error']}")
            else:
                st.success("Campanha atualizada com sucesso!")
                time.sleep(2)
                st.rerun()

        if col2.button("Remover Campanha"):
            if delete_campaign(campaign):
                st.rerun()


def summary_cenario(campaign):
    st.markdown("#### Registro de Campanha - Cenários Conquistados")
    if len(campaign["cenario_campaigns"])==0:
        st.warning("Não há cenários conquistados por enquanto.")

    else:
        for c in campaign["cenario_campaigns"]:
            name = c['cenario']['name']
            points = c['total_final_points']
            with st.container(border=True):
                col1, col2 = st.columns([9,1])
                col1.markdown(f"###### {name}")
                col2.markdown(f"##### {points}")
                expander_cenario(campaign, c, mode="view")


        st.markdown("---")
        _, col2, col3 = st.columns([6,3,1])
        col2.markdown("###### Total de Pontos")
        col3.markdown(f"##### {campaign['points']}")
        st.markdown("---")

    if st.toggle("Adicionar Cenário"):
        expander_cenario(campaign, {}, mode='create')


def expander_cenario(campaign, cenario_campaign, mode='create'):

    cenarios = api.get_cenarios()
    if mode == 'create':
        cenarios = [c for c in cenarios if c['id'] not in [i['cenario_id'] for i in campaign['cenario_campaigns']]]
        cenario = st.selectbox("Selecione um cenário", cenarios, format_func=lambda x: x["name"])
        cenario_players = [{'player': p['player']} for p in campaign['players']]

        if cenario is None:
            st.warning("Todos cenários disponíveis já foram conquistados")
            return

    else:
        cenario = cenario_campaign['cenario']
        cenario_players = cenario_campaign['cenario_players']
        
    with st.expander(label=f"Detalhes", expanded=mode=='create'):

        edit_toggle = st.toggle("Habilitar Edição", key=f"edit_cenario_{campaign['id']}_{cenario['id']}") if mode != 'create' else True
        col1, col2, col3, col4, col5, col6, col7, col8 = st.columns([2,2,1,2,1,2,1,2])

        col1.markdown("Jogador")
        col2.markdown("Ameaça Final")
        col4.markdown("Herói Derrotados")
        col6.markdown("Dano Heróis")
        col8.markdown("Total")

        for i in range(len(cenario_players)):
            cenario_players[i] = get_cenario_players(campaign, cenario_players[i], cenario, edit=edit_toggle)

        total_jogadores = sum([i['total'] for i in cenario_players])

        st.markdown("---")

        col1, col2, col3, col4, col5, col6, col7 = st.columns([2,1,2,1,2,1,2])
        col1.markdown("Total Jogadores")
        col3.markdown("Pontos de Vitória")
        col5.markdown("Rodadas")
        col7.markdown("Pontuação Final")

        with st.container(border=True):
            c1, c2, c3, c4, c5, c6, c7 = st.columns([2,1,2,1,2,1,2])
            c1.markdown(f"### {total_jogadores}")
            
            c2.markdown("### -")
            victory = c3.number_input("Pontos de Vitória", min_value=0, max_value=100,
                                      step=1, label_visibility='collapsed',
                                      value=cenario_campaign.get('victory_points', 0), key=f"victory_points_{cenario['id']}",
                                      disabled=not edit_toggle,
                                      )

            c4.markdown("### +")
            rounds = c5.number_input("Rodadas", min_value=1, max_value=50,
                                     step=1, label_visibility='collapsed',
                                     value=cenario_campaign.get('rounds', 1), key=f"rounds_{cenario['id']}",
                                     disabled=not edit_toggle)
            
            c6.markdown("### =")
            total_cenario = total_jogadores - victory + (rounds * 10)
            c7.markdown(f"### {total_cenario}")

        cenario_campaign.update({
                    "campaign_id":campaign['id'],
                    "cenario_id":cenario['id'],
                    "cenario_players":cenario_players,
                    "total_points":total_jogadores,
                    "victory_points":victory,
                    "rounds":rounds,
                })

        if mode == 'create':
            if create_cenario(cenario_campaign, cenario_players):
                time.sleep(2)
                st.rerun()

        else:
            if edit_toggle:
                if edit_remove_cenario(campaign, cenario_campaign, cenario_players):
                    time.sleep(2)
                    st.rerun()


def get_cenario_players(campaign, cenario_player, cenario, edit=True):
    with st.container(border=True):

        p = cenario_player
        c1, c2, c3, c4, c5, c6, c7, c8 = st.columns([2,2,1,2,1,2,1,2])
        
        c1.markdown(p['player']['name'])
        threat_player = c2.number_input("", min_value=20, max_value=50,step=1,
                                key=f"final_threat_{cenario['id']}_{p['player']['id']}",
                                label_visibility='collapsed',
                                value=p.get('final_threat', 20),
                                disabled=not edit,
                                )
        
        c3.markdown("### +")
        heroes_defeated = c4.number_input("", min_value=0, max_value=3,step=1,
                                key=f"heroes_defeated_{cenario['id']}_{p['player']['id']}",
                                label_visibility='collapsed',
                                value=p.get('threat_hero_defeated', 0),
                                disabled=not edit,
                                )
        
        c5.markdown("### +")
        heroes_damage = c6.number_input("", min_value=0, max_value=50,step=1,
                                key=f"heroes_damage_{cenario['id']}_{p['player']['id']}",
                                label_visibility='collapsed',
                                value=p.get('heroes_damage', 0),
                                disabled=not edit,
                                )
        
        c7.markdown("### =")
        total = threat_player + heroes_defeated + heroes_damage
        c8.markdown(f"### {total}")

        cenario_player.update({
            "campaign_id":campaign['id'],
            "player_id":p['player']['id'],
            "player":p['player'],
            "cenario_id":cenario['id'],
            "cenario":cenario,
            "final_threat":threat_player,
            "threat_hero_defeated":heroes_defeated,
            "heroes_damage":heroes_damage,
            "total": total
        })

        return cenario_player


def create_cenario(cenario_campaign, cenario_players):
    if st.button("Criar Cenário"):
        for p in cenario_players:
            resp = api.create_cenario_player(**p)
            if "error" in resp:
                st.error(f"Erro ao criar cenário para jogador: {resp['error']}")
                return

        resp = api.create_cenario_campaign(**cenario_campaign)
        if "error" in resp:
            st.error(f"Erro ao criar cenário: {resp['error']}")
        else:
            st.success("Cenário criado com sucesso!")
            return True


def edit_remove_cenario(campaign, cenario_campaign, cenario_players):
    col1, _, col3 = st.columns([1,2,1])
    if col1.button("Salvar Cenário"):

        for p in cenario_players:
            resp = api.put_cenario_player(**p)
            if "error" in resp:
                st.error(f"Erro ao atualizar cenário para jogador: {resp['error']}")
                return

        resp = api.put_cenario_campaign(**cenario_campaign)
        if "error" in resp:
            st.error(f"Erro ao atualizar cenário: {resp['error']}")
        else:
            st.success("Cenário atualizado com sucesso!")
            return True

    if col3.button("Remover Cenário"):
        resp = api.delete_cenario_campaign(**cenario_campaign)
        if "error" in resp:
            st.error(f"Erro ao remover cenário: {resp['error']}")
            return
        else:
            st.success("Cenário removido com sucesso!")
            return True