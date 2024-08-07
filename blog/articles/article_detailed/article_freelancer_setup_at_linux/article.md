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
- [Wine 9.0](<https://github.com/Kron4ek/Wine-Builds/releases/download/9.0/wine-9.0-amd64.tar.xz>)
- [dxvk2.4 with direct 8 support](<https://github.com/doitsujin/dxvk/releases/tag/v2.4>)

## 1. Setup Lutris

Visit list of [lutris installating instructions](<https://lutris.net/downloads>)
Check your OS and install Lutris accordingly. Article guarantees all this stuff will work with lutris 0.5.16, how it will work with other versions is not guaranteed.

- Since we used Kubuntu, we downloaded [deb package](<https://github.com/lutris/lutris/releases/tag/v0.5.16>) for it
- and used `sudo apt install ./lutris_0.5.16_all.deb` for its installation

if you are able to launch Lutris and see it, the installation is succeful

![]({{.StaticRoot}}article_freelancer_setup_at_linux/installed_lutris.png)

Verify that u installed Lutris `0.5.16` in Lutris About. `lutris --version` for CLI.
If we will use newer Lutris version, the guide will be updated.

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
It becomes active only if i alt tab and return. Ensure on game enter, u can open chat by "Enter" before flying from planet!

## 3. Setup Freelancer and Wine Prefix

Visit the [Freelancer installing link](<https://discoverygc.com/forums/showthread.php?tid=126999>) and load Freelancer.iso

![]({{.StaticRoot}}article_freelancer_setup_at_linux/freelancer_install.png)

For fast speed recommended using torrentink link and loading via Transmission (or Qbittorrent). Unpack the iso with archive program to some folder (i used Ark app coming as default in Kubuntu).

After that u need to create locally installed Game in Lutris. Name in whatever way u wish.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/freelancer_install_wine_prefix.png)

- The important part u need to select 32 `Prefix architecture` inside (even if your OS has installed 64 bit Wine, it still supports 32 bit running)
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

- [Download DxVK v2.4](https://github.com/doitsujin/dxvk/releases/tag/v2.4) and unpack `d3d8.dll` and `d3d9.dll`
into your Freelancer/EXE folder

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dxvk_install_manual.png)

open for Lutris application `Freelancer Online` (or how u named it) options menu and open `Wine Configuration` and turn on overrides for `d3d8.dll` and `d3d9.dll`

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dxvk_install_manual_dll_overrides.png)

if u did this step correctly, Alt tabbing will work correctly regardless of how many times u alt tabbed at full screen.

You can optionally add Environment variable `DXVK_HUD=devinfo,fps,api` as described here in docs https://github.com/doitsujin/dxvk?tab=readme-ov-file#dll-dependencies

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dxvk_install_manual_confirmation.png)

if u did this step, later in game u will see turned on d3d8 in a HUD overlay for confirmation

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

`By default you don't need this step, if u followed Step "3. Setup Freelancer and Wine Prefix" step correctly!`

If you copied game from windows 10 instead of installing in Lutris/Wine, then u need to put the fonts manually in windotws/Fonts folder

![]({{.StaticRoot}}article_freelancer_setup_at_linux/fonts_setup1.png)

U can find them in the freelancer ISO

![]({{.StaticRoot}}article_freelancer_setup_at_linux/fonts_setup_source.png)

## Extra info - Legacy reliable Alt-tab switch.

`By default you don't need this step, if u followed Step DxVK 2.4 setup correctly!`

Current version with `## 6. Setup DxVk` having manual d3d8/d3d9 from dxvk makes alt tab working reliably at it is.
Having this section only out of legacy how this problem was solved in the past.

Older instruction suggested installing gamescope separately and turn on it in Lutris app configuration.
It should make alt-tab working reliably.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/gamescope_install.png)

P.S. Alternatively can be used Virtual Desktop feature.

## Extra info - d3d8to9 v1.12.0

The guide used dxvk 2.4 with direct 8 support to launch the game, but [d3d8to9 v1.12.0](https://github.com/crosire/d3d8to9/releases/tag/v1.12.0) can be used as alternative.

## Extra info - vendoring dependencies

Just in case providing copies of all main dependencies, vendored in for this articles, in case they aren't longer downloadable.

- [Lutris 0.5.16]({{.StaticRoot}}article_freelancer_setup_at_linux/vendored_freel_stuff/lutris_0.5.17_all.deb)
- [Wine 9.0]({{.StaticRoot}}article_freelancer_setup_at_linux/vendored_freel_stuff/wine-9.0-amd64.tar.xz)
- [dxvk2.4 with direct 8 support]({{.StaticRoot}}article_freelancer_setup_at_linux/vendored_freel_stuff/dxvk-2.4.tar.gz)

## Extra info - possible unexpected dependencies.

as it was mentioned before lutris/now 0.5.16 version is currently used and it may be turning on some unexpected other important stuff.
Just in case showing page with all app default settings. This information may be will be helpful for future maintenance.

![]({{.StaticRoot}}article_freelancer_setup_at_linux/unexpected_dependencies.png)

## Extra info - Limit Frame rate

Some Discovery players need limiting their frame rate.
You can do it by inserting desired framerate into Env variables of application `DXVK_FRAME_RATE=150` for example

![]({{.StaticRoot}}article_freelancer_setup_at_linux/frame_limit_install.png)

You can validate it works, by also injecting env var `DXVK_HUD=devinfo,fps,api` to see FPS in real time

## Extra info - Migrating from legacy Wine Proton 8.0.4

For those of you who used Wine Proton 8.0.4 before, you can reuse same created Wine Prefix
- You need only to repeat step `## 2. Setup Wine` to install new wine
- Specify it in your App Runner options
- and repeat step with Winetricks installation (specifically `msacm32` and `directplay`) and ensure dll overrides are present.

## Extra info - useful links

- [Anouncements and communicating with players for current tutorial](<https://discoverygc.com/forums/showthread.php?tid=204034&pid=2339308>)
- [Tutorial from Missfit](<https://discoverygc.com/forums/showthread.php?tid=173057>). Very outdated, but still useful for some bits
- [Tutorial from LieOfDamashi](<https://discoverygc.com/forums/showthread.php?tid=202325&highlight=Linux>). Relatively fresh and about Lutris too.
- [Tutorial from Corile](<https://discoverygc.com/forums/showthread.php?tid=147190&highlight=Linux>). Recently was used

## Acknowlegments

- Emawind (Source of truth. This guide would not be possible without this person)
    - he mentioned using it with Steam Deck.
- Darkwind (A.k.a me: Interrogating source of truth, following the instructions and writing this article)
    - Your Kubuntu user.
