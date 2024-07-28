package handler

type Handler struct{}

type NewServerOptions struct{}

func NewServer(opts NewServerOptions) *Handler {
	return &Handler{}
}
