# Aplikasi Restoran

Aplikasi ini adalah aplikasi backend untuk mengelola menu restoran, pesanan, dan ulasan dari pengguna. Aplikasi ini dibuat menggunakan bahasa Go (Golang). Aplikasi mendukung berbagai peran pengguna, termasuk Admin, Chef, dan Customer, masing-masing dengan izin tertentu.

## Fitur

- **Autentikasi Pengguna**: Pengguna dapat login dan logout dari sistem.
- **Manajemen Pesanan**: Customer dapat membuat pesanan, sedangkan Admin dan Chef dapat memperbarui status pesanan.
- **Manajemen Menu**: Admin dapat menambahkan item menu baru.
- **Sistem Penilaian**: Customer dapat memberikan rating pada pesanan yang telah selesai.
- **Kontrol Akses Berdasarkan Peran**: Izin yang berbeda untuk peran Admin, Chef, dan Customer.

## Penggunaan

Setelah aplikasi dijalankan, Anda akan diminta untuk memasukkan endpoint tertentu sebagai perintah. Berikut cara menggunakan setiap endpoint:

### Daftar Endpoint

- **login** : Autentikasi pengguna.
- **logout** : Logout dari sesi saat ini.
- **add-order** : Customer dapat membuat pesanan baru.
- **update-status** : Admin atau Chef dapat memperbarui status pesanan.
- **get-order** : Pengguna dapat melihat detail pesanan sesuai perannya.
- **add-menu** : Admin dapat menambahkan item menu baru.
- **delete-order** : Admin dapat menghapus pesanan.
- **add-rating** : Customer dapat memberikan rating pada pesanan yang selesai.
- **get-rating** : Mendapatkan ulasan, dapat diakses oleh Admin, Chef, dan Customer.


1. **login**

**_Izin Berdasarkan Peran_**

- **Admin**:
  - Dapat melihat semua pesanan, menambah dan menghapus pesanan, menambahkan item menu, dan melihat semua ulasan.
- **Chef**:
  - Dapat melihat semua pesanan dan memperbarui status pesanan.
- **Customer**:
  - Dapat membuat pesanan, melihat pesanan miliknya, dan memberikan rating pada pesanan yang selesai atau dibatalkan.

```json
{
  "username": "admin",
  "password": "hashed_password1"
}
```

```json
{
  "username": "chef",
  "password": "hashed_password2"
}
```

```json
{
  "username": "customer1",
  "password": "hashed_password3"
}
```

2. **add-order** Customer dapat membuat pesanan baru.

```json
{
  "items": [{ "menu_item_id": 1, "quantity": 3 }],
  "discount_id": 0
}
```

3. **update-status** Admin atau Chef dapat memperbarui status pesanan.

```json
{
  "id": 1,
  "status": "Canceled"
}
```

4. **add-menu** Admin dapat menambahkan item menu baru.

```json
{
  "name": "Tacos",
  "price": 12
}
```

5. **delete-order**: Admin dapat menghapus pesanan.

```json
{
  "id": 1
}
```

6. **add-rating**: Customer dapat memberikan rating pada pesanan yang selesai.

```json
{
  "order_id": 1,
  "rating": 4,
  "comment": "Good Food!"
}
```
