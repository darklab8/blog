# Long term maintained software

## The article that sparkled the ideas

While reading article [cold blooded software](<https://dubroy.com/blog/cold-blooded-software/>), we can see an interesting analog comparing some projects and tech to cold blooded software and some to hot burning ones.

It recounted next qualities for cold blood projects:
- Depends on boring tech
- Vendors in dependencies

And that according to author makes project easy maintainable in a long run for dozen of years with big gaps between development efforts.

I really liked the reptile analogy where we compare long term maintaince projects to turtles which can be frozen and reawakened later, but i think his ideas can be extended much further than two qualities he mentioned.

Lets place as next goals what we wish for maintaing our projects in terms of multiple years of lifetime with huge time passed between development active phases:
- increasing maintanability to maximum
- decreasing toll for updates to minimum
- increasing project code readability to maximum.

Lets take some examples first for clarity

## Super hot burning project

For such project we will take project written in ReactJS

- it will be depended on dozens of different node.js libs, each one having its own updating cycles.
- if we will not be regularily launching CI and discovering in time small things broken, then in a year amount of things accumulated can reach too huge amount.
- It will be hard distinguishing between all the breakings, what needs to be fixed first in order to make the project again operational.

I agree with author of the article above, that many projects will be abandoned by authors and / or will stop working with new versions of different dependencies.

I will add here that if not having in this project typization increasing readability of a project, even author of the project will completely forget all hidden details how the code was working in a few months.

## Super cold project

Lets take project made in Golang.
- Due to the specifics of how golang is made, it is not needing a lot of libraries to operate in the first place. A lot of stuff is present already in std libraries and golang is often not needing extra libraries to make a plenty of things.
- Due to the spirit of best practices and attitudes flowing through its community, it is common to be writing solutions using as least external dependencies as possible
- Static typing in addition increases readability of project considerably. it will not be hard to reread a code of a project in a year and figuring out all its internal working, just because when u type, u manage to write code that describes itself well for everything that is going on [about it in a small function](https://mvysny.github.io/code-locality-and-ability-to-navigate/). You often don't need rereading through full project and figuring out all the possible data mutations in order to understand how the project is operating (A common thing needed in javascript / python projects)
- The developers of Golang language promise backward compatibility, and develop features in a very careful steady amount.

All those things make high chance that project in golang u try rebuild in a year of your absence, will be still compilable and operating still correctly and u will be able easily return from where u left it.

## Revisioning qualities of cold blooded

### 1. Vendoring in dependencies - Yes.

I agree with article thought about it. That's an obvious plus, since in a few years any dependency can dissapear for some reason.

### 2. Using boring tech - No, Using Stable technologies.

I disagree with article regarding using boring tech, because author is not considering increased code complexity to maintain such project.
We could take javascript here. A very boring tech commonly used in browsers across everywhere.
Yet, it is rapidly changing in its major versions for Node.js compilers and its standards.
Due to lack of any typing it will be very hard to reread the code of any long term made project in javascript, and therefore hard to return or continue.

Or we could take even more ridiculous example of choosing boring technology like Assembly. Yet building in it something sane to maintain will be very challenging adventure.

I think it is important choosing static typed stable technlogy first! Like Golang. Or even Java.
All those languages are self documenting themselves to maximum and very stable in amount of changes.

We could continue thought here, and say that usage of Docker despite it being not boring will be completely justified for long term project.
It follows very stable standard of containers. It is made in stable tech.
It will add to our long term project important amount of self documentation for all system level dependencies and how to build it!
Very valid to use if we go for any web related project at least.

### 3. Extra - Static typing.

I will repeat here that one of most important qualities is static typing for our long term projects
Because it allows us in a not ambigious way to write very self documented code.
We will be able easily reread what is going and continue working with it.

But... there is certain danger involved on this path, from the point that static typing needs to be stable technology too.
Like it is for Golang or Java and etc.
If we will choose such hot burning stuff like types for Javascript or Python... things are going to be potentially way more problematic.
We need at least to make a choice towards typing system that will not dissapear itself from language for sure in future years.

### 4. Extra - Unit testing.

Typing is only part of documentation. We can do more.
Unit tests are part of the documentation showing how your code can be used.
Checking throughly that it works for all aspects u will forget in a year.
Ideally our language should be having inbuilt unit testing framework that is not requiring extra dependencies though.
That will help keeping unit testing especially cold blooded

### 5. Extra - Auto generated documentation

Typing, unit testing, it is all part of documentation, which can be linked into self generated docs from the code.
Golang godocs tooling, or python sphinx autodoc stuff, it all helps us throughly documenting our code as a code.
That can add additional layer of helping easier to maintain the code.
Documentation matters. In a year u will not remember any difficult enough decision u made regarding the code.
if u will be able to reread easily your previous decision and continue from that, u decrease toll further for long term maintanance.
Unit tests can be literally linked in both godocs and sphinx autodoc as examples into documentation.
U can help yourself by organizing unit tests and their file on purpose for linking to docs

Static typing, unit testing and auto generated documentation like godocs/ sphinx autodoc they all synergize to add very protected from deprecating documentation.
Because all of it will remain part of the code and code does not lie.

U need to be careful in using comments and documenting with such tools though. Make sure to keep your documentation DRY as a code.
Don't repeat yourself in documenting efforts. it will help upkeeping it easier.
Also don't document/comment stuff that is too easily descriable with better types / rewritten code logic / provided unit test unless it is absolutely necessary.
YAGNI principle is applyable for documentation too :)

### 6. Extra - Automated periodic warming up.

The original author rejected CI usage because they can be bought or ran out of money.
I disagree with him on this point, because CI is usually requring very little code effort for pet projects.
But it adds additional quality to self documentation of the project how to build, test and run it.
I scavenged old projects sometimes only with the help of broken, yet present CI code before.

U can help your long term maintaned project further with asigning [cron job scheduled](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#schedule) periodic CI triggering, like once in a week or month.
That will help you tremendously in a bit warming up project and discovering little accumulated problems before they snowballed into big amount.

Yes with low amount of dependencies it should not be big issue, but with maintaing stuff through many years, anything can break apart.
Plus as we remember we need extra help due to potentially using extra dependencies that very greatly help to keep our code simple and sane.

### 7. Extra - Sacrifice reasonably low dependnencies towards simpler and type safer code.

Keeping code too boring will lead towards too much complex code. For example writing this blog site, i could have went for fully html/css only
and only my self written solutions for static building.
I made sacrifice of low dependencies towards simpler code through the usage of [templ go](https://github.com/a-h/templ) library

<p align="center">
  <img src="{{.StaticRoot}}cold_blood/templ_go_demo.gif"/>
</p>

This library helps me to keep very simple and easy to reread code that is certainly enjoyable to maintain. Also makes very easy customizing blog for some extra features.
Evaluate if brought extra technlogy will help you to maintain project easier because it made the project better type safed, unit tested of auto documented, and then make the choice if price the paid for this extra dependency helps more for long term mantaince than payment for it.

### 8. Extra - Keeping lower amount of dependencies

As much as possible u need to withold yourself from adding extra dependency
unless it is completely justified for one of the reasons i mentioned before or another one.
Usually custom made solutions can be prefered for long term maintained software because it will remain with minimal amount of code fitting exactly for your situation. They replace need for having extra libs.
When u build commercial hot burning software it can fine using libraries for every sneeze, because devs are able to keep up with their changes and plus they win a lot from using already made solutions over remaking frameworks on their own, but for long term maintained stuff solution developed in house is way more preferable, unless some reason justifies using already made third party library

## Conclusion.

Those are all qualities for long term maintained projects with big gaps between development active phases.
If u will follow some of them for active projects, u will still have achieved decreased toll for maintanance and magntitudes faster new developer onboardings. Not all of those qualities benefit every situation though, as they were written with keeping in mind long term maintained software by extremely small amount of developers.
