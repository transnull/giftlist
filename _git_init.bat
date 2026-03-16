@echo off
cd /d c:\Users\Forest\WorkBuddy\20260316104955

echo [1/6] git config...
git config user.email "transnull@outlook.com"
git config user.name "transnull"

echo [2/6] git init...
git init

echo [3/6] git add...
git add .

echo [4/6] git commit...
git commit -m "feat: initial release - Gift Wishlist app with Go + Vue 3 + Docker support"

echo Done.
