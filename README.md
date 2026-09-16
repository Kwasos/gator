# Gator

Gator is a command-line RSS feed aggregator written in Go. It lets you register 
users, follow RSS feeds, periodically scrape new posts, and browse them from the terminal.

## Requirements

- [Go](https://go.dev/) (1.21+ recommended)
- [PostgreSQL](https://www.postgresql.org/)

## Installation

Install the CLI using `go install`:

\`\`\`bash
go install github.com/Kwasos/gator@latest
\`\`\`

## Configuration

Gator reads a config file located at `~/.gatorconfig.json`. Example:

\`\`\`json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
\`\`\`

Run the database migrations (using something like `goose`) before starting.

## Commands

| Command | Description |
|---|---|
| `register <name>` | Register a new user |
| `login <name>` | Log in as an existing user |
| `reset` | Reset the database (danger, this is a complete wipe!) |
| `users` | List all registered users |
| `addfeed <name> <url>` | Add a new RSS feed and follow it |
| `feeds` | List all feeds |
| `follow <url>` | Follow an existing feed |
| `following` | List feeds the current user follows |
| `unfollow <url>` | Unfollow a feed |
| `agg <duration>` | Continuously scrape feeds at the given interval (e.g. `1m`, `30s`) |
| `browse [limit]` | Browse recent posts from followed feeds (default limit: 2) |

## Example Usage

\`\`\`bash
gator register alice
gator login alice
gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
gator agg 1m
gator browse 5
\`\`\`

## License

MIT
