# Post installation

I recommend going into your wine prefix created location, into user `~/apps/freelancer_related/winreprefix_freelancer_vanilla/drive_c/users/steamuser`. Wine by default creates `Documents` folder as symlink. i recommend removing symlink and just having as real folder. That is the best decision for multiple Freelancer installations, because they all very conflictful regarding already present save games if they have different installed mods and it will help to keep them separate and not crashing with weird errors.

Go to your dedicated wine prefix created for this Freelancer vanilla or mod (We create separate Wine prefixes for each!)
- `cd drive_c/users/steamuser`
- `rm Documents ; mkdir Documents && chmod a+rw Documents`

WARNING: if you crash with weird errors 000c22fe, 000c21de on New Game, that can be because your "My Games" is still symlinked to single place. Rename/delete existing ones to ensure u stat a fresh one not touched by other Freelancer mods. Or use advice above and delete symlink and create Documents folder as recommended



