#!/bin/bash
set -e
OUTPUT_DIR="./out"

function echo_task() {
    echo
    echo -e "\e[34mTask: \e[36m$1\e[39m"
}

rm -rf "${OUTPUT_DIR}"
mkdir -p "${OUTPUT_DIR}"

echo_task "Resolving dependencies"
go get -t -v ./...

echo_task "Running tests"
go test -v ./...

echo_task "Building"
go build -o ./out/bsearch

BSEARCH_VERSION=$("${OUTPUT_DIR}"/bsearch --version | sed -e 's/bsearch\ version //')
echo_task "Determined version: ${BSEARCH_VERSION}"

echo_task "Preparing debian package"
PACKAGE_DIRECTORY="${OUTPUT_DIR}/bsearch_${BSEARCH_VERSION}"

mkdir -p "${PACKAGE_DIRECTORY}/usr/local/bin" "${PACKAGE_DIRECTORY}/usr/share/man/man1"
cp "${OUTPUT_DIR}/bsearch" "${PACKAGE_DIRECTORY}/usr/local/bin/bsearch"
cp "./bsearch.1" "${PACKAGE_DIRECTORY}/usr/share/man/man1/bsearch.1"
gzip "${PACKAGE_DIRECTORY}/usr/share/man/man1/bsearch.1"

mkdir -p "${PACKAGE_DIRECTORY}/DEBIAN"

cat << EOF > "${PACKAGE_DIRECTORY}/DEBIAN/control"
Package: bsearch
Version: ${BSEARCH_VERSION}
Section: base
Priority: optional
Architecture: amd64
Maintainer: James Ridgway <myself@james-ridgway.co.uk>
Homepage: https://www.james-ridgway.co.uk
Description: A utility for binary searching a sorted file for lines that start with the search key.
EOF

echo_task "Building debian package"
dpkg-deb --build "${PACKAGE_DIRECTORY}"
DEBIAN_FILE="${OUTPUT_DIR}/bsearch_${BSEARCH_VERSION}.deb"


if git tag -l --contains HEAD | grep -oP "^\d+.\d+.\d+$"; then

    RELEASE_TAG=$(git tag -l --contains HEAD | grep -oP "^\d+.\d+.\d+$")
    echo_task "Detected release version ${RELEASE_TAG}"

    if [ -n "${BINTRAY_API_KEY}" ]; then

        echo_task "Preparing to release"

        if [ -z "${GPG_PASSPHRASE}" ]; then
            echo "ERROR: GPG_PASSPHRASE not set!"
            exit 1
        fi

        echo_task "Publishing"
        curl -T "${DEBIAN_FILE}" -ujamesridgway:"${BINTRAY_API_KEY}" -H " X-GPG-PASSPHRASE: ${GPG_PASSPHRASE}" "https://api.bintray.com/content/jamesridgway/debian/bsearch/${BSEARCH_VERSION}/bsearch_${BSEARCH_VERSION}.deb;deb_distribution=xenial;deb_component=main;deb_architecture=amd64;publish=1"
    fi

else

    echo_task "Skipping release process, no tag found"
fi

echo
echo -e "\e[32mBuild Complete!\e[39m"