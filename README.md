# Thai Baht Text Converter

แปลงตัวเลขทศนิยมเป็นข้อความภาษาไทยสกุลเงินบาท

## ตัวอย่าง

| Input | Output |
|-------|--------|
| 1234 | หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน |
| 33333.75 | สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์ |

## Requirements

- Go 1.23+

## วิธี Run

```bash
# 1. Clone project
git clone https://github.com/yourusername/baht-converter
cd baht-converter

# 2. Download dependencies
go mod tidy

# 3. Run
go run .
```

## โครงสร้างโปรเจกต์

```
baht-converter/
├── main.go        # จุดเริ่มต้นโปรแกรม
├── converter.go   # logic แปลงตัวเลขเป็นภาษาไทย
├── go.mod
└── go.sum
```