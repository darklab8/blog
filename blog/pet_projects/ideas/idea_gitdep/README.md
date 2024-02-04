// A person has to deal with C++ dependencies


Ergh, at this point may be making small scripting tool to download and track dependencies sounds like not a bad choice \
Using git checkout under the hood \
It will match how Terraform reusable modules are shared \
It downloads dependency matched in tag version , it automatically reloads them if tag in requirements no longer matches downloaded (edited) \
Pretty simple concept fully based on git tags and git used for cloning (edited) \ 
For that matter I could have written such tool without any dependencies working for any language completely in language agnostic way \
In Golang pretty simple  (edited) \
Sounds like a nice little project. Very small effort for high gain \
All dependencies will be declared in toolname.yml \

Under dependencies: tag \

Darkwind — Yesterday at 11:47 PM \
Since golang has available Inbuilt git client, result will be not using system level deps at all \
