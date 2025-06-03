# 🏦 GoBank

> *"The only way to make sense out of change is to plunge into it, move with it, and join the dance."* - Alan Watts

Welcome to **GoBank** - where your money is safer than a squirrel's nuts in winter! 🐿️ This is my personal playground for mastering Go backend development, because let's face it, someone's gotta keep those zeros and ones in line.

[![Go Version](https://img.shields.io/badge/Go-1.24.2-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![Fiber](https://img.shields.io/badge/Fiber-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://gofiber.io/)
[![Test Coverage](https://img.shields.io/badge/Coverage-90%25-brightgreen?style=for-the-badge)](https://github.com/fr13nd230/gobank)

## 🎯 What's This All About?

GoBank is my humble attempt at building a banking system in Go - think of it as a digital piggy bank, but with more code and fewer coins rattling around. The main goal? To get cozy with the **Fiber framework** and prove that Go can handle financial shenanigans better than my actual bank handles my overdraft fees! 💸

*"Code is like humor. When you have to explain it, it's bad."* - Cory House

## 🛠️ Tech Stack

- **Go 1.24.2** - Because life's too short for slow languages
- **Fiber Framework** - Fast, Express-inspired web framework (coming soon to a codebase near you!)
- **PostgreSQL** - Where your data lives rent-free
- **Docker** - Containerization magic ✨
- **SQLC** - Type-safe SQL generation
- **golang-migrate** - Database migration wizardry

## 📦 Dependencies

```go
// The usual suspects
github.com/google/uuid v1.6.0          // For unique IDs (like snowflakes, but digital)
github.com/jackc/pgx/v5 v5.7.5         // PostgreSQL driver that actually works
github.com/joho/godotenv v1.5.1        // Because hardcoding secrets is for amateurs
github.com/stretchr/testify v1.10.0    // Testing made bearable
```

**Note:** Fiber framework will join the party once I stop procrastinating and actually install it! 🎉

## 🚀 Getting Started

### Prerequisites

Make sure you have these installed, or this journey ends before it begins:

- **Go 1.24.2+** - [Download here](https://golang.org/dl/)
- **Docker & Docker Compose** - [Get Docker](https://www.docker.com/get-started)
- **golang-migrate** - Install with: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

### Installation

1. **Clone this masterpiece:**
   ```bash
   git clone https://github.com/fr13nd230/gobank.git
   cd gobank
   ```

2. **Set up your environment:**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials (no, 'password123' is not secure)
   ```

3. **Fire up the containers:**
   ```bash
   make cmp-up
   ```

4. **Run migrations:**
   ```bash
   make migrate-up
   ```

5. **Start the server:**
   ```bash
   make run
   ```

*"It's not a bug, it's a feature!"* - Every developer ever

## 🎮 Available Commands

Our Makefile is like a Swiss Army knife, but for code:

| Command | Description | Fun Factor |
|---------|-------------|------------|
| `make build` | Build the application | 🔨 Hammer time! |
| `make test` | Run tests with coverage | 🧪 Science! |
| `make run` | Start the server | 🚀 Blast off! |
| `make cmp-up` | Start Docker containers | 🐳 Wake the whale! |
| `make cmp-down` | Stop Docker containers | 😴 Nap time! |
| `make createdb` | Create database | 🗄️ Birth of data |
| `make dropdb` | Drop database | 💥 Digital demolition |
| `make connectdb` | Connect to database | 🔌 Phone home |
| `make migrate-up` | Run migrations up | ⬆️ Leveling up |
| `make migrate-down` | Run migrations down | ⬇️ Going backward |
| `make fixdirty` | Fix dirty migrations | 🧹 Spring cleaning |
| `make gensqlc` | Generate SQLC code | 🎭 Code magic |

## 🤝 Contributing

Contributing to GoBank is easier than explaining why you need another banking app:

### Branch Naming Convention
- `feature/awesome-new-thing` - For new features that'll blow minds
- `fix/that-annoying-bug` - For squashing bugs like a pro
- `config/environment-setup` - For configuration changes
- `test/unit-coverage` - For test improvements

### Commit Message Format
Follow this pattern or face the wrath of inconsistent git history:

```
GB | purpose(scope): Your brilliant message here

Examples:
GB | feat(auth): Add user authentication with JWT tokens
GB | fix(database): Resolve connection pool leak issue  
GB | test(unit): Creating unit tests for account service
GB | config(docker): Update PostgreSQL version in compose file
```

### The Sacred Rules

1. **Always open a Pull Request** - No direct pushes to main (we're civilized here)
2. **Unit Tests are MANDATORY** - If it's not tested, it doesn't exist
3. **Minimum 90% test coverage** - Because 89% is for quitters
4. **Code reviews are your friend** - Two pairs of eyes are better than one

*"Any fool can write code that a computer can understand. Good programmers write code that humans can understand."* - Martin Fowler

## 🧪 Testing

We take testing seriously here (unlike my diet):

```bash
make test
```

This will run all tests with coverage reporting. If you're not hitting 90% coverage, you're not trying hard enough! 📈

## 🐳 Docker Support

Because "it works on my machine" is not a valid deployment strategy:

- **Development**: `make cmp-up`
- **Cleanup**: `make cmp-down`

## 🗄️ Database

PostgreSQL is our database of choice because:
- It's reliable (unlike my sleep schedule)
- It handles complex queries (unlike my brain on Monday mornings)
- It's open source (like my admiration for good code)

## 📄 License

This project is licensed under the MIT License - because sharing is caring, and lawyers are expensive.

## 🎯 Final Words

*"The best time to plant a tree was 20 years ago. The second best time is now."* - Chinese Proverb

The same goes for learning Go - start now, thank yourself later!

---

**Built with ❤️ and probably too much coffee by [fr13nd230](https://github.com/fr13nd230)**

*Remember: In code we trust, but we test everything else!* 🚀
