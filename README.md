# Mock Data Generator CLI
Program untuk membuat data dummy berdasarkan template yang diberikan.

## Instalasi
1. Install dengan Go 1.22
```
go install github.com/fastcampus-backend-golang/mockdata@latest 
```

2. Siapkan file template sebagai input, contoh:
```
{
    "nama": "name",
    "tanggalLahir": "date",
    "alamat": "address",
    "telepon": "phone"
}
```

3. Jalankan program
```
mockdata --input ./input.json --output ./output.json
```