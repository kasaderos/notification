#!/bin/bash

# Create events_percolations index
curl -X PUT "localhost:9200/events_percolations" \
  -H "Content-Type: application/json" \
  -d '{
    "mappings": {
      "properties": {
        "query": {
          "type": "percolator"
        },
        "user_id": {
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
