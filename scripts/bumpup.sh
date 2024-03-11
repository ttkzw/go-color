#!/bin/sh

if ! gobump up -w; then
  echo "Aborted"
  exit
fi

git add color.go
version=v$(gobump show -r)
git commit -m "Release ${version}"
git tag -a "${version}" -m "Release ${version}"
git push origin "${version}"
git push
exit
