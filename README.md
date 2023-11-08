# Mesquite
A full stack web framework in Go built for rapidly prototyping majestic monoliths

Traditional web development is overly complicated for the majority of projects. This framework offers a robust way to quickly build websites with modern ideas.

## Technologies
- Support for HTMX to make responsive user experiences
- SQLite database for speed an simplicity (Yes, this is a good idea in most cases. More info below)
- HTTP/2 to send all of the files needed for a page at once
- Written in Go for speed and ecosystem

Everything is modular so you can use a different database/router/frontend if your needs change

## Why use Mesquite over the standard library?
The standard library in Go is awesome. But you still end up writing a lot of boilerplate code to do common actions like routing, serving templates and logging. The intent of Mesquite is to make a simple, beginner friendly, fullstack framework that is fast to build with. Similar to Django and Ruby on Rails.

## Should we really be using SQLite in production?
There is a growing movement towards using SQLite for production websites. A great primer for this is Ben Johnson's [I'm All-In on Server-Side SQLite](https://fly.io/blog/all-in-on-sqlite-litestream/).
Lately, some higher profile companies like [Tailscale](https://tailscale.com/blog/database-for-2022/) and [EpicWeb.dev](https://kentcdodds.com/blog/i-migrated-from-a-postgres-cluster-to-distributed-sqlite-with-litefs) are using SQLite in production.

## How does this differ from Django and Flask?
Mesquite sits in the middle of these two frameworks in terms of supporting use cases. Ideally, Mesquite will have common use cases within the same framework and documentation to support it. But unlike Django, you can write your own code to replace pieces of Mesquite. Additionally, Mesquite follows the architecture philosophy of Flask where the user writes everything as they need it (as opposed to auto-generating code).


## Roadmap
- [ ] Routing
- [ ] Tempting
- [ ] Static files
- [ ] Config for running in dev, prod ect.
- [ ] Hot reload (Including HTML edits?)
- [ ] HTTP/2 Send multiple assets to the client during the first request? Can we get this when compiling the templates?
- [ ] SQL database wrapper (without cgo)
- [ ] User input (Forms/CSRF)
- [ ] Authentication
- [ ] Create a tutorial like the flask mega tutorial. Build a basic family social media site? (Something a lot of people could make use of) Or polls app
- [ ] Lite sail-like SQLite backup and restore
- [ ] Database migrations? (Like Django or flask-migrate?)
- [ ] Performance profiling/optimization


Ideas:
- Build step for tailwind or something similar (So NPM is not needed)
- Server side analytics and something like Django admin?
- Use ko.build? It looks neat

---

Slow is smooth, smooth is fast
