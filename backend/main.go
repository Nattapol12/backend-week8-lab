package main

import (
    "log"
    "net/http"

    "kitchen/handler"
    "kitchen/mw"
    "kitchen/store"
)

func main() {
    h := &handler.Handler{Store: store.NewMemoryStore()}

var app http.Handler = http.TimeoutHandler(h.Routes(), time.Second,
    "ครัวใช้เวลานานเกินไป")
app = mw.Logging(mw.Recovery(mw.CORS(app)))

log.Println("ครัวเปิดที่ :8080")
http.ListenAndServe(":8080", app)
}