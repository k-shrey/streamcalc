# Stream Statistics Calculator

A RESTful API to calculate real-time statistics for financial instruments in a 60-second window

## Usage

1) Make sure you have Go installed (tested on version 1.17.4, Windows 21H2 x64)  
2) Clone the repo and run

```bash
 go build
```

## Endpoints

**Send data:** `/tick POST  `  
Request format:
```json
{
   "instrument": "string",
   "value": "float",
   "timestamp": "epoch/int64",
}
```
**Get statistics:** `/stats/{instrument} GET  `  
*Example*: `/stats/GOOGL  `  
Use `/stats/ALL` for stats across all instruments   
Response format:  
```json
{
   "avg": "float",
   "min": "float",
   "max": "float",
   "count": "int64"
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)