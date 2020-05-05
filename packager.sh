#!/bin/bash

set -ex

OUT_DIR=$(pwd)/dist
rm -rf $OUT_DIR
mkdir -p $OUT_DIR

cd backend
go build -o $OUT_DIR/main .

GOOS=windows GOARCH=386 \
  CGO_ENABLED=1 \
  CXX=i686-w64-mingw32-g++ \
  CC=i686-w64-mingw32-gcc \
  go build -o $OUT_DIR/main.exe .

cd -

cd frontend
npm i
npm run build
cd ..

cd backend
zip $OUT_DIR/release.zip config.json
cd $OUT_DIR
zip release.zip main main.exe
cd ..
cd ./frontend
zip -r $OUT_DIR/release.zip dist
