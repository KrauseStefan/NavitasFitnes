
cd websrc

#npm install
npm run clean

Start-Process "npm" "run watch"

cd ..

Start-Process "goapp" "serve ./app-engine/"

cd ipn-simulator
Start-Process "npm start"
cd ..
