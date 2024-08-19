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

Some players need limiting their frame rate.
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

