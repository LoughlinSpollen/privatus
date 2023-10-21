#! /usr/bin/env sh

if [[ -z "$TENANCY_API_HTTP_HOST" ]]; then
    TENANCY_API_HTTP_HOST="localhost"
fi

stringToBase64() {
    local string=$1
    local encoded=$(echo "$string" | base64)
    echo "$encoded"
}

json() {
    json=$(echo '{
            "id":       "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
            "ml-model": "Python code here"
        }' | jq --compact-output '.')
    echo "$json"
}

createTenancy () {
    local data=$(json)
    data=$(stringToBase64 "$data")

    local status=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Content-Type: application/vnd.api+json" -H "Accept: application/vnd.api+json"  -d $data http://localhost:8080/v1/tenancy)

    if [ "$status" -eq 201 ]; then
        echo "Success: status code $response"
    else
        echo "Error: Expected 201 but got $response"
        exit 1
    fi
}

createTenancy