#!/bin/bash

imageName=$1

repoName=$(echo $imageName | cut -d '/' -f 2 | cut -d ":" -f 1 | tr -d '"')
imageTag=$(echo $imageName | cut -d '/' -f 2 | cut -d ":" -f 2 | tr -d '"')
imageDigest=$(aws ecr batch-get-image --repository-name $repoName --image-ids imageTag=$imageTag | jq '.images[]' | jq --slurp -r '.[0].imageId.imageDigest')

imageTags=( $(aws ecr batch-get-image --repository-name $repoName --image-ids imageDigest=$imageDigest | jq  -c '.images[].imageId.imageTag' ) )
result=()

for tag in "${imageTags[@]}"
do
  tag=$(echo $tag | tr -d '"')
  if [ $tag != $imageTag ]; then
    result+=($tag)
  fi
done

json_array() {
  echo -n '['
  while [ $# -gt 0 ]; do
    x=${1//\\/\\\\}
    echo -n \"${x//\"/\\\"}\"
    [ $# -gt 1 ] && echo -n ', '
    shift
  done
  echo ']'
}

echo "{ \"result\" : $(json_array "${result[@]}") }"
