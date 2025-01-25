# docker-wine is a compatability layer between Windows and Linux.  Allows you to run Windows programs on Linux machines.
FROM scottyhardy/docker-wine:latest

EXPOSE 2302/udp

ENV WINEARCH="win32"

WORKDIR /app

# Install dependencies
RUN /usr/bin/entrypoint xvfb-run winetricks -q --force vcrun2022 vcrun6sp6 riched30 directplay

COPY EULABypass.reg .
COPY Freelancer .
COPY fl.py .

ENV SERVER_NAME="FLServer On Docker"
ENV SERVER_DESCRIPTION="FLServer is running on docker, that's crazy!"
ENV SERVER_PASSWORD=""
ENV SERVER_CAPACITY="16"
ENV SERVER_ALLOW_NEW_PLAYERS="true"
ENV SERVER_INTERNET_ACCESSIBLE="true"
ENV SERVER_ENABLE_PVP="true"
ENV SERVER_NEWS="1234"

ENTRYPOINT [ "/usr/bin/entrypoint", "python3", "fl.py" ]