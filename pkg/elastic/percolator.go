package elastic

// // IndexManager handles Elasticsearch index operations
// type IndexManager struct {
// 	client *elasticsearch.Client
// }

// // NewIndexManager creates a new index manager
// func NewIndexManager(client *elasticsearch.Client) *IndexManager {
// 	return &IndexManager{client: client}
// }

// // SetupIndexes creates the necessary Elasticsearch indexes
// func (im *IndexManager) SetupIndexes(ctx context.Context) error {
// 	// Create events_percolator index
// 	if err := im.createEventsPercolatorIndex(ctx); err != nil {
// 		return fmt.Errorf("failed to create events_percolator index: %w", err)
// 	}

// 	// Create events index
// 	if err := im.createEventsIndex(ctx); err != nil {
// 		return fmt.Errorf("failed to create events index: %w", err)
// 	}

// 	return nil
// }

// // createEventsPercolatorIndex creates the percolator index for events queries
// func (im *IndexManager) createEventsPercolatorIndex(ctx context.Context) error {
// 	mapping := map[string]interface{}{
// 		"mappings": map[string]interface{}{
// 			"properties": map[string]interface{}{
// 				"query": map[string]interface{}{
// 					"type": "percolator",
// 				},
// 				"domain": map[string]interface{}{
// 					"type": "keyword",
// 				},
// 				"user_id": map[string]interface{}{
// 					"type": "keyword",
// 				},
// 				"rule_id": map[string]interface{}{
// 					"type": "keyword",
// 				},
// 				"title": map[string]interface{}{
// 					"type": "text",
// 				},
// 				"content": map[string]interface{}{
// 					"type": "text",
// 				},
// 			},
// 		},
// 	}

// 	return im.createIndex(ctx, "events_percolator", mapping)
// }

// // createEventsIndex creates the events index for storing events articles
// func (im *IndexManager) createEventsIndex(ctx context.Context) error {
// 	mapping := map[string]interface{}{
// 		"mappings": map[string]interface{}{
// 			"properties": map[string]interface{}{
// 				"domain": map[string]interface{}{
// 					"type": "keyword",
// 				},
// 				"url": map[string]interface{}{
// 					"type": "keyword",
// 				},
// 				"title": map[string]interface{}{
// 					"type": "text",
// 				},
// 				"content": map[string]interface{}{
// 					"type": "text",
// 				},
// 				"published_at": map[string]interface{}{
// 					"type": "date",
// 				},
// 			},
// 		},
// 	}

// 	return im.createIndex(ctx, "events", mapping)
// }

// // createIndex creates an Elasticsearch index with the given mapping
// func (im *IndexManager) createIndex(ctx context.Context, indexName string, mapping map[string]interface{}) error {
// 	var buf bytes.Buffer
// 	if err := json.NewEncoder(&buf).Encode(mapping); err != nil {
// 		return fmt.Errorf("error encoding mapping: %w", err)
// 	}

// 	res, err := im.client.Indices.Create(
// 		indexName,
// 		im.client.Indices.Create.WithBody(&buf),
// 		im.client.Indices.Create.WithContext(ctx),
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error creating index: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.IsError() {
// 		body, _ := io.ReadAll(res.Body)
// 		return fmt.Errorf("error response: %s", string(body))
// 	}

// 	return nil
// }

// // PercolatorService handles percolator operations
// type PercolatorService struct {
// 	client *elasticsearch.Client
// }

// // NewPercolatorService creates a new percolator service
// func NewPercolatorService(client *elasticsearch.Client) *PercolatorService {
// 	return &PercolatorService{client: client}
// }

// // RegisterQuery registers a percolator query for a notification rule
// func (ps *PercolatorService) RegisterQuery(ctx context.Context, rule *model.NotificationRule) error {
// 	query := ps.buildQuery(rule)

// 	doc := map[string]interface{}{
// 		"query":   query,
// 		"rule_id": rule.ID.String(),
// 		"user_id": "", // Will be set when we have user context
// 		"domain":  strings.Join(rule.Rule.Sources.Domains, ","),
// 		"title":   strings.Join(rule.Rule.Keywords, " "),
// 		"content": rule.Rule.Prompt,
// 	}

// 	var buf bytes.Buffer
// 	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
// 		return fmt.Errorf("error encoding document: %w", err)
// 	}

// 	res, err := ps.client.Index(
// 		"events_percolator",
// 		&buf,
// 		ps.client.Index.WithDocumentID(rule.ID.String()),
// 		ps.client.Index.WithContext(ctx),
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error indexing percolator query: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.IsError() {
// 		body, _ := io.ReadAll(res.Body)
// 		return fmt.Errorf("error response: %s", string(body))
// 	}

// 	return nil
// }

// // PercolateEvents checks if events matches any registered queries
// func (ps *PercolatorService) PercolateEvents(ctx context.Context, events *model.Event) ([]string, error) {
// 	eventsDoc := map[string]interface{}{
// 		"domain":  events.Domain,
// 		"url":     events.URL,
// 		"title":   events.Title,
// 		"content": events.Content,
// 	}

// 	percolateQuery := map[string]interface{}{
// 		"percolate": map[string]interface{}{
// 			"field":    "query",
// 			"document": eventsDoc,
// 		},
// 	}

// 	var buf bytes.Buffer
// 	if err := json.NewEncoder(&buf).Encode(percolateQuery); err != nil {
// 		return nil, fmt.Errorf("error encoding percolate query: %w", err)
// 	}

// 	res, err := ps.client.Search(
// 		ps.client.Search.WithIndex("events_percolator"),
// 		ps.client.Search.WithBody(&buf),
// 		ps.client.Search.WithContext(ctx),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("error performing percolate search: %w", err)
// 	}
// 	defer res.Body.Close()

// 	if res.IsError() {
// 		body, _ := io.ReadAll(res.Body)
// 		return nil, fmt.Errorf("error response: %s", string(body))
// 	}

// 	var result map[string]interface{}
// 	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
// 		return nil, fmt.Errorf("error decoding response: %w", err)
// 	}

// 	// Extract matching rule IDs
// 	var ruleIDs []string
// 	if hits, ok := result["hits"].(map[string]interface{}); ok {
// 		if hitList, ok := hits["hits"].([]interface{}); ok {
// 			for _, hit := range hitList {
// 				if hitMap, ok := hit.(map[string]interface{}); ok {
// 					if source, ok := hitMap["_source"].(map[string]interface{}); ok {
// 						if ruleID, ok := source["rule_id"].(string); ok {
// 							ruleIDs = append(ruleIDs, ruleID)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return ruleIDs, nil
// }

// // buildQuery builds an Elasticsearch query from a notification rule
// func (ps *PercolatorService) buildQuery(rule *model.NotificationRule) map[string]interface{} {
// 	query := map[string]interface{}{
// 		"bool": map[string]interface{}{
// 			"must": []map[string]interface{}{},
// 		},
// 	}

// 	// Add domain filter if specified
// 	if len(rule.Rule.Sources.Domains) > 0 {
// 		query["bool"].(map[string]interface{})["must"] = append(
// 			query["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
// 			map[string]interface{}{
// 				"terms": map[string]interface{}{
// 					"domain": rule.Rule.Sources.Domains,
// 				},
// 			},
// 		)
// 	}

// 	// Add keyword filters if specified
// 	if len(rule.Rule.Keywords) > 0 {
// 		keywordQuery := map[string]interface{}{
// 			"bool": map[string]interface{}{
// 				"should": []map[string]interface{}{},
// 			},
// 		}

// 		for _, keyword := range rule.Rule.Keywords {
// 			keywordQuery["bool"].(map[string]interface{})["should"] = append(
// 				keywordQuery["bool"].(map[string]interface{})["should"].([]map[string]interface{}),
// 				map[string]interface{}{
// 					"multi_match": map[string]interface{}{
// 						"query":  keyword,
// 						"fields": []string{"title", "content"},
// 					},
// 				},
// 			)
// 		}

// 		query["bool"].(map[string]interface{})["must"] = append(
// 			query["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
// 			keywordQuery,
// 		)
// 	}

// 	return query
// }
