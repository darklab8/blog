# Intro

In every software engineer path always comes a need to practice what he learned.
It is not possible just to go through tutorials, online courses and books only, as they remain only dead knowledge.
A person wishing to become a software engineer should transform the dead knowledge to practice.
But how to pick which projects to do?

Also, some food for thought in addition, how much are we supposed to practice? Practice in self-studies should take not less than 50%, and reach even up to 90% of your time and further. I personally found myself at the beginning [a lot picking theory]({{.SiteRoot}}favourite.html#CodeCompleteAPracticalHandbookofSoftwareConstruction) only and practicing all the stuff at work (that made for me a balanced proportion that 10-15% of theory come with 85-90% of practice). But as I learned a plethora of code quality things that the majority of the world is rarely ever using anyway, i reached the need pretty much even in my free time mostly practicing only too for the desired stuff, despite having practice at work. What's difference between practice at work and home? At work, I use what i have to. At home i use the most desired quality tech, the one i most enjoy to work with, and would prefer to get better enough to make part of stuff i am comfortable to work with.

# UI tools

I find sparkling sparkling with ideas, if I know a good UI tool.
- To implement CLI
- to implement Desktop GUI
- to implement TUI
- to implement web interface
    - which can be Discord bot too

Every tool we build, usually needs some end interface to interact with a user.

There are lists for ecosystem of every language, that are googlable like this "LanguageName Awesome curated list of libraries"
- [For python](<https://github.com/vinta/awesome-python#devops-tools>)
- [For golang](<https://github.com/avelino/awesome-go>)
- [For Java](<https://github.com/akullpp/awesome-java>)

Having tools to build UI interface, we have some "entry point" to plan how it will be looking for our user.

(P.S.technically we approach the stage formalized for devs as gathering requirements and building user usage case scenarios, this stuff actually has [theory]({{.SiteRoot}}favourite.html#SystemsAnalysisandDesign) about it, but i would prefer to recommend for starting people to read [Code Complete]({{.SiteRoot}}favourite.html#CodeCompleteAPracticalHandbookofSoftwareConstruction) first as more novice-friendly book to start and covering a lot of things that exist in programming), 

# Someone needs it, at least you

The best pet projects are the ones someone will be using and someone actually needs.
It is really pointless to make "business-like needed project in a vacuum". It could work as part of tutorial completion, but it has no life beyond its first implementation.
A good project has users 
- at minimum one, you. If you are not using it, how could you persuade other users that it will have a valid usability?
    - You could build CLI tool you need in development
    - you could build any interface daily task organizer for example too.
    - or you could find yourself one day wishing where to present projects you made, and writing words as articles once and then able to link them instead of repeating  again and again, and thus making [blog site]({{.SiteRoot}}) like the one this article in.
- Also by having at least one user you, it is ensured that someone will be sending bug fix requests, maintenance requests and may be even feature requests to the author

For example, i started with CLI instrument (because it is kind of easiest interface also too)
and went to build [autogit](<{{.SiteRoot}}pet_projects.html#autogit>) tool. Which makes for me easier to write [git conventional commits](<https://www.conventionalcommits.org/en/v1.0.0/>) and based on that, the tool generates changelogs i utilize to make releases for products. It helps me [to communicate with users and other developers using my tools]({{.SiteRoot}}article/git_conventional_commits.html).

Ideally, you are part of some community.
- For example, gaming community, minecraft, or in my case space simulator Freelancer 2003 became more part of my life
- Possibly part of some tool/framework, like Django, kubernetes, and etc.
- It can be any community that has specific problems and needs. As part of it as user, you can recognize needs the community have and offer application/program/solution that can do better in specific case than everything else the community has.

For me personally unlimited amount of things to do appeared because space game Freelancer community has accumulated in 20 years many many tools, but they became bad, deprecated, lost code and etc, and there is just a simple need to... rebuild things in a quality way. As long as i build employing at least [unit testing]({{.SiteRoot}}favourite.html#UnitTestingPrinciplesPracticesandPatterns) and giving a thought from the position of a user of this community, what is needed and what i miss for a comfortable life, i have always projects to do.
- As part of this community, I recognized need for Discord integration through a bot, and rewrote all broken bots community had [to my own bot]({{.SiteRoot}}pet_projects.html#fl-darkbot)
- As part of this community, I recognized need for new [game info parsing tool]({{.SiteRoot}}pet_projects.html#fl-darkstat) that shows game data in a comfortable way for user with added things that became relevant in latest modding develoment
    - As member of community i recognized which features existing tools miss, and what is needed to be built! For future details when the tool was made and my features were implemented... community just voiced all the extra feature requests it wishes.

Other ideas appear for me because i just... use different tools with their own ecosystem and find things i miss.
- I love ~~terraform~~ opentofu, and this makes me tempted to build [my own terraform provider](https://registry.terraform.io/browse/providers) which will extend what i will be able to do with this tool
- I use kubernetes, and i find its over yamling... not nice to my style of coding, for that reason tried for fun making ArgoCD plugin for support of a Cuelang language that addresses this issue.

I love playing minecraft, and that makes me opened to infinity a plethora of minecraft modding projects. Including as just infrastructure project to run minecraft server with mods.

# Lets summarize things we said.

- Ideally you build projects that someone actually needs (at least you. Just a single user makes all the difference. U will be able to write articles about the written program if u are at least the user of it)
- Your project should come preferably from a heart. Build something the most according to your interests, even if it is smth related to game modding or some other very silly topic.
    - Building smth that others need for work/dev related stuff is even more nicer to do eventually, as u are even more contributing back to society.
- For you TO SEE what the community/some tool needs in terms of development, u should be PART OF IT. As user that uses it too, u will face the limitations the users have and u will be able to recognize what can be done to do better, what can be addressed with programmatic solution!
    - As part of a community you are able to make research what other solutions already present there and how your solution can be having things better than others already present there. (the most obvious sign, the other solutions are broken, or not working for your specific usage cases, or no longer maintained.)
- It is a good entry point to learn which CLI, TUI, GUI, Web UI tools exist and to learn them in order to understand how to build interface for an end user. Some tool can be running with UI as Discord bot, other can be mods and having as user interface game itself.
- It will be very nice if u will build eventually projects with [unit testing]({{.SiteRoot}}favourite.html#TestDrivenDevelopmentByExample). U could learn [theory about it here]({{.SiteRoot}}favourite.html#UnitTestingPrinciplesPracticesandPatterns) why it is important. If to say shortly though, if u will do for project above 50% coverage in unit testing, and it will be possible to cover it up to 80-90%... it is almost guaranteed that your project has sufficient code quality to evolve and be maintained easily through years, through your entire career with you (if the need in the project will persist through this dozens years of time)
    - If your use dynamic type danguage , some static typing can be especially nice (Mypy/Pyright for python, Typescript for js). That increases likely hood that if u open your code in a year or two, u will be able to read what is going on there and to clean up the now shitty code (as your skill improve in a year or two, it is very unevitable at a start often seeing past code as quite bad), make it clean again and continue with new life to it (instead of rewriting from zero... again. P.S. i rewrote my discord bot three times from zero before it became charming)
- Make sure you don't go through [Tutorial hell](https://www.linkedin.com/pulse/escaping-tutorial-hell-guide-progress-your-learning-journey-jatasra-dvdgf). Practice makes perfection, and in pet projects, you will be challenged far more than any tutorials, or online courses offer. So all the time input into projects is justified.
