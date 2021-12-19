package route

type Router interface {
	SourceUrl(state string) string
	LoginUrl(state string) string
}
