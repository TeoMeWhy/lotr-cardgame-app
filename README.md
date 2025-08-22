# Senhor dos Aneis Card Game

Um app para registro de campanha

Todo trabalho foi desenvolvido em live no canal [Téo Me Why](https://twitch.tv/teomewhy).

> [!WARNING]
> Finalidade apenas de aprendizado e uso pessoal.

## Setup

Há dois componentes nessa aplicação

### Backend

Nosso backend foi feito em GoLang, para isso, você precisará tê-lo instalado para compilar o código.

```bash
cd src/lotr_api
go run main.go
```

### Frontend

O frontend foi construido em python com a biblioteca `Streamlit`. Assim, é necessário ter o ambiente Python configurado com as devidas bibliotecas.


```bash
cd src/lotr_app
pip install -r requirements.txt
streamlit run app.py
```

## Schema

O schema de banco de dados foi desenhado a partir das entidades identificadas na documentação do `Livro de Regras Base` para preenchimento da ficha de Registro de Capanha.

<img src="schema.png">