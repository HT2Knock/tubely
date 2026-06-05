package main

import "net/http"

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir(s.cfg.FilePathRoot)))
	mux.Handle("/app/", appHandler)

	assetsHandler := http.StripPrefix("/assets", http.FileServer(http.Dir(s.cfg.AssetsRoot)))
	mux.Handle("/assets/", cacheMiddleware(assetsHandler))

	mux.HandleFunc("POST /api/login", s.handlerLogin)
	mux.HandleFunc("POST /api/refresh", s.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", s.handlerRevoke)

	mux.HandleFunc("POST /api/users", s.handlerUsersCreate)

	mux.HandleFunc("POST /api/videos", s.handlerVideoMetaCreate)
	mux.HandleFunc("POST /api/thumbnail_upload/{videoID}", s.handlerUploadThumbnail)
	mux.HandleFunc("POST /api/video_upload/{videoID}", s.handlerUploadVideo)
	mux.HandleFunc("GET /api/videos", s.handlerVideosRetrieve)
	mux.HandleFunc("GET /api/videos/{videoID}", s.handlerVideoGet)
	mux.HandleFunc("GET /api/thumbnails/{videoID}", s.handlerThumbnailGet)
	mux.HandleFunc("DELETE /api/videos/{videoID}", s.handlerVideoMetaDelete)

	mux.HandleFunc("POST /admin/reset", s.handlerReset)

	return mux
}
