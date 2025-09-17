package commands

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"gitlab.com/MadhurSahu/gator/internal/database"
	"gitlab.com/MadhurSahu/gator/internal/state"
)

func handlerBrowseFeeds(s *state.State, cmd Command, user database.User) error {
	limit := 2

	if len(cmd.Args) > 0 {
		if val, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = val
		}
	}

	posts, err := s.DB.GetUserPosts(context.Background(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s (%s)\n - %s\n - %s\n----------\n", post.Title, prettySince(post.PublishedAt.Time), post.Description, post.Url)
	}

	return nil
}

func prettySince(t time.Time) string {
	d := time.Since(t)
	switch {
	case d < time.Minute:
		return fmt.Sprintf("%ds ago", int(d.Seconds()))
	case d < time.Hour:
		return fmt.Sprintf("%dm ago", int(d.Minutes()))
	case d < 24*time.Hour:
		return fmt.Sprintf("%dh ago", int(d.Hours()))
	default:
		return fmt.Sprintf("%dd ago", int(d.Hours()/24))
	}
}
