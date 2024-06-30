# Notes

## Dependecies

### Databases

- Neo4j - primary query database
- MongoDB - bulk writing storage / potential cache
- Redis - session cache / query cache

## Deployment Strategy

Just use AWS, dude. It's going to cost money regardless, and you don't have your own servers. Plus, Neo4j is available on AWS, and hosting on AWS will give you good industry experience.

- Neo4j, MongoDB, and Redis will be deployed on Amazon EC2 instances
- Backend API will be hosted on an EC2 instance as well
- Frontend assets can be delivered by Netlify or Cloudfront
- Media assets can be stored in S3

### Payment processing

- Stripe
  - Supports Apple Pay and Google Pay
- (Paypal?)

## Neo4j Internal Structure

- <a href="https://neo4j.com/docs/operations-manual/current/tutorial/tutorial-composite-database/#tutorial-composite-database-get-results">Composite Database Query Docs</a>

- Neo4j will be set up as a composite database
- There will be an unsharded user-centric database that will be federated as necessary
  - Data stored in this user database will include:
    - User nodes
    - Creator nodes
    - Genre nodes
    - Tag nodes
    - Bookshelf nodes
    - HAS_WRIT_GENRE relationships - Creator -> Genre
    - HAS_WRIT_TAG relationships - Creator -> Tag
    - IS_CREATOR relationships - User -> Creator
    - FOLLOWS relationships - User -> Creator
    - SUBSCRIBED relationships - User -> Creator
    - BLOCKED relationship - User -> Creator
      - This one will be tricky to model. Because the query we want to run will be
      ```
      MATCH (u:User {uId: $userId})
      MATCH (w:Writing) [...]
      WHERE NOT (w) - [:CREATED_BY] -> (c:Creator) <- [:BLOCKED] - (u)
      ```
      - The trouble with this is that we are going to be running these queries on the sharded datasets, but this core relationship will be specified in the user database
      - One solution to this is to just replicate the blocked creators whenever a new user node is added to a database
      - This would entail getting the blocked creators from the main user database, then MERGEing them into the new database as needed.
      - It's possible that this would require round trips to and from the database - first to check if the user node exists in the sharded database - if they don't, run the query.
      - This could also simply be handled client side by loading a users blocked creators, then checking if a piece of content has that creator id and name using a hash table.
    - AUTHOR_IN_LIB relationships - User -> Creator
    - HAS_BOOKSHELF relationships - User -> Bookshelf
- There will be a sharded content-centric database into which user and creator information will be federated
  - Data stored in this database will include:
    - Writing nodes
    - Tag nodes
    - Donation nodes
    - Federated User nodes
    - Federated Creator nodes
      - Any Creator node added will also have the parent user node added
    - Federated Bookshelf nodes
    - CREATED_BY relationships - Creator -> Writing
    - IS_CREATOR relationships - User -> Creator
    - DONATED_TO_CREATOR relationships - User -> Creator
    - DONATED_TO_WRIT relationships - User -> Writing
    - HAS_TAG relationships - Writing -> Tag
    - LIKED relationships - User -> Writing
    - WRIT_IN_LIB - User -> Writing
    - ON_BOOKSHELF - Bookshelf -> Writing
- There will be a single ALL TIME database in which content that reaches above a certain absolute rank will be included
  - It will basically be the same as any of the sharded content-centric databases, except there is only one of it and it's for all time content

### Novels, Series, and Collections

- In each of these media types, many different pieces of writing will be connected together. These pieces of writing could be posted in different years, and sharded across different time periods.
- Novels:
  - A novel will have many chapters, which could be published across multiple years
  - We want to search for novels by their update date, rather than just their writing date.
  - Whenever a novel is updated in a new year, we add its node to that new year.
    - Do we duplicate the data across all the years in which its present?
    - That would be the best way to preserve query integrity
    - No, actually, we want the "year" of the novel to be the year that it is finally finished.
    - So, we'll copy the novel to the new year, and delete the old node.
    - We will need to connect the novel to all of its old tags in the new db.
  - we don't actually have to store any of the chapters in neo - we can store them all in Mongo
- Series:
  - A series would be a series of novels, or something that exists within an overarching universe or world.
  - We could have a separate series db that is a single instance db
  - This could also be where we store collections
  - There could also be fictional universe nodes that a series or collection could belong to.

## Querying

This is how the following queries will be run.

- Querying for writing:
  - The most recent database will be targeted, unless a prior year is specified.
- Loading Donations:
  - For a creator user:
    - Will query the central database for all of the user's creator profiles
    - Users will be able to view donations by year (which is how the database is sharded)
    - Will then use those creator profiles to query the past year's database
    - Will then load all of the donations for each of those creators, with a Limit
      - Each donation will show who the donation is from (nickname and username)
    - Will also present a total of all donations
    - Shows them their current earnings of the year as well as a breakdown of the donations
  - For a user:
    - Will query the specific year's database for donations from the user.
    - Will also load which creation and creator the donation was given to.
- Loading Like Writing:
  - Will query the latest year's database first, then query previous databases to load content
- Loading Library:
  - The user will have a variety of options regarding how to view their library
  - The will need the following resources:
    - All creators in library
    - All writing in library
    - All bookshelves in library
  - They can organize their library by
    - Author
    - Bookshelf
    - Writing Type
    - Alphabetical
  - They can also search their library
  - Query:
    - Simple version:
      - When a user visits any portion of their library, their whole library is loaded.
      - This would load all bookshelves and creators from the central user db
      - This would load each piece of writing and its associated creator and bookshelf IDs from across all sharded databases
      - Once all this data is loaded, the view update logic will be handled client side.
      - There is a high potential for caching this data, either in MongoDB or Redis.
        - This cache could be updated each time a user adds or removes something from their library.
          - This approach could result in stale data, but also only updates library caches when absolutely necessary.
