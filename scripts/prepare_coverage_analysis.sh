PLATFORM=$(uname -s | tr '[A-Z]' '[a-z]')
ARCH=$(uname -m | sed -e 's/x86_/amd/g')
CC_DL_BASE_URL=https://codeclimate.com/downloads/test-reporter/test-reporter-latest-
DOWNLOAD_LINK=$CC_DL_BASE_URL$PLATFORM-$ARCH

curl \
    -L $DOWNLOAD_LINK > ./cc-test-reporter \
&& chmod +x ./cc-test-reporter \
&& ./cc-test-reporter before-build \
&& go install github.com/jandelgado/gcov2lcov@latest