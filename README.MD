```
nginx 配置
log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
'$status $body_bytes_sent "$http_referer" '
'"$http_user_agent" "$http_x_forwarded_for" '
'"$request_time"';

access_log /home/wwwlogs/www.phpshjgame.com_access_log.log main;
```
