package services

import (
	"github.com/nanpipat/nanpipat-agnos-backend/internal/core"
	"github.com/nanpipat/nanpipat-agnos-backend/models"
	"github.com/nanpipat/nanpipat-agnos-backend/repo"
	"github.com/nanpipat/nanpipat-agnos-backend/requests"
	"github.com/nanpipat/nanpipat-agnos-backend/views"
)

type IPasswordService interface {
	LogPasswordCheck(payload *requests.PasswordRequest) (*views.PasswordResponse, error)
}

type passwordService struct {
	ctx core.IContext
}

func NewPasswordService(ctx core.IContext) IPasswordService {
	return &passwordService{
		ctx: ctx,
	}
}

func (s *passwordService) LogPasswordCheck(payload *requests.PasswordRequest) (*views.PasswordResponse, error) {
	steps := calculateSteps(payload.InitPassword)

	logs := &models.PasswordLog{
		BaseModel: models.NewBaseModel(),
		Request:   payload.InitPassword,
		Response:  steps,
	}

	err := repo.New[models.PasswordLog](s.ctx).Create(logs)
	if err != nil {
		return nil, err
	}

	// ส่ง response ผ่านการเรียกจาก db เพื่อเช็คการทำงานของ db ว่าถูกต้องหรือไม่
	log, err := repo.New[models.PasswordLog](s.ctx).FindOne("id = ?", logs.ID)
	if err != nil {
		return nil, err
	}

	// ถ้าไม่เช็ตจาก db ก็ได้เหมือนกัน ก็จะเป็นแบบนี้ที่ return steps กลับไปเลย
	// return &views.PasswordResponse{
	// 	NumOfSteps: steps,
	// }, nil

	return &views.PasswordResponse{
		NumOfSteps: log.Response,
	}, nil
}

func calculateSteps(password string) int {
	lengthSteps := adjustLength(password)
	charTypeSteps := ensureCharacterTypes(password)
	repeatSteps := removeRepeatingChars(password)

	n := len(password)
	if n < 6 {
		return max(lengthSteps, charTypeSteps+repeatSteps)
	}
	// สำหรับรหัสผ่านที่ยาวเกิน 20 ตัว รวมทุกขั้นตอนเข้าด้วยกัน
	return lengthSteps + charTypeSteps + repeatSteps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func adjustLength(password string) int {
	n := len(password)
	if n < 6 {
		return 6 - n // ต้องเพิ่มตัวอักษรให้ครบ 6 ตัว
	}
	if n > 20 {
		return 1 // ต้องลบตัวอักษรให้เหลือ 20 ตัว (นับเป็น 1 ขั้นตอน)
	}
	return 0
}

func ensureCharacterTypes(password string) int {
	hasLower, hasUpper, hasDigit := false, false, false

	for _, char := range password {
		switch {
		case 'a' <= char && char <= 'z':
			hasLower = true
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case '0' <= char && char <= '9':
			hasDigit = true
		}
	}

	missing := 0
	if !hasLower {
		missing++
	}
	if !hasUpper {
		missing++
	}
	if !hasDigit {
		missing++
	}

	return missing
}

func removeRepeatingChars(password string) int {
	repeats := 0
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i] == password[i+2] {
			repeats++
			i += 2
		}
	}
	return repeats
}
