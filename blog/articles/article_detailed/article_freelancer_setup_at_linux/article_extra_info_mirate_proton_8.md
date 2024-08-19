## Extra info - Migrating from legacy Wine Proton 8.0.4

For those of you who used Wine Proton 8.0.4 before, you can reuse same created Wine Prefix
- You need only to repeat step `## 2. Setup Wine` to install new wine
- Specify it in your App Runner options
- and repeat step with Winetricks installation (specifically `msacm32` and `directplay`) and ensure dll overrides are present.
