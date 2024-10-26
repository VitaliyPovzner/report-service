# Go Metrics Aggregation Service

## Overview

The **Go Metrics Aggregation Service** is a lightweight microservice designed to aggregate and analyze metrics based on user-defined dimensions and metrics. This service enables users to extract valuable insights from their data over specified date ranges, facilitating better decision-making and performance tracking.

## Features

- **Dynamic Querying**: Users can specify dimensions and metrics to group data dynamically.
- **Date Range Filtering**: The service allows filtering of data between two dates, enhancing the relevance of the analysis.
- **Breakdown Support**: Enables users to request data breakdowns based on specific criteria for more granular insights.


## Functionality

### API Endpoint

- **POST /aggregate**
  
  This endpoint accepts a JSON payload to perform the aggregation. The payload should include the following fields:

  ```json
  {
    "dimensions": ["string"],
    "metrics": ["string"],
    "dateFrom": "YYYY-MM-DD",
    "dateTo": "YYYY-MM-DD",
    "breakdown": "string"
  }
