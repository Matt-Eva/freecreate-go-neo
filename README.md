# FreeCreate

FreeCreate is a web-based hosting platform for fiction writers. It is currently under active development.

This repository contains its backend code, written in Golang.

You can view the Frontend repository <a href="https://github.com/Matt-Eva/freecreate-react">here</a>.

## Build

FreeCreate's backend is built with Go, and currently uses Gorilla Mux for routing.

FreeCreate stores its data in the following databases:

- Neo4j - main query database.
- MongoDB - for storing long-form written content, formatted as JSON.
- Redis - query cache and session store.

While Auth is not yet implemented, FreeCreate will use server-side sessions via Gorilla Sessions and Redis, to securely store and serve user session data.

## Features

- FreeCreate's initial media hosting options including short stories, novelettes, novellas, novels, and collections and fictional universes, in which other writing types can be aggregated.

- Scalable storage of queryable content via Neo4j sharding and federation.

  - A central user database will hold user related data
  - Writing and donation based data will be stored in sharded databases, and leverage data federation to implement node / graph relationships.
  - Rationale:
    - Posted content and donations have no definitive cap - new content and new donations can be posted / given indefinitely
    - In order to anticipate and accommodate for increased scale, a strategic horizontal scaling architecture is needed
    - Complex queries are run within a single shard in the cluster - only simple queries with limited data may need to be run on multiple shards, which can execute in parallel.
  - Downsides
    - This approach introduces database query and management complexity
    - Some data duplication is necessary, which will make certain data updates slower and more fragile.
    - However, these costs are outweighed by the freedom of scaling beyond a single machine's capabilities.

- A custom-built rich text editor that outputs JSON

  - Provides basic formatting, and can include link and image embedding when the platform is ready to offer such features.
  - Content can easily be stored in MongoDB and easily parsed for rendering.

- Scalable storage of written content via MongoDB

  - JSON formatted content can easily be slotted directly into MongoDB
  - Collection can be sharded on Author name and content id, allowing quick query access and seamless scaling.

- Redis query cache

  - Frequently accessed data will be stored in Redis cache, accelerating response times and freeing Noe4j from running complex queries as frequently.

- While donations are not yet enabled, future payment processing will be handled by Stripe.
