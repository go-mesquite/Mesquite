# Mesquite
A full stack web framework in Go built for rapidly prototyping HTMX monoliths

Traditional web development is overly complicated for the majority of projects. This framework offers a robust way to quickly build reliable websites.

> [!WARNING]  
> This framework is under heavy development and is not ready for production

## Technologies
- Built for HTMX
- A local database that natively supports Go data types
- HTTP/2 to send all of the files needed for a page at once
- Written in Go for speed and ecosystem

Everything is modular so you can use a different database/router/frontend if your needs change

Mesquite is built with the architecture of MVCS: model view controller service [The Model View Controller Pattern – MVC Architecture and Frameworks Explained](https://www.freecodecamp.org/news/the-model-view-controller-pattern-mvc-architecture-and-frameworks-explained/)

### Why use Mesquite over the standard library?
The standard library in Go is awesome. But you still end up writing a lot of boilerplate code to do common actions like routing, serving templates and logging. The intent of Mesquite is to make a simple, beginner friendly, fullstack framework that is fast to build with. Similar to Django and Ruby on Rails.

### Should we really be using SQLite in production?
There is a growing movement towards using SQLite for production websites. A great primer for this is Ben Johnson's [I'm All-In on Server-Side SQLite](https://fly.io/blog/all-in-on-sqlite-litestream/).
Lately, some higher profile companies like [Tailscale](https://tailscale.com/blog/database-for-2022/) and [EpicWeb.dev](https://kentcdodds.com/blog/i-migrated-from-a-postgres-cluster-to-distributed-sqlite-with-litefs) are using SQLite in production.

### How does this differ from Django and Flask?
Mesquite sits in the middle of these two frameworks in terms of supporting use cases. Ideally, Mesquite will have common use cases within the same framework and documentation to support it. But unlike Django, you can write your own code to replace pieces of Mesquite. Additionally, Mesquite follows the architecture philosophy of Flask where the user writes everything as they need it (as opposed to auto-generating code).

## Design decisions
### Routing
Mesquite has it's own basic router to reduce dependencies.
It is compatible with standard `net.http` so it can be swapped for others like `gorilla/mux`


## Roadmap
- [x] Routing
- [ ] Templates (Or should I call them views?)
- [ ] Static files
- [ ] Config for running in dev, prod ect.
- [ ] Reload on save
- [ ] User input (Forms/CSRF)
- [ ] Authentication
- [ ] Add the functionality for multiple routers and reverse lookups. Like Flask's url_for()
- [ ] Create a tutorial like the flask mega tutorial. Build a basic family social media site? (Something a lot of people could make use of) Or polls app
- [ ] Performance profiling/optimization

---

Inspired in by: https://paulgraham.com/avg.html

---

### Notes for Jackson since this is still very rough
Embrace DDD with the model layer?:
- Each Model has it's own file in /models. These are used as objects with Structs and have getters and setters.
- Each object/struct is part of the domain. They may correspond to one or more database tables.
- This eliminates the need for a service/utility layer that is hard to enforce other developers to use
- How should we deal with actions that could be optimized? Advocating for a bulk create instead of running a single create SQL statement over and over? (From the controller perspective)
- Conventional wisdom says that business logic should not be in the database layer. But I think it makes sense in this case
- Or use MVCD. Add a domain layer instead of a service layer. Models would just be CREATE TABLEs and migrations
- Use https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/ to kep files in the binary?