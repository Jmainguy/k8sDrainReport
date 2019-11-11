#!/bin/bash
# Set REPO To update
export REPO="jmainguy/opsMarina"
export DIRNAME=$(echo $REPO | awk -F'/' '{print $2}')

git clone "https://$GITHUB_TOKEN@github.com/$REPO.git"
cd $DIRNAME

# Set release tag
LASTRELEASE=$(curl --silent "https://$GITHUB_TOKEN@api.github.com/repos/$REPO/releases/latest" | jq -r .tag_name)
echo "LASTRELEASE: $LASTRELEASE"
TODAY=$(date +"%Y-%m-%d")
echo "TODAY: $TODAY"
LASTRELEASEDAY=$(echo $LASTRELEASE | awk -F"-" '{print $1"-"$2"-"$3}')
if [[ $TODAY != $LASTRELEASEDAY ]]; then
    export TRAVIS_TAG=$(echo "$TODAY-1")
else
    N=$(echo $LASTRELEASE | awk -F'-' '{print $NF}')
    ((N++))
    export TRAVIS_TAG=$(echo "$TODAY-$N")
fi
echo "TRAVIS_TAG: $TRAVIS_TAG"

# Tag and Push
git tag $TRAVIS_TAG
git push origin $TRAVIS_TAG
