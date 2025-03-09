package routes

import (
    "encoding/json"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
    "backend/controllers"
    "backend/database"
)

func SetupRouter() http.Handler {
    // Charge les variables d'environnement
    database.LoadEnv()
    allowedOrigins := strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
    mux := http.NewServeMux()

    // Route racine "/"
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }

        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Bienvenue sur l'API Moyens Logistiques !",
        })
    })

    // Routes "/categories" et "/categories/{id}/..."
    mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/categories" && r.Method == http.MethodGet {
            controllers.GetCategories(w, r)
            return
        }
        http.NotFound(w, r)
    })

    mux.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
        parts := strings.Split(r.URL.Path, "/")
        if len(parts) == 4 && r.Method == http.MethodGet {
            switch parts[3] {
            case "accessoires_defauts":
                controllers.GetAccessoiresDefauts(w, r)
                return
            case "accessoires":
                controllers.GetAccessoiresForCategorie(w, r)
                return
            }
        }

        if len(parts) == 3 && r.Method == http.MethodGet {
            controllers.GetCategorie(w, r)
            return
        }

        http.NotFound(w, r)
    })

    // Routes "/moyens"
    mux.HandleFunc("/moyens", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/moyens" && r.Method == http.MethodGet {
            controllers.GetMoyens(w, r)
            return
        }
    })

    // Routes "/moyens/{id}" et "/moyens/{id}/accessoires"
    mux.HandleFunc("/moyens/", func(w http.ResponseWriter, r *http.Request) {
        parts := strings.Split(r.URL.Path, "/")
        if len(parts) == 3 && r.Method == http.MethodGet {
            controllers.GetMoyen(w, r)
            return
        }
        if len(parts) == 4 && parts[3] == "accessoires" && r.Method == http.MethodGet {
            controllers.GetAccessoiresForMoyen(w, r)
            return
        }
        http.NotFound(w, r)
    })

    // Routes "/accessoires" et "/accessoires/{id}"
    mux.HandleFunc("/accessoires", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/accessoires" && r.Method == http.MethodGet {
            controllers.GetAccessoires(w, r)
            return
        }
    })

    mux.HandleFunc("/accessoires/", func(w http.ResponseWriter, r *http.Request) {
        parts := strings.Split(r.URL.Path, "/")
        if len(parts) == 3 && r.Method == http.MethodGet {
            controllers.GetAccessoire(w, r)
            return
        }
        http.NotFound(w, r)
    })

    // ✅ Appliquer `corsMiddleware` globalement à toutes les routes
    return corsMiddleware(allowedOrigins, mux.ServeHTTP)
}


// Middleware CORS personnalisé
func corsMiddleware(allowedOrigins []string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        origin := r.Header.Get("Origin")
        allowed := false

        origin = strings.TrimSpace(origin)
        for _, o := range allowedOrigins {
            if strings.TrimSpace(o) == origin {
                allowed = true
                break
            }
        }

        if allowed {
            w.Header().Set("Access-Control-Allow-Origin", origin)
            w.Header().Set("Access-Control-Allow-Methods", "GET")
            w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
            w.Header().Set("Access-Control-Max-Age", strconv.Itoa(int((12*time.Hour).Seconds())))
        }

        // Réponse immédiate en cas de requête OPTIONS
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next(w, r)
    }
}