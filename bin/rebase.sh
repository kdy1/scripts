#!/usr/bin/env bash
set -eu

git fetch origin

# Determine merge base with main
MERGE_BASE=$(git merge-base HEAD origin/main)
COMMIT_COUNT=$(git rev-list --count "$MERGE_BASE..HEAD")

# Squash all PR commits into one
if [ "$COMMIT_COUNT" -gt 1 ]; then
  # Use PR title if available, otherwise first commit subject
  MSG=$(gh pr view --json title -q .title 2>/dev/null) || \
    MSG=$(git log --format="%s" --reverse "$MERGE_BASE..HEAD" | head -n 1)

  git reset --soft "$MERGE_BASE"
  git commit -m "$MSG"
  echo "Squashed $COMMIT_COUNT commits into 1"
elif [ "$COMMIT_COUNT" -eq 1 ]; then
  echo "Already a single commit"
else
  echo "No commits to squash"
  exit 1
fi

# Rebase onto latest origin/main using claude
claude -p 'Do git rebase origin/main' --allow-tools 'Bash,WebFetch,WebSearch'

# Force push
git push --force-with-lease
