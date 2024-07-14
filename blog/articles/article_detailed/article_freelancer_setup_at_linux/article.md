## Intro

This article describes installation of [Freelancer Discovery](<https://discoverygc.com/>) online space simulator for Linux. We show specifics for Kubuntu 22.04 LTS, but it should be usable for any Linux as long as you are able to install Lutris and having reasonable CPU architecture (we go with amd64). We organize article in paragraphs having next structure:

- How to start step/stage of installation
- What u see as progress and warning about possible pitifals
- Confirming that the step was completed succesfully.

In theory the same guide can be adopted for any Freelancer/mod installation. As well as serve as inspiration to how dockerize its server. There are [wine docker images](<https://hub.docker.com/r/scottyhardy/docker-wine/>), and it would be interesting to make it happen.

The guide was written on 2024.07.14 and last time checked it is all working on 2024.07.14.

## Dependencies

Writing last versions of dependencies with which it was possible to make it happen:

- `Kubuntu 22.04.3 LTS`
- `lutris/now 0.5.16 all [installed,upgradable to: 0.5.17]` ( output of `$ apt list --installed | grep "lutris"`)
- [Wine proton 8.0-4](<https://github.com/Kron4ek/Wine-Builds/releases/tag/proton-8.0-4>)
- [dxvk2.4 with direct 8 support](<https://github.com/doitsujin/dxvk/releases/tag/v2.4>)

## 1. Setup Lutris

Visit list of [lutris installating instructions](<https://lutris.net/downloads>)
Check your OS and install Utris accordingly.

- Since we used Kubuntu, we downloaded [deb package](<https://github.com/lutris/lutris/releases/tag/v0.5.16>) for it
- and used `sudo apt install ./lutris_0.5.16_all.deb` for its installation

if you are able to launch Lutris and see it, the installation is succeful

![]({{.StaticRoot}}article_freelancer_setup_at_linux/installed_lutris.png)

We installed Lutris `0.5.16` at the moment of last attempt.

## 2. Setup Wine Proton

There are many wine versions, but this one is the last that keep the UI text from not cutting out
and also keeps intro video playing nicely with Freelancer Vanilla.

- [Download Wine Proton](<https://github.com/Kron4ek/Wine-Builds/releases/tag/proton-8.0-4>)

![]({{.StaticRoot}}article_freelancer_setup_at_linux/wine_proton_installing.png)

- find folder with Lutris wines and unpack archive in to it (i unpacked `wine-proton-8.0-4-amd64.tar.xz` since i have amd64)

the folder inside expects to look like this

![]({{.StaticRoot}}article_freelancer_setup_at_linux/wine_proton_expected.png)

Relaunch lutris, u should see it in discovered wines

![]({{.StaticRoot}}article_freelancer_setup_at_linux/wine_proton_expected2.png)

## 3. Setup Freelancer and Wine Prefix

Visit the [Freelancer installing link](<https://discoverygc.com/forums/showthread.php?tid=126999>) and load Freelancer.iso

![]({{.StaticRoot}}article_freelancer_setup_at_linux/freelancer_install.png)

For fast speed recommended using torrentink link and loading via Transmission (or Qbittorrent). Unpack the iso with archive program to some folder (i used Ark app coming as default in Kubuntu).

After that u need to create locally installed Game in Lutris. Name in whatever way u wish.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/freelancer_install_wine_prefix.png)

- The important part u need to select 32 `Prefix architecture` inside (even if your OS has installed 64 bit Wine Proton, it still supports 32 bit running)
- And u need to select any empty dedicated folder for `Wine prefix` (I chose `/home/naa/apps/freelancer_related/wine_prefix_freelancer_online`)
    - that ensures a fresh copy of Windows Filesystem emulation will be used with dedication to our Freelancer Discovery setup.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/freelancer_install_game_options.png)

- For `executable` set path to unarchived SETUP.exe from freelancer.iso: like this `/home/naa/apps/freelancer_related/freelancer_iso/SETUP.EXE`
- Save configuration and launch this lutris app in order to launch Freelancer setup. Make installation. To default address express installation is okay.

## 4. Setup winetricks installable dependencies.

Open Lutris bash console for the current application

```
winetricks msacm32
winetricks dotnet40
winetricks dotnet45
winetricks directplay
winetricks webdings
```

and install this list of dependencies

![]({{.StaticRoot}}article_freelancer_setup_at_linux/winetricks_install.png)

in the process u will see different stuff that resembles errors. that's okay

![]({{.StaticRoot}}article_freelancer_setup_at_linux/winetricks_progress1.png)

![]({{.StaticRoot}}article_freelancer_setup_at_linux/winetricks_progress2.png)

After installation. Go to your `Wine Prefix` folder and confirm in file `winetricks.log` (for me it is at the path `/home/naa/apps/freelancer_related/wine_prefix_freelancer_online/winetricks.log`), that u have shown as installed next stuff:

```
msacm32
remove_mono internal
winxp
dotnet40
remove_mono internal
dotnet45
directplay
webdings
```

## 5. Setup Discovery Freelancer

Visit the [Freelancer installing link](<https://discoverygc.com/forums/showthread.php?tid=126999>) again and load Discovery Installer.

Change in Lutris app configuration `executable` path to 

![]({{.StaticRoot}}article_freelancer_setup_at_linux/discovery_install.png)

- launch installation

![]({{.StaticRoot}}article_freelancer_setup_at_linux/installer_picture.png)

- select path to vanilla freelancer, or skip if used default
- pick path for Discovery installing destination. I prefered to install to `C:\Discovery\` path for less path headache.
- After installation completed. Change Lutris app `executable` to Discovery Launcher path. ( `/home/naa/apps/freelancer_related/wine_prefix_freelancer_online/drive_c/Discovery/DSLauncher.exe` for me )

![]({{.StaticRoot}}article_freelancer_setup_at_linux/discovery_pick_launcher_executable.png)

## 6. Setup DxVk

- [Download DxVK](https://github.com/doitsujin/dxvk/releases/tag/v2.4) and unpack to your lutrix dxvk folder.
I unpackaged to `v2.4` folder for the same convention used as for other installed dxvk

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dxvk_install.png)

Relaunch Lutris and check that u can select in app Runner options 2.4 version

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dxvk_confirmation.png)

## 7. Setup dll override

u need extra variable into Lutris app configuration in Runners options. key `msacm32` and value `n,b`

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dll_override_setup.png)

It is possible to do the same by opening `Wine Configuration` menu for the app and selecting the dll override there

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dll_override_setup_2.png)

## 8. Launching game

U should be having already selected DSLauncher.exe as main `executable`

![]({{.StaticRoot}}article_freelancer_setup_at_linux/discovery_pick_launcher_executable.png)

Launch the game :) You are ready. Close notes, agree to license, patch to latest and launch as normal.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/game_launch.png)

## Extra info - Fonts.

Thie step is not need if u used this guide main instruction and used Lutris/Wine to install the game properly as described in `3. Setup Freelancer and Wine Prefix` step.
If you copied game from windows 10 instead of installing in Lutris/Wine, then u need to put the fonts manually in windotws/Fonts folder

![]({{.StaticRoot}}article_freelancer_setup_at_linux/fonts_setup1.png)

U can find them in the freelancer ISO

![]({{.StaticRoot}}article_freelancer_setup_at_linux/fonts_setup_source.png)

## Extra info - Reliable Alt-tab switch.

In general alt tab works. and u can play from window after turning it on through Alt-Enter.
But sometimes alt tabing breaks the game. In this case u can install gamescope separately and turn on it in Lutris app configuration.
It should make alt-tab working reliably.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/gamescope_install.png)

P.S. Alternatively can be used Virtual Desktop feature, but it is not available anymore in the proton version. This is why we go with gamescope.

## Extra info - d3d8to9 v1.12.0

The guide used dxvk 2.4 with direct 8 support to launch the game, but [d3d8to9 v1.12.0](https://github.com/crosire/d3d8to9/releases/tag/v1.12.0) can be used as alternative.

## Extra info - vendoring dependencies

Just in case providing copies of all main dependencies, vendored in for this articles, in case they aren't longer downloadable.

- [Lutris 0.5.16]({{.StaticRoot}}article_freelancer_setup_at_linux/vendored_freel_stuff/lutris_0.5.17_all.deb)
- [Wine proton 8.0-4]({{.StaticRoot}}article_freelancer_setup_at_linux/vendored_freel_stuff/wine-proton-8.0-4-amd64.tar.xz)
- [dxvk2.4 with direct 8 support]({{.StaticRoot}}article_freelancer_setup_at_linux/vendored_freel_stuff/dxvk-2.4.tar.gz)

## Extra info - useful links

- [This forum page](<https://discoverygc.com/forums/showthread.php?tid=173057>) contains legacy instruction from other person and has anouncements regarding this upaded version and possible comments from other people. May be helpful for other bits to know.

## Acknowlegments

- Emawind (Source of truth. This guide would not be possible without this person)
    - he mentioned using it with Steam Deck.
- Darkwind (A.k.a me: Interrogating source of truth, following the instructions and writing this article)
    - Your Kubuntu user.
