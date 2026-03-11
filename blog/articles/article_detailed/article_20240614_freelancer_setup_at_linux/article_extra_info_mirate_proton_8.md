## Extra info - Migrating from any previous Wine version

For those of you who used Wine Proton 8.0.4 before Or Wine 9.0, you can reuse same created Wine Prefix
- You need only to repeat step `## 2. Setup Wine` to install new wine
- Specify it in your App Runner options
- and repeat step with Winetricks installation (specifically `msacm32` and `directplay`) and ensure dll overrides are present.

Full clean reinstall is still recommended though.