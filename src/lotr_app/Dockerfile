FROM python:3.13

WORKDIR /app/src/lotr_app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

CMD ["streamlit", "run", "app.py", "--server.port", "3001"]