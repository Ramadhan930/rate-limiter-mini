RATE-LIMITER-MINI dengan Go

Hari pertama:

Target: HTTP server yang merespons "hello" di endpoint/hello

langkah:
1. import library "fmt", dan "net/http"
2. buat func handler untuk "/hello"
3. di func handler, tulis response "hello"
4. di main daftarkan route/hello ke handler
5. di main jalankan server di post 8080

Hari kedua: 

Target: server yang membatasi request per IP

logika:
- setiap request masuk, ambil IP-nya.
- simpan IP di map: berapa kali dia request.
- jika >5 request dalam 1 menit -> kembalikan 429.
- jika <5 -> izinkan akses.

alur pikir:
1. buat map: map[string]int -> IP sebagai key, counter sebagai value
2. di handler:
    a. ambil ip dari r.RemoteAddr
    b. cek apakah IP ada di map
    c. jika tidak ada, inisialisali dengan 1
    d. jika ada, tambah 1
    e. jika counter > 5, kirim 429
    f. jika < 5, kirim "hello"

Hari ketiga:

Target: Rate limiter yang reset otomatis setiap 1 menit.

aturan:
- Setiap IP maksimal 5 request per menit.
- Setelah 1 menit, counter di-reset.
- Request berikutnya mulai dari 1 lagi.

alur pikir:
1. Buat struct RateLimitInfo { Count int; LastReset time.Time }
2. Buat map: map[string]RateLimitInfo
3. Di handler:
   a. Ambil IP dari r.RemoteAddr
   b. Ambil info dari map (jika ada)
   c. Jika info.LastReset kosong ATAU sudah lewat 1 menit:
      - Reset: info.Count = 0, info.LastReset = time.Now()
   d. info.Count++
   e. Jika info.Count > 5 → 429
   f. Simpan kembali info ke map
   g. Jika aman → tampilkan pesan

Hari keempat:
Setiap request dicatat ke file log.txt dengan format:
```text
[2026-10-05 14:30:22] IP: 127.0.0.1 | Request ke-3 | Status: OK
[2026-10-05 14:30:25] IP: 127.0.0.1 | Request ke-7 | Status: BLOCKED
```
Alur kerja:
1. Buat function terpisah: writeLog(ip string, count int, status string)
2. Di function:
   a. Buka file log.txt (append + create)
   b. Format string: [waktu] IP: xxx | Request ke-x | Status: xxx
   c. Tulis ke file
   d. Tutup file
3. Di handler rateLimiter:
   a. Jika OK → panggil writeLog(ip, info.Count, "OK")
   b. Jika BLOCKED → panggil writeLog(ip, info.Count, "BLOCKED")

Hari kelima:
Server yang bisa:
1. Membaca query parameter dari URL.
2. Mendeteksi pola berbahaya (SQL Injection sederhana).
3. Memblokir IP yang terdeteksi menyerang.
4. Menyimpan blacklist IP di memory.

Alur kerja: 
1. Buat map global: blacklist (map[string]bool)
2. Buat function: isSQLInjection(input string) bool
   - Cek apakah input mengandung pola berbahaya
   - Return true jika mencurigakan
3. Di handler rateLimiter, URUTANNYA:
   a. Ambil IP
   b. CEK BLACKLIST DULU → jika ada, langsung 403 Forbidden
   c. Ambil query parameter "name"
   d. Cek isSQLInjection(name) → jika ya, tambah ke blacklist + 403
   e. Baru lanjut ke rate limiting (logika Hari 3-4)
4. Function isSQLInjection:
   - Lowercase input
   - Cek pola: "or 1=1", "union select", "--", "' or"
   - Return true jika ada yang cocok