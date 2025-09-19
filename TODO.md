# Hot-Reload Optionen für Go-Projekte

## 1. Air (empfohlen)
- Beobachtet Go-Dateien, Templates und CSS.  
- Startet den Server automatisch neu bei Änderungen.  
- Installation: `go install github.com/cosmtrek/air@latest`  
- Nutzung: `air` im Projektordner.

## 2. CompileDaemon
- Minimalistisches Tool, überwacht nur Go-Dateien.  
- Startet den Server automatisch neu bei Änderungen.  
- Installation: `go install github.com/githubnemo/CompileDaemon@latest`  
- Nutzung: `CompileDaemon -command="go run main.go"`

## 3. Reflex
- Flexibles Tool, kann beliebige Dateien überwachen.  
- Startet bei Änderungen den angegebenen Befehl neu.  
- Installation: `go install github.com/cespare/reflex@latest`  
- Nutzung (Beispiel für Go + HTML):  
  ```bash
  reflex -r '\.go$|\.html$' -s -- sh -c 'go run main.go'
