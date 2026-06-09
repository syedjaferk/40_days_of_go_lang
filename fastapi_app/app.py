from fastapi import FastAPI
from pydantic import BaseModel


class Item(BaseModel):
    name: str
    age: int
    phone: int


app = FastAPI()


@app.get("/home")
def home():
    return {"message": "Hello World"}


@app.get("/item/{item_id}")  # Path Paramater
def read_item(item_id: int):
    # DB logic
    return {"message": "Hello World", "item_id": item_id}


@app.get("/search")
def search(q: str = "", name: str = ""):
    return {"query": q, "name": name}


@app.post("/items/")
def create_item(item: Item):
    return {"item": item}
