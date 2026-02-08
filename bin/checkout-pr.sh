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

if [ $# -lt 1 ]; then
  err "Usage: checkout-pr.sh <PR_NUMBER>"
  exit 1
fi

PR_NUMBER="$1"

echo ""
step "Fetching PR #${PR_NUMBER} info..."
BRANCH=$(gh pr view "$PR_NUMBER" --json headRefName -q .headRefName)
ok "Branch: ${CYAN}${BRANCH}${RESET}"

echo ""
step "Fetching origin..."
git fetch origin
ok "Fetched"

echo ""
step "Checking out ${CYAN}${BRANCH}${RESET}..."
git checkout -B "$BRANCH" "origin/$BRANCH"
ok "Checked out"

echo ""
echo -e "${GREEN}${BOLD}Done!${RESET} PR #${PR_NUMBER} checked out as ${CYAN}${BRANCH}${RESET}."
echo ""
