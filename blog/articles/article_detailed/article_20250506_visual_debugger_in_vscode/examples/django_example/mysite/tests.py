from django.test import Client

def test_our_view() -> None:
    c = Client()
    response = c.get("/polls/")
    assert response.status_code == 200
