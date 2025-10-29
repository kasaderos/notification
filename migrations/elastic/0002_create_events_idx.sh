#!/bin/bash

# Create events index
curl -X PUT "localhost:9200/events" \
  -H "Content-Type: application/json" \
  -d '{
    "mappings": {
      "properties": {
        "url": {
          "type": "keyword"
        },
        "domain": {
          "type": "keyword"
        },
        "content": {
          "type": "text"
        }
      }
    }
  }'
