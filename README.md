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
