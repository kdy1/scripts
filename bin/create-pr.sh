#!/usr/bin/env bash
set -eu

# Colors
BOLD='\033[1m'
DIM='\033[2m'
GREEN='\033[32m'
BLUE='\033[34m'
YELLOW='\033[33m'
RED='\033[31m'
CYAN='\033[36m'
RESET='\033[0m'

step() { echo -e "${BLUE}${BOLD}==>${RESET}${BOLD} $1${RESET}"; }
info() { echo -e "    ${DIM}$1${RESET}"; }
ok()   { echo -e "    ${GREEN}✔${RESET} $1"; }
warn() { echo -e "    ${YELLOW}⚠${RESET} $1"; }
err()  { echo -e "    ${RED}✖${RESET} $1"; }

BRANCH=$(git branch --show-current)

# Guard: don't create PR from main
echo ""
if [ "$BRANCH" = "main" ] || [ "$BRANCH" = "master" ]; then
  err "Cannot create PR from ${CYAN}${BRANCH}${RESET}"
  exit 1
fi

# Show branch info
step "Preparing PR for ${CYAN}${BRANCH}${RESET}"
COMMIT_COUNT=$(git rev-list --count "origin/main..HEAD" 2>/dev/null || echo "?")
info "${COMMIT_COUNT} commit(s) ahead of origin/main"

# Create PR via Claude
echo ""
step "Creating PR via Claude..."
claude -p 'Create a PR. Commit files if necessary.' --allowed-tools 'Bash'
ok "PR created"

echo ""
echo -e "${GREEN}${BOLD}Done!${RESET} PR created for ${CYAN}${BRANCH}${RESET}."
echo ""
