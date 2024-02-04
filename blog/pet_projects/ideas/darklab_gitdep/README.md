// A person has to deal with C++ dependencies

...
Ergh, at this point may be making small scripting tool to download and track dependencies sounds like not a bad choice
[11:39 PM]
Using git checkout under the hood
[11:40 PM]
It will match how Terraform reusable modules are shared
[11:40 PM]
It downloads dependency matched in tag version , it automatically reloads them if tag in requirements no longer matches downloaded (edited)
[11:42 PM]
Pretty simple concept fully based on git tags and git used for cloning (edited)
[11:45 PM]
For that matter I could have written such tool without any dependencies working for any language completely in language agnostic way
In Golang pretty simple  (edited)
[11:45 PM]
Sounds like a nice little project. Very small effort for high gain
[11:46 PM]
All dependencies will be declared in toolname.yml

Under dependencies: tag

Darkwind — Yesterday at 11:47 PM
Since golang has available Inbuilt git client, result will be not using system level deps at all