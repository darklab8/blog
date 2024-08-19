
## 6. Setup DxVk

- [Download DxVK v2.4](https://github.com/doitsujin/dxvk/releases/tag/v2.4) and unpack `d3d8.dll` and `d3d9.dll` (from 32bit version)
into your Freelancer/EXE folder 

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dxvk_install_manual.png)

open for Lutris application `Freelancer Online` (or how u named it) options menu and open `Wine Configuration` and turn on overrides for `d3d8.dll` and `d3d9.dll`

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dxvk_install_manual_dll_overrides.png)

if u did this step correctly, Alt tabbing will work correctly regardless of how many times u alt tabbed at full screen.

You can optionally add Environment variable `DXVK_HUD=devinfo,fps,api` as described here in docs https://github.com/doitsujin/dxvk?tab=readme-ov-file#dll-dependencies

![]({{.StaticRoot}}article_freelancer_setup_at_linux/dxvk_install_manual_confirmation.png)

if u did this step, later in game u will see turned on d3d8 in a HUD overlay for confirmation
