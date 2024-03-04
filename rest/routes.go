package rest

import "net/http"

func addRoutes(
	mux *http.ServeMux,
	logger Logger,
	//config Config,
	// tenantsStore *TenantsStore,
	// commentsStore *CommentsStore,
	// conversationService *ConversationService,
	// chatGPTService *ChatGPTService,
	// authProxy *authProxy,
) {
	//mux.Handle("/api/v1/", handleTenantsGet(logger, tenantsStore))
	//mux.Handle("/oauth2/", handleOAuth2Proxy(logger, authProxy))
	mux.HandleFunc("GET /healthz", healthzHandler)
	mux.Handle("GET /", http.NotFoundHandler())
	mux.Handle("POST /upload", HandleFileReceiver())
}
