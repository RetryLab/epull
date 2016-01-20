#!/bin/bash
set -o pipefail

init() {
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

op_image() {
  # get action and image name passed
  ACTION="${1}"
  IMAGE="${2}"

  # check the action and act accordingly
  case ${ACTION} in
    push)
      # push image
      echo -e "${INFO} Pushing ${IMAGE}"
      (docker push ${IMAGE} && echo -e "${OK} Successfully ${ACTION}ed ${IMAGE}\n") || catch_push_pull_error "push" "${IMAGE}"
      ;;
    pull)
      # pull image
      echo -e "${INFO} Pulling ${IMAGE}"
      (docker pull ${IMAGE} && echo -e "${OK} Successfully ${ACTION}ed ${IMAGE}\n") || catch_push_pull_error "pull" "${IMAGE}"
      ;;
  esac
}

retag_image() {
  # get source and destination image names passed
  SOURCE_IMAGE="${1}"
  DESTINATION_IMAGE="${2}"

  # retag image
  (docker tag -f ${SOURCE_IMAGE} ${DESTINATION_IMAGE} && echo -e "${OK} ${V1_REGISTRY}/${i} > ${V2_REGISTRY}/${i}") || catch_retag_error "${SOURCE_IMAGE}" "${DESTINATION_IMAGE}"
}

init

echo -e "${INFO} epull work started"
echo -e "${INFO} TARGET_URI : $TARGET_URI"
echo -e "${INFO} ORI_URI : $ORI_URI" 

if [ "$TARGET_URI" == "" -o "$ORI_URI" == "" ]; then
  echo -e "${ERROR} invalid URI"
  exit 1
fi

echo -e "${INFO} pull image $ORI_URI"
op_image "pull" "$ORI_URI"

echo -e "${INFO} retag..."
retag_image "$ORI_URI" "$TARGET_URI" 

echo -e "${INFO} push..."
op_image "push" "$TARGET_URI"

echo -e "${OK} Finished!"

