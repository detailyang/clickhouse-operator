#!/bin/bash

# Start new release branch


#
# Increments version represented as x.y.z
# $1: version itself
# $2: number of part: 0 – major, 1 – minor, 2 – patch
#
function increment_version() {
    local version="${1}"
    local what="${2}"

    local delimiter="."
    local array=($(echo "${version}" | tr "${delimiter}" '\n'))

    # Increment desired part
    array[${what}]=$((array[${what}]+1))
    # Zero all following parts
    if [[ ${what} -lt 2 ]]; then array[2]=0; fi
    if [[ ${what} -lt 1 ]]; then array[1]=0; fi
    # Provide result
    echo $(local IFS=${delimiter}; echo "${array[*]}")
}

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/go_build_config.sh"

CUR_RELEASE=$(cat "${SRC_ROOT}/release")
NEW_RELEASE_MAJOR=$(increment_version "${CUR_RELEASE}" 0)
NEW_RELEASE_MINOR=$(increment_version "${CUR_RELEASE}" 1)
NEW_RELEASE_PATCH=$(increment_version "${CUR_RELEASE}" 2)
echo "Starting new release."
echo "Current release: ${CUR_RELEASE}"
echo "What would you like to start? Possible options:"
echo "  1     - new MAJOR version: ${NEW_RELEASE_MAJOR}"
echo "  2     - new MINOR version: ${NEW_RELEASE_MINOR}"
echo "  3     - new PATCH version: ${NEW_RELEASE_PATCH}"
echo "  x.y.z - in case you'd like to start something completely different just write required version"
echo -n "Enter command choice (1, 2, 3) or custom release (x.y.z): "
read COMMAND
# Trim EOL from the command received
COMMAND=$(echo "${COMMAND}" | tr -d '\n\t\r ')
echo "Provided command is: ${COMMAND}"
echo -n "Which means we are going to "

case "${COMMAND}" in
    "1")
        NEW_RELEASE="${NEW_RELEASE_MAJOR}"
        echo "start new MAJOR release: ${NEW_RELEASE}"
        ;;
    "2")
        NEW_RELEASE="${NEW_RELEASE_MINOR}"
        echo "start new MINOR release: ${NEW_RELEASE}"
        ;;
    "3")
        NEW_RELEASE="${NEW_RELEASE_PATCH}"
        echo "start new PATCH release: ${NEW_RELEASE}"
        ;;
    *)
        NEW_RELEASE="${COMMAND}"
        echo "start new CUSTOM release: ${NEW_RELEASE}"
        ;;
esac

read -p "Press enter to start new release"
echo "Starting new release: ${NEW_RELEASE}"

# Create release branch
git branch "${NEW_RELEASE}"
git checkout "${NEW_RELEASE}"
# Append cur release to the head of releases list
cat "${SRC_ROOT}/release" "${SRC_ROOT}/releases" > "${SRC_ROOT}/releases_tmp" && mv "${SRC_ROOT}/releases_tmp" "${SRC_ROOT}/releases"
# And write new release to release file
echo "${NEW_RELEASE}" > "${SRC_ROOT}/release"

# Commit new branch
COMMIT=$(cd "${SRC_ROOT}" && git add . && git commit -m "${NEW_RELEASE}")
echo ${COMMIT}

# Some niceness
echo "Releases:"
cat "${SRC_ROOT}/release"
head -n3 "${SRC_ROOT}/releases"

echo "git status:"
git status
