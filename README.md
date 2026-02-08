# 🎓 Student Management API
> **Project Status:** 🟢 Active | **Language:** Go (Golang) | **Database:** SQLite

---

## 🎨 Overview
โปรเจกต์ RESTful API สำหรับระบบจัดการข้อมูลนักเรียน พัฒนาด้วยภาษา **Go** และใช้ **SQLite** เป็นระบบฐานข้อมูล เน้นความเร็ว เรียบง่าย และโครงสร้างโค้ดที่เป็นระเบียบ

---

## ✨ Features
* ✅ **CRUD Operations:** รองรับการ เพิ่ม, อ่าน, แก้ไข และลบข้อมูลครบถ้วน
* ✅ **Auto-Database Setup:** สร้างไฟล์ Database และ Table ให้อัตโนมัติเมื่อเริ่มรัน
* ✅ **JSON Response:** สื่อสารข้อมูลผ่านรูปแบบมาตรฐาน JSON
* ✅ **Clean Architecture:** แยกส่วนการทำงานชัดเจนผ่าน Struct Handlers

---

## 🛠 Tech Stack
| Component | Technology |
| :--- | :--- |
| **Backend** | ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) |
| **Database** | ![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white) |
| **Tools** | Git, cURL, Postman |

---

## 📁 Project Structure
```bash
.
├── main.go           # ⚡ Entry point & API Logic
├── go.mod            # 📦 Module dependencies
├── students.db       # 💾 SQLite Database file (Auto-generated)
└── README.md         # 📄 Project documentation
```

---

🚀 การเริ่มต้นใช้งาน (Getting Started)
1. การติดตั้ง (Installation)
```bash
# คัดลอกโปรเจกต์จาก Repository
git clone <your-repo-url>

# ติดตั้ง Library ต่างๆ ที่จำเป็น
go mod download
```
2. การรันเซิร์ฟเวอร์ (Running the Server)
```bash
go run main.go
```

---

ข้อมูล API (API Documentation)
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| GET | /students | ดึงข้อมูลนักเรียนทั้งหมด |
| GET | /students/{id} | ค้นหานักเรียนด้วย ID |
| POST | /students | เพิ่มข้อมูลนักเรียนใหม่ |
| PUT | /students/{id} | แก้ไขข้อมูลนักเรียน |
| DELETE | /students/{id} | ลบข้อมูลนักเรียน |

---

ตัวอย่าง Request Body (JSON)
```bash
{
  "id": "66090001",
  "name": "Somchai Jaidee",
  "major": "Computer Science",
  "gpa": 3.75
}
```
