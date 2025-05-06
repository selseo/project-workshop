**Theme: Core Banking API** 

**ส่วนที่ 1: Customer Management & Authentication (การจัดการลูกค้าและการยืนยันตัวตน) - 7 Features**

1.  **Feature:** Onboard New Customer (เปิดบัญชีลูกค้าใหม่)
    *   **Endpoint:** `POST /customers/onboard`
    *   **รายละเอียด:** รับข้อมูลลูกค้า (ชื่อ, นามสกุล, เลขบัตรประชาชน, เบอร์โทร, อีเมล, ที่อยู่) และรหัสผ่านเริ่มต้น, สร้าง customer ID, บันทึกลงฐานข้อมูล
2.  **Feature:** Customer Login (ลูกค้าเข้าสู่ระบบ)
    *   **Endpoint:** `POST /auth/login`
    *   **รายละเอียด:** รับ customer ID/username และรหัสผ่าน, ตรวจสอบข้อมูล, และคืน JWT token หากสำเร็จ
3.  **Feature:** Get Current Customer Profile (ดูข้อมูลลูกค้าปัจจุบัน)
    *   **Endpoint:** `GET /customers/me` (ต้องใช้ token)
    *   **รายละเอียด:** คืนข้อมูลลูกค้าที่ login อยู่
4.  **Feature:** Update Customer Contact Information (อัปเดตข้อมูลติดต่อลูกค้า)
    *   **Endpoint:** `PUT /customers/me/contact` (ต้องใช้ token)
    *   **รายละเอียด:** อนุญาตให้ลูกค้าอัปเดตเบอร์โทร, อีเมล
5.  **Feature:** Customer Change Password (ลูกค้าเปลี่ยนรหัสผ่าน)
    *   **Endpoint:** `PUT /auth/change-password` (ต้องใช้ token, รับ old_password, new_password)
    *   **รายละเอียด:** ตรวจสอบรหัสผ่านเก่า และอัปเดตรหัสผ่านใหม่
6.  **Feature (Bank Staff):** Search Customers by Name or ID (ค้นหาลูกค้าด้วยชื่อหรือ ID - สำหรับพนักงาน)
    *   **Endpoint:** `GET /staff/customers/search?q={query}` (ต้องใช้ token พนักงาน)
    *   **รายละเอียด:** ค้นหาลูกค้าด้วยชื่อ, นามสกุล, หรือ customer ID
7.  **Feature (Bank Staff):** View Specific Customer Details (ดูรายละเอียดลูกค้าที่ระบุ - สำหรับพนักงาน)
    *   **Endpoint:** `GET /staff/customers/{customerId}` (ต้องใช้ token พนักงาน)
    *   **รายละเอียด:** แสดงข้อมูลลูกค้าแบบละเอียดสำหรับพนักงาน

**ส่วนที่ 2: Account Management (การจัดการบัญชี) - 7 Features**

8.  **Feature:** Open New Savings Account (เปิดบัญชีเงินฝากออมทรัพย์ใหม่)
    *   **Endpoint:** `POST /accounts/savings` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** สร้างเลขที่บัญชีออมทรัพย์ใหม่ผูกกับลูกค้า, กำหนดสถานะเริ่มต้น
9.  **Feature:** List Customer's Accounts (ดูรายการบัญชีทั้งหมดของลูกค้า)
    *   **Endpoint:** `GET /customers/me/accounts` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** คืนรายการบัญชีทั้งหมด (เลขที่, ประเภท, สถานะ) ของลูกค้าที่ login อยู่
10. **Feature:** Get Specific Account Details (ดูรายละเอียดบัญชีที่ระบุ)
    *   **Endpoint:** `GET /accounts/{accountId}` (ต้องใช้ token ลูกค้า, ตรวจสอบความเป็นเจ้าของบัญชี)
    *   **รายละเอียด:** คืนข้อมูลบัญชี (เลขที่บัญชี, ประเภท, ยอดคงเหลือ, สถานะ)
11. **Feature:** Get Account Balance (ตรวจสอบยอดเงินในบัญชีที่ระบุ)
    *   **Endpoint:** `GET /accounts/{accountId}/balance` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** คืนยอดเงินคงเหลือในบัญชี
12. **Feature:** Set Account Nickname (ตั้งชื่อเล่นให้บัญชี)
    *   **Endpoint:** `PUT /accounts/{accountId}/nickname` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** ลูกค้าตั้งชื่อเล่นให้บัญชีเพื่อง่ายต่อการจดจำ
13. **Feature:** Request Account Statement (ขอบัญชีสรุปยอด)
    *   **Endpoint:** `POST /accounts/{accountId}/statements/request` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** รับช่วงวันที่, สร้างคำขอรายการเดินบัญชี (อาจจะแค่บันทึกคำขอ ไม่ได้ generate file จริง)
14. **Feature (Bank Staff):** Close Account (ปิดบัญชี - สำหรับพนักงาน)
    *   **Endpoint:** `DELETE /staff/accounts/{accountId}` (ต้องใช้ token พนักงาน)
    *   **รายละเอียด:** ตรวจสอบเงื่อนไข (เช่น ยอดคงเหลือเป็นศูนย์), เปลี่ยนสถานะบัญชีเป็น 'closed'

**ส่วนที่ 3: Transaction Processing (การประมวลผลธุรกรรม) - 5 Features**

15. **Feature (Bank Staff/ATM):** Deposit Funds into Account (ฝากเงินเข้าบัญชี)
    *   **Endpoint:** `POST /transactions/deposit`
    *   **รายละเอียด:** รับ `accountId`, `amount`, อัปเดตยอดคงเหลือในบัญชี, สร้างรายการธุรกรรม
16. **Feature:** Withdraw Funds from Account (ถอนเงินจากบัญชี)
    *   **Endpoint:** `POST /transactions/withdraw` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** รับ `accountId`, `amount`, ตรวจสอบยอดคงเหลือ, อัปเดตยอดคงเหลือ, สร้างรายการธุรกรรม
17. **Feature:** Transfer Funds (Intra-bank) (โอนเงินระหว่างบัญชีภายในธนาคาร)
    *   **Endpoint:** `POST /transactions/transfer/internal` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** รับ `fromAccountId`, `toAccountId`, `amount`, `note`, ตรวจสอบยอด, ทำการโอน, สร้างรายการธุรกรรม 2 รายการ
18. **Feature:** Get Account Transaction History (ดูประวัติการทำธุรกรรมของบัญชี)
    *   **Endpoint:** `GET /accounts/{accountId}/transactions` (ต้องใช้ token ลูกค้า, query params: `limit`, `offset`, `startDate`, `endDate`, `type`)
    *   **รายละเอียด:** คืนรายการธุรกรรมของบัญชีนั้นๆ พร้อม pagination และ filter
19. **Feature:** Get Specific Transaction Details (ดูรายละเอียดธุรกรรมที่ระบุ)
    *   **Endpoint:** `GET /transactions/{transactionId}` (ต้องใช้ token ลูกค้า หรือพนักงาน)
    *   **รายละเอียด:** คืนรายละเอียดของธุรกรรมนั้นๆ (ประเภท, จำนวน, วันที่, บัญชีที่เกี่ยวข้อง)

**ส่วนที่ 4: Loan Management (การจัดการสินเชื่อ - แบบง่าย) - 5 Features**

20. **Feature:** Customer Applies for a Personal Loan (ลูกค้ายื่นขอสินเชื่อส่วนบุคคล)
    *   **Endpoint:** `POST /loans/personal/apply` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** รับ `amount_requested`, `purpose`, `income_details`, สร้าง loan application record
21. **Feature:** Customer Checks Loan Application Status (ลูกค้าตรวจสอบสถานะคำขอสินเชื่อ)
    *   **Endpoint:** `GET /customers/me/loan-applications/{applicationId}` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** คืนสถานะ (pending, approved, rejected, more_info_required) ของคำขอสินเชื่อ
22. **Feature (Bank Staff):** List All Pending Loan Applications (ดูรายการคำขอสินเชื่อที่รออนุมัติ - สำหรับพนักงาน)
    *   **Endpoint:** `GET /staff/loans/applications?status=pending` (ต้องใช้ token พนักงาน)
    *   **รายละเอียด:** คืนรายการคำขอสินเชื่อทั้งหมดที่สถานะเป็น 'pending'
23. **Feature (Bank Staff):** Update Loan Application Status (อนุมัติ/ปฏิเสธ คำขอสินเชื่อ - สำหรับพนักงาน)
    *   **Endpoint:** `PUT /staff/loans/applications/{applicationId}/status` (ต้องใช้ token พนักงาน)
    *   **รายละเอียด:** รับสถานะใหม่ (approved, rejected) และเหตุผล, ถ้า approved อาจมีการสร้าง loan account
24. **Feature:** Make Loan Repayment (ชำระค่างวดสินเชื่อ)
    *   **Endpoint:** `POST /loans/{loanId}/repayments` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** รับ `accountId` (บัญชีที่ใช้ชำระ), `amount`, บันทึกการชำระ, อัปเดตยอดหนี้คงค้าง

**ส่วนที่ 5: Card Management (การจัดการบัตร - แบบง่าย) - 6 Features**

25. **Feature:** Customer Requests a New Debit Card for an Account (ลูกค้าขอบัตรเดบิตใหม่สำหรับบัญชี)
    *   **Endpoint:** `POST /accounts/{accountId}/cards/request-debit` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** สร้างคำขอบัตรเดบิตใหม่ผูกกับบัญชี
26. **Feature:** Activate New Card (เปิดใช้งานบัตรใหม่)
    *   **Endpoint:** `PUT /cards/{cardPANLast4Digits}/activate` (ต้องใช้ token ลูกค้า, รับ `activationCode`)
    *   **รายละเอียด:** เปลี่ยนสถานะบัตรเป็น 'active' (ใช้เลขท้าย 4 ตัวของบัตรในการระบุเพื่อความง่าย)
27. **Feature:** Customer Blocks Card Temporarily (ลูกค้าระงับบัตรชั่วคราว)
    *   **Endpoint:** `PUT /cards/{cardPANLast4Digits}/block` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** เปลี่ยนสถานะบัตรเป็น 'temporarily_blocked'
28. **Feature:** Customer Unblocks Card (ลูกค้าปลดระงับบัตร)
    *   **Endpoint:** `PUT /cards/{cardPANLast4Digits}/unblock` (ต้องใช้ token ลูกค้า)
    *   **รายละเอียด:** เปลี่ยนสถานะบัตรกลับเป็น 'active' (จาก 'temporarily_blocked')
29. **Feature (Bank Staff):** View Card Details and Status (ดูรายละเอียดและสถานะบัตร - สำหรับพนักงาน)
    *   **Endpoint:** `GET /staff/cards/search?panLast4={digits}` (ต้องใช้ token พนักงาน)
    *   **รายละเอียด:** คืนข้อมูลบัตร (ประเภท, สถานะ, บัญชีที่ผูก, วันหมดอายุบางส่วน)
30. **Feature (Bank Staff):** Report Card Lost/Stolen (แจ้งบัตรหาย/ถูกขโมย - สำหรับพนักงาน)
    *   **Endpoint:** `PUT /staff/cards/{cardPANLast4Digits}/report-lost-stolen` (ต้องใช้ token พนักงาน)
    *   **รายละเอียด:** เปลี่ยนสถานะบัตรเป็น 'permanently_blocked' หรือ 'lost_stolen'

** หมายเหตุ **

*   **Scaffolding:** เตรียมโครงสร้างโปรเจกต์, database connection, middleware พื้นฐาน (เช่น auth), และ Mocks/Interfaces สำหรับ Dependencies ที่อาจจะยังไม่เสร็จสมบูรณ์ เพื่อให้แต่ละคนเริ่มงานได้เร็ว
*   **Data Models:** กำหนด Data Models (Structs) หลักๆ ให้ชัดเจนตั้งแต่ต้น (Customer, Account, Transaction, LoanApplication, Card) เพื่อลดความขัดแย้ง
*   **API Contract:** ใช้ OpenAPI/Swagger เพื่อกำหนด API contract ให้ทุกคนเห็นภาพเดียวกัน
*   **Helper Functions:** เตรียม Utility/Helper functions ที่ใช้บ่อยๆ (เช่น การ parse request, การสร้าง response chuẩn)
*   **Focus:** ให้แต่ละคนโฟกัสที่ business logic ของ feature ตัวเองเป็นหลักก่อน
*   **Dependencies:** Feature บางอย่างอาจต้องรอ feature อื่น (เช่น การทำธุรกรรมต้องมีบัญชีก่อน) วางแผนลำดับการทำ หรือใช้ mock data ไปก่อน
