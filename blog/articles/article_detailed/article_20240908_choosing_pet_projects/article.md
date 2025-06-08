# Intro

Every software engineer needs to practice what they have learned. 
Unfortunately, it is impossible to learn swimming through reading, as it would remain dead, passive knowledge until the skill is practiced. Similarly,  when mastering software engineering,  going through tutorials, online courses, and books is not the ultimate answer, as they share the same quality. Therefore, one wishing to become a software engineer should transform dead knowledge into practice. But how do you pick the projects for that?

Here is some food for thought. Off the top of your head, how much are we supposed to practice? From my experience, practice in self-studies should take no less than 50% and reach even up to 90% of your time and further. In the beginning, I preferred to [pick theory]({{.SiteRoot}}favourite.html#CodeCompleteAPracticalHandbookofSoftwareConstruction) only and practice at work (that made a balanced proportion of 10-15% of theory coming with 85-90% of practice). But as I learned a plethora of code quality elements that the majority of the world population rarely ever uses, I came to practice the stuff I needed even in my free time, despite having practice at work.
What’s the difference between practice at work and home? At work, I use what I have to. At home, I use something I enjoy working with the most and would prefer to get comfortable with enough to make it a part of my professional skillset.

# UI tools

I find sparkling sparkling with ideas, if I know a good UI tool.
- To implement CLI
- to implement Desktop GUI
- to implement TUI
- to implement a web interface
    - which can be Discord bot too

Every tool we build usually needs the end interface to interact with a user.

There are lists for ecosystems of every language that are googlable like this: "LanguageName Awesome curated list of libraries"
Here are some of the examples: 
- [For python](<https://github.com/vinta/awesome-python#devops-tools>)
- [For golang](<https://github.com/avelino/awesome-go>)
- [For Java](<https://github.com/akullpp/awesome-java>)

Having tools to build UI interfaces, we have an “entry point” to plan how it will look for our user.

(P.S. Technically, we approached the stage formalized for devs as gathering requirements and building user-usage case scenarios. Here is some [theory]({{.SiteRoot}}favourite.html#SystemsAnalysisandDesign) on the topic; however, if there was someone who was not too proficient, I’d recommend reading [Code Complete]({{.SiteRoot}}favourite.html#CodeCompleteAPracticalHandbookofSoftwareConstruction) first, as it is more novice-friendly and covers a lot of things that exist in programming).

# Someone needs it, at least you

The best pet projects are the ones someone actually needs and will be using. It is pointless to make a “business-like needed project in a vacuum." Even though it could work as a part of a tutorial completion, it has no life beyond its first implementation. A good project has users:

- at minimum one, you. If you are not using it, how could you convince others of its valid usability?
    - You could build a CLI tool you need in development
    - You could build any interface daily task organizer, for example, too.
    - ...or you could find yourself one day wondering where to present projects you’ve made and writing articles every now and then to have something to refer to instead of repeating things over and over again, thus making a [blog site]({{.SiteRoot}}) much like this one.
- Also, by having at least one user, which is you, it is ensured that someone will be sending bug fix requests, maintenance requests, and maybe even feature requests to the author. 
- As a more trivial reason, having users also ensures you will get some feedback and encouragement to continue. This is important as it helps with motivation to go further than you originally intended, and leads to product maturity.


For example, i started with CLI instrument (because it is kind of easiest interface also too)
and went to build [autogit](<{{.SiteRoot}}pet_projects.html#autogit>) tool. Which makes for me easier to write [git conventional commits](<https://www.conventionalcommits.org/en/v1.0.0/>) and based on that, the tool generates changelogs i utilize to make releases for products. It helps me [to communicate with users and other developers using my tools]({{.SiteRoot}}article/git_conventional_commits.html).

Ideally, you are part of some community.
- For example, the gaming community, Minecraft, or in my case, space simulator Freelancer 2003, became a part of my life. Another example is Starsector, a game that has a rich modding community in Java.
- You could also be a part of some tool/framework, like Django, Kubernetes, etc.
- It can be any community that focuses on specific problems and needs. As a participant, you can recognize the needs the community has and offer even a revolutionary solution in various forms, including an application or a program.

I found an unlimited number of things to do at the space game Freelancer community, as in 20 years they have accumulated a huge number of tools but failed to uphold the quality standards and maintain the code for an array of them; so the tools got bad, deprecated, lost code, etc., so there was a need to... rebuild things in a quality way. As long as I make sure to employ at least [unit testing]({{.SiteRoot}}favourite.html#UnitTestingPrinciplesPracticesandPatterns) and give it a thought from the perspective of a community member about what is needed and what I miss for a comfortable life, I always have a project to do.

- As a member of this community, I recognized the need for Discord integration through a bot and rewrote all the broken bots the community had into the [one and only bot]({{.SiteRoot}}pet_projects.html#fl-darkbot)
- As part of this community, I recognized need for new [game info parsing tool]({{.SiteRoot}}pet_projects.html#fl-darkstat) that shows game data in a comfortable way for user with added things that became relevant in latest modding develoment
    - Additionally, I recognized which features existing tools were missing and what needed to be built! Eventually, when the tool was made and my features were implemented... the community just voiced all the extra feature requests.

Other ideas come to me just because I use different tools with their own ecosystems and find things I miss.
- I love ~~terraform~~ opentofu, and this makes me tempted to build [my own terraform provider](https://registry.terraform.io/browse/providers) which will extend what i will be able to do with this tool
- I use Kubernetes, and I find it's overly yamling… not my style of coding, which is why I, just for fun, tried making an ArgoCD plugin to support a Cuelang language that addresses this issue.

I love playing Minecraft and Starsector, which introduces me to the infinity of their modding projects, including infrastructure projects to run Minecraft servers with mods.

# Lets summarize things we said.

- Ideally, you build projects that someone actually needs (at least you. Just a single user makes all the difference. You will be able to write articles about the written program if you are at least the user of it.)
- Your project should come preferably from the heart. Build something the most according to your interests, even if it is smth related to game modding or some other very silly topic.
    - Building something that others need for work/dev-related stuff is an even nicer thing to do, as you are, quite literally, contributing back to society.
- For you TO SEE what a community some tool needs in terms of development, you should be A PART OF IT. As a user, you will face the limitations, and you will be able to recognize what can be done better and what can be addressed with a programmatic solution.
    - As part of a community, you can research what other solutions already exist and how your solution could be better. The most obvious signs are that the other solutions are broken, not working for your specific usage cases, or no longer maintained.
- It is a good entry point to find out which CLI, TUI, GUI, and Web UI tools exist and to learn them to understand how to build an interface for an end user. Some tools can be running with UI as a Discord bot; others can be mods and have a user interface game itself.
- It will be very nice if you start building projects with [unit testing]({{.SiteRoot}}favourite.html#TestDrivenDevelopmentByExample). You could learn the theory and why it is important [here]({{.SiteRoot}}favourite.html#UnitTestingPrinciplesPracticesandPatterns). If to say shortly, if you do above 50% coverage in unit testing for a project (or maybe even cover it up to 80–90%), it is almost guaranteed that your project will have sufficient code quality to evolve and be maintained easily through years, or even your entire career (if the need in the project will persist through this dozens of years)
    - If you use a dynamic type language, some static typing can be especially nice (Mypy/Pyright for Python, Typescript for JS). That increases the likelihood of you opening your code in a year or two, and being able to read it, and grants you the ability to fix the newly found problems in it  (as, hopefully, your skill improves in a year or two; you quite inevitably start seeing your past code as quite bad as you progress with your skill), make it clean again, and continue with new life to it (instead of starting from square one... again). P.S. I rewrote my Discord bot three times from zero before it became charming.
- Make sure you don’t go through tutorial hell. Practice makes perfect, and you will be challenged far more in your pet projects than any tutorials or online courses offer. 
So all the time you dedicate to your projects is justified.
