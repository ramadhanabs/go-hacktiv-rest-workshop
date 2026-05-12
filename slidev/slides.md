---
theme: none
title: Build Your First Backend API in 90 Minutes
info: |
  Back End Golang Workshop
  Speaker: Ramadhana Bagus — Full-stack Engineer
class: text-center
highlighter: shiki
drawings:
  persist: false
transition: slide-left
mdc: true
fonts:
  sans: DM Sans
  mono: JetBrains Mono
---

<div style="position: absolute; inset: 0; background: url('https://images.unsplash.com/photo-1618318328245-ff25be2ec69e?q=80&w=1920&auto=format&fit=crop') center/cover no-repeat; z-index: 0;"></div>
<div style="position: absolute; inset: 0; background: rgba(0,0,0,0.8); z-index: 1;"></div>
<div style="position: relative; z-index: 2;">

# Build Your First **Backend API** in 90 Minutes

Back End Golang Workshop

<div class="pt-12">
  <span class="px-4 py-2 rounded-lg" style="background: #1a1a1a; border: 1px solid #2a2a2a; color: #b0b0b0;">
    Ramadhana Bagus — Full-stack Engineer
  </span>
</div>

<div class="pt-6" style="display: flex; gap: 16px; justify-content: center; align-items: center;">
  <a href="https://www.linkedin.com/in/ramadhanabagus/" target="_blank" style="color: #b0b0b0 !important; text-decoration: none !important; display: flex; align-items: center; gap: 6px;">
    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433a2.062 2.062 0 0 1-2.063-2.065 2.064 2.064 0 1 1 2.063 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/></svg>
    ramadhanabagus
  </a>
  <a href="https://github.com/ramadhanabs" target="_blank" style="color: #b0b0b0 !important; text-decoration: none !important; display: flex; align-items: center; gap: 6px;">
    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0C5.374 0 0 5.373 0 12c0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23A11.509 11.509 0 0 1 12 5.803c1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576C20.566 21.797 24 17.3 24 12c0-6.627-5.373-12-12-12z"/></svg>
    ramadhanabs
  </a>
</div>

</div>

---
layout: center
class: text-center
---

# Sebelum Mulai — Isi Survey Dulu!

<div class="pt-4">
  <img src="https://api.qrserver.com/v1/create-qr-code/?size=220x220&data=https://survey.bagus.icu" alt="QR Code Survey" style="width: 220px; height: 220px; margin: 0 auto; border-radius: 12px; border: 2px solid #2a2a2a; background: white; padding: 8px;" />
</div>

<div class="pt-4" style="font-size: 1.1em;">

Scan QR atau buka **survey.bagus.icu**

</div>

<v-click>

<div style="background: #1a1a1a; border: 1px solid #2a2a2a; border-radius: 8px; padding: 12px 24px; display: inline-block; margin-top: 12px; font-size: 0.85em; color: #b0b0b0;">
Hasil survey akan muncul di slide berikutnya secara live
</div>

</v-click>

---
layout: iframe
url: https://survey.bagus.icu/presenter?embed
---

---
layout: center
---

# 🤔 "Kan udah ada AI, ngapain coding manual?"

<v-click>

### Karena AI itu GPS — kamu tetap harus bisa nyetir.

</v-click>

<v-click>

- AI bisa generate code, tapi **kamu** yang harus validasi hasilnya
- Kalau error, kamu harus tau **dimana** dan **kenapa**
- Session ini = fondasi supaya kamu jadi **pilot**, bukan penumpang

</v-click>

---
layout: two-cols
---

# Apa yang Kita Bangun Hari Ini

Satu REST API lengkap dari nol, pakai **Go** 🐹

::right::

<v-clicks>

✅ REST API sederhana (GET & POST)

✅ Hands-on setup project dari nol

✅ Memahami request–response & JSON

✅ Auto-dokumentasi API pakai Swagger

✅ Testing & basic performance check

</v-clicks>

---

# Kenapa Go?

<v-clicks>

- **Simple syntax** — lebih mudah dipelajari dari C/C++
- **Compiled & fast** — performa mendekati C, jauh di atas Python/Node
- **Built-in concurrency** — goroutines bikin handling ribuan request jadi gampang
- **Batteries included** — `net/http` udah cukup buat bikin API, tanpa framework
- **Satu binary** — deploy tinggal copy file, no dependency hell

</v-clicks>

---
layout: section
---

# Part 1
## Konsep Request–Response & REST API

---

# Apa itu REST API?

**REST** = Representational State Transfer

<v-clicks>

- Arsitektur standar untuk komunikasi antar aplikasi lewat **HTTP**
- Client kirim **request**, server bales **response**
- Data biasanya dalam format **JSON**
- Setiap resource punya **URL** unik (endpoint)

</v-clicks>

<v-click>

### Analogi sederhana

> Kamu (client) pesan makanan di restoran. Kamu bilang ke pelayan (HTTP) mau apa (request), dapur (server) masak, pelayan bawa hasilnya ke kamu (response).

</v-click>

---

# Selain REST, Ada Apa Lagi?

<div class="grid grid-cols-2 gap-3 mt-1">

<div v-click="1" class="proto-card rest">
<div class="card-title">REST</div>
<div class="card-sub">Hypertext Transfer Protocol</div>
<div class="anim-area">
<div class="node node-client">CLIENT</div>
<div class="pipe">
<div class="packet rest-req"><span class="packet-label">GET /users</span></div>
<div class="packet rest-res"><span class="packet-label">200 OK</span></div>
</div>
<div class="node node-server">SERVER</div>
</div>
<div class="desc">Request-response — <strong>1 tanya, 1 jawab</strong></div>
<div class="use-case">Koneksi dibuka → request → response → ditutup</div>
<div><span class="tag">CRUD</span><span class="tag">Web API</span><span class="tag">Mobile</span></div>
</div>

<div v-click="2" class="proto-card gql">
<div class="card-title">GraphQL</div>
<div class="card-sub">Query Language for APIs</div>
<div class="anim-area">
<div class="node node-client">CLIENT</div>
<div class="pipe">
<div class="packet gql-req"><span class="packet-label">query { }</span></div>
<div class="packet gql-res1"><span class="packet-label">name</span></div>
<div class="packet gql-res2"><span class="packet-label">email</span></div>
<div class="packet gql-res3"><span class="packet-label">avatar</span></div>
</div>
<div class="node node-server">SERVER</div>
</div>
<div class="desc">Client <strong>pilih sendiri</strong> field yang dimau</div>
<div class="use-case">1 request → exact fields, no over/under-fetching</div>
<div><span class="tag">Flexible Query</span><span class="tag">Frontend-driven</span></div>
</div>

<div v-click="3" class="proto-card ws">
<div class="card-title">WebSocket</div>
<div class="card-sub">Full-Duplex Persistent Connection</div>
<div class="anim-area">
<div class="node node-client">CLIENT</div>
<div class="pipe ws-pipe">
<div class="packet ws-pkt ws-out1"></div>
<div class="packet ws-pkt ws-out2"></div>
<div class="packet ws-pkt ws-in1"></div>
<div class="packet ws-pkt ws-in2"></div>
</div>
<div class="node node-server">SERVER</div>
</div>
<div class="desc">Koneksi <strong>dua arah, real-time</strong></div>
<div class="use-case">Koneksi tetap terbuka — kirim/terima kapan saja</div>
<div><span class="tag">Chat</span><span class="tag">Live Notif</span><span class="tag">Game</span></div>
</div>

<div v-click="4" class="proto-card grpc">
<div class="card-title">gRPC / SSE</div>
<div class="card-sub">Server-Initiated Streaming</div>
<div class="anim-area">
<div class="node node-client">CLIENT</div>
<div class="pipe">
<div class="packet grpc-init"><span class="packet-label">subscribe</span></div>
<div class="packet grpc-pkt grpc-s1"></div>
<div class="packet grpc-pkt grpc-s2"></div>
<div class="packet grpc-pkt grpc-s3"></div>
<div class="packet grpc-pkt grpc-s4"></div>
<div class="packet grpc-pkt grpc-s5"></div>
</div>
<div class="node node-server">SERVER</div>
</div>
<div class="desc">Server <strong>push stream</strong> terus-menerus</div>
<div class="use-case">Client subscribe 1x → server kirim data tanpa diminta</div>
<div><span class="tag">Microservices</span><span class="tag">Live Feed</span><span class="tag">Ticker</span></div>
</div>

</div>

---

# Pilih Tools yang Tepat

REST bisa cover **mayoritas** kebutuhan, tapi **bukan berarti cocok untuk semua**

<div class="flex gap-6 mt-4">
<div class="flex-1">

### Pakai REST

<v-clicks>

- CRUD data (buku, user, produk)
- Form submission
- Upload file
- Integrasi antar service yang simpel
- Public API untuk third-party

</v-clicks>

</div>
<div class="flex-1">

### Jangan pakai REST

<v-clicks>

- Chat app real-time → **WebSocket**
- Live score / stock ticker → **SSE**
- Dashboard yang butuh banyak data spesifik → **GraphQL**
- Komunikasi microservice high-throughput → **gRPC**

</v-clicks>

</div>
</div>

<v-click>

<div style="background: #1a1a1a; border: 1px solid #2a2a2a; border-radius: 8px; padding: 12px 16px; margin-top: 16px; font-size: 0.85em;">

**Contoh salah:** Bikin platform chat pakai REST → client harus polling tiap detik → boros resource, delay tinggi. Pakai **WebSocket** = koneksi sekali, pesan langsung sampai.

</div>

</v-click>

---

# Apa itu JSON?

**JSON** = JavaScript Object Notation — format teks untuk **menyimpan & mengirim data**

<div class="flex gap-6 mt-4">
<div class="flex-1">

### Data di dunia nyata

<div style="background: #1a1a1a; border: 1px solid #2a2a2a; border-radius: 8px; padding: 16px; font-size: 0.85em;">

**Nama:** Budi Santoso

**Umur:** 25

**Kota:** Jakarta

**Hobi:** Coding, Membaca

</div>

</div>
<div class="flex-1">

<v-click>

### Dalam format JSON

```json
{
  "nama": "Budi Santoso",
  "umur": 25,
  "kota": "Jakarta",
  "hobi": ["Coding", "Membaca"]
}
```

</v-click>

</div>
</div>

<v-click>

> 💡 JSON itu seperti **KTP digital** — data terstruktur yang bisa dibaca manusia & mesin

</v-click>

---

# Aturan JSON

<div class="flex gap-6">
<div class="flex-1">

### Tipe data yang didukung

```json
{
  "string": "teks dalam kutip",
  "number": 42,
  "decimal": 3.14,
  "boolean": true,
  "null": null,
  "array": [1, 2, 3],
  "object": {
    "nested": "bisa dalam-dalam"
  }
}
```

</div>
<div class="flex-1">

<v-click>

### Aturan penting

- Key **harus** pakai **tanda kutip ganda** `"`
- Tidak boleh trailing comma
- Tidak ada komentar

</v-click>

<v-click>

### Yang salah

```json
{
  nama: "Budi",        // key tanpa kutip
  "umur": 25,          // trailing comma
}
```

</v-click>

</div>
</div>

---

# JSON dalam REST API

<div class="flex gap-6 mt-2">
<div class="flex-1">

### Client kirim JSON (Request Body)

<div style="background: #1a1a1a; border: 1px solid #2a2a2a; border-radius: 8px; padding: 12px;">
<div style="font-size: 0.65em; color: #666; margin-bottom: 6px;">POST /books</div>

```json
{
  "title": "Atomic Habits",
  "author": "James Clear"
}
```

</div>

</div>
<div style="display: flex; align-items: center; font-size: 1.5em; color: #666; padding: 0 8px;">→</div>
<div class="flex-1">

<v-click>

### Server bales JSON (Response Body)

<div style="background: #1a1a1a; border: 1px solid #2a2a2a; border-radius: 8px; padding: 12px;">
<div style="font-size: 0.65em; color: #666; margin-bottom: 6px;">201 Created</div>

```json
{
  "id": 3,
  "title": "Atomic Habits",
  "author": "James Clear"
}
```

</div>

</v-click>

</div>
</div>

<v-click>

> 💡 JSON jadi **bahasa universal** antara client & server — semua bahasa pemrograman bisa baca & tulis JSON

</v-click>

---

# HTTP Methods

| Method | Fungsi | Contoh |
|--------|--------|--------|
| `GET` | Ambil data | Lihat daftar buku |
| `POST` | Kirim data baru | Tambah buku baru |
| `PUT` | Update seluruh data | Edit semua field buku |
| `PATCH` | Update sebagian data | Edit judul buku saja |
| `DELETE` | Hapus data | Hapus buku |

<v-click>

> 💡 Hari ini kita fokus ke **GET** dan **POST** — dua method paling fundamental

</v-click>

---

# Bagaimana GET Request Bekerja?

<div class="flex gap-6 items-start">
<div class="flex-1">

<div class="seq-diagram">
<div class="seq-node left">Browser</div>
<div class="seq-node right">Go Server</div>
<div class="seq-lifeline left"></div>
<div class="seq-lifeline right"></div>

<div class="seq-arrow" style="top: 76px;">
<div class="seq-line to-right" style="animation: seqLineGrow 4s ease-in-out infinite;"></div>
<div class="seq-dot to-right"></div>
<div class="seq-label above" style="left: 30%;">GET /books</div>
<div class="seq-label tag below" style="left: 15%;">Accept: application/json</div>
</div>

<div class="seq-self s1" style="top: 136px;"></div>
<div class="seq-self-label s1" style="top: 138px; right: calc(10% + 80px);">Proses handler</div>

<div class="seq-self s2" style="top: 176px;"></div>
<div class="seq-self-label s2" style="top: 178px; right: calc(10% + 80px);">Ambil data</div>

<div class="seq-arrow" style="top: 226px;">
<div class="seq-line to-left" style="animation: seqLineGrowDelay 4s ease-in-out infinite; border-top: 2px dashed #666; height: 0;"></div>
<div class="seq-dot to-left"></div>
<div class="seq-label above" style="left: 35%;">200 OK</div>
<div class="seq-label tag below" style="left: 10%;">Content-Type: application/json</div>
</div>
</div>

</div>
<div class="flex-1 text-sm pt-4">

**1. Browser kirim request**
- Method: `GET`
- URL: `http://localhost:8080/books`
- Minta data dalam format JSON

**2. Server terima & proses**
- Router cocokkan URL ke handler
- Handler ambil data dari memory/DB

**3. Server kirim response**
- Status `200 OK` = berhasil
- Body berisi array of books dalam JSON

</div>
</div>

---

# Bagaimana POST Request Bekerja?

<div class="flex gap-6 items-start">
<div class="flex-1">

<div class="seq-diagram">
<div class="seq-node left">Browser</div>
<div class="seq-node right">Go Server</div>
<div class="seq-lifeline left"></div>
<div class="seq-lifeline right"></div>

<div class="seq-arrow" style="top: 76px;">
<div class="seq-line to-right" style="animation: seqLineGrow 5s ease-in-out infinite;"></div>
<div class="seq-dot to-right" style="animation: seqDotRight 5s ease-in-out infinite;"></div>
<div class="seq-label above" style="left: 25%;">POST /books + JSON</div>
<div class="seq-label tag below" style="left: 5%;">{"title":"Atomic Habits","author":"James Clear"}</div>
</div>

<div class="seq-self s3" style="top: 136px;"></div>
<div class="seq-self-label s3" style="top: 138px; right: calc(10% + 80px);">Decode JSON</div>

<div class="seq-self s4" style="top: 171px;"></div>
<div class="seq-self-label s4" style="top: 173px; right: calc(10% + 80px);">Validasi</div>

<div class="seq-self s5" style="top: 206px;"></div>
<div class="seq-self-label s5" style="top: 208px; right: calc(10% + 80px);">Simpan</div>

<div class="seq-arrow" style="top: 256px;">
<div class="seq-line to-left" style="animation: seqLineGrowDelay 5s ease-in-out infinite; border-top: 2px dashed #666; height: 0;"></div>
<div class="seq-dot to-left" style="animation: seqDotLeft 5s ease-in-out infinite;"></div>
<div class="seq-label above" style="left: 30%;">201 Created</div>
<div class="seq-label tag below" style="left: 5%;">{"id":3,"title":"Atomic Habits"}</div>
</div>
</div>

</div>
<div class="flex-1 text-sm pt-4">

**1. Browser kirim request + data**
- Method: `POST`
- Body berisi JSON data buku baru
- Header: `Content-Type: application/json`

**2. Server terima & proses**
- Decode JSON dari request body
- Validasi: apakah format benar?
- Simpan buku baru ke storage

**3. Server kirim response**
- Status `201 Created` = resource baru dibuat
- Body berisi data buku + ID baru

</div>
</div>

---

# Raw HTTP vs DevTools — Apa Bedanya?

<div class="flex gap-8">
<div class="flex-1">

### Raw HTTP (yang sebenarnya dikirim)

```http
POST /books HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Accept: application/json
User-Agent: Mozilla/5.0
Content-Length: 52

{"title":"Atomic Habits","author":"James Clear"}
```

<v-click>

```http
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 May 2025 07:00:00 GMT

{"id":3,"title":"Atomic Habits","author":"James Clear"}
```

</v-click>

</div>
<div class="flex-1">

<v-click>

### DevTools (tampilan yang kamu lihat)

<div style="background: #1a1a1a; border: 1px solid #2a2a2a; border-radius: 8px; padding: 12px; font-size: 0.75em; font-family: monospace;">
<div style="color: #666; border-bottom: 1px solid #2a2a2a; padding-bottom: 8px; margin-bottom: 8px;">
<span style="color: #4ade80;">POST</span> &nbsp; /books &nbsp; <span style="color: #4ade80;">201</span> &nbsp; 3ms
</div>
<div style="margin-bottom: 8px;">
<div style="color: #666;">Request Headers</div>
<div style="color: #e0e0e0;">Content-Type: application/json</div>
<div style="color: #e0e0e0;">Accept: application/json</div>
</div>
<div style="margin-bottom: 8px;">
<div style="color: #666;">Request Payload</div>
<div style="color: #e0e0e0;">{ title: "Atomic Habits", author: "James Clear" }</div>
</div>
<div>
<div style="color: #666;">Response</div>
<div style="color: #e0e0e0;">{ id: 3, title: "Atomic Habits", author: "James Clear" }</div>
</div>
</div>

</v-click>

<v-click>

> DevTools cuma **visualisasi** dari raw HTTP — di balik layar, yang dikirim tetap **plain text** seperti sebelah kiri

</v-click>

</div>
</div>

---

# Anatomi HTTP Request

<div class="flex gap-8">
<div class="flex-1">

```http {none|1|2-5|7}{at:1}
POST /books HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Accept: application/json
Authorization: Bearer eyJhbGci...

{"title":"Atomic Habits","author":"James Clear"}
```

</div>
<div class="flex-1 text-sm">

<div v-click="1">

**Request Line** — method + path + versi HTTP

</div>

<div v-click="2">

**Headers** — metadata tentang request

</div>

<div v-click="3">

**Body** — data yang dikirim (opsional)

</div>

</div>
</div>

---

# Versi HTTP — Dari 1.0 sampai 3

| Versi | Tahun | Cara Kerja | Kelebihan |
|-------|-------|------------|-----------|
| **HTTP/1.0** | 1996 | 1 request = 1 koneksi baru | Simpel, tapi lambat |
| **HTTP/1.1** | 1997 | Koneksi bisa di-reuse (keep-alive) | Standar paling lama dipakai |
| **HTTP/2** | 2015 | Multiplexing — banyak request dalam 1 koneksi | Lebih cepat, binary protocol |
| **HTTP/3** | 2022 | Pakai QUIC (UDP), bukan TCP | Tercepat, tahan packet loss |

<v-click>

### Yang paling umum dipakai?

**HTTP/1.1** — masih jadi default di hampir semua API dan tools (curl, Postman, Go `net/http`)

**HTTP/2** — sudah umum di browser & CDN modern (Google, Cloudflare)

</v-click>

<v-click>

> 💡 Untuk workshop ini kita pakai **HTTP/1.1** — format yang kamu lihat di raw request (`POST /books HTTP/1.1`) adalah versi ini

</v-click>

---

# Request Headers — Metadata Request

Headers = **instruksi tambahan** yang dikirim bersama request

| Header | Fungsi | Contoh |
|--------|--------|--------|
| `Content-Type` | Format data yang **dikirim** | `application/json` |
| `Accept` | Format data yang **diminta** | `application/json` |
| `Authorization` | Identitas / token autentikasi | `Bearer eyJhbGci...` |
| `User-Agent` | Info tentang client | `Mozilla/5.0`, `curl/8.1` |
| `Host` | Domain tujuan | `localhost:8080` |

<v-click>

> 💡 Analoginya: headers itu **amplop surat** — berisi info pengirim, penerima, dan jenis isi — sebelum surat dibuka

</v-click>

---

# Request Body — Data yang Dikirim

Body = **isi utama** yang kamu kirim ke server

<div class="flex gap-8">
<div class="flex-1">

### Kapan ada body?

| Method | Body? |
|--------|-------|
| `GET` | Tidak ada |
| `POST` | Ada |
| `PUT` | Ada |
| `PATCH` | Ada |
| `DELETE` | Jarang |

</div>
<div class="flex-1">

### Contoh body (JSON)

```json
{
  "title": "Atomic Habits",
  "author": "James Clear"
}
```

<v-click>

### Format body yang umum

- **JSON** — paling populer untuk API
- **Form Data** — upload file / form HTML
- **XML** — legacy systems

</v-click>

</div>
</div>

<v-click>

> 💡 `Content-Type` header ngasih tau server: "body ini formatnya JSON ya, tolong di-parse sesuai"

</v-click>

---

# HTTP Status Codes yang Penting

| Code | Arti | Kapan Dipakai |
|------|-------|---------------|
| `200` | OK | GET berhasil |
| `201` | Created | POST berhasil bikin resource baru |
| `400` | Bad Request | JSON salah / input invalid |
| `404` | Not Found | Endpoint nggak ada |
| `405` | Method Not Allowed | POST ke endpoint yang cuma terima GET |
| `500` | Internal Server Error | Ada bug di server |

<v-click>

> 💡 Rule of thumb: **2xx** = sukses, **4xx** = salah client, **5xx** = salah server

</v-click>

---

# Response Wrapper Pattern

Best practice: bungkus response dalam format konsisten

```go
type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
```

<v-click>

Contoh response sukses:

```json
{
  "status": "success",
  "message": "Books retrieved",
  "data": [
    {"id": 1, "title": "The Go Programming Language", "author": "..."}
  ]
}
```

</v-click>

<v-click>

Contoh response error:

```json
{
  "status": "error",
  "message": "Invalid JSON format"
}
```

</v-click>

---

# Response Tanpa Wrapper — Kenapa Buruk?

<div class="flex gap-6">
<div class="flex-1">

### Tanpa wrapper

```json
[
  {"id": 1, "title": "Clean Code"},
  {"id": 2, "title": "The Go Programming Language"}
]
```

<v-click>

Kalau error?

```json
"something went wrong"
```

</v-click>

<v-click>

- Format **berubah-ubah** (array vs string vs object)
- Client harus **tebak** struktur response
- Susah bikin error handling yang konsisten

</v-click>

</div>
<div class="flex-1">

<v-click>

### Dengan wrapper

```json
{
  "status": "success",
  "message": "Books retrieved",
  "data": [
    {"id": 1, "title": "Clean Code"},
    {"id": 2, "title": "The Go Programming Language"}
  ]
}
```

```json
{
  "status": "error",
  "message": "Something went wrong"
}
```

- Format **selalu sama** — predictable
- Cek `status` dulu, baru proses `data`
- Frontend & mobile dev senang

</v-click>

</div>
</div>

---
layout: section
---

# Part 2
## Setup Project dari Nol

---

# Prerequisites

Pastikan sudah terinstall:

```bash
# Cek Go version
go version
# Output: go version go1.22.x ...

# Cek code editor (VS Code recommended)
code --version
```

<v-click>

### Buat project baru

```bash
mkdir my-first-api
cd my-first-api
go mod init my-first-api
```

</v-click>

<v-click>

> 💡 `go mod init` = mirip `npm init` di Node.js — bikin module baru

</v-click>

---

# Struktur Project

```
my-first-api/
├── go.mod          # dependency management
├── go.sum          # dependency checksums
└── main.go         # entry point kita
```

<v-click>

### Hello World Server

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Backend! 🚀")
    })

    fmt.Println("Server running di :8080")
    http.ListenAndServe(":8080", nil)
}
```

</v-click>

---

# Jalankan Server Pertama Kamu

```bash
go run main.go
# Server running di :8080
```

<v-click>

Buka browser → `http://localhost:8080`

Atau test pakai curl:

```bash
curl http://localhost:8080
# Hello, Backend! 🚀
```

</v-click>

<v-click>

### 🎉 Selamat — kamu udah punya web server!

Sekarang kita bikin jadi proper REST API.

</v-click>

---
layout: section
---

# Part 3
## Membuat REST API — GET & POST

---

# Data Model

Kita bikin API untuk manage **daftar buku** 📚

```go
type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// In-memory storage (simpan di RAM dulu, belum database)
var books = []Book{
    {ID: 1, Title: "The Go Programming Language", Author: "Donovan & Kernighan"},
    {ID: 2, Title: "Clean Code", Author: "Robert C. Martin"},
}
```

<v-click>

> 💡 **json tag** (`json:"id"`) = ngasih tau Go gimana format field ini pas jadi JSON

</v-click>

---

# GET Endpoint — Ambil Semua Buku

```go
func getBooks(w http.ResponseWriter, r *http.Request) {
    // Set header supaya client tau ini JSON
    w.Header().Set("Content-Type", "application/json")

    // Encode slice books jadi JSON, langsung tulis ke response
    json.NewEncoder(w).Encode(books)
}
```

<v-click>

Register handler di `main()`:

```go
func main() {
    http.HandleFunc("/books", getBooks)

    fmt.Println("Server running di :8080")
    http.ListenAndServe(":8080", nil)
}
```

</v-click>

<v-click>

```bash
curl http://localhost:8080/books
# [{"id":1,"title":"The Go Programming Language","author":"Donovan & Kernighan"}, ...]
```

</v-click>

---

# POST Endpoint — Tambah Buku Baru

```go
func addBook(w http.ResponseWriter, r *http.Request) {
    // Cek method — harus POST
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var newBook Book

    // Decode JSON dari request body
    err := json.NewDecoder(r.Body).Decode(&newBook)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Auto-generate ID
    newBook.ID = len(books) + 1
    books = append(books, newBook)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // 201
    json.NewEncoder(w).Encode(newBook)
}
```

---

# Test POST Endpoint

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": "Atomic Habits", "author": "James Clear"}'
```

<v-click>

Response:

```json
{
  "id": 3,
  "title": "Atomic Habits",
  "author": "James Clear"
}
```

</v-click>

<v-click>

Verifikasi — GET lagi:

```bash
curl http://localhost:8080/books
# Sekarang ada 3 buku ✅
```

</v-click>

---
layout: section
---

# Part 3.5
## Enhance — Kode yang Lebih Production-Ready

---

# Masalah di Kode Kita Sekarang

<v-clicks>

- ❌ **Routing nggak jelas** — GET & POST dicampur, harus cek `r.Method` manual
- ❌ **Bisa tambah buku tanpa title/author** — nggak ada validasi input
- ❌ **Response nggak konsisten** — sukses beda format sama error
- ❌ **Nggak pakai Response Wrapper** — padahal di Part 1 udah belajar
- ❌ **Error handling pakai `http.Error`** — cuma plain text, bukan JSON

</v-clicks>

<v-click>

> Kita udah punya API yang **jalan**, sekarang kita bikin **bener**

</v-click>

---

# Step 1 — Terapkan Response Wrapper

Ingat `APIResponse` dari Part 1? Sekarang kita pakai beneran:

```go
type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, status int, resp APIResponse) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(resp)
}
```

<v-click>

> 💡 `sendJSON` = helper function supaya nggak nulis 3 baris yang sama berulang-ulang

</v-click>

---

# Step 2 — Go 1.22 Method Routing

Go 1.22+ bisa langsung define **method di pattern** — nggak perlu cek `r.Method` manual:

<div class="flex gap-6">
<div class="flex-1">

### Sebelum ❌

```go
http.HandleFunc("/books", getBooks)

// Di dalam handler:
if r.Method != http.MethodGet {
    // manual check setiap handler 😩
}
```

</div>
<div class="flex-1">

<v-click>

### Sesudah ✅ (Go 1.22+)

```go
http.HandleFunc("GET /books", getBooks)
http.HandleFunc("POST /books", addBook)
```

- Method **salah**? Go otomatis bales `405 Method Not Allowed`
- Handler jadi **bersih** — fokus ke logic, bukan routing

</v-click>

</div>
</div>

<v-click>

### Handler jadi lebih clean

```go
func getBooks(w http.ResponseWriter, r *http.Request) {
    // Nggak perlu cek method — Go yang handle
    sendJSON(w, http.StatusOK, APIResponse{
        Status: "success", Message: "Books retrieved", Data: books,
    })
}
```

</v-click>

---

# Step 3 — Validasi Input pakai Struct Tags

POST tanpa validasi = bahaya:

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": ""}'
# Buku tanpa judul dan tanpa author masuk ke storage! 😱
```

<v-click>

### ❌ Cara manual — nggak scalable

```go
if newBook.Title == "" || newBook.Author == "" { ... }
// Kalau field-nya 20? Nulis 20 if? 😵
```

</v-click>

<v-click>

### ✅ Cara Go — pakai `validate` struct tags

```bash
go get github.com/go-playground/validator/v10
```

```go
type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title" validate:"required"`
    Author string `json:"author" validate:"required"`
}

var validate = validator.New()
```

> 💡 Sama seperti `json:"..."`, tag `validate:"required"` langsung nempel di struct — **satu sumber kebenaran**

</v-click>

---

# Validasi Struct Tags — Cara Pakainya

```go
var newBook Book
json.NewDecoder(r.Body).Decode(&newBook)

// Satu baris — validasi semua field sekaligus
if err := validate.Struct(newBook); err != nil {
    sendJSON(w, http.StatusBadRequest, APIResponse{
        Status:  "error",
        Message: err.Error(),
    })
    return
}
```

<v-click>

### Tag yang sering dipakai

| Tag | Fungsi | Contoh |
|-----|--------|--------|
| `required` | Wajib diisi, nggak boleh zero value | `validate:"required"` |
| `min` / `max` | Panjang minimum/maksimum | `validate:"min=3,max=100"` |
| `email` | Harus format email valid | `validate:"required,email"` |
| `oneof` | Harus salah satu dari pilihan | `validate:"oneof=fiction non-fiction"` |

</v-click>

<v-click>

> 💡 Tambah field baru? Tinggal tambahin tag di struct — nggak perlu edit handler sama sekali

</v-click>

---

# Step 4 — Refactor POST Endpoint (Full)

```go
func addBook(w http.ResponseWriter, r *http.Request) {
    var newBook Book

    // 1. Decode JSON
    if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
        sendJSON(w, http.StatusBadRequest, APIResponse{
            Status: "error", Message: "Invalid JSON format",
        })
        return
    }

    // 2. Validasi otomatis berdasarkan struct tags
    if err := validate.Struct(newBook); err != nil {
        sendJSON(w, http.StatusBadRequest, APIResponse{
            Status: "error", Message: err.Error(),
        })
        return
    }

    // 3. Simpan & respond
    newBook.ID = len(books) + 1
    books = append(books, newBook)

    sendJSON(w, http.StatusCreated, APIResponse{
        Status: "success", Message: "Book created", Data: newBook,
    })
}
```

> Handler cuma urus **logic** — routing & method checking udah di-handle Go

---

# Perbandingan Response: Sebelum vs Sesudah

<div class="flex gap-6">
<div class="flex-1">

### Sebelum ❌

GET sukses:
```json
[{"id":1,"title":"..."}]
```

POST error:
```
Invalid JSON
```

Format **beda-beda**, client harus tebak

</div>
<div class="flex-1">

<v-click>

### Sesudah ✅

GET sukses:
```json
{
  "status": "success",
  "message": "Books retrieved",
  "data": [{"id":1,"title":"..."}]
}
```

POST error:
```json
{
  "status": "error",
  "message": "Key: 'Book.Title' Error:Field
    validation for 'Title' failed on
    the 'required' tag"
}
```

Format **selalu sama** — predictable!

</v-click>

</div>
</div>

---

# Kode Lengkap Setelah Enhancement

```go {*}{maxHeight:'380px'}
package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/go-playground/validator/v10"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title" validate:"required"`
    Author string `json:"author" validate:"required"`
}

type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

var books = []Book{
    {ID: 1, Title: "The Go Programming Language", Author: "Donovan & Kernighan"},
    {ID: 2, Title: "Clean Code", Author: "Robert C. Martin"},
}

var validate = validator.New()

func sendJSON(w http.ResponseWriter, status int, resp APIResponse) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(resp)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    sendJSON(w, http.StatusOK, APIResponse{
        Status: "success", Message: "Books retrieved", Data: books,
    })
}

func addBook(w http.ResponseWriter, r *http.Request) {
    var newBook Book
    if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
        sendJSON(w, http.StatusBadRequest, APIResponse{
            Status: "error", Message: "Invalid JSON format",
        })
        return
    }
    if err := validate.Struct(newBook); err != nil {
        sendJSON(w, http.StatusBadRequest, APIResponse{
            Status: "error", Message: err.Error(),
        })
        return
    }
    newBook.ID = len(books) + 1
    books = append(books, newBook)
    sendJSON(w, http.StatusCreated, APIResponse{
        Status: "success", Message: "Book created", Data: newBook,
    })
}

func main() {
    // Go 1.22+ method routing — method salah = otomatis 405
    http.HandleFunc("GET /books", getBooks)
    http.HandleFunc("POST /books", addBook)

    fmt.Println("Server running di :8080")
    http.ListenAndServe(":8080", nil)
}
```

---

# Test Semua Skenario

```bash
# ✅ GET sukses
curl http://localhost:8080/books

# ✅ POST sukses
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": "Atomic Habits", "author": "James Clear"}'

# ❌ POST tanpa title — validator tangkap
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"author": "James Clear"}'

# ❌ JSON invalid
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d 'ini bukan json'

# ❌ Method salah — Go otomatis bales 405
curl -X DELETE http://localhost:8080/books
# 405 Method Not Allowed (tanpa kita tulis kode apapun!)
```

<v-click>

> Go 1.22 routing handle method mismatch, validator handle input — handler kita cuma fokus ke **business logic**

</v-click>

---

# net/http vs Framework — Kapan Harus Upgrade?

Kita pakai `net/http` (standard library) — ini **bukan kekurangan**, ini **fondasi**

<v-click>

### Yang harus kita bikin sendiri pakai net/http

- ✍️ Helper `sendJSON` — encode response manual
- ✍️ Validasi — install `validator` sendiri
- ✍️ Middleware — logging, CORS, auth harus di-chain manual
- ✍️ Route groups — nggak ada `/api/v1/...` grouping
- ✍️ Path params — `GET /books/{id}` baru ada di Go 1.22, parsing manual

</v-click>

<v-click>

> Untuk belajar, `net/http` itu **perfect** — kamu paham apa yang terjadi di balik layar. Tapi untuk production, ada tools yang bikin hidup lebih mudah.

</v-click>

---

# Framework Populer di Go

<div class="flex gap-6">
<div class="flex-1">

### Gin — Yang Paling Populer

<v-clicks>

- **77k+ GitHub stars** — komunitas terbesar
- Built-in JSON binding + validation
- Middleware ecosystem yang kaya (CORS, auth, rate limit)
- Performa sangat cepat (httprouter-based)
- Cocok untuk: **REST API, microservices**

</v-clicks>

</div>
<div class="flex-1">

<v-click>

### Alternatif lain

| Framework | Keunggulan |
|-----------|------------|
| **Echo** | Minimalis, performa setara Gin |
| **Fiber** | Terinspirasi Express.js, paling cepat |
| **Chi** | Lightweight, 100% kompatibel net/http |

</v-click>

<v-click>

> **Rekomendasi:** Mulai dengan **Gin** — dokumentasi lengkap, banyak tutorial, dan jadi standar industri di ekosistem Go

</v-click>

</div>
</div>

---

# Perbandingan: net/http vs Gin

<div class="flex gap-6">
<div class="flex-1">

### net/http (yang kita pakai)

```go
func addBook(w http.ResponseWriter,
    r *http.Request) {
    var newBook Book
    err := json.NewDecoder(r.Body).
        Decode(&newBook)
    if err != nil {
        sendJSON(w, 400, APIResponse{
            Status:  "error",
            Message: "Invalid JSON",
        })
        return
    }
    if err := validate.Struct(newBook);
        err != nil {
        sendJSON(w, 400, APIResponse{
            Status:  "error",
            Message: err.Error(),
        })
        return
    }
    // ... simpan & response
}
```

</div>
<div class="flex-1">

<v-click>

### Gin

```go
func addBook(c *gin.Context) {
    var newBook Book

    // Bind + validasi dalam 1 langkah
    if err := c.ShouldBindJSON(&newBook);
        err != nil {
        c.JSON(400, gin.H{
            "status":  "error",
            "message": err.Error(),
        })
        return
    }

    // ... simpan & response
    c.JSON(201, gin.H{
        "status":  "success",
        "message": "Book created",
        "data":    newBook,
    })
}
```

</v-click>

</div>
</div>

<v-click>

> `ShouldBindJSON` = decode JSON + validasi struct tags **dalam 1 baris** — nggak perlu `validator` terpisah

</v-click>

---
layout: section
---

# Part 4
## Dokumentasi API dengan Swagger

---

# Install Swag CLI

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

<v-click>

Tambah dependency:

```bash
go get github.com/swaggo/http-swagger
go get github.com/swaggo/swag
```

</v-click>

---

# Tambah Swagger Annotations

Response kita pakai `APIResponse` wrapper — annotations harus reflect itu:

```go
// @Summary      Get all books
// @Description  Ambil semua daftar buku
// @Tags         books
// @Produce      json
// @Success      200  {object}  APIResponse{data=[]Book}
// @Router       /books [get]
func getBooks(w http.ResponseWriter, r *http.Request) {
    // ... kode yang sama
}
```

<v-click>

```go
// @Summary      Add a new book
// @Description  Tambah buku baru ke collection
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      Book  true  "Book object"
// @Success      201   {object}  APIResponse{data=Book}
// @Failure      400   {object}  APIResponse
// @Router       /books [post]
func addBook(w http.ResponseWriter, r *http.Request) {
    // ... kode yang sama
}
```

</v-click>

<v-click>

> 💡 `APIResponse{data=[]Book}` = Swagger tau response dibungkus wrapper, field `data` isinya array of Book

</v-click>

---

# Tambah Main Annotations

Swagger butuh **metadata aplikasi** — taruh di atas fungsi `main()`:

```go
// @title           My First API
// @version         1.0
// @description     REST API buku sederhana pakai Go

// @host            localhost:8080
// @BasePath        /
func main() {
    // ...
}
```

<v-click>

> 💡 Annotations ini jadi judul, deskripsi, dan base URL di Swagger UI — tanpa ini, Swagger nggak tau API kamu tentang apa

</v-click>

---

# Generate & Serve Swagger UI

```bash
# Generate docs dari annotations
swag init
```

<v-click>

Ini akan bikin folder `docs/` berisi file Swagger JSON.

Mount Swagger UI di `main.go`:

```go
import httpSwagger "github.com/swaggo/http-swagger"
import _ "my-first-api/docs" // generated docs

// @title           My First API
// @version         1.0
// @description     REST API buku sederhana pakai Go
// @host            localhost:8080
// @BasePath        /
func main() {
    http.HandleFunc("GET /books", getBooks)
    http.HandleFunc("POST /books", addBook)

    // Swagger UI endpoint
    http.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

    fmt.Println("Server running di :8080")
    fmt.Println("Swagger UI → http://localhost:8080/swagger/")
    http.ListenAndServe(":8080", nil)
}
```

</v-click>

<v-click>

### 🎉 Buka `http://localhost:8080/swagger/` — API kamu sekarang punya dokumentasi interaktif!

</v-click>

---
layout: section
---

# Part 5
## Testing & Performance

---

# Unit Test dengan httptest

Buat file `main_test.go`:

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetBooks(t *testing.T) {
    req := httptest.NewRequest("GET", "/books", nil)
    w := httptest.NewRecorder()

    getBooks(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected 200, got %d", w.Code)
    }

    if w.Header().Get("Content-Type") != "application/json" {
        t.Error("Expected JSON content type")
    }
}
```

---

# Jalankan Test

```bash
go test -v
```

<v-click>

Output:

```
=== RUN   TestGetBooks
--- PASS: TestGetBooks (0.00s)
PASS
ok      my-first-api    0.003s
```

</v-click>

<v-click>

### Quick Performance Check

```bash
# Install hey (HTTP load generator)
go install github.com/rakyll/hey@latest

# Blast 1000 requests, 50 concurrent
hey -n 1000 -c 50 http://localhost:8080/books
```

> 💡 Go bisa handle **ribuan request per detik** bahkan di laptop biasa — goroutines FTW

</v-click>

---
layout: center
class: text-center
---

# Recap — Apa yang Kamu Bangun Hari Ini

<v-clicks>

✅ Setup Go project dari nol

✅ REST API dengan GET & POST endpoint

✅ Paham alur request–response & JSON

✅ Auto-generated Swagger documentation

✅ Unit testing & performance benchmarking

</v-clicks>

---
layout: center
---

# Next Steps 🚀

<v-clicks>

- 🗄️ **Database** — Ganti in-memory storage dengan PostgreSQL
- 🔐 **Authentication** — Tambah JWT middleware
- 🐳 **Docker** — Containerize API kamu
- ☁️ **Deploy** — Push ke cloud (Railway, Fly.io, VPS)
- 🤖 **AI-Assisted Dev** — Pakai AI buat accelerate, bukan replace

</v-clicks>

---
layout: center
class: text-center
---

# Terima Kasih! 🙏

<div style="color: #ffffff; font-size: 1.2em; font-weight: 600;">Ramadhana Bagus</div>
<div style="color: #666666;">Full-stack Engineer</div>

<br>

### Ada pertanyaan?

<br>

<div style="background: #1a1a1a; border: 1px solid #2a2a2a; border-radius: 12px; padding: 16px 24px; display: inline-block;">
  🎓 Dapatkan penawaran spesial program <strong>Night Bootcamp Hacktiv8</strong> hari ini!
</div>
