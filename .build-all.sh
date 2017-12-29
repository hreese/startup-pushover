BIN_FILE_NAME_PREFIX=$1
PROJECT_DIR=$2
BIN_FILE_NAME_PREFIX=$1PROJECT_DIR=$2PLATFORMS=linux/386 linux/amd64 linux/arm linux/arm64 linux/mips linux/mips64 linux/mips64le linux/mipsle linux/ppc64 linux/ppc64le
for PLATFORM in $PLATFORMS; do
        GOOS=${PLATFORM%/*}
        GOARCH=${PLATFORM#*/}
        FILEPATH="$PROJECT_DIR/artifacts/${GOOS}-${GOARCH}"
        #echo $FILEPATH
        mkdir -p $FILEPATH
        BIN_FILE_NAME="$FILEPATH/${BIN_FILE_NAME_PREFIX}"
        #echo $BIN_FILE_NAME
        if [[ "${GOOS}" == "windows" ]]; then BIN_FILE_NAME="${BIN_FILE_NAME}.exe"; fi
        CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o ${BIN_FILE_NAME}"
        #echo $CMD
        echo "${CMD}"
        eval $CMD || FAILURES="${FAILURES} ${PLATFORM}"
done
