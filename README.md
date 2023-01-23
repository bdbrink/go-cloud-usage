# What this does

leverage aws cost usage service to generate data in json format using golang

## Why

The AWS docs for golang sdk is quite poor in comparsion to other lanuguages, this is a simple query to get s3 metrics to serve as boiler plate, because I literally can't find any examples online.

## Future

The idea with this was to query unused buckets that had no API activity (GET/PUTs) and mark them for deletion by adding a tag. A dry run would first deny all access to the bucket then delete when their is no users determined at a later date.

## Reponse syntax

Ex reponse from running in aws account

```
{
          Keys: ["WriteVersioningProps"],
          Metrics: {
            UsageQuantity: {
              Amount: "341890.0077954377",
              Unit: "N/A"
            }
          }
        }
      ],
      TimePeriod: {
        End: "2023-01-01",
        Start: "2022-12-01"
      },
```