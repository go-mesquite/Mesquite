# Mesquite
A simple full stack web library in Go

Traditional web development is more complicated than it should be for the majority of projects. This library offers a simplified way to build websites with modern ideas.

## Supported technologies
- Go for speed and portability
- HTMX for a responsive user experience
- SQLite database for speed an simplicity (Yes, this is a good idea in most cases)

Everything is modular so you can use a different database/router/frontend if your needs change

## Should we really be using SQLite in production?
There is a growing movement towards using SQLite for production websites. A great primer for this is Ben Johnson's [I'm All-In on Server-Side SQLite](https://fly.io/blog/all-in-on-sqlite-litestream/).
Lately, some higher profile companies like [Tailscale](https://tailscale.com/blog/database-for-2022/) and [EpicWeb.dev](https://kentcdodds.com/blog/i-migrated-from-a-postgres-cluster-to-distributed-sqlite-with-litefs) are using SQLite in production.

## Why use Mesquite over the standard library?
The standard library in Go is awesome. But you still end up writing a lot of boilerplate code to do common actions like routing, serving templates and logging. The intent of Mesquite is to make a simple, beginner friendy, fullstack library that is fast to build with. Similar to Django and Ruby on Rails.

## How does this differ from Django and Flask?
Mesquite sits in the middle of these two frameworks in terms of supporting use cases. Ideally, Mesquite will have common use cases within the same library and documentation to support it. But unlike Django, you can write your own code to replace pieces of Mesquite. Additionaly, Mesquite follows the architecture philosophy of Flask where the user writes everything as they need it (as opposed to auto-generating code).


## Roadmap
- [ ] Routing
- [ ] Templating
- [ ] Statfiles
- [ ] Hot reload
- [ ] SQL database (without cgo)
- [ ] User input (Forms/CSRF)
- [ ] Authentication
- [ ] Lightsail-like SQLite backup and restore
- [ ] Tailwind? (without NPM)
- [ ] API? (Make an official way to build REST APIs?)
- [ ] Something like Django admin?


---

Slow is smooth, smooth is fast
