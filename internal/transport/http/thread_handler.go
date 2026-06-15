package http

import (
    "json"
    "net/http"

    "github.com/lipkerton/wildcard/internal/domain
    "github.com/lipkerton/wildcard/internal/transport/dto"
    "github.com/lipkerton/wildcard/internal/service"
)

type ThreadHandler struct {
    svc *service.ThreadService
}

func NewThreadHandler (svc *service.ThreadService) {
    return &ThreadHandler{svc: svc}
}

// Create - POST /threads
func (h *ThreadHandler) Create(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    var req dto.CreateThreadRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "некорректный JSON"})
    	return
    }

    thread, err := h.svc.Create(req.Subject, req.Body, req.Author, req.Password)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        switch {
        case errors.Is(err, domain.ErrorEmptyBody):
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        case errors.Is(err, domain.ErrorEmptyPassword):
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        default:
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]string{"error": "внутренняя ошибка сервера!"})
        }
    }
}
