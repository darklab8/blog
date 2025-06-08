FROM python:3.13-bookworm

ENV PYTHONUNBUFFERED 1
ENV PYTHONDONTWRITEBYTECODE 1
WORKDIR /code

# gettext for `python3 manage.py compilemessages`
RUN apt update && apt install -y \ 
    gettext \
    && rm -rf /var/lib/apt/lists/*

COPY ./requirements.txt ./constraints.txt ./
RUN pip install -r requirements.txt -c constraints.txt

COPY manage.py pytest.ini ./
COPY mysite mysite

RUN ./manage.py collectstatic --noinput --clear
CMD ./manage.py runserver
