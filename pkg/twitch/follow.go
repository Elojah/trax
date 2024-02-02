package twitch

type FollowFilter struct {
	FromID *string
	ToID   *string

	After *string
	First *string
}
