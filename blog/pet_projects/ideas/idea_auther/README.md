# Description

Idea for pet project

- implement authentificating application for central management of users for access to admin/moderator interfaces
- SimpleBackendHTMLApp is a simple backend application in Django with its Django Admin interface
- OAuthAgent1 is a golang redirecter, and verifier of authentifcation.
- Microservice1, and admin interface that needs to be protected

# Diagrams

```mermaid
flowchart TD
    Microservice1 --> OAuthAgent1
    Microservice2 --> OAuthAgent2
    Microservice3 --> OAuthAgent3

    SimpleBackendHTMLApp --> UserDatabase

    OAuthAgent1 --> SimpleBackendHTMLApp
    OAuthAgent2 --> SimpleBackendHTMLApp
    OAuthAgent3 --> SimpleBackendHTMLApp
```

```mermaid
sequenceDiagram
    User->>OAuthAgent1: I want to be authentificated

    OAuthAgent1->>SimpleBackendHTMLApp: Get user for authentification and return to me back
    SimpleBackendHTMLApp ->> UserDatabase: Is user existing like that?
    UserDatabase ->> SimpleBackendHTMLApp: Yup ot does / No it does not

    SimpleBackendHTMLApp ->> OAuthAgent1: Sending user back to you with cookie of his rights

    OAuthAgent1 ->> OAuthAgent1: Reading from cookie permissions

    OAuthAgent1 -->> Microservice1: Sending user to Admin Interface / or rejecting
```