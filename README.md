# Go Metrics Aggregation Service

## Overview

The **Go Metrics Aggregation Service** is a lightweight microservice designed to aggregate and analyze metrics based on user-defined dimensions and metrics. This service enables users to extract valuable insights from their data over specified date ranges, facilitating better decision-making and performance tracking. It supports dynamic querying, date range filtering, and data breakdowns for granular insights.

## Features

- **Dynamic Querying**: Users can specify dimensions and metrics to group data dynamically, allowing for flexible and customizable data analysis.
- **Date Range Filtering**: The service allows filtering of data between two dates, enhancing the relevance of the analysis by focusing on specific time periods.
- **Breakdown Support**: Enables users to request data breakdowns based on specific criteria, such as time intervals or custom dimensions, for more detailed insights.
- **Custom Filters**: Supports a variety of filter operations to refine data selection, including:
  - `eq`: Equals
  - `not_eq`: Not equals
  - `less_eq`: Less than or equal to
  - `less`: Less than
  - `gr_eq`: Greater than or equal to
  - `gr`: Greater than
  - `cont`: Contains (LIKE)
  - `not_cont`: Does not contain (NOT LIKE)
  - `starts`: Starts with (LIKE)
  - `in`: In a list
  - `not_in`: Not in a list
- **SQL Generation**: Automatically generates SQL queries based on user input, ensuring efficient data retrieval and processing.

## Functionality

### API Endpoint

- **POST /report/example-report**
  
  This endpoint accepts a JSON payload to perform the aggregation. The payload should include the following fields:

  ```json
  {
    "dimensions": ["string"],  // List of dimensions to group data by
    "metrics": ["string"],     // List of metrics to aggregate
    "dateFrom": "YYYY-MM-DD",  // Start date for the data range
    "dateTo": "YYYY-MM-DD",    // End date for the data range
    "breakdown": "string",     // Criteria for data breakdown (e.g., daily, weekly)
    "filters": [               // Optional filters to apply to the data
      {
        "operand": "string",   // The field to filter on
        "operator": "string",  // The operation to perform (e.g., eq, not_eq, in)
        "value": "string"      // The value to compare against
      }
    ]
  }
  ```

### Example Input

Here is an example of a JSON payload that can be sent to the `/report/example-report` endpoint:

```json
{
  "dimensions": ["test", "category"],
  "metrics": ["total_value", "total_count"],
  "dateFrom": "2024-05-11",
  "dateTo": "2024-05-12",
  "breakdown": "daily",
  "filters": [
    {
      "operand": "total_count",
      "operator": "gr_eq",
      "value": "100"
    }
  ]
}
```