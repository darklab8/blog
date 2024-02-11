# Description

[Personal blog site](https://darklab8.github.io/github.com/darklab8/blog/)

# Code architecture choices

- We try to follow [our own advices](https://darklab8.github.io/github.com/darklab8/blog/article/long_term_maintained_software.html) for long term maintaned software
- the project is intended for 10+ years lifetime

# Features

- Static Site Generator, automatically deployed with Github CI to github pages
  - Having low amount of dependencies
- Templ templating for most of front operations and code reusage.
  - Golang is usable during this templating
- Offers markdown usage for article writing / or when else it is needed
- Important links offer on hover archived versions of a page
- Dark and Light theme switch (Dark is default)
- Mermaid.js diagrams
- Highlight.js is attached for code blocks (switches to dark mode for dark theme)
- Flexibly customizable further to any other desired features to add.
