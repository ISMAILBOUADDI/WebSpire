# **WebSpire Framework**

WebSpire is a lightweight, extensible web framework for Go that provides dynamic routing, named routes, route grouping, and middleware support. Designed for simplicity and scalability, WebSpire is ideal for building web applications with custom logic.

---

## **Features**

- **Dynamic Routes**: Use placeholders in routes to capture and process URL parameters.
- **Named Routes**: Generate URLs programmatically using route names.
- **Route Groups**: Organize related routes under a common prefix.
- **Middleware Support**: Apply global or route-specific middleware for tasks like authentication, logging, or preprocessing.

---

## **Installation**

Clone the repository and include it in your Go project:

```bash
git clone https://github.com/ISMAILBOUADDI/WebSpire.git
cd WebSpire
```

---

## **Usage**

### **Basic Routing**

Define routes with WebSpire's `AddRoute` method:

```go
router.AddRoute("GET", "/", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
    w.Write([]byte("Welcome to WebSpire!"))
}, "home")
```

### **Dynamic Routes**

Capture URL parameters dynamically:

```go
router.AddRoute("GET", "/user/{id}", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
    w.Write([]byte("User ID: " + params["id"]))
}, "user.show")
```

### **Route Groups**

Group related routes under a common prefix:

```go
router.Group("/admin", func(router *core.Router) {
    router.AddRoute("GET", "/dashboard", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
        w.Write([]byte("Admin Dashboard"))
    }, "admin.dashboard")
})
```

### **Middleware**

Add global or route-specific middleware:

```go
// Global middleware
router.Use(func(w http.ResponseWriter, r *http.Request, next func(http.ResponseWriter, *http.Request)) {
    fmt.Println("Global Middleware")
    next(w, r)
})

// Route-specific middleware
authMiddleware := func(w http.ResponseWriter, r *http.Request, next func(http.ResponseWriter, *http.Request)) {
    if r.Header.Get("Authorization") != "Bearer secret" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    next(w, r)
}

router.AddRoute("GET", "/dashboard", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
    w.Write([]byte("Dashboard"))
}, "dashboard", authMiddleware)
```

---

## **Testing**

1. Run the server:

   ```bash
   go run main.go
   ```

2. Visit the routes in your browser or use tools like Postman or curl:
    - `/` for the welcome page.
    - `/user/{id}` for dynamic routes.
    - `/dashboard` for protected routes (requires `Authorization: Bearer secret` header).

---

## **Future Enhancements**

1. **Rate Limiting**: Implement middleware to limit the number of requests per client.
2. **Request Validation**: Add input validation middleware for cleaner request handling.
3. **Response Compression**: Include middleware for gzip compression of responses.
4. **Error Handling Middleware**: Provide centralized error handling.
5. **View Rendering**: Integrate template engines for dynamic HTML rendering.
6. **WebSocket Support**: Extend support for WebSocket connections.

---

## **Roadmap**

1. **Database Integration**: Add support for database drivers (MySQL, PostgreSQL, SQLite).
2. **CLI Tooling**: Provide a command-line tool for scaffolding routes, middleware, and models.
3. **Session Management**: Include middleware for managing user sessions securely.
4. **Dependency Injection**: Simplify application structure with built-in DI containers.
5. **API Versioning**: Facilitate backward compatibility with API versioning support.

---

## **Contributing**

Feel free to fork this repository and submit pull requests with improvements or new features.