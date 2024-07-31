# Notes

## Databases

- Neo4j - primary query database
- MongoDB - bulk writing storage / potential cache
- Redis - session cache / query cache

## Deployment Strategy

Just use AWS, dude. It's going to cost money regardless, and you don't have your own servers. Plus, Neo4j is available on AWS, and hosting on AWS will give you good industry experience.

- Neo4j, MongoDB, and Redis will be deployed on Amazon EC2 instances
- Backend API will be hosted on an EC2 instance as well
- Frontend assets can be delivered by Netlify or Cloudfront
- Media assets can be stored in S3

## Payment processing

- Stripe
  - Supports Apple Pay and Google Pay
- (Paypal?)

## Neo4j Internal Structure

### Non-Sharded architecture

All of these problems go away if we simply don't shard Neo4j.

If we use an Amazon Ec2 i4i.metal instance, we get almost a terabyte of memory and potentially up to 30 TB of ssd storage.

So, I think it's best to ditch this sharding strategy and just focus on leveraging the full potential of Neo4j.

However, we should still have an all-time database for all-time queries. Otherwise we'll simply have too much data to comb through.

### Sharded Architecture

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
    - problem - we would need to copy all donation nodes and liked relationships to the new db as well. This could be millions.
  - we don't actually have to store any of the chapters in neo - we can store them all in Mongo
- Series:
  - A series would be a series of novels, or something that exists within an overarching universe or world.
  - We could have a separate series db that is a single instance db
  - This could also be where we store collections
  - There could also be fictional universe nodes that a series or collection could belong to.

### Donations

- There is a problem with this current model - if a specific piece of content is relegated to a specific year, we won't be able to get the most recent donations for that piece of content
- We might instead need to create a separate donations db where we can track donations given to creators / users, specific pieces of content, and who donated them.
- Ok, we have a problem
  - We want a creator to be able to quickly and easily see how many donations they have received for a specific piece of content
  - We also want a user to be able to quickly and easily see who they have donated to, and how much.
  - This means we want creator nodes, creation nodes, and user nodes, to all be connected to a single donation.
  - However, we are sorting writing nodes by the date they were posted, and a donation could be given to a piece of writing in a year later than it was posted.
  - But, we want to donations to be presented in the order in which they have been given - so we want them organized by latest year as well.
  - If we store donations in their own independent database, and just federate creator, user, and creation nodes
    based on id, we could easily see an individual donation as a line by line item.

### Querying

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

## Rank Indexing

In order to improve the performance of order_by queries, we want to have indexes for rank and relative rank.
However, these are updated every time there is a new like, new read, new donation, etc. This high frequency update
of an indexed field could be quite detrimental to the platform, as neo4j would basically be constantly updating the index.

So, we want to instead update the rank and rel_rank fields of an item periodically.

When updating the rank or rel_rank of a node, we want to put a lock on the affected node.

There are a few approaches for doing this - setting a time window, or setting a number of updates.

We can actually do both

WHERE w.updateTime < timeFrame OR w.rankIncrease < rankIncreaseFrame

We also want to check when the item was posted - things that are posted more recently should be updated more frequently.

For most recent items, we should set the rankChange to 100?

post < day: rankChange === 1000, timeFrame == 30 seconds // 1 donation
day < posted < week : rankChange === 10000, timeFrame == 1 minute // 10 donations
week < posted < month: rankChange === 100000, timeFrame == 5 minutes // 100 donations
month < posted: rankChange === 100000, timeFrame == 10 minutes // 1000 donations

Can apply `apoc.lock()` to lock node for writes. <a href="https://claude.ai/chat/278ba8ac-a132-4e9d-99b6-ea80b117e6f7">reference</a>

## Caching

It could actually be easier to just use MongoDB for the query cache. We can set the "genre" to be the combo of specific keys
and can create a composite index with the time frame.
We could then simply store the cached queries as lists in the document.

## Mongo Sharding

## Chapters & Draft_Chapters

Shard key: creator name / creator id / title / neo id

Overarching

## Weighting

Read - 1
Like - 10
List - 10
Library - 100
Donation - 1000

## Authentication
