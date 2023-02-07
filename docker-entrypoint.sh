echo "nohup ./server > server.log 2>&1 &" | bash;
service nginx start;
tail -f /var/log/nginx/access.log