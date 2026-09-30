RATE-LIMITER-MINI dengan Go

Hari pertama:

Target: HTTP server yang merespons "hello" di endpoint/hello

langkah:
- import library "fmt", dan "net/http"
- buat func handler untuk "/hello"
- di func handler, tulis response "hello"
- di main daftarkan route/hello ke handler
- di main jalankan server di post 8080

Hari kedua: 

Target: server yang membatasi request per IP

logika:
- setiap request masuk, ambil IP-nya.
- simpan IP di map: berapa kali dia request.
- jika >5 request dalam 1 menit -> kembalikan 429.
- jika <5 -> izinkan akses.

alur pikir:
- buat map: map[string]int -> IP sebagai key, counter sebagai value
- di handler:
    - ambil ip dari r.RemoteAddr
    - cek apakah IP ada di map
    - jika tidak ada, inisialisali dengan 1
    - jika ada, tambah 1
    - jika counter > 5, kirim 429
    - jika < 5, kirim "hello"