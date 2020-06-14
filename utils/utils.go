package utils

import (
	"crypto/rand"
	"fmt"
	"mime/multipart"
	"net/smtp"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

const (
	JPG = "image/jpeg"
	PNG = "image/png"
)

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type H map[string]interface{}

func CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	return err
}

func GetExpiredInterval() int {
	intrvl, _ := strconv.Atoi(os.Getenv("AD_EXPIRED_INTERVAL"))
	return intrvl
}

func GetMode() int {
	mode, _ := strconv.Atoi(os.Getenv("MODE"))
	return mode
}

func GetAmountFee() float64 {
	fee, _ := strconv.ParseFloat(os.Getenv("AMOUNT_FEE"), 64)
	return fee
}

func GetImageDir() string {
	return os.Getenv("IMAGE_DIR")
}

func GetMaxSize() int64 {
	size, _ := strconv.Atoi(os.Getenv("IMAGE_SIZE"))
	return int64(size)
}

func GetPerPage() int {
	perpage, _ := strconv.Atoi(os.Getenv("PER_PAGE"))
	return perpage
}

func GetLimitOffset(page int, offset int) (int, int) {
	limit := (page - 1) * offset
	return limit, offset
}

func IsFiletypeAccepted(files []*multipart.FileHeader) bool {
	// We only have to pass the file header = first 261 bytes

	for _, file := range files {
		f, _ := file.Open()
		defer f.Close()
		mime, err := mimetype.DetectReader(f)

		isAllowed := (mime.String() == JPG || mime.String() == PNG) && err == nil
		if !isAllowed {
			return false
		}
	}

	return true
}

func GetFileExtension(filename string) string {
	return path.Ext(filename)
}

func GenerateNewFile(filename string) string {

	random, _ := RandomFile16Char()
	extension := GetFileExtension(filename)
	return random + extension
}

func IsFileSizeAccepted(filesize int64) bool {
	maxSize := GetMaxSize()
	fsizeKb := filesize / 1024
	return fsizeKb < maxSize
}

func RandomFile16Char() (s string, err error) {
	b := make([]byte, 8)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	s = fmt.Sprintf("%x", b)
	return
}

func ValidateEmail(email string) bool {
	if !EmailRegex.MatchString(email) {
		return false
	}
	return true
}

func GenerateFolder() (string, error) {
	parentDir := GetImageDir()
	year, month, day := time.Now().Date()
	pathdir := fmt.Sprintf("%s/%s/%s/%s/", parentDir, strconv.FormatInt(int64(year), 10), strconv.FormatInt(int64(month), 10), strconv.FormatInt(int64(day), 10))

	err := CreateDirectory(pathdir)
	if err != nil {
		return "", err
	}

	return pathdir, nil
}

func SendMail(to []string, message string) error {
	body := "From: " + os.Getenv("MAIL_USERNAME") + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + os.Getenv("MAIL_SUBJECT") + "\n\n" +
		message

	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))
	smtpAddr := fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT"))

	err := smtp.SendMail(smtpAddr, auth, os.Getenv("MAIL_USERNAME"), to, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
