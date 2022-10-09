echo "Kill PM2"
pm2 kill

echo "Jump to app folder"
cd ~/qubo-backend

echo "Update app from Git"
git pull

echo "Build Go App"
go build main.go

echo "Start in background"
cd ~
pm2 start ecosystem.config.js