#!/bin/bash

TABLE="Users"
KEY_NAME="UserId"

# すべてのアイテムのキーを取得
KEYS=$(aws dynamodb scan --table-name $TABLE --projection-expression "$KEY_NAME" --output json | jq -c '.Items[]')

# 一度に25件ずつ削除リクエストを作成
BATCH=()
COUNT=0

for key in $KEYS; do
    BATCH+=("{\"DeleteRequest\": {\"Key\": $key}}")
    COUNT=$((COUNT+1))
    
    if [ $COUNT -eq 25 ]; then
        # JSON 配列に変換してリクエストファイルを作成
        REQUEST=$(printf '%s\n' "${BATCH[@]}" | jq -s '{ "'$TABLE'": . }')
        echo "$REQUEST" > delete-items.json
        aws dynamodb batch-write-item --request-items file://delete-items.json
        BATCH=()
        COUNT=0
    fi
done

# バッチに残ったアイテムがあれば削除
if [ $COUNT -gt 0 ]; then
    REQUEST=$(printf '%s\n' "${BATCH[@]}" | jq -s '{ "'$TABLE'": . }')
    echo "$REQUEST" > delete-items.json
    aws dynamodb batch-write-item --request-items file://delete-items.json
fi

# リクエストファイルを削除
rm delete-items.json