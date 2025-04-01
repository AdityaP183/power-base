### **HeroHub - To-Do List**

#### ✅ **1. Project Setup**
- [x] Create a new Go module (`go mod init herohub`).
- [x] Define folder structure: `/cmd`, `/internal`, `/api`, `/docs`, `/web`.

---

#### 🔨 **2. Backend (Golang API)**
- [x] **Initialize Server (main.go in `/cmd/` folder)**
  - [x] Set up HTTP router (using `net/http` or `mux`).

- [x] **Define Models (in `/internal/models/hero.go`)**
  - [x] `Hero`
  - [x] `PowerStats`
  - [x] `Appearance`
  - [x] `Biography`
  - [x] `Connections`

- [ ] **Business Logic Layer (in `/internal/core/hero.go`)**
  - [ ] Functions to handle hero data.
  - [ ] Functions for CRUD operations.

- [ ] **Storage Layer (in `/internal/storage/datastore.go`)**
  - [ ] Load & save data (CSV/JSON/Database).

- [ ] **API Handlers (in `/api/handlers/`)**
  - [ ] `hero.go` (CRUD operations for heroes).
  - [ ] `search.go` (Search functionality).

- [ ] **Middleware (in `/api/middleware/`)
  - [ ] Logging Middleware (`logger.go`).
  - [ ] Authentication Middleware (if needed).

---

#### 🔖 **3. Documentation (in `/docs/`)**
- [ ] Create OpenAPI Specification (`openapi.yaml`).
- [ ] Write Markdown Documentation (`README.md`).
- [ ] Generate Swagger Docs (if using).

---

#### 🌐 **4. Frontend & Docs Website (in `/web/`)**
- [ ] **Static Site Setup**
  - [ ] Create `index.html` (Homepage for documentation & data browsing).
  - [ ] Add `static/` folder for CSS, JS, images.
  - [ ] Add `templates/` for server-side rendering (if necessary).

- [ ] **API Documentation Integration**
  - [ ] Display data from API (`/api/handlers/hero.go`).
  - [ ] Generate UI for browsing heroes.

---

#### 📦 **5. Testing & Deployment**
- [ ] Write Unit Tests (`*_test.go` files) for all handlers & core logic.
- [ ] Set up CI/CD (GitHub Actions or similar).
- [ ] Deploy Backend (e.g., Render, Railway, or your preferred service).
- [ ] Deploy Frontend (Static site hosting like Vercel, Netlify, etc.).


