aws dynamodb put-item \
    --table-name Users \
    --item file://user-item.json \
    --return-consumed-capacity TOTAL \
    --return-item-collection-metrics SIZE
