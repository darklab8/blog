#!/usr/bin/env python3

from pathlib import Path
import subprocess
import os
import sys
import ctypes
import struct

class Config(ctypes.Structure):
    _fields_ = [
        ("unk", ctypes.c_char * 4),
        ("name", ctypes.c_char * 66),
        ("description", ctypes.c_char * 258),
        ("password", ctypes.c_char * 34),
        ("capacity", ctypes.c_char),
        ("allow_new_players", ctypes.c_char),
        ("internet_accessible", ctypes.c_char * 4),
        ("pvp_enabled", ctypes.c_char)
    ]

config = Config(
    unk = struct.pack("i", 3),
    name = os.environ.get("SERVER_NAME", "FLServer Dockerised")[0:64].encode('utf-8'),
    description = os.environ.get("SERVER_DESCRIPTION", "FLServer running within docker.")[0:256].encode('utf-8'),
    password = os.environ.get("SERVER_PASSWORD", "")[0:32].encode('utf-8'),
    capacity = int(os.environ.get("SERVER_CAPACITY", "16")),
    allow_new_players = struct.pack("?", os.environ.get("SERVER_ALLOW_NEW_PLAYERS", "true").lower() == "true"),
    internet_accessible = struct.pack("i", 32 if os.environ.get("SERVER_INTERNET_ACCESSIBLE", "true").lower() == "true" else 0),
    pvp_enabled = struct.pack("?", os.environ.get("SERVER_ENABLE_PVP", "true").lower() == "true")
)

path_str = "/home/wineuser/.wine/drive_c/users/wineuser/Documents/My Games/Freelancer/Accts/MultiPlayer"
Path(path_str).mkdir(parents=True, exist_ok=True)

with open(path_str + "/FLServer.cfg", "wb") as data:
    data.write(ctypes.string_at(ctypes.byref(config), ctypes.sizeof(Config)))

news = os.environ.get("SERVER_NEWS")
if news is not None:
    news_bytes = news.encode('utf-8')
    with open(path_str + "/motd.dat", "wb") as data:
        data.write(news_bytes)

# We have to install directplay manually due to registrations not being done correctly within the dockerfile
process = subprocess.Popen(["xvfb-run", "winetricks", "-q", "--force", "directplay"], 
                           stdout=sys.stdout, stderr=sys.stderr)
process.wait()

process = subprocess.Popen(["xvfb-run", "wineconsole", "regedit", "/app/EULABypass.reg"], 
                           stdout=sys.stdout, stderr=sys.stderr)
process.wait()

# Run FLServer with the provided config
result = subprocess.Popen(["xvfb-run", "wine", "/app/EXE/FLServer.exe", "/c" ], 
                          cwd="/app/EXE", stdout=sys.stdout, stderr=sys.stderr)
result.wait()