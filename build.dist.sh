#!/usr/bin/env bash
# bash
BIN_FILE=ms-email-graphql-api

read -p 'platform: ' platform

# platforms=("darwin/amd64" "linux/amd64" "linux/386" "windows/amd64" "windows/386")
platforms=($platform)

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$BIN_FILE
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo "build \"$output_name\" for $GOOS/$GOARCH ..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    OS_DIR=$GOOS'-'$GOARCH
    if [ -f "$output_name" ]; then
        if [ -d "dist/$OS_DIR/" ]; then
            rm -rf dist/$OS_DIR/
            sleep 2
        fi

        mkdir -p dist/$OS_DIR/conf/casbin
        mkdir -p dist/$OS_DIR/conf/data
        mkdir -p dist/$OS_DIR/interface/http-apps/graphql/echo/graphql/schema

        mv $output_name dist/$OS_DIR/

        cp conf/config.yaml dist/$OS_DIR/conf/
        cp conf/casbin/casbin_rbac_rest_model.conf dist/$OS_DIR/conf/casbin/
        cp conf/data/test-data.yaml dist/$OS_DIR/conf/data/
        
        cp -R www dist/$OS_DIR/
        cp -R interface/http-apps/graphql/echo/graphql/schema/schema.graphql dist/$OS_DIR/interface/http-apps/graphql/echo/graphql/schema/
    fi
done