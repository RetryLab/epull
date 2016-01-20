#!/bin/bash

## usage
#
# This script is meant for quick & easy install via:
#   curl -sSL https://epull-api.yunpro.cn/i/epull.sh > /usr/local/bin/epull ; chmod +x  /usr/local/bin/epull
# or:
#   wget -qO- https://epull-api.yunpro.cn/i/epull.sh  > /usr/local/bin/epull ; chmod +x  /usr/local/bin/epull
#
# verify :  epull gcr.io/google_containers/pause
#
API_SERVER="https://epull-api.yunpro.cn"

init(){
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
}

pull_retag(){
	ORI=$1
	TARGET=$2
	docker pull $TARGET
	echo -e "${OK} docker pull $TARGET"
	docker tag -f $TARGET $ORI
	echo -e "${OK} docker tag $TARGET $ORI"
	echo -e "${OK} Finish! Have a nice day :)"
}

init
ORI=$1
if [ "$ORI" = "" ]; then
	echo -e "${NOTICE} No image is specified"
	exit 1
fi

echo -e "${INFO} epull is working"
result=$(curl -sf -X POST  --form "uri=${ORI}" "$API_SERVER/v1/tool/encode-uri")
resultString=$(echo  $result |tr "\n" " ")
resultArr=($resultString)

if [ "${resultArr[0]}" != "SUCCESS" ]; then
	echo -e "\n${ERROR} Invalid Request to Server"
	exit 1
fi

TARGET=${resultArr[1]}
V2URI=${resultArr[2]}

echo -e "${INFO} Processing $TARGET"
echo -e "${INFO} Docker Distribution URI: $V2URI"

result=$(curl -sf "$V2URI" |grep schemaVersion)

if [ "$result" != "" ]; then
	pull_retag $ORI $TARGET
	exit 0
fi

echo -e "${INFO} No cache in server side. Downloading $ORI, Please Wait..."
for i in `seq 1 100`;
do
	#trigger porter-worker 
    curl -sf -X POST  -F ori=${ORI} -F target=${TARGET} $API_SERVER/v1/worker/
    sleep 10
    echo "waiting : $((i*10))s"
    result=$(curl -sf "$V2URI" |grep schemaVersion)
	if [ "$result" != "" ]; then
		pull_retag $ORI $TARGET
		exit 0
	fi
done    

echo -e "${ERROR} TIME OUT, Please retry!"


