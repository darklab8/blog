
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
