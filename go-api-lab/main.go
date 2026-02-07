package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import driver แบบพิเศษ (ใช้ _)
)

// --- ส่วนที่ 1: นิยามข้อมูล (Structs) ---
type Student struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Major string  `json:"major"`
	GPA   float64 `json:"gpa"`
}

type StudentHandler struct {
	DB *sql.DB
}

// --- ส่วนที่ 2: ฟังก์ชันหลัก (Main Function) ---
func main() {
	// 1. เชื่อมต่อ Database
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // ปิดฐานข้อมูลเมื่อจบโปรแกรม

	// 2. สร้างตารางถ้ายังไม่มี
	db.Exec(`
    CREATE TABLE IF NOT EXISTS students (
        id TEXT PRIMARY KEY,
        name TEXT,
        major TEXT,
        gpa REAL
    )
    `)

	// 3. เตรียม Handler
	handler := &StudentHandler{DB: db}

	// 4. กำหนดเส้นทาง (Routes)
	http.HandleFunc("/students", handler.studentsHandler)

	http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.studentByIDHandler(w, r)
		case http.MethodPut:
			handler.updateStudentHandler(w, r)
		case http.MethodDelete:
			handler.deleteStudentHandler(w, r)
		}
	})

	// 5. เริ่มรัน Server
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// --- ส่วนที่ 3: ตรรกะของ API (Handler Methods) ---

// สำหรับ GET ทั้งหมด และ POST ข้อมูลใหม่
func (h *StudentHandler) studentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// ตรรกะดึงข้อมูลทั้งหมด (ตาม Step E1-Fixed)
		rows, _ := h.DB.Query("SELECT id, name, major, gpa FROM students")
		defer rows.Close()
		var students []Student
		for rows.Next() {
			var s Student
			rows.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
			students = append(students, s)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)

	} else if r.Method == http.MethodPost {
		// ตรรกะการเพิ่มข้อมูล (ตาม Step E3)
		var s Student
		json.NewDecoder(r.Body).Decode(&s)
		h.DB.Exec("INSERT INTO students (id, name, major, gpa) VALUES (?, ?, ?, ?)",
			s.Id, s.Name, s.Major, s.GPA)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)
	}
}

// สำหรับ GET รายคนตาม ID
func (h *StudentHandler) studentByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Path[len("/students/"):]
		row := h.DB.QueryRow("SELECT id, name, major, gpa FROM students WHERE id = ?", id)
		var s Student
		err := row.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
		if err != nil {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s)
	}
}

// เพิ่มต่อท้ายใน main.go (ภายใต้ StudentHandler)
func (h *StudentHandler) updateStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		// 1. ดึง ID จาก URL เหมือนตอน GET by ID
		id := r.URL.Path[len("/students/"):]

		// 2. แกะข้อมูลใหม่ที่ส่งมาใน Body
		var s Student
		json.NewDecoder(r.Body).Decode(&s)

		// 3. สั่ง Database ให้ Update ข้อมูล
		_, err := h.DB.Exec(
			"UPDATE students SET name = ?, major = ?, gpa = ? WHERE id = ?",
			s.Name, s.Major, s.GPA, id,
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s)
	}
}

func (h *StudentHandler) deleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id := r.URL.Path[len("/students/"):]

		_, err := h.DB.Exec("DELETE FROM students WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusNoContent) // ส่งรหัส 204 (ลบสำเร็จแต่ไม่มีข้อมูลส่งกลับ)
	}
}
