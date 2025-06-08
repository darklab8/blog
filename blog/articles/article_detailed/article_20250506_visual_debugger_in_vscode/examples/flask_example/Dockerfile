FROM python:3.13-bookworm

ENV PYTHONUNBUFFERED 1
ENV PYTHONDONTWRITEBYTECODE 1
WORKDIR /code

COPY ./requirements.txt ./constraints.txt ./
RUN pip install -r requirements.txt -c constraints.txt

COPY app.py app_test.py pytest.ini ./

CMD python3 app.py
