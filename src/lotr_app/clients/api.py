# %%

import requests

def get_campaigns(**kwargs):
    resp = requests.get("http://lotr_api:3000/campaigns", params=kwargs)
    return resp.json()


def create_campaign(**kwargs):
    resp = requests.post("http://lotr_api:3000/campaigns", json=kwargs)
    return resp.json()


def put_campaign(**kwargs):
    resp = requests.put("http://lotr_api:3000/campaigns", json=kwargs)
    return resp.json()


def delete_campaign(**kwargs):
    resp = requests.delete("http://lotr_api:3000/campaigns", json=kwargs)
    return resp.json()


def get_cenario_players(**kwargs):
    resp = requests.get("http://lotr_api:3000/cenario_players", params=kwargs)
    return resp.json()


def create_player_to_play(**kwargs):
    resp = requests.post("http://lotr_api:3000/player_to_play", json=kwargs)
    return resp.json()


def put_player_to_play(**kwargs):
    resp = requests.put("http://lotr_api:3000/player_to_play", json=kwargs)
    return resp.json()


def get_players(**kwargs):
    resp = requests.get("http://lotr_api:3000/players", params=kwargs)
    return resp.json()


def create_player(**kwargs):
    resp = requests.post("http://lotr_api:3000/players", json=kwargs)
    return resp.json()


def update_player(**kwargs):
    resp = requests.put("http://lotr_api:3000/players", json=kwargs)
    return resp.json()


def delete_player(**kwargs):
    resp = requests.delete("http://lotr_api:3000/players", json=kwargs)
    return resp.json()


def get_decks(**kwargs):
    resp = requests.get("http://lotr_api:3000/decks", params=kwargs)
    return resp.json()


def create_deck(**kwargs):
    resp = requests.post("http://lotr_api:3000/decks", json=kwargs)
    return resp.json()


def update_deck(**kwargs):
    resp = requests.put("http://lotr_api:3000/decks", json=kwargs)
    return resp.json()


def delete_deck(**kwargs):
    resp = requests.delete("http://lotr_api:3000/decks", json=kwargs)
    return resp.json()


def get_cards(**kwargs):
    resp = requests.get("http://lotr_api:3000/cards", params=kwargs)
    return resp.json()


def create_card(**kwargs):
    resp = requests.post("http://lotr_api:3000/cards", json=kwargs)
    return resp.json()


def put_card(**kargs):
    resp = requests.put("http://lotr_api:3000/cards", json=kargs)
    return resp.json()


def delete_card(**kwargs):
    resp = requests.delete(f"http://lotr_api:3000/cards", json=kwargs)
    return resp.json()


def get_cenarios(**kwargs):
    resp = requests.get("http://lotr_api:3000/cenarios", params=kwargs)
    return resp.json()


def create_cenario(**kwargs):
    resp = requests.post("http://lotr_api:3000/cenarios", json=kwargs)
    return resp.json()


def put_cenarios(**kwargs):
    resp = requests.put("http://lotr_api:3000/cenarios", json=kwargs)
    return resp.json()

def delete_cenario(**kwargs):
    resp = requests.delete("http://lotr_api:3000/cenarios", json=kwargs)
    return resp.json()


def create_cenario_player(**kwargs):
    resp = requests.post("http://lotr_api:3000/cenario_players", json=kwargs)
    return resp.json()


def put_cenario_player(**kwargs):
    resp = requests.put("http://lotr_api:3000/cenario_players", json=kwargs)
    return resp.json()


def get_collections(**kwargs):
    resp = requests.get("http://lotr_api:3000/collections", params=kwargs)
    return resp.json()


def create_collection(**kwargs):
    resp = requests.post("http://lotr_api:3000/collections", json=kwargs)
    return resp.json()


def put_collection(**kwargs):
    resp = requests.put("http://lotr_api:3000/collections", json=kwargs)
    return resp.json()


def create_cenario_campaign(**kwargs):
    resp = requests.post("http://lotr_api:3000/cenario_campaigns", json=kwargs)
    return resp.json()


def put_cenario_campaign(**kwargs):
    resp = requests.put("http://lotr_api:3000/cenario_campaigns", json=kwargs)
    return resp.json()


def delete_cenario_campaign(**kwargs):
    resp = requests.delete("http://lotr_api:3000/cenario_campaigns", json=kwargs)
    return resp.json()