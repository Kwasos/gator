package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Kwasos/gator/internal/database"

	"github.com/Kwasos/gator/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	dbQueries := database.New(db)

	appState := state{cfg: &cfg, db: dbQueries}

	cmds := commands{
		commands: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	if len(os.Args) < 2 {
		fmt.Println("missing argument")
		os.Exit(1)
	}

	commandName := os.Args[1]
	argsSlice := os.Args[2:]
	fullCommand := command{Name: commandName, Args: argsSlice}
	err = cmds.run(&appState, fullCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
