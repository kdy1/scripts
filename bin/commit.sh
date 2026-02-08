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

# Check for changes
echo ""
step "Checking working tree..."
if git diff --quiet && git diff --cached --quiet && [ -z "$(git ls-files --others --exclude-standard)" ]; then
  err "No changes to commit"
  exit 1
fi

STAGED=$(git diff --cached --stat | tail -1)
UNSTAGED=$(git diff --stat | tail -1)
UNTRACKED=$(git ls-files --others --exclude-standard | wc -l | tr -d ' ')

if [ -n "$STAGED" ]; then
  info "Staged: ${STAGED}"
fi
if [ -n "$UNSTAGED" ]; then
  info "Unstaged: ${UNSTAGED}"
fi
if [ "$UNTRACKED" -gt 0 ]; then
  info "Untracked: ${UNTRACKED} file(s)"
fi

# Commit via Claude
echo ""
step "Committing via Claude..."
claude -p 'Do git commit' --allowed-tools 'Bash'
ok "Committed"

echo ""
echo -e "${GREEN}${BOLD}Done!${RESET}"
echo ""
