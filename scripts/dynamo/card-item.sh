aws dynamodb put-item \
    --table-name Cards \
    --item file://card-item.json \
    --return-consumed-capacity TOTAL \
    --return-item-collection-metrics SIZE
