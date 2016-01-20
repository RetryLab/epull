#!/bin/bash

## usage
#
# This script is meant for quick & easy install via:
#   curl -sSL https://epull-api.yunpro.cn/i/epull.sh > /usr/local/bin/epull ; chmod +x  /usr/local/bin/epull
# or:
#   wget -qO- https://epull-api.yunpro.cn/i/epull.sh  > /usr/local/bin/epull ; chmod +x  /usr/local/bin/epull
#

# Color Info
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CLEAR='\033[0m'
# pre-configure ok, warning, and error output
OK="[${GREEN}OK${CLEAR}]"
INFO="[${BLUE}INFO${CLEAR}]"
NOTICE="[${YELLOW}!!${CLEAR}]"
ERROR="[${RED}ERROR${CLEAR}]"

echo -e "${INFO} epull is working"

ORI=$1
API_SERVER="https://epull-api.yunpro.cn"
if [ "$ORI" = "" ]; then
	ORI="gcr.io/google_containers/pause:0.9.3"
	echo -e "${NOTICE} No image is specified"
	exit 1
fi

echo -e '${INFO} curl -sf -X POST  --form "uri=${ORI}" "$API_SERVER/v1/tool/encode-uri'
result=$(curl -sf -X POST  --form "uri=${ORI}" "$API_SERVER/v1/tool/encode-uri")

resultString=$(echo  $result |tr "\n" " ")

resultArr=($resultString)

if [ "${resultArr[0]}" != "SUCCESS" ]; then
	echo -e "\n${ERROR} Invalid Request to Server"
	exit 1
fi

targetUri=${resultArr[1]}
v2Uri=${resultArr[2]}

echo -e "${INFO} epull $targetUri"
echo -e "${INFO} epull $v2Uri"

result=$(curl -sf "$v2Uri" |grep schemaVersion)

updateRequired=true
if [ "$result" == "" ]; then
	updateRequired=true
else
	docker pull $targetUri
	echo -e "${OK} docker pull $targetUri"
	docker tag -f $targetUri $ORI
	echo -e "${OK} docker tag $targetUri $ORI"
	echo -e "${OK} Finish!"
	exit 0
fi

echo -e "${INFO} downloading $ORI, Please wait..."
for i in `seq 1 100`;
do
    curl -sf -X POST  -F ori=${ORI} -F target=${targetUri} $API_SERVER/v1/worker/
    sleep 9
    echo "waiting : $((i*10))s"
    result=$(curl -sf "$v2Uri" |grep schemaVersion)
	updateRequired=true
	if [ "$result" == "" ]; then
		sleep 1		
	else
		docker pull $targetUri
		echo -e "${OK} docker pull $targetUri"
		docker tag -f $targetUri $ORI
		echo -e "${OK} docker tag $targetUri $ORI"
		echo -e "${OK} Finish!"
		exit 0
	fi
done    

