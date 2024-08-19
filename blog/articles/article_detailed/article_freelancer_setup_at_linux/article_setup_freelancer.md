
## 3. Setup Freelancer and Wine Prefix

Visit the [Freelancer installing link](<https://discoverygc.com/forums/showthread.php?tid=126999>) and load Freelancer.iso

![]({{.StaticRoot}}article_freelancer_setup_at_linux/freelancer_install.png)

For fast speed recommended using torrentink link and loading via Transmission (or Qbittorrent). Unpack the iso with archive program to some folder (i used Ark app coming as default in Kubuntu).

After that u need to create locally installed Game in Lutris. Name in whatever way u wish.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/freelancer_install_wine_prefix.png)

- The important part u need to select 32 `Prefix architecture` inside (even if your OS has installed 64 bit Wine, it still supports 32 bit running)
- And u need to select any empty dedicated folder for `Wine prefix` (I chose `/home/naa/apps/freelancer_related/wine_prefix_freelancer_online`)
    - that ensures a fresh copy of Windows Filesystem emulation will be used with dedication to our Freelancer setup.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/freelancer_install_game_options.png)

- For `executable` set path to unarchived SETUP.exe from freelancer.iso: like this `/home/naa/apps/freelancer_related/freelancer_iso/SETUP.EXE`
- Save configuration and launch this lutris app in order to launch Freelancer setup. Make installation. To default address express installation is okay.
