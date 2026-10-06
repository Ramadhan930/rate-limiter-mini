# Go WAF Rate Limiter

A lightweight Web Application Firewall (WAF) built with Go, featuring IP-based rate limiting and basic SQL Injection detection.

## Features

- **IP-based rate limiting** — Maximum 6 requests per 10 seconds per IP
- **Automatic counter reset** — Counter resets after the time window expires
- **SQL Injection detection** — Basic pattern matching for common SQLi payloads
- **Automatic IP blacklisting** — IPs detected attacking are blocked (in-memory, cleared on server restart)
- **Request logging** — All requests and events are logged to `log.txt`

## Tech Stack

- **Language:** Go (Golang)
- **Standard Library:** `net/http`, `os`, `strings`, `time`
- **Architecture:** Modular packages (`handler`, `detector`, `logger`, `model`)

## Project Structure

```text
rate-limiter-mini/
├── main.go               # Entry point
├── handler/
│   └── handler.go        # HTTP handler with rate limiting logic
├── detector/
│   └── detector.go       # SQL Injection detection
├── logger/
│   └── logger.go         # File-based logging
├── model/
│   └── model.go          # Data structures
└── go.mod
```

## How to Run

1. Clone this repository:

   ```bash
   git clone https://github.com/Ramadhan930/rate-limiter-mini.git
   cd rate-limiter-mini
   ```

2. Run the server:

   ```bash
   go run main.go
   ```

3. Test with a browser or `curl`:

   ```bash
   # Normal request
   curl "http://localhost:8080/hello?name=Budi"

   # SQL Injection attempt (will be blocked)
   curl -G "http://localhost:8080/hello" --data-urlencode "name=1' OR '1'='1"
   ```

## How It Works

1. **Request comes in** → the client IP is extracted from `r.RemoteAddr`
2. **Blacklist check** → if the IP is blacklisted, return `403 Forbidden`
3. **SQL Injection check** → if the input matches dangerous patterns, blacklist the IP and return `403 Forbidden`
4. **Rate limit check** → if the IP exceeds 6 requests per 10 seconds, return `429 Too Many Requests`
5. **Allowed** → return the response and log the request

## Log Format

Each request is appended to `log.txt`:

```text
[2026-10-05 14:30:22] IP: 127.0.0.1 | Request ke-3 | Status: OK
[2026-10-05 14:30:25] IP: 127.0.0.1 | Request ke-7 | Status: BLOCKED
```

## What I Learned

This project was built as part of a 7-day learning sprint to master Go backend fundamentals. Key concepts practiced:

- HTTP server with `net/http`
- State management with `map` and `struct`
- Time-based logic with `time.Now()` and `time.Since()`
- File I/O with `os.OpenFile` and `defer`
- Modular package structure in Go
- Basic security concepts (rate limiting, SQLi detection, blacklisting)

## Future Improvements

- [ ] Replace in-memory storage with Redis
- [ ] Add more attack patterns (XSS, Path Traversal)
- [ ] Add greylist for suspicious (but not confirmed) IPs
- [ ] Store logs in PostgreSQL instead of a file
- [ ] Add monitoring dashboard

## Author

**Muhammad Gilang Ramadhan**

- GitHub: [@your-username](https://github.com/Ramadhan930)
- LinkedIn: [linkedin.com/in/m-Mr](https://linkedin.com/in/m-Mr)

## License

This project is open source and available under the MIT License.