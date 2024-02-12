touch /app/server.log;
echo "nohup ./server > server.log 2>&1 &" | bash;
tail -f /app/server.log;
