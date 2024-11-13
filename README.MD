# Project Technical Architecture

This project demonstrates a modern Go web application using:
- [chi router](https://github.com/go-chi/chi) for HTTP routing
- [templ](https://github.com/a-h/templ) for type-safe HTML templates
- [tailwindcss](https://tailwindcss.com/) for styling

## Key Features

- Hot reload development workflow
- Type-safe HTML templates with templ
- Modern CSS with tailwindcss
- Clean architecture with chi router
- Parallel process management with Make

## Development Setup

### Prerequisites

- Go 1.21+
- Node.js & npm
- iTerm2 (for split pane development view)
- Make

### Project Structure

.
├── cmd/
│ └── server/ # Main application entry
├── internal/
│ ├── handlers/ # HTTP handlers
│ └── views/ # Templ templates
├── public/ # Static assets
├── tailwind.config.js # Tailwind configuration
└── Makefile # Development automation

### Development Workflow

The project uses a sophisticated development setup managed through `make dev` command that:

1. Launches 4 parallel processes in split iTerm2 panes:
   - Go server with hot reload (using `air`)
   - Templ template watcher & compiler
   - Tailwind CSS watcher & compiler
   - Browser-sync for automatic browser refresh

### Makefile Implementation

The `Makefile` implements parallel process management:


makefile
dev:
# Launch iTerm2 split panes
osascript ./scripts/iterm.applescript
# Run processes in parallel
make watch-server & \
make watch-templates & \
make watch-css & \
make browser-sync
watch-server:
air -c .air.toml
watch-templates:
templ generate --watch
watch-css:
tailwindcss -i ./styles/input.css -o ./public/styles.css --watch
browser-sync:
browser-sync start --proxy "localhost:3000" --files "public//"

### Hot Reload Implementation

The development environment provides hot reload for:

- Go code changes (via `air`)
- Templ template changes (via `templ generate --watch`)
- CSS changes (via `tailwindcss --watch`)
- Browser auto-refresh (via `browser-sync`)

When any file changes:
1. Relevant watcher detects the change
2. Files are recompiled if needed
3. Server restarts if necessary
4. Browser automatically refreshes

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   npm install
   ```
3. Start development server:
   ```bash
   make dev
   ```

This will launch the development environment with split panes in iTerm2 showing all running processes.

## Production Build

For production deployment:

```bash
make build
```

This will:
1. Generate templ templates
2. Build optimized CSS
3. Compile Go binary



