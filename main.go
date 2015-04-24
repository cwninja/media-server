package main
import "fmt"
import "log"
import "net/http"
import "os"
import "path/filepath"

import "github.com/gorilla/mux"

func main() {
  listen := os.Getenv("MEDIA_SERVER_LISTEN")
  mediaDir := os.Getenv("MEDIA_SERVER_MEDIA")

  if listen == "" {
    log.Fatal("MEDIA_SERVER_LISTEN must be set")
  }

  if mediaDir == "" {
    log.Fatal("MEDIA_SERVER_MEDIA must be set")
  }

  router := mux.NewRouter()

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to media-server.")
  })

  router.HandleFunc("/{key:[0-9a-f]+}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["key"]

    matches, err := filepath.Glob(filepath.Join(mediaDir, fmt.Sprintf("%s.*", key)))
    if err != nil {
      w.WriteHeader(500)
      fmt.Fprintf(w, "Error globbing files - %s", err)
      return
    }

    if len(matches) > 0 {
      http.ServeFile(w, r, matches[0])
      return
    }

    filename := filepath.Join(mediaDir, key)
    if _, err := os.Stat(filename); err == nil {
      http.ServeFile(w, r, filename)
      return
    }

    w.WriteHeader(404)
    fmt.Fprintf(w, "File not found")
  })

  http.Handle("/", router)

  log.Fatal(http.ListenAndServe(listen, nil))
}
