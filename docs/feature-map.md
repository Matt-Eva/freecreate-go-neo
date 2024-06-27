# Feature-Map

[] - Connect to databases
[] - Seed databases (including cache)
[] - Implement Search (including cache)
[] - Create User
[] - User Auth
[] - Edit User
[] - Delete User
[] - Create Creator
[] - Edit Creator
[] - Delete Creator
[] - Create Writing
[] - Edit Writing
[] - Delete Writing
[] - Like Writing
[] - Add Writing to Library
[] - Add Writing to Reading List
[] - Add Writing to Bookshelf
[] - Follow Creator
[] - Subscribe to Creator
[] - Save Reading (allow reader to pick up from where left off)

## Search

Information input will be as follows:

- Writing or Writer?
- Writing:
  - Genres
  - Writing Type
  - Date Posted
  - Tags
  - Name
- Writer:
  - Genres
  - Tags
  - Name

Logical flow will be as follows:

- Writing or Writer:
  - Writing: Writing Search
    - No Tags or Title, Date Posted not Most Recent or Specific Year: Cache Search
      - Query Redis Cache for Cached search results
    - Tags, Title, or Date Posted Most Recent or Specific Year: Neo Search
      - Date Posted Most Recent
        - Query current year database, order by date
      - Date Posted All Time
        - Query all time database, order by rank and rel rank
      - Date Posted Specific Year
        - Query specific year database, order by rank and rel rank
      - None of above
        - Query current year database, order by rank and rel rank
