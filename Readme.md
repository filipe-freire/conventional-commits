# Conventional Commits CLI!

Long in the past are the days of chaotic commits! With this CLI tool, you can easily pick from a list what kind of change you are making & provide your commit message easily ✨

Here are all the commit options available:

- 💎 feat: (ex: a complete and fully fledged feature)
- 🎁 add: (ex: little additions like installing a package, small improvements, tests)
- 🆙 update: (ex: updating a package to a certain version)
- 🔧 ref: (ex: when executing refactors)
- ⏳ wip: (ex: short for 'work in progress', when your change isn't quite there yet, but good enough to commit)
- 🔥 delete: (ex: when getting rid of files, packages, tests, redundant code, comments...)
- 🧹 chore: (ex: fixing linting errors, updating docs, migrating from an API to another due to breaking changes..)
- 🐛 bugfix: (ex: when you just fixed that nasty bug breaking production, you rockstar!)

## Motivation

The CLI is written in Typescript (through Deno), Go and Rust for 2 main reasons:

- To have a feel for how simple/complex they'd be to write depending on the language chosen;
- To compare the binary size of all the 3 languages, and more easily weigh the drawbacks of one VS the other. Because shipping roughly 90mb just for a CLI _may_ be overkill (looking at you Deno 😉)
