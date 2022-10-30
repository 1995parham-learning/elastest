# Elastic APM (Application Performance Monitoring) Playground

## Introduction

Testing logging into elastic with elastic APM server.

## Configuration

Please set up these environment variables for connection string to elastic:

| Variable                   | Default Value           | Usage  |
| -------------------------- |:-----------------------:|:----- |
| ELASTIC_APM_SERVER_URL     | <http://localhost:8200> | The URL for your Elastic APM server. |
| ELASTIC_APM_SERVER_TIMEOUT | 30s                     | The timeout for requests made to your Elastic APM server. |
| ELASTIC_APM_SECRET_TOKEN   |                         | This string is used to ensure that only your agents can send data to your APM server. |
