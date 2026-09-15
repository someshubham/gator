# Gator

A small command-line RSS aggregator and feed manager written in Go.

Features

- Register and manage local users
- Add and list feeds
- Follow/unfollow feeds for a user
- Periodic aggregator that fetches feeds and stores posts in a PostgreSQL database
- Uses sqlc-generated database queries (see internal/database and sql/schema)

## Quickstart

Requirements

- Go 1.26+
- PostgreSQL
- (Optional, for regenerating DB code) sqlc

1. Prepare the database

- Create a PostgreSQL database for Gator, e.g.:

  createdb gator

- Apply the schema files located in sql/schema in order. Example using psql:

  psql -d gator -f sql/schema/001_users.sql
  psql -d gator -f sql/schema/002_feeds.sql
  psql -d gator -f sql/schema/003_feed_follows.sql
  psql -d gator -f sql/schema/004_add_fetched_at_feeds.sql
  psql -d gator -f sql/schema/005_posts.sql

2. Create a config file
   Gator expects a configuration file named .gatorconfig.json in the current user's home directory.
   Create it with at least the db_url field. Example (~/.gatorconfig.json):

{
"db_url": "postgresql://gatoruser:password@localhost:5432/gator?sslmode=disable",
"current_user_name": ""
}

When you register a user, the CLI will update current_user_name for convenience.

3. Build

cd /path/to/gator
go build -o gator

Or install with module-aware install:

go install github.com/someshubham/gator@latest

4. Common commands / Usage

The CLI expects a command name as the first argument. Some available commands:

- register <name>
  Create a new user and set them as the current user.

- login <name>
  Set the current user (reads from the DB).

- reset
  Clears the DB (for development/testing).

- users
  List all users.

- addfeed <name> <url>
  Add a feed under the current user and automatically follow it.

- feeds
  List all feeds and their owner.

- follow <feed-url>
  Follow a feed for the current user by feed URL.

- unfollow <feed-url>
  Unfollow a feed for the current user.

- following
  List feeds followed by the current user.

- browse
  Browse posts for the current user's followed feeds.

- agg
  Start the aggregator loop that periodically fetches the next feed and stores posts. This runs an infinite loop (ticker based).

Example workflow

# build and run

go build -o gator
./gator register alice
./gator addfeed "Example Blog" "https://example.com/rss"
./gator feeds
./gator agg # run the aggregator in a terminal or background process

## Development notes

- This project uses sqlc (sqlc.yaml present) to generate type-safe DB access. Generated files are checked in under internal/database. To regenerate, install sqlc and run:

  sqlc generate

- Database queries and schema are under the sql/ directory.
- The configuration loader expects a JSON file at ~/.gatorconfig.json. The Config struct is in internal/config.
- The aggregator periodically selects the next feed to fetch using GetNextFeedToFetch and marks feeds as fetched by updating fetched_at.

## Testing

There are no automated tests included in the repository at the moment. For manual testing, use a local PostgreSQL instance and the workflow above.

## Contributing

Contributions are welcome. Please open issues or pull requests. If adding features that modify the database schema, add a new SQL file to sql/schema and update any sqlc mappings as needed.

## License

No license file included. If publishing this project, add a LICENSE file to indicate terms.

## Contact

For questions or contributions, open an issue or pull request on the repository.
