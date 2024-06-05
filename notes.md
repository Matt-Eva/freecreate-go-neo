# Notes

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
    - HAS_TAG relationships - Creator -> Tag
    - IS_CREATOR relationships - User -> Creator
    - FOLLOWS relationships - User -> Creator
    - SUBSCRIBED relationships - User -> Creator
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

## Querying

This is how the following queries will be run.

- Querying for writing:
  - The most recent database will be targeted, and the genre
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
