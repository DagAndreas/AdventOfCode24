
#!/bin/bash

# === Configuration ===
OLD_EMAIL="dag.folvell.ext@siemens.com"          # Replace with your incorrect email
NEW_NAME="Dag Andreas Foss Folvell"                # Replace with your correct name
NEW_EMAIL="daanfofo@gmail.com"  # Replace with your correct email
# =======================

# Ensure the script is run from the repository root
if [ ! -d ".git" ]; then
    echo "Error: This script must be run from the root of a Git repository."
    exit 1
fi

echo "Backing up current state..."
git branch backup-before-email-change

echo "Rewriting commit history to update email addresses..."

git filter-branch --env-filter '
if [ "$GIT_COMMITTER_EMAIL" = "'"$OLD_EMAIL"'" ]
then
    export GIT_COMMITTER_NAME="'"$NEW_NAME"'"
    export GIT_COMMITTER_EMAIL="'"$NEW_EMAIL"'"
fi
if [ "$GIT_AUTHOR_EMAIL" = "'"$OLD_EMAIL"'" ]
then
    export GIT_AUTHOR_NAME="'"$NEW_NAME"'"
    export GIT_AUTHOR_EMAIL="'"$NEW_EMAIL"'"
fi
' --tag-name-filter cat -- --branches --tags

echo "Commit history successfully rewritten."
echo "You may need to force push to update the remote repository:"
echo "  git push --force --tags origin 'refs/heads/*'"
