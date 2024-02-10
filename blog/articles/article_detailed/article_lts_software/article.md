# Long term maintained software

## The article that sparkled the ideas

While reading the article {{.LinkColdBloodedSoftware}}, we can see an interesting analog comparing some projects and tech to cold-blooded software and some to hot burning ones.

It recounted the next qualities of cold blood projects:
- Depends on boring tech
- Vendors in dependencies

That according to the author makes a project easily maintainable in the long run for dozens of years with big gaps between development efforts.

I really liked the reptile analogy where we compare long-term maintenance projects to turtles which can be frozen and reawakened later, but I think his ideas can be extended much further than the two qualities he mentioned.

Let's start with next aims in terms of what we wish for maintaining our projects in terms of multiple years of lifetime with huge time passed between development active phases:
- increasing maintainability to maximum
- decreasing toll for updates to minimum
- increasing project code readability to a maximum.

Let's now take some examples first for clarity:

## Super hot burning project

For such project, we will take a project written in ReactJS

- it will be dependent on dozens of different node.js libs, each one having its own updating cycles.
- if we will not be regularly launching CI and discovering in time small things broken, then in a year amount of things accumulated can reach too huge amount.
- It will be hard to distinguish between all the breakings, and what needs to be fixed first in order to make the project again operational.

I agree with the author of the article above, that many projects will be abandoned by authors and / or will stop working with new versions of different dependencies.

I will add here that if not having in the project typization increasing readability of a project, even the author of the project will completely forget all the hidden details of how the code was working in a few months.

## Super cold project

Let's take a project made in Golang.
- Due to the specifics of how golang is made, it does not need a lot of libraries to operate in the first place. A lot of stuff is present already in std libraries and golang is often not needing extra libraries to make a plenty of things.
- Due to the spirit of best practices and attitudes flowing through its community, it is common to be writing solutions using as least external dependencies as possible
- Static typing in addition increases readability of project considerably. it will not be hard to reread a code of a project in a year and figure out all its internal working, just because when u type, u manage to write code that describes itself well for everything that is going on [about it in a small function](https://mvysny.github.io/code-locality-and-ability-to-navigate/). You often don't need rereading through the full project and figure out all the possible data mutations to understand how the project is operating (A common thing needed in javascript / python projects)
- The developers of Golang language promise backward compatibility, and develop features in a very careful steady amount.

All those things make a high chance that a project in golang u try rebuild in a year of your absence, will be still compilable and operating still correctly and u will be able easily to return from where u left it.

## Revisioning qualities of cold-blooded

### 1. Vendoring in dependencies - Yes.

I agree with the article's thought about it. That's an obvious plus, since in a few years any dependency can disappear for some reason.

### 2. Using boring tech - No, Using Stable technologies.

I disagree with the article regarding using boring tech, because author is not considering increased code complexity to maintain such project.
We could take javascript here. A very boring tech commonly used in browsers across everywhere.
Yet, it is rapidly changing in its major versions for Node.js compilers and its standards.
Due to lack of any typing, it will be very hard to reread the code of any long-term made project in javascript, and therefore hard to return or continue.

Or we could take an even more drastic example of choosing boring technology like Assembly. Yet building in it something sane to maintain will be a very challenging adventure.

I think it is important to choose static typed stable technology first! Like Golang. Or even Java.
All those languages are self-documenting themselves to maximum and very stable in amount of changes.

We could continue thought here, and say that the usage of Docker despite it being not boring will be completely justified for long-term project.
It follows very stable standard of containers. It is made in stable tech.
It will add to our long term project important amount of self documentation for all system level dependencies and how to build it!
Very valid to use if we go for any web related project at least.
And most importantly it will freeze all system-level dependencies to saved image in Docker Registry and allow reraise our long-term project without requiring to figure out what got broken in a year or during upgrades.

### 3. Extra - Static typing.

I will repeat here that one of the most important qualities is static typing for our long-term projects
Because it allows us in a not ambiguous way to write very self documented code.
We will be able to easily reread what is going on and continue working with it.

But... there is a certain danger involved on this path, from the point that static typing needs to be stable technology too.
Like it is for Golang or Java and etc.
If we choose such hot-burning stuff like types for Javascript or Python... things are going to be potentially way more problematic.
We need at least to make a choice towards a typing system that will not disappear from a language with a high chance in future years.

### 4. Extra - Unit testing.

Typing is only part of the documentation. We can do more.
Unit tests are part of the documentation showing how your code can be used.
Checking thoroughly that it works for all aspects u will forget in a year.
Ideally, our language should have having inbuilt unit testing framework that does not require extra dependencies though.
That will help keeping unit testing especially cold blooded

### 5. Extra - Auto generated documentation

Typing, unit testing, it is all part of documentation, which can be linked into self-generated docs from the code.
Golang godocs tooling, or python sphinx autodoc stuff, it all helps us thoroughly documenting our code as a code.
That can add an additional layer of helping easier to maintain the code.
Documentation matters. In a year u will not remember any difficult enough decision u made regarding the code.
if u will be able to reread easily your previous decision and continue from that, u decrease toll further for long term maintanance.
Unit tests can be literally linked in both godocs and sphinx autodoc as examples into documentation.
U can help yourself by organizing unit tests and their file on purpose for linking to docs

Static typing, unit testing and auto generated documentation like godocs/ sphinx autodoc, they all synergize to add very protected from deprecating documentation.
Because all of it will remain part of the code and code does not lie.

U need to be careful in using comments and documenting with such tools though. Make sure to keep your documentation DRY as a code.
Don't repeat yourself in documenting efforts. it will help upkeeping it easier.
Also don't document/comment stuff that is too easily descriable with better types / rewritten code logic / provided unit test unless it is absolutely necessary.
YAGNI principle is applicable for documentation too :)

### 6. Extra - Automated periodic warming up.

The original author rejected CI usage because they can be bought or ran out of money.
I disagree with him on this point, because CI usually requires very little code effort for pet projects.
But it adds additional quality to self documentation of the project and how to build, test, and run it.
I scavenged old projects sometimes only with the help of broken, yet present CI code before.

U can help your long-term maintaned project further with assigning [cron job scheduled](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#schedule) periodic CI triggering, like once in a week or month.
That will help you tremendously in a bit warming up project and discovering little accumulated problems before they snowballed into big amount.

Yes with low amount of dependencies, it should not be big issue, but with maintaining stuff through many years, anything can break apart.
Plus as we remember we need extra help due to potentially using extra dependencies that very greatly help to keep our code simple and sane.

### 7. Extra - Sacrifice reasonably low dependencies towards simpler and type safer code.

Keeping code too boring will lead towards too much complex code. For example writing this blog site, I could have went for fully html/css only
and only my self written solutions for static building.
I made a sacrifice of low dependencies towards simpler code through the usage of [templ go](https://github.com/a-h/templ) library

<p align="center">
  <img src="{{.StaticRoot}}cold_blood/templ_go_demo.gif"/>
</p>

This library helps me to keep very simple and easily reread code that is certainly enjoyable to maintain. Also makes it very easy customizable blog for some extra features.
Evaluate if brought extra technology will help you to maintain project easier because it made the project better type safed, unit tested of auto documented, and then make the choice if the price the paid for this extra dependency helps more for long term mantaince than payment for it.

### 8. Extra - Keeping a lower amount of dependencies

As much as possible u need to withhold yourself from adding extra dependency
unless it is completely justified for one of the reasons i mentioned before or another one.
Usually custom-made solutions can be preferred for long term maintained software because they will remain with a minimal amount of code fitting exactly for your situation. They replace the need for having extra libs.
When u build commercial hot burning software it can be fine using libraries for every sneeze, because devs are able to keep up with their changes and plus they win a lot from using already made solutions over remaking frameworks on their own, but for long-term maintained stuff solution developed in house is way more preferable, unless some reason justifies using already made third party library

## Conclusion.

Those are all qualities for long-term maintained projects with big gaps between development active phases. So it should be kept in mind, as it will not fit every usage case.
If u will follow some of them for active projects, u will still have achieved decreased toll for maintenance and magnitudes faster new developer onboardings. Or at least most of the mentioned points will help towards this goal, but not all.
