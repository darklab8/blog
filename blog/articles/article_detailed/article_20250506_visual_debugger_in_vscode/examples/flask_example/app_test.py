import pytest
from app import app
from flask import Flask
from flask.testing import FlaskClient

@pytest.fixture()
def application() -> Flask:
    app.config.update({
        "TESTING": True,
    })
    yield app


@pytest.fixture()
def client(application: Flask) -> FlaskClient:
    return application.test_client()

def test_request_example(client: FlaskClient) -> None:
    response = client.get("/")
    assert b"<p>Hello, World!</p>" in response.data
