package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/someshubham/gator/internal/database"
)

func handlerBrowse(s *state, c command, usr database.User) error {
	var limit int
	if len(c.args) == 0 {
		limit = 2
	}
	limit, err := strconv.Atoi(c.args[0])
	if err != nil {
		return err
	}

	posts, err := s.db.GePostsForUser(context.Background(), database.GePostsForUserParams{
		UserID: usr.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title.String)
	}

	return nil
}
