#!/bin/bash
set -e

# Tailwind Watch starten (im Hintergrund)
npx tailwindcss -i ./src/app/css/input.css -o ./src/app/static/css/output.css --watch &

# PID merken, damit wir ihn später ggf. beenden können
TAILWIND_PID=$!

# Air starten (Go Hot-Reload)
cd src/app/
air

# Wenn Air beendet wird, auch Tailwind stoppen
kill $TAILWIND_PID
