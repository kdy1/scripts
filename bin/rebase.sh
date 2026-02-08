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

echo ""
step "Fetching origin..."
git fetch origin
ok "Fetched"

# Determine merge base with main
MERGE_BASE=$(git merge-base HEAD origin/main)
COMMIT_COUNT=$(git rev-list --count "$MERGE_BASE..HEAD")
BRANCH=$(git branch --show-current)

echo ""
step "Squashing commits on ${CYAN}${BRANCH}${RESET}"
info "${COMMIT_COUNT} commit(s) ahead of origin/main"

# Squash all PR commits into one
if [ "$COMMIT_COUNT" -gt 1 ]; then
  # Use PR title if available, otherwise first commit subject
  MSG=$(gh pr view --json title -q .title 2>/dev/null) || \
    MSG=$(git log --format="%s" --reverse "$MERGE_BASE..HEAD" | head -n 1)

  git reset --soft "$MERGE_BASE"
  git commit -m "$MSG"
  ok "Squashed ${COMMIT_COUNT} commits → 1"
  info "Message: ${MSG}"
elif [ "$COMMIT_COUNT" -eq 1 ]; then
  ok "Already a single commit"
else
  err "No commits to squash"
  exit 1
fi

# Rebase onto latest origin/main using claude
echo ""
step "Rebasing onto ${CYAN}origin/main${RESET} via Claude..."
claude -p 'Do git rebase origin/main' --allow-tools 'Bash,WebFetch,WebSearch'
ok "Rebased"

# Force push
echo ""
step "Pushing to ${CYAN}${BRANCH}${RESET}..."
git push --force-with-lease
ok "Pushed"

echo ""
echo -e "${GREEN}${BOLD}Done!${RESET} Branch ${CYAN}${BRANCH}${RESET} rebased and pushed."
echo ""
