 
Redid my VPS test box with lxde-core for a lighter footprint:

sudo apt update && sudo apt upgrade
sudo apt install lxde-core lxterminal xrdp xorgxrdp firewalld wine winetricks cabextract
sudo nano /etc/xrdp/startwm.sh
comment out (#):
test -x /etc/X11/Xsession && exec /etc/X11/Xsession
exec /bin/sh /etc/X11/Xsession
add at the bottom: "startlxde" and save
sudo nano /etc/xrdp/xrdp.ini
change port= to a non standard port and save
sudo systemctl restart xrdp
sudo firewall-cmd --add-port=49974/tcp --permanent <---- The number being your non standard port
sudo firewall-cmd --reload
to save a bit of ram, stop the host from booting into the gui with: systemctl set-default multi-user.target
dpkg --add-architecture i386 && apt-get update && apt-get install wine32:i386
RDP in at this point, since you need a GUI to use winetricks
winetricks vcrun6 riched30 directplay vcrun2022
wine start /Path/To/FLServer.exe
