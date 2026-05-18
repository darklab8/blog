## FLhook server
void catgirl rglx — 5/7/26, 2:52 AM
okay. i think the server starts and is working. i stashed my wineprefix and made a new one, and with winetricks installed the following:
winetricks -q directplay vcrun2026 riched30 vcrun6sp6 dotnet48 win10
(that last winetricks verb switches the prefix back to windows 10, which i think some winetricks instances forget to do at the end of installing dotnet48)
after that simply running wine flserver.exe in the EXE directory started the server just fine
vcrun2026's description on their website seems to imply that it includes everything from the 2015 runtimes forward. but this is microsoft we're talking about so who the hell really knows what that includes.


## Flstat with infocards showing

Darkwind The Dark Duck — 11/28/25, 9:23 PM
AHAHAHAHAHAHAHA
Image
I DID IT
i found the dependency for old flstat to run in Wine correctly with infocard showing
(Also i started more to belive into chatgpt because it was a paired debugging session with it)
winetricks -q ie8 helped. may be just winetricks -q mshtml should be enough too (that thing was installed with it)
winetricks -q wine_gecko potentially could work or be broken
apperently flstat depends on Html ActiveX control bla bla bla to transform its xml to ui 
tempted to be not lazy and writing it into some mini guide at my web site 
( launched it succesfully in proton 8.0.4 amd64 wine, but technically should run in wine9 fine too and etc)
hehe, old flstat essentially depends kind of onto Internet Explorer (its dependency more like but still funny) 
i could tell btw, that chatgpt was full of advices what to try to make any freelancer tool working. Potentially has rich amount of scanned pages for it somehow inside. 

## You are arch user installing Disco

Darkwind The Dark Duck — 1:25 AM
Manjaro is Arch based... you have same Arch issue then
Žugr — 9/18/25, 1:03 AM
It was simply pacman -S wine-gecko
Žugr — 9/18/25, 1:04 AM
EndeavorOS
try this
i should add this into guide, since it looks to be consistent issue potentially then (the guide was written with presumption it is executed on Ubuntu based system, not Arch based) 
some other Arch people are present, and could be pinged for extra advice