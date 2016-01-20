#!/bin/bash
##  curl -sSL https://epull-api.yunpro.cn/i/install.sh |sh
##
echo "Downloading"
curl -sSL https://epull-api.yunpro.cn/i/epull.sh > /usr/local/bin/epull 
chmod +x  /usr/local/bin/epull
echo "That's it ! Finished!"
echo "\n\n Try input : "
echo "epull gcr.io/google_containers/pause "
