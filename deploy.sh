#!/bin/bash
cd frontend
npm run build
cd ..
flyctl deploy --env GIN_MODE=release