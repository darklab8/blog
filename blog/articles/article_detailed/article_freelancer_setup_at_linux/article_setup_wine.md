## 2. Setup Wine

There are many wine versions, but this one is the last that keep the UI text from not cutting out
and also keeps intro video playing nicely with Freelancer Vanilla.

- [Download Wine-9.0](<https://github.com/Kron4ek/Wine-Builds/releases/download/9.0/wine-9.0-amd64.tar.xz>)

![]({{.StaticRoot}}article_freelancer_setup_at_linux/wine_installing.png)

- find folder with Lutris wines (at Kubuntu `$HOME/.local/share/lutris/runners/wine`) and unpack archive in to it (i unpacked `wine-9.0-amd64` since i have amd64)
- P.S. if you use Steam Deck, the possible expected path is known to look like `$HOME/.var/app/net.lutris.Lutris/data/lutris/runners/wine`

the folder inside expects to look like this

![]({{.StaticRoot}}article_freelancer_setup_at_linux/wine_expected.png)

Relaunch lutris, u should see it in discovered wines

![]({{.StaticRoot}}article_freelancer_setup_at_linux/wine_expected2.png)

P.S. Known issues for Wine 9.0. For me the keyboard is not available right away at the game entering.
It becomes active only if i alt tab and return. Ensure on game enter, u can open chat by "Enter" before flying from planet! Quick alt tabbing in game helps me to fix it.