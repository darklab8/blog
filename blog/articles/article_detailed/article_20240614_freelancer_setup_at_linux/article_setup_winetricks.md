
## 4. Setup winetricks installable dependencies.

Open Lutris bash console for the current application

```
winetricks msacm32
// winetricks dotnet40 // i suspect not needed at all any longer if u install dotnet48, try not to install
// winetricks dotnet45 // deprecated disco dotnet, no longer needed
// some level of issues with dotnet48 is present, ensure to deinstall old one dotnet45 first and install new one with --force flag
// for deinstallation of dotnet45, run `wine deinstaller` in Bash lutris console`, in same console where we input winetricks commands
// https://github.com/Winetricks/winetricks/issues/2159
// https://www.reddit.com/r/winehq/comments/eqx1uu/wine_doesnt_see_mscoreedll/
winetricks --force dotnet48 // latest in use by Discovery Freelancer
winetricks directplay
winetricks webdings
```

and install this list of dependencies

![]({{.StaticRoot}}article_20240614_freelancer_setup_at_linux/winetricks_install.png)

in the process u will see different stuff that resembles errors. that's okay

![]({{.StaticRoot}}article_20240614_freelancer_setup_at_linux/winetricks_progress1.png)

![]({{.StaticRoot}}article_20240614_freelancer_setup_at_linux/winetricks_progress2.png)

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
